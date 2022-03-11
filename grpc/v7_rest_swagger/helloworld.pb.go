// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: helloworld.proto

package main

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StringMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StringMessage) Reset() {
	*x = StringMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMessage) ProtoMessage() {}

func (x *StringMessage) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMessage.ProtoReflect.Descriptor instead.
func (*StringMessage) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *StringMessage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_helloworld_proto protoreflect.FileDescriptor

var file_helloworld_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x98, 0x01,
	0x0a, 0x0b, 0x52, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x7d, 0x12, 0x42, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22, 0x05,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x6d,
	0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_proto_rawDescOnce sync.Once
	file_helloworld_proto_rawDescData = file_helloworld_proto_rawDesc
)

func file_helloworld_proto_rawDescGZIP() []byte {
	file_helloworld_proto_rawDescOnce.Do(func() {
		file_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_proto_rawDescData)
	})
	return file_helloworld_proto_rawDescData
}

var file_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_helloworld_proto_goTypes = []interface{}{
	(*StringMessage)(nil), // 0: main.StringMessage
}
var file_helloworld_proto_depIdxs = []int32{
	0, // 0: main.RestService.Get:input_type -> main.StringMessage
	0, // 1: main.RestService.Post:input_type -> main.StringMessage
	0, // 2: main.RestService.Get:output_type -> main.StringMessage
	0, // 3: main.RestService.Post:output_type -> main.StringMessage
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_helloworld_proto_init() }
func file_helloworld_proto_init() {
	if File_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_proto_goTypes,
		DependencyIndexes: file_helloworld_proto_depIdxs,
		MessageInfos:      file_helloworld_proto_msgTypes,
	}.Build()
	File_helloworld_proto = out.File
	file_helloworld_proto_rawDesc = nil
	file_helloworld_proto_goTypes = nil
	file_helloworld_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RestServiceClient is the client API for RestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RestServiceClient interface {
	// 首先为 gRPC 定义了 Get 和 Post 方法，然后通过元扩展语法在对应的方法后添加路由信息。
	//其中 “/get/{value}” 路径对应的是 Get 方法，{value} 部分对应参数中的 value 成员，结果通过 json 格式返回。
	//Post 方法对应 “/post” 路径，body 中包含 json 格式的请求信息。
	Get(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
	Post(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type restServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRestServiceClient(cc grpc.ClientConnInterface) RestServiceClient {
	return &restServiceClient{cc}
}

func (c *restServiceClient) Get(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/main.RestService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restServiceClient) Post(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/main.RestService/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestServiceServer is the server API for RestService service.
type RestServiceServer interface {
	// 首先为 gRPC 定义了 Get 和 Post 方法，然后通过元扩展语法在对应的方法后添加路由信息。
	//其中 “/get/{value}” 路径对应的是 Get 方法，{value} 部分对应参数中的 value 成员，结果通过 json 格式返回。
	//Post 方法对应 “/post” 路径，body 中包含 json 格式的请求信息。
	Get(context.Context, *StringMessage) (*StringMessage, error)
	Post(context.Context, *StringMessage) (*StringMessage, error)
}

// UnimplementedRestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRestServiceServer struct {
}

func (*UnimplementedRestServiceServer) Get(context.Context, *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRestServiceServer) Post(context.Context, *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}

func RegisterRestServiceServer(s *grpc.Server, srv RestServiceServer) {
	s.RegisterService(&_RestService_serviceDesc, srv)
}

func _RestService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.RestService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestServiceServer).Get(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestService_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestServiceServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.RestService/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestServiceServer).Post(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _RestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.RestService",
	HandlerType: (*RestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RestService_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _RestService_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}