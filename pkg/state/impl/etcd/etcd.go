// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package etcd provides an implementation of state.State in etcd.
package etcd

import (
	"context"
	"fmt"
	"sort"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/state"
	"github.com/cosi-project/runtime/pkg/state/impl/store"
)

// State implements state.CoreState.
type State struct {
	cli       *clientv3.Client
	marshaler store.Marshaler
}

// NewState creates new State with default options.
func NewState(cli *clientv3.Client) *State {
	return NewStateWithOptions(cli)
}

// NewStateWithOptions returns state builder function with options.
func NewStateWithOptions(cli *clientv3.Client, opts ...StateOption) *State {
	options := DefaultStateOptions()

	for _, opt := range opts {
		opt(&options)
	}

	return &State{
		cli:       cli,
		marshaler: store.ProtobufMarshaler{},
	}
}

// Get a resource.
func (st *State) Get(ctx context.Context, resourcePointer resource.Pointer, opts ...state.GetOption) (resource.Resource, error) { //nolint:ireturn
	var options state.GetOptions

	for _, opt := range opts {
		opt(&options)
	}

	etcdKey := etcdKeyFromPointer(resourcePointer)

	resp, err := st.cli.Get(ctx, etcdKey)
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, ErrNotFound(resourcePointer)
	}

	return st.unmarshalResource(resp.Kvs[0])
}

// List resources.
func (st *State) List(ctx context.Context, resourceKind resource.Kind, opts ...state.ListOption) (resource.List, error) {
	var options state.ListOptions

	for _, opt := range opts {
		opt(&options)
	}

	etcdKey := etcdKeyPrefixFromKind(resourceKind)

	resp, err := st.cli.Get(ctx, etcdKey, clientv3.WithPrefix())
	if err != nil {
		return resource.List{}, err
	}

	resources := make([]resource.Resource, 0, len(resp.Kvs))

	for _, kv := range resp.Kvs {
		res, err := st.unmarshalResource(kv)
		if err != nil {
			return resource.List{}, err
		}

		if !options.LabelQuery.Matches(*res.Metadata().Labels()) {
			continue
		}

		resources = append(resources, res)
	}

	return resource.List{
		Items: resources,
	}, nil
}

// Create a resource.
func (st *State) Create(ctx context.Context, resource resource.Resource, opts ...state.CreateOption) error {
	resource = resource.DeepCopy()

	var options state.CreateOptions

	for _, opt := range opts {
		opt(&options)
	}

	etcdKey := etcdKeyFromPointer(resource.Metadata())

	if err := resource.Metadata().SetOwner(options.Owner); err != nil {
		return err
	}

	data, err := st.marshaler.MarshalResource(resource)
	if err != nil {
		return err
	}

	resp, err := st.cli.Txn(ctx).If(
		clientv3.Compare(clientv3.Version(etcdKey), "=", 0), // not exists check
	).Then(
		clientv3.OpPut(etcdKey, string(data)),
	).Commit()
	if err != nil {
		return err
	}

	if !resp.Succeeded {
		return ErrAlreadyExists(resource.Metadata())
	}

	return err
}

// Update a resource.
func (st *State) Update(ctx context.Context, curVersion resource.Version, newResource resource.Resource, opts ...state.UpdateOption) error {
	newResource = newResource.DeepCopy()

	options := state.DefaultUpdateOptions()

	for _, opt := range opts {
		opt(&options)
	}

	etcdKey := etcdKeyFromPointer(newResource.Metadata())

	newResource.Metadata().SetUpdated(time.Now())

	data, err := st.marshaler.MarshalResource(newResource)
	if err != nil {
		return err
	}

	getResp, err := st.cli.Get(ctx, etcdKey)
	if err != nil {
		return err
	}

	if len(getResp.Kvs) == 0 {
		return ErrNotFound(newResource.Metadata())
	}

	curResource, err := st.unmarshalResource(getResp.Kvs[0])
	if err != nil {
		return err
	}

	expectedVersion := int64(curVersion.Value())

	etcdVersion := getResp.Kvs[0].Version

	if !curVersion.Equal(resource.VersionUndefined) && etcdVersion != expectedVersion {
		return ErrVersionConflict(curResource.Metadata(), expectedVersion, etcdVersion)
	}

	if curResource.Metadata().Owner() != options.Owner {
		return ErrOwnerConflict(curResource.Metadata(), curResource.Metadata().Owner())
	}

	if options.ExpectedPhase != nil && curResource.Metadata().Phase() != *options.ExpectedPhase {
		return ErrPhaseConflict(curResource.Metadata(), *options.ExpectedPhase)
	}

	txnResp, err := st.cli.Txn(ctx).If(
		clientv3.Compare(clientv3.Version(etcdKey), "=", etcdVersion),
	).Then(
		clientv3.OpPut(etcdKey, string(data)),
	).Else(
		clientv3.OpGet(etcdKey),
	).Commit()
	if err != nil {
		return err
	}

	if txnResp.Succeeded {
		return nil
	}

	var foundVersion int64

	txnGetKvs := txnResp.Responses[0].GetResponseRange().Kvs
	if len(txnResp.Responses[0].GetResponseRange().Kvs) > 0 {
		foundVersion = txnGetKvs[0].Version
	}

	return ErrVersionConflict(newResource.Metadata(), expectedVersion, foundVersion)
}

