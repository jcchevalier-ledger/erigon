// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: remote/kv.proto

package remote

import (
	context "context"
	types "github.com/ledgerwatch/erigon/erigon-lib/gointerfaces/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	KV_Version_FullMethodName      = "/remote.KV/Version"
	KV_Tx_FullMethodName           = "/remote.KV/Tx"
	KV_StateChanges_FullMethodName = "/remote.KV/StateChanges"
	KV_Snapshots_FullMethodName    = "/remote.KV/Snapshots"
	KV_Range_FullMethodName        = "/remote.KV/Range"
	KV_DomainGet_FullMethodName    = "/remote.KV/DomainGet"
	KV_HistoryGet_FullMethodName   = "/remote.KV/HistoryGet"
	KV_IndexRange_FullMethodName   = "/remote.KV/IndexRange"
	KV_HistoryRange_FullMethodName = "/remote.KV/HistoryRange"
	KV_DomainRange_FullMethodName  = "/remote.KV/DomainRange"
)

// KVClient is the client API for KV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KVClient interface {
	// Version returns the service version number
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*types.VersionReply, error)
	// Tx exposes read-only transactions for the key-value store
	//
	// When tx open, client must receive 1 message from server with txID
	// When cursor open, client must receive 1 message from server with cursorID
	// Then only client can initiate messages from server
	Tx(ctx context.Context, opts ...grpc.CallOption) (KV_TxClient, error)
	StateChanges(ctx context.Context, in *StateChangeRequest, opts ...grpc.CallOption) (KV_StateChangesClient, error)
	// Snapshots returns list of current snapshot files. Then client can just open all of them.
	Snapshots(ctx context.Context, in *SnapshotsRequest, opts ...grpc.CallOption) (*SnapshotsReply, error)
	// Range [from, to)
	// Range(from, nil) means [from, EndOfTable)
	// Range(nil, to)   means [StartOfTable, to)
	// If orderAscend=false server expecting `from`<`to`. Example: Range("B", "A")
	Range(ctx context.Context, in *RangeReq, opts ...grpc.CallOption) (*Pairs, error)
	// Temporal methods
	DomainGet(ctx context.Context, in *DomainGetReq, opts ...grpc.CallOption) (*DomainGetReply, error)
	HistoryGet(ctx context.Context, in *HistoryGetReq, opts ...grpc.CallOption) (*HistoryGetReply, error)
	IndexRange(ctx context.Context, in *IndexRangeReq, opts ...grpc.CallOption) (*IndexRangeReply, error)
	HistoryRange(ctx context.Context, in *HistoryRangeReq, opts ...grpc.CallOption) (*Pairs, error)
	DomainRange(ctx context.Context, in *DomainRangeReq, opts ...grpc.CallOption) (*Pairs, error)
}

type kVClient struct {
	cc grpc.ClientConnInterface
}

func NewKVClient(cc grpc.ClientConnInterface) KVClient {
	return &kVClient{cc}
}

