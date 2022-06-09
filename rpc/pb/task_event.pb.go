// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: task_event.proto

package pb

import (
	context "context"
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

type TaskEventResp_StatusCode int32

const (
	TaskEventResp_SUCCESS TaskEventResp_StatusCode = 0
	TaskEventResp_FAILURE TaskEventResp_StatusCode = -1
)

// Enum value maps for TaskEventResp_StatusCode.
var (
	TaskEventResp_StatusCode_name = map[int32]string{
		0:  "SUCCESS",
		-1: "FAILURE",
	}
	TaskEventResp_StatusCode_value = map[string]int32{
		"SUCCESS": 0,
		"FAILURE": -1,
	}
)

func (x TaskEventResp_StatusCode) Enum() *TaskEventResp_StatusCode {
	p := new(TaskEventResp_StatusCode)
	*p = x
	return p
}

func (x TaskEventResp_StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskEventResp_StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_task_event_proto_enumTypes[0].Descriptor()
}

func (TaskEventResp_StatusCode) Type() protoreflect.EnumType {
	return &file_task_event_proto_enumTypes[0]
}

func (x TaskEventResp_StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskEventResp_StatusCode.Descriptor instead.
func (TaskEventResp_StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_task_event_proto_rawDescGZIP(), []int{2, 0}
}

// The request message
type TaskEventReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType string            `protobuf:"bytes,1,opt,name=eventType,proto3" json:"eventType,omitempty"`
	Data      *TaskEventReqData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TaskEventReq) Reset() {
	*x = TaskEventReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskEventReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskEventReq) ProtoMessage() {}

func (x *TaskEventReq) ProtoReflect() protoreflect.Message {
	mi := &file_task_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskEventReq.ProtoReflect.Descriptor instead.
func (*TaskEventReq) Descriptor() ([]byte, []int) {
	return file_task_event_proto_rawDescGZIP(), []int{0}
}

func (x *TaskEventReq) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *TaskEventReq) GetData() *TaskEventReqData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TaskEventReqData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcessID         string `protobuf:"bytes,1,opt,name=processID,proto3" json:"processID,omitempty"`
	ProcessInstanceID string `protobuf:"bytes,2,opt,name=processInstanceID,proto3" json:"processInstanceID,omitempty"`
	NodeDefKey        string `protobuf:"bytes,3,opt,name=nodeDefKey,proto3" json:"nodeDefKey,omitempty"`
	TaskID            string `protobuf:"bytes,4,opt,name=taskID,proto3" json:"taskID,omitempty"` // 如果是会签则taskID是多个，英文逗号间隔的字符串
}

func (x *TaskEventReqData) Reset() {
	*x = TaskEventReqData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskEventReqData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskEventReqData) ProtoMessage() {}