// Destroy a resource.
func (st *State) Destroy(ctx context.Context, resourcePointer resource.Pointer, opts ...state.DestroyOption) error {
	var options state.DestroyOptions

	for _, opt := range opts {
		opt(&options)
	}

	etcdKey := etcdKeyFromPointer(resourcePointer)

	resp, err := st.cli.Get(ctx, etcdKey)
	if err != nil {
		return err
	}

	if len(resp.Kvs) == 0 {
		return ErrNotFound(resourcePointer)
	}

	etcdVersion := resp.Kvs[0].Version

	curResource, err := st.unmarshalResource(resp.Kvs[0])
	if err != nil {
		return err
	}

	if curResource.Metadata().Owner() != options.Owner {
		return ErrOwnerConflict(curResource.Metadata(), curResource.Metadata().Owner())
	}

	if !curResource.Metadata().Finalizers().Empty() {
		return ErrPendingFinalizers(*curResource.Metadata())
	}

	txnResp, err := st.cli.Txn(ctx).If(
		clientv3.Compare(clientv3.Version(etcdKey), "=", etcdVersion),
	).Then(
		clientv3.OpDelete(etcdKey),
	).Else(
		clientv3.OpGet(etcdKey),
	).Commit()
	if err != nil {
		return err
	}

	if txnResp.Succeeded {
		return nil
	}

	var foundVersion int64

	txnGetKvs := txnResp.Responses[0].GetResponseRange().Kvs
	if len(txnResp.Responses[0].GetResponseRange().Kvs) > 0 {
		foundVersion = txnGetKvs[0].Version
	}

	return ErrVersionConflict(resourcePointer, etcdVersion, foundVersion)
}

// Watch a resource.
func (st *State) Watch(ctx context.Context, resourcePointer resource.Pointer, ch chan<- state.Event, opts ...state.WatchOption) error {
	var options state.WatchOptions

	for _, opt := range opts {
		opt(&options)
	}

	if options.TailEvents > 0 {
		return ErrUnsupported("tailEvents")
	}

	etcdKey := etcdKeyFromPointer(resourcePointer)

	var (
		curResource resource.Resource
		err         error
	)

	getResp, err := st.cli.Get(ctx, etcdKey)
	if err != nil {
		return err
	}

	var initialEvent state.Event

	if len(getResp.Kvs) > 0 {
		curResource, err = st.unmarshalResource(getResp.Kvs[0])
		if err != nil {
			return err
		}

		initialEvent.Resource = curResource
		initialEvent.Type = state.Created
	} else {
		initialEvent.Resource = resource.NewTombstone(
			resource.NewMetadata(
				resourcePointer.Namespace(),
				resourcePointer.Type(),
				resourcePointer.ID(),
				resource.VersionUndefined,
			))
		initialEvent.Type = state.Destroyed
	}

	revision := getResp.Header.Revision

	watchCh := st.cli.Watch(ctx, etcdKey, clientv3.WithPrevKV(), clientv3.WithRev(revision))

	go func() {
		select {
		case <-ctx.Done():
			return
		case ch <- initialEvent:
		}

		for {
			var watchResponse clientv3.WatchResponse

			select {
			case <-ctx.Done():
				return
			case watchResponse = <-watchCh:
			}

			if watchResponse.Canceled {
				return
			}

			for _, etcdEvent := range watchResponse.Events {
				event, err := st.convertEvent(etcdEvent)
				if err != nil {
					// skip the event
					continue
				}

				select {
				case <-ctx.Done():
					return
				case ch <- *event:
				}
			}
		}
	}()

	return nil
}