func (c *kVClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*types.VersionReply, error) {
	out := new(types.VersionReply)
	err := c.cc.Invoke(ctx, KV_Version_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Tx(ctx context.Context, opts ...grpc.CallOption) (KV_TxClient, error) {
	stream, err := c.cc.NewStream(ctx, &KV_ServiceDesc.Streams[0], KV_Tx_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &kVTxClient{stream}
	return x, nil
}

type KV_TxClient interface {
	Send(*Cursor) error
	Recv() (*Pair, error)
	grpc.ClientStream
}

type kVTxClient struct {
	grpc.ClientStream
}

func (x *kVTxClient) Send(m *Cursor) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kVTxClient) Recv() (*Pair, error) {
	m := new(Pair)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kVClient) StateChanges(ctx context.Context, in *StateChangeRequest, opts ...grpc.CallOption) (KV_StateChangesClient, error) {
	stream, err := c.cc.NewStream(ctx, &KV_ServiceDesc.Streams[1], KV_StateChanges_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &kVStateChangesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KV_StateChangesClient interface {
	Recv() (*StateChangeBatch, error)
	grpc.ClientStream
}

type kVStateChangesClient struct {
	grpc.ClientStream
}

func (x *kVStateChangesClient) Recv() (*StateChangeBatch, error) {
	m := new(StateChangeBatch)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kVClient) Snapshots(ctx context.Context, in *SnapshotsRequest, opts ...grpc.CallOption) (*SnapshotsReply, error) {
	out := new(SnapshotsReply)
	err := c.cc.Invoke(ctx, KV_Snapshots_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Range(ctx context.Context, in *RangeReq, opts ...grpc.CallOption) (*Pairs, error) {
	out := new(Pairs)
	err := c.cc.Invoke(ctx, KV_Range_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) DomainGet(ctx context.Context, in *DomainGetReq, opts ...grpc.CallOption) (*DomainGetReply, error) {
	out := new(DomainGetReply)
	err := c.cc.Invoke(ctx, KV_DomainGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) HistoryGet(ctx context.Context, in *HistoryGetReq, opts ...grpc.CallOption) (*HistoryGetReply, error) {
	out := new(HistoryGetReply)
	err := c.cc.Invoke(ctx, KV_HistoryGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) IndexRange(ctx context.Context, in *IndexRangeReq, opts ...grpc.CallOption) (*IndexRangeReply, error) {
	out := new(IndexRangeReply)
	err := c.cc.Invoke(ctx, KV_IndexRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) HistoryRange(ctx context.Context, in *HistoryRangeReq, opts ...grpc.CallOption) (*Pairs, error) {
	out := new(Pairs)
	err := c.cc.Invoke(ctx, KV_HistoryRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) DomainRange(ctx context.Context, in *DomainRangeReq, opts ...grpc.CallOption) (*Pairs, error) {
	out := new(Pairs)
	err := c.cc.Invoke(ctx, KV_DomainRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVServer is the server API for KV service.
// All implementations must embed UnimplementedKVServer
// for forward compatibility
type KVServer interface {
	// Version returns the service version number
	Version(context.Context, *emptypb.Empty) (*types.VersionReply, error)
	// Tx exposes read-only transactions for the key-value store
	//
	// When tx open, client must receive 1 message from server with txID
	// When cursor open, client must receive 1 message from server with cursorID
	// Then only client can initiate messages from server
	Tx(KV_TxServer) error
	StateChanges(*StateChangeRequest, KV_StateChangesServer) error
	// Snapshots returns list of current snapshot files. Then client can just open all of them.
	Snapshots(context.Context, *SnapshotsRequest) (*SnapshotsReply, error)
	// Range [from, to)
	// Range(from, nil) means [from, EndOfTable)
	// Range(nil, to)   means [StartOfTable, to)
	// If orderAscend=false server expecting `from`<`to`. Example: Range("B", "A")
	Range(context.Context, *RangeReq) (*Pairs, error)
	// Temporal methods
	DomainGet(context.Context, *DomainGetReq) (*DomainGetReply, error)
	HistoryGet(context.Context, *HistoryGetReq) (*HistoryGetReply, error)
	IndexRange(context.Context, *IndexRangeReq) (*IndexRangeReply, error)
	HistoryRange(context.Context, *HistoryRangeReq) (*Pairs, error)
	DomainRange(context.Context, *DomainRangeReq) (*Pairs, error)
	mustEmbedUnimplementedKVServer()
}

// UnimplementedKVServer must be embedded to have forward compatible implementations.
type UnimplementedKVServer struct {
}

func (UnimplementedKVServer) Version(context.Context, *emptypb.Empty) (*types.VersionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedKVServer) Tx(KV_TxServer) error {
	return status.Errorf(codes.Unimplemented, "method Tx not implemented")
}
func (UnimplementedKVServer) StateChanges(*StateChangeRequest, KV_StateChangesServer) error {
	return status.Errorf(codes.Unimplemented, "method StateChanges not implemented")
}
func (UnimplementedKVServer) Snapshots(context.Context, *SnapshotsRequest) (*SnapshotsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snapshots not implemented")
}
func (UnimplementedKVServer) Range(context.Context, *RangeReq) (*Pairs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Range not implemented")
}
func (UnimplementedKVServer) DomainGet(context.Context, *DomainGetReq) (*DomainGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DomainGet not implemented")
}
func (UnimplementedKVServer) HistoryGet(context.Context, *HistoryGetReq) (*HistoryGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HistoryGet not implemented")
}
func (UnimplementedKVServer) IndexRange(context.Context, *IndexRangeReq) (*IndexRangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IndexRange not implemented")
}
func (UnimplementedKVServer) HistoryRange(context.Context, *HistoryRangeReq) (*Pairs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HistoryRange not implemented")
}
func (UnimplementedKVServer) DomainRange(context.Context, *DomainRangeReq) (*Pairs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DomainRange not implemented")
}
func (UnimplementedKVServer) mustEmbedUnimplementedKVServer() {}

// UnsafeKVServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KVServer will
// result in compilation errors.
type UnsafeKVServer interface {
	mustEmbedUnimplementedKVServer()
}

func RegisterKVServer(s grpc.ServiceRegistrar, srv KVServer) {
	s.RegisterService(&KV_ServiceDesc, srv)
}

func _KV_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KV_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Tx_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KVServer).Tx(&kVTxServer{stream})
}

type KV_TxServer interface {
	Send(*Pair) error
	Recv() (*Cursor, error)
	grpc.ServerStream
}

type kVTxServer struct {
	grpc.ServerStream
}

func (x *kVTxServer) Send(m *Pair) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kVTxServer) Recv() (*Cursor, error) {
	m := new(Cursor)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _KV_StateChanges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StateChangeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KVServer).StateChanges(m, &kVStateChangesServer{stream})
}

type KV_StateChangesServer interface {
	Send(*StateChangeBatch) error
	grpc.ServerStream
}

type kVStateChangesServer struct {
	grpc.ServerStream
}

func (x *kVStateChangesServer) Send(m *StateChangeBatch) error {
	return x.ServerStream.SendMsg(m)
}

func _KV_Snapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Snapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KV_Snapshots_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Snapshots(ctx, req.(*SnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Range_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Range(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KV_Range_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Range(ctx, req.(*RangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_DomainGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DomainGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).DomainGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KV_DomainGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).DomainGet(ctx, req.(*DomainGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_HistoryGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).HistoryGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KV_HistoryGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).HistoryGet(ctx, req.(*HistoryGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_IndexRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).IndexRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KV_IndexRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).IndexRange(ctx, req.(*IndexRangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_HistoryRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryRangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).HistoryRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KV_HistoryRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).HistoryRange(ctx, req.(*HistoryRangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_DomainRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DomainRangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).DomainRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KV_DomainRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).DomainRange(ctx, req.(*DomainRangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// KV_ServiceDesc is the grpc.ServiceDesc for KV service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KV_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remote.KV",
	HandlerType: (*KVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _KV_Version_Handler,
		},
		{
			MethodName: "Snapshots",
			Handler:    _KV_Snapshots_Handler,
		},
		{
			MethodName: "Range",
			Handler:    _KV_Range_Handler,
		},
		{
			MethodName: "DomainGet",
			Handler:    _KV_DomainGet_Handler,
		},
		{
			MethodName: "HistoryGet",
			Handler:    _KV_HistoryGet_Handler,
		},
		{
			MethodName: "IndexRange",
			Handler:    _KV_IndexRange_Handler,
		},
		{
			MethodName: "HistoryRange",
			Handler:    _KV_HistoryRange_Handler,
		},
		{
			MethodName: "DomainRange",
			Handler:    _KV_DomainRange_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tx",
			Handler:       _KV_Tx_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StateChanges",
			Handler:       _KV_StateChanges_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "remote/kv.proto",
}