func (x *TaskEventReqData) ProtoReflect() protoreflect.Message {
	mi := &file_task_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskEventReqData.ProtoReflect.Descriptor instead.
func (*TaskEventReqData) Descriptor() ([]byte, []int) {
	return file_task_event_proto_rawDescGZIP(), []int{1}
}

func (x *TaskEventReqData) GetProcessID() string {
	if x != nil {
		return x.ProcessID
	}
	return ""
}

func (x *TaskEventReqData) GetProcessInstanceID() string {
	if x != nil {
		return x.ProcessInstanceID
	}
	return ""
}

func (x *TaskEventReqData) GetNodeDefKey() string {
	if x != nil {
		return x.NodeDefKey
	}
	return ""
}

func (x *TaskEventReqData) GetTaskID() string {
	if x != nil {
		return x.TaskID
	}
	return ""
}

// The response message
type TaskEventResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code TaskEventResp_StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=proto.TaskEventResp_StatusCode" json:"code,omitempty"`
	Msg  string                   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *TaskEventRespData       `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TaskEventResp) Reset() {
	*x = TaskEventResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskEventResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskEventResp) ProtoMessage() {}

func (x *TaskEventResp) ProtoReflect() protoreflect.Message {
	mi := &file_task_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskEventResp.ProtoReflect.Descriptor instead.
func (*TaskEventResp) Descriptor() ([]byte, []int) {
	return file_task_event_proto_rawDescGZIP(), []int{2}
}

func (x *TaskEventResp) GetCode() TaskEventResp_StatusCode {
	if x != nil {
		return x.Code
	}
	return TaskEventResp_SUCCESS
}

func (x *TaskEventResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *TaskEventResp) GetData() *TaskEventRespData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TaskEventRespData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExecuteType string `protobuf:"bytes,1,opt,name=executeType,proto3" json:"executeType,omitempty"` // 执行类型：endExecution，endProcess
}

func (x *TaskEventRespData) Reset() {
	*x = TaskEventRespData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskEventRespData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskEventRespData) ProtoMessage() {}

func (x *TaskEventRespData) ProtoReflect() protoreflect.Message {
	mi := &file_task_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskEventRespData.ProtoReflect.Descriptor instead.
func (*TaskEventRespData) Descriptor() ([]byte, []int) {
	return file_task_event_proto_rawDescGZIP(), []int{3}
}

func (x *TaskEventRespData) GetExecuteType() string {
	if x != nil {
		return x.ExecuteType
	}
	return ""
}

var File_task_event_proto protoreflect.FileDescriptor

var file_task_event_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0c, 0x54, 0x61, 0x73,
	0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x96, 0x01, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x66,
	0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x44,
	0x65, 0x66, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x22, 0xb5, 0x01,
	0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x33, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x2f, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x01, 0x22, 0x35, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x32, 0x41, 0x0a, 0x09,
	0x54, 0x61, 0x73, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_task_event_proto_rawDescOnce sync.Once
	file_task_event_proto_rawDescData = file_task_event_proto_rawDesc
)

func file_task_event_proto_rawDescGZIP() []byte {
	file_task_event_proto_rawDescOnce.Do(func() {
		file_task_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_event_proto_rawDescData)
	})
	return file_task_event_proto_rawDescData
}

var file_task_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_task_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_task_event_proto_goTypes = []interface{}{
	(TaskEventResp_StatusCode)(0), // 0: proto.TaskEventResp.StatusCode
	(*TaskEventReq)(nil),          // 1: proto.TaskEventReq
	(*TaskEventReqData)(nil),      // 2: proto.TaskEventReqData
	(*TaskEventResp)(nil),         // 3: proto.TaskEventResp
	(*TaskEventRespData)(nil),     // 4: proto.TaskEventRespData
}
var file_task_event_proto_depIdxs = []int32{
	2, // 0: proto.TaskEventReq.data:type_name -> proto.TaskEventReqData
	0, // 1: proto.TaskEventResp.code:type_name -> proto.TaskEventResp.StatusCode
	4, // 2: proto.TaskEventResp.data:type_name -> proto.TaskEventRespData
	1, // 3: proto.TaskEvent.Event:input_type -> proto.TaskEventReq
	3, // 4: proto.TaskEvent.Event:output_type -> proto.TaskEventResp
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_task_event_proto_init() }
func file_task_event_proto_init() {
	if File_task_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_task_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskEventReq); i {
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
		file_task_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskEventReqData); i {
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
		file_task_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskEventResp); i {
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
		file_task_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskEventRespData); i {
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
			RawDescriptor: file_task_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_event_proto_goTypes,
		DependencyIndexes: file_task_event_proto_depIdxs,
		EnumInfos:         file_task_event_proto_enumTypes,
		MessageInfos:      file_task_event_proto_msgTypes,
	}.Build()
	File_task_event_proto = out.File
	file_task_event_proto_rawDesc = nil
	file_task_event_proto_goTypes = nil
	file_task_event_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TaskEventClient is the client API for TaskEvent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskEventClient interface {
	Event(ctx context.Context, in *TaskEventReq, opts ...grpc.CallOption) (*TaskEventResp, error)
}

type taskEventClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskEventClient(cc grpc.ClientConnInterface) TaskEventClient {
	return &taskEventClient{cc}
}

func (c *taskEventClient) Event(ctx context.Context, in *TaskEventReq, opts ...grpc.CallOption) (*TaskEventResp, error) {
	out := new(TaskEventResp)
	err := c.cc.Invoke(ctx, "/proto.TaskEvent/Event", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskEventServer is the server API for TaskEvent service.
type TaskEventServer interface {
	Event(context.Context, *TaskEventReq) (*TaskEventResp, error)
}

// UnimplementedTaskEventServer can be embedded to have forward compatible implementations.
type UnimplementedTaskEventServer struct {
}

func (*UnimplementedTaskEventServer) Event(context.Context, *TaskEventReq) (*TaskEventResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Event not implemented")
}

func RegisterTaskEventServer(s *grpc.Server, srv TaskEventServer) {
	s.RegisterService(&_TaskEvent_serviceDesc, srv)
}

func _TaskEvent_Event_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskEventServer).Event(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TaskEvent/Event",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskEventServer).Event(ctx, req.(*TaskEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskEvent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TaskEvent",
	HandlerType: (*TaskEventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Event",
			Handler:    _TaskEvent_Event_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task_event.proto",
}