// WatchKind all resources by type.
//
//nolint:gocyclo,cyclop,gocognit
func (st *State) WatchKind(ctx context.Context, resourceKind resource.Kind, ch chan<- state.Event, opts ...state.WatchKindOption) error {
	var options state.WatchKindOptions

	for _, opt := range opts {
		opt(&options)
	}

	matches := func(res resource.Resource) bool {
		return options.LabelQuery.Matches(*res.Metadata().Labels())
	}

	if options.TailEvents > 0 {
		return ErrUnsupported("tailEvents")
	}

	etcdKey := etcdKeyPrefixFromKind(resourceKind)

	var bootstrapList []resource.Resource

	var revision int64

	if options.BootstrapContents {
		getResp, err := st.cli.Get(ctx, etcdKey, clientv3.WithPrefix())
		if err != nil {
			return err
		}

		revision = getResp.Header.Revision

		for _, kv := range getResp.Kvs {
			res, err := st.unmarshalResource(kv)
			if err != nil {
				return err
			}

			if !matches(res) {
				continue
			}

			bootstrapList = append(bootstrapList, res)
		}

		sort.Slice(bootstrapList, func(i, j int) bool {
			return bootstrapList[i].Metadata().ID() < bootstrapList[j].Metadata().ID()
		})
	}

	watchCh := st.cli.Watch(ctx, etcdKey, clientv3.WithPrefix(), clientv3.WithPrevKV(), clientv3.WithRev(revision))

	go func() {
		sentBootstrapEventSet := make(map[string]struct{}, len(bootstrapList))

		// send initial contents if they were captured
		for _, res := range bootstrapList {
			select {
			case ch <- state.Event{
				Type:     state.Created,
				Resource: res,
			}:
			case <-ctx.Done():
				return
			}

			sentBootstrapEventSet[res.Metadata().String()] = struct{}{}
		}

		bootstrapList = nil

		for {
			var watchResponse clientv3.WatchResponse

			select {
			case <-ctx.Done():
				return
			case watchResponse = <-watchCh:
			}

			if watchResponse.Canceled {
				return
			}

			for _, etcdEvent := range watchResponse.Events {
				event, err := st.convertEvent(etcdEvent)
				if err != nil {
					// skip the event
					continue
				}

				switch event.Type {
				case state.Created, state.Destroyed:
					if !matches(event.Resource) {
						// skip the event
						continue
					}
				case state.Updated:
					oldMatches := matches(event.Old)
					newMatches := matches(event.Resource)

					switch {
					// transform the event if matching fact changes with the update
					case oldMatches && !newMatches:
						event.Type = state.Destroyed
						event.Old = nil
					case !oldMatches && newMatches:
						event.Type = state.Created
						event.Old = nil
					case newMatches && oldMatches:
						// passthrough the event
					default:
						// skip the event
						continue
					}
				}

				if !(event.Type == state.Destroyed) {
					_, alreadySent := sentBootstrapEventSet[event.Resource.Metadata().String()]
					if alreadySent {
						continue
					}
				}

				select {
				case <-ctx.Done():
					return
				case ch <- *event:
				}
			}
		}
	}()

	return nil
}

func (st *State) unmarshalResource(kv *mvccpb.KeyValue) (resource.Resource, error) {
	res, err := st.marshaler.UnmarshalResource(kv.Value)
	if err != nil {
		return nil, err
	}

	res.Metadata().SetVersion(resource.NewVersion(uint64(kv.Version)))

	return res, err
}

func (st *State) convertEvent(etcdEvent *clientv3.Event) (*state.Event, error) {
	if etcdEvent == nil {
		return nil, fmt.Errorf("etcd event is nil")
	}

	var (
		current  resource.Resource
		previous resource.Resource
		err      error
	)

	if etcdEvent.Kv != nil && etcdEvent.Type != clientv3.EventTypeDelete {
		current, err = st.unmarshalResource(etcdEvent.Kv)
		if err != nil {
			return nil, err
		}
	}

	if etcdEvent.PrevKv != nil {
		previous, err = st.unmarshalResource(etcdEvent.PrevKv)
		if err != nil {
			return nil, err
		}
	}

	if etcdEvent.Type == clientv3.EventTypeDelete {
		return &state.Event{
			Resource: previous,
			Type:     state.Destroyed,
		}, nil
	}

	eventType := state.Updated
	if etcdEvent.IsCreate() {
		eventType = state.Created
	}

	return &state.Event{
		Resource: current,
		Old:      previous,
		Type:     eventType,
	}, nil
}
