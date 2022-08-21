package etcd_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.etcd.io/etcd/server/v3/embed"
	"go.etcd.io/etcd/server/v3/etcdserver/api/v3client"

	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/resource/protobuf"
	"github.com/cosi-project/runtime/pkg/state"
	"github.com/cosi-project/runtime/pkg/state/conformance"
	"github.com/cosi-project/runtime/pkg/state/impl/etcd"
)

func TestInterfaces(t *testing.T) {
	t.Parallel()

	assert.Implements(t, (*state.CoreState)(nil), new(etcd.State))
}

func TestEtcdConformance(t *testing.T) {
	t.Parallel()

	err := protobuf.RegisterResource(conformance.PathResourceType, &conformance.PathResource{})
	if err != nil {
		t.Fatalf("failed to register resource: %v", err)
	}

	withEtcd(t, func(s state.State) {
		suite.Run(t, &conformance.StateSuite{
			State:      s,
			Namespaces: []resource.Namespace{"default"},
		})
	})
}

func withEtcd(t *testing.T, f func(state.State)) {
	tempDir := t.TempDir()

	cfg := embed.NewConfig()
	cfg.Dir = tempDir

	e, err := embed.StartEtcd(cfg)
	if err != nil {
		t.Fatalf("failed to start etcd: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	select {
	case <-e.Server.ReadyNotify():
	case <-ctx.Done():
		t.Fatalf("etcd failed to start")
	}

	defer e.Close()

	cli := v3client.New(e.Server)

	defer cli.Close() //nolint:errcheck

	st := state.WrapCore(etcd.NewState(cli))

	f(st)
}
