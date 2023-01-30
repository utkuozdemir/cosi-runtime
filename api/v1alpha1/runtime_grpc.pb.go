// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: v1alpha1/runtime.proto

package v1alpha1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ControllerRuntimeClient is the client API for ControllerRuntime service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControllerRuntimeClient interface {
	// RegisterController registers the controller and establishes a token for ControllerAdapter calls.
	//
	// RegisterController builds initial set of inputs and outputs for the controller.
	// If there's a conflict on inputs or outputs, RegisterController fails.
	// RegisterController enforces unique controller names.
	RegisterController(ctx context.Context, in *RegisterControllerRequest, opts ...grpc.CallOption) (*RegisterControllerResponse, error)
	// Start the controller runtime.
	//
	// Start should only be called once.
	// Once the runtime is started, controllers start receiving reconcile events via
	// the ControllerAdapter APIs.
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	// Stop the controller runtime.
	//
	// Stop should only be called once.
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
}

type controllerRuntimeClient struct {
	cc grpc.ClientConnInterface
}

func NewControllerRuntimeClient(cc grpc.ClientConnInterface) ControllerRuntimeClient {
	return &controllerRuntimeClient{cc}
}

func (c *controllerRuntimeClient) RegisterController(ctx context.Context, in *RegisterControllerRequest, opts ...grpc.CallOption) (*RegisterControllerResponse, error) {
	out := new(RegisterControllerResponse)
	err := c.cc.Invoke(ctx, "/cosi.runtime.ControllerRuntime/RegisterController", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerRuntimeClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/cosi.runtime.ControllerRuntime/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerRuntimeClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/cosi.runtime.ControllerRuntime/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControllerRuntimeServer is the server API for ControllerRuntime service.
// All implementations must embed UnimplementedControllerRuntimeServer
// for forward compatibility
type ControllerRuntimeServer interface {
	// RegisterController registers the controller and establishes a token for ControllerAdapter calls.
	//
	// RegisterController builds initial set of inputs and outputs for the controller.
	// If there's a conflict on inputs or outputs, RegisterController fails.
	// RegisterController enforces unique controller names.
	RegisterController(context.Context, *RegisterControllerRequest) (*RegisterControllerResponse, error)
	// Start the controller runtime.
	//
	// Start should only be called once.
	// Once the runtime is started, controllers start receiving reconcile events via
	// the ControllerAdapter APIs.
	Start(context.Context, *StartRequest) (*StartResponse, error)
	// Stop the controller runtime.
	//
	// Stop should only be called once.
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	mustEmbedUnimplementedControllerRuntimeServer()
}

// UnimplementedControllerRuntimeServer must be embedded to have forward compatible implementations.
type UnimplementedControllerRuntimeServer struct {
}

func (UnimplementedControllerRuntimeServer) RegisterController(context.Context, *RegisterControllerRequest) (*RegisterControllerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterController not implemented")
}
func (UnimplementedControllerRuntimeServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedControllerRuntimeServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedControllerRuntimeServer) mustEmbedUnimplementedControllerRuntimeServer() {}

// UnsafeControllerRuntimeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControllerRuntimeServer will
// result in compilation errors.
type UnsafeControllerRuntimeServer interface {
	mustEmbedUnimplementedControllerRuntimeServer()
}

func RegisterControllerRuntimeServer(s grpc.ServiceRegistrar, srv ControllerRuntimeServer) {
	s.RegisterService(&ControllerRuntime_ServiceDesc, srv)
}

func _ControllerRuntime_RegisterController_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterControllerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerRuntimeServer).RegisterController(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosi.runtime.ControllerRuntime/RegisterController",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerRuntimeServer).RegisterController(ctx, req.(*RegisterControllerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerRuntime_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerRuntimeServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosi.runtime.ControllerRuntime/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerRuntimeServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerRuntime_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerRuntimeServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosi.runtime.ControllerRuntime/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerRuntimeServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ControllerRuntime_ServiceDesc is the grpc.ServiceDesc for ControllerRuntime service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ControllerRuntime_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosi.runtime.ControllerRuntime",
	HandlerType: (*ControllerRuntimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterController",
			Handler:    _ControllerRuntime_RegisterController_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _ControllerRuntime_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _ControllerRuntime_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1alpha1/runtime.proto",
}

// ControllerAdapterClient is the client API for ControllerAdapter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControllerAdapterClient interface {
	// ReconcileEvents sends a message on each reconcile event to the controller.
	//
	// Controller is supposed to called 'ReconcileEvents' on start and run reconcile loop
	// each time there is an event received.
	ReconcileEvents(ctx context.Context, in *ReconcileEventsRequest, opts ...grpc.CallOption) (ControllerAdapter_ReconcileEventsClient, error)
	QueueReconcile(ctx context.Context, in *QueueReconcileRequest, opts ...grpc.CallOption) (*QueueReconcileResponse, error)
	// UpdateInputs updates the list of controller inputs.
	//
	// This call replaces the list of inputs with the new list.
	// For any new inputs a separate watch is established and reconcile
	// events are sent for new inputs.
	UpdateInputs(ctx context.Context, in *UpdateInputsRequest, opts ...grpc.CallOption) (*UpdateInputsResponse, error)
	// Get a resource.
	//
	// Resource should be either input or ouput of the controller.
	Get(ctx context.Context, in *RuntimeGetRequest, opts ...grpc.CallOption) (*RuntimeGetResponse, error)
	// List resources.
	//
	// Resource should be either input or ouput of the controller.
	List(ctx context.Context, in *RuntimeListRequest, opts ...grpc.CallOption) (ControllerAdapter_ListClient, error)
	// WatchFor a specific resource state.
	//
	// Resource should be either input or ouput of the controller.
	WatchFor(ctx context.Context, in *RuntimeWatchForRequest, opts ...grpc.CallOption) (*RuntimeWatchForResponse, error)
	// Create a new resource.
	//
	// Resource should be an output of the controller.
	Create(ctx context.Context, in *RuntimeCreateRequest, opts ...grpc.CallOption) (*RuntimeCreateResponse, error)
	// Update a resource.
	//
	// Up-to-date current version should be specified for the update to succeed.
	//
	// Resource should be an output of the controller, for shared outputs
	// resource should be owned by the controller.
	Update(ctx context.Context, in *RuntimeUpdateRequest, opts ...grpc.CallOption) (*RuntimeUpdateResponse, error)
	// Teardown marks a resource as going through the teardown phase.
	//
	// Teardown phase notifies other controllers using the resource as a strong input
	// that the resource is going away, and once cleanup is done, finalizer should be removed
	// which unblocks resource destruction.
	//
	// Resource should be an output of the controller, for shared outputs
	// resource should be owned by the controller.
	Teardown(ctx context.Context, in *RuntimeTeardownRequest, opts ...grpc.CallOption) (*RuntimeTeardownResponse, error)
	// Destroy a resource.
	//
	// Resource should have no finalizers to be destroyed.
	//
	// Resource should be an output of the controller, for shared outputs
	// resource should be owned by the controller.
	Destroy(ctx context.Context, in *RuntimeDestroyRequest, opts ...grpc.CallOption) (*RuntimeDestroyResponse, error)
	// AddFinalizer adds a finalizer on the resource.
	//
	// Resource should be a strong input of the controller.
	AddFinalizer(ctx context.Context, in *RuntimeAddFinalizerRequest, opts ...grpc.CallOption) (*RuntimeAddFinalizerResponse, error)
	// RemoveFinalizer remove a finalizer from the resource.
	//
	// Resource should be a strong input of the controller.
	RemoveFinalizer(ctx context.Context, in *RuntimeRemoveFinalizerRequest, opts ...grpc.CallOption) (*RuntimeRemoveFinalizerResponse, error)
}

type controllerAdapterClient struct {
	cc grpc.ClientConnInterface
}

func NewControllerAdapterClient(cc grpc.ClientConnInterface) ControllerAdapterClient {
	return &controllerAdapterClient{cc}
}

func (c *controllerAdapterClient) ReconcileEvents(ctx context.Context, in *ReconcileEventsRequest, opts ...grpc.CallOption) (ControllerAdapter_ReconcileEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ControllerAdapter_ServiceDesc.Streams[0], "/cosi.runtime.ControllerAdapter/ReconcileEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &controllerAdapterReconcileEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ControllerAdapter_ReconcileEventsClient interface {
	Recv() (*ReconcileEventsResponse, error)
	grpc.ClientStream
}

type controllerAdapterReconcileEventsClient struct {
	grpc.ClientStream
}

func (x *controllerAdapterReconcileEventsClient) Recv() (*ReconcileEventsResponse, error) {
	m := new(ReconcileEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controllerAdapterClient) QueueReconcile(ctx context.Context, in *QueueReconcileRequest, opts ...grpc.CallOption) (*QueueReconcileResponse, error) {
	out := new(QueueReconcileResponse)
	err := c.cc.Invoke(ctx, "/cosi.runtime.ControllerAdapter/QueueReconcile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerAdapterClient) UpdateInputs(ctx context.Context, in *UpdateInputsRequest, opts ...grpc.CallOption) (*UpdateInputsResponse, error) {
	out := new(UpdateInputsResponse)
	err := c.cc.Invoke(ctx, "/cosi.runtime.ControllerAdapter/UpdateInputs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerAdapterClient) Get(ctx context.Context, in *RuntimeGetRequest, opts ...grpc.CallOption) (*RuntimeGetResponse, error) {
	out := new(RuntimeGetResponse)
	err := c.cc.Invoke(ctx, "/cosi.runtime.ControllerAdapter/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerAdapterClient) List(ctx context.Context, in *RuntimeListRequest, opts ...grpc.CallOption) (ControllerAdapter_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &ControllerAdapter_ServiceDesc.Streams[1], "/cosi.runtime.ControllerAdapter/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &controllerAdapterListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ControllerAdapter_ListClient interface {
	Recv() (*RuntimeListResponse, error)
	grpc.ClientStream
}

type controllerAdapterListClient struct {
	grpc.ClientStream
}

func (x *controllerAdapterListClient) Recv() (*RuntimeListResponse, error) {
	m := new(RuntimeListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controllerAdapterClient) WatchFor(ctx context.Context, in *RuntimeWatchForRequest, opts ...grpc.CallOption) (*RuntimeWatchForResponse, error) {
	out := new(RuntimeWatchForResponse)
	err := c.cc.Invoke(ctx, "/cosi.runtime.ControllerAdapter/WatchFor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerAdapterClient) Create(ctx context.Context, in *RuntimeCreateRequest, opts ...grpc.CallOption) (*RuntimeCreateResponse, error) {
	out := new(RuntimeCreateResponse)
	err := c.cc.Invoke(ctx, "/cosi.runtime.ControllerAdapter/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerAdapterClient) Update(ctx context.Context, in *RuntimeUpdateRequest, opts ...grpc.CallOption) (*RuntimeUpdateResponse, error) {
	out := new(RuntimeUpdateResponse)
	err := c.cc.Invoke(ctx, "/cosi.runtime.ControllerAdapter/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerAdapterClient) Teardown(ctx context.Context, in *RuntimeTeardownRequest, opts ...grpc.CallOption) (*RuntimeTeardownResponse, error) {
	out := new(RuntimeTeardownResponse)
	err := c.cc.Invoke(ctx, "/cosi.runtime.ControllerAdapter/Teardown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerAdapterClient) Destroy(ctx context.Context, in *RuntimeDestroyRequest, opts ...grpc.CallOption) (*RuntimeDestroyResponse, error) {
	out := new(RuntimeDestroyResponse)
	err := c.cc.Invoke(ctx, "/cosi.runtime.ControllerAdapter/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerAdapterClient) AddFinalizer(ctx context.Context, in *RuntimeAddFinalizerRequest, opts ...grpc.CallOption) (*RuntimeAddFinalizerResponse, error) {
	out := new(RuntimeAddFinalizerResponse)
	err := c.cc.Invoke(ctx, "/cosi.runtime.ControllerAdapter/AddFinalizer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerAdapterClient) RemoveFinalizer(ctx context.Context, in *RuntimeRemoveFinalizerRequest, opts ...grpc.CallOption) (*RuntimeRemoveFinalizerResponse, error) {
	out := new(RuntimeRemoveFinalizerResponse)
	err := c.cc.Invoke(ctx, "/cosi.runtime.ControllerAdapter/RemoveFinalizer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControllerAdapterServer is the server API for ControllerAdapter service.
// All implementations must embed UnimplementedControllerAdapterServer
// for forward compatibility
type ControllerAdapterServer interface {
	// ReconcileEvents sends a message on each reconcile event to the controller.
	//
	// Controller is supposed to called 'ReconcileEvents' on start and run reconcile loop
	// each time there is an event received.
	ReconcileEvents(*ReconcileEventsRequest, ControllerAdapter_ReconcileEventsServer) error
	QueueReconcile(context.Context, *QueueReconcileRequest) (*QueueReconcileResponse, error)
	// UpdateInputs updates the list of controller inputs.
	//
	// This call replaces the list of inputs with the new list.
	// For any new inputs a separate watch is established and reconcile
	// events are sent for new inputs.
	UpdateInputs(context.Context, *UpdateInputsRequest) (*UpdateInputsResponse, error)
	// Get a resource.
	//
	// Resource should be either input or ouput of the controller.
	Get(context.Context, *RuntimeGetRequest) (*RuntimeGetResponse, error)
	// List resources.
	//
	// Resource should be either input or ouput of the controller.
	List(*RuntimeListRequest, ControllerAdapter_ListServer) error
	// WatchFor a specific resource state.
	//
	// Resource should be either input or ouput of the controller.
	WatchFor(context.Context, *RuntimeWatchForRequest) (*RuntimeWatchForResponse, error)
	// Create a new resource.
	//
	// Resource should be an output of the controller.
	Create(context.Context, *RuntimeCreateRequest) (*RuntimeCreateResponse, error)
	// Update a resource.
	//
	// Up-to-date current version should be specified for the update to succeed.
	//
	// Resource should be an output of the controller, for shared outputs
	// resource should be owned by the controller.
	Update(context.Context, *RuntimeUpdateRequest) (*RuntimeUpdateResponse, error)
	// Teardown marks a resource as going through the teardown phase.
	//
	// Teardown phase notifies other controllers using the resource as a strong input
	// that the resource is going away, and once cleanup is done, finalizer should be removed
	// which unblocks resource destruction.
	//
	// Resource should be an output of the controller, for shared outputs
	// resource should be owned by the controller.
	Teardown(context.Context, *RuntimeTeardownRequest) (*RuntimeTeardownResponse, error)
	// Destroy a resource.
	//
	// Resource should have no finalizers to be destroyed.
	//
	// Resource should be an output of the controller, for shared outputs
	// resource should be owned by the controller.
	Destroy(context.Context, *RuntimeDestroyRequest) (*RuntimeDestroyResponse, error)
	// AddFinalizer adds a finalizer on the resource.
	//
	// Resource should be a strong input of the controller.
	AddFinalizer(context.Context, *RuntimeAddFinalizerRequest) (*RuntimeAddFinalizerResponse, error)
	// RemoveFinalizer remove a finalizer from the resource.
	//
	// Resource should be a strong input of the controller.
	RemoveFinalizer(context.Context, *RuntimeRemoveFinalizerRequest) (*RuntimeRemoveFinalizerResponse, error)
	mustEmbedUnimplementedControllerAdapterServer()
}

// UnimplementedControllerAdapterServer must be embedded to have forward compatible implementations.
type UnimplementedControllerAdapterServer struct {
}

func (UnimplementedControllerAdapterServer) ReconcileEvents(*ReconcileEventsRequest, ControllerAdapter_ReconcileEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReconcileEvents not implemented")
}
func (UnimplementedControllerAdapterServer) QueueReconcile(context.Context, *QueueReconcileRequest) (*QueueReconcileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueueReconcile not implemented")
}
func (UnimplementedControllerAdapterServer) UpdateInputs(context.Context, *UpdateInputsRequest) (*UpdateInputsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInputs not implemented")
}
func (UnimplementedControllerAdapterServer) Get(context.Context, *RuntimeGetRequest) (*RuntimeGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedControllerAdapterServer) List(*RuntimeListRequest, ControllerAdapter_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedControllerAdapterServer) WatchFor(context.Context, *RuntimeWatchForRequest) (*RuntimeWatchForResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WatchFor not implemented")
}
func (UnimplementedControllerAdapterServer) Create(context.Context, *RuntimeCreateRequest) (*RuntimeCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedControllerAdapterServer) Update(context.Context, *RuntimeUpdateRequest) (*RuntimeUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedControllerAdapterServer) Teardown(context.Context, *RuntimeTeardownRequest) (*RuntimeTeardownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Teardown not implemented")
}
func (UnimplementedControllerAdapterServer) Destroy(context.Context, *RuntimeDestroyRequest) (*RuntimeDestroyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedControllerAdapterServer) AddFinalizer(context.Context, *RuntimeAddFinalizerRequest) (*RuntimeAddFinalizerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFinalizer not implemented")
}
func (UnimplementedControllerAdapterServer) RemoveFinalizer(context.Context, *RuntimeRemoveFinalizerRequest) (*RuntimeRemoveFinalizerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFinalizer not implemented")
}
func (UnimplementedControllerAdapterServer) mustEmbedUnimplementedControllerAdapterServer() {}

// UnsafeControllerAdapterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControllerAdapterServer will
// result in compilation errors.
type UnsafeControllerAdapterServer interface {
	mustEmbedUnimplementedControllerAdapterServer()
}

func RegisterControllerAdapterServer(s grpc.ServiceRegistrar, srv ControllerAdapterServer) {
	s.RegisterService(&ControllerAdapter_ServiceDesc, srv)
}

func _ControllerAdapter_ReconcileEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReconcileEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControllerAdapterServer).ReconcileEvents(m, &controllerAdapterReconcileEventsServer{stream})
}

type ControllerAdapter_ReconcileEventsServer interface {
	Send(*ReconcileEventsResponse) error
	grpc.ServerStream
}

type controllerAdapterReconcileEventsServer struct {
	grpc.ServerStream
}

func (x *controllerAdapterReconcileEventsServer) Send(m *ReconcileEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ControllerAdapter_QueueReconcile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueReconcileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerAdapterServer).QueueReconcile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosi.runtime.ControllerAdapter/QueueReconcile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerAdapterServer).QueueReconcile(ctx, req.(*QueueReconcileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerAdapter_UpdateInputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInputsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerAdapterServer).UpdateInputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosi.runtime.ControllerAdapter/UpdateInputs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerAdapterServer).UpdateInputs(ctx, req.(*UpdateInputsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerAdapter_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuntimeGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerAdapterServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosi.runtime.ControllerAdapter/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerAdapterServer).Get(ctx, req.(*RuntimeGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerAdapter_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RuntimeListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControllerAdapterServer).List(m, &controllerAdapterListServer{stream})
}

type ControllerAdapter_ListServer interface {
	Send(*RuntimeListResponse) error
	grpc.ServerStream
}

type controllerAdapterListServer struct {
	grpc.ServerStream
}

func (x *controllerAdapterListServer) Send(m *RuntimeListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ControllerAdapter_WatchFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuntimeWatchForRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerAdapterServer).WatchFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosi.runtime.ControllerAdapter/WatchFor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerAdapterServer).WatchFor(ctx, req.(*RuntimeWatchForRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerAdapter_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuntimeCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerAdapterServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosi.runtime.ControllerAdapter/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerAdapterServer).Create(ctx, req.(*RuntimeCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerAdapter_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuntimeUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerAdapterServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosi.runtime.ControllerAdapter/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerAdapterServer).Update(ctx, req.(*RuntimeUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerAdapter_Teardown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuntimeTeardownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerAdapterServer).Teardown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosi.runtime.ControllerAdapter/Teardown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerAdapterServer).Teardown(ctx, req.(*RuntimeTeardownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerAdapter_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuntimeDestroyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerAdapterServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosi.runtime.ControllerAdapter/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerAdapterServer).Destroy(ctx, req.(*RuntimeDestroyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerAdapter_AddFinalizer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuntimeAddFinalizerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerAdapterServer).AddFinalizer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosi.runtime.ControllerAdapter/AddFinalizer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerAdapterServer).AddFinalizer(ctx, req.(*RuntimeAddFinalizerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerAdapter_RemoveFinalizer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuntimeRemoveFinalizerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerAdapterServer).RemoveFinalizer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosi.runtime.ControllerAdapter/RemoveFinalizer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerAdapterServer).RemoveFinalizer(ctx, req.(*RuntimeRemoveFinalizerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ControllerAdapter_ServiceDesc is the grpc.ServiceDesc for ControllerAdapter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ControllerAdapter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosi.runtime.ControllerAdapter",
	HandlerType: (*ControllerAdapterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueueReconcile",
			Handler:    _ControllerAdapter_QueueReconcile_Handler,
		},
		{
			MethodName: "UpdateInputs",
			Handler:    _ControllerAdapter_UpdateInputs_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ControllerAdapter_Get_Handler,
		},
		{
			MethodName: "WatchFor",
			Handler:    _ControllerAdapter_WatchFor_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ControllerAdapter_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ControllerAdapter_Update_Handler,
		},
		{
			MethodName: "Teardown",
			Handler:    _ControllerAdapter_Teardown_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _ControllerAdapter_Destroy_Handler,
		},
		{
			MethodName: "AddFinalizer",
			Handler:    _ControllerAdapter_AddFinalizer_Handler,
		},
		{
			MethodName: "RemoveFinalizer",
			Handler:    _ControllerAdapter_RemoveFinalizer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReconcileEvents",
			Handler:       _ControllerAdapter_ReconcileEvents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "List",
			Handler:       _ControllerAdapter_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1alpha1/runtime.proto",
}
