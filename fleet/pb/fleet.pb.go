// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: fleet/pb/fleet.proto

package pb

import (
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

type AgentByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentID string `protobuf:"bytes,1,opt,name=agentID,proto3" json:"agentID,omitempty"`
	OwnerID string `protobuf:"bytes,2,opt,name=ownerID,proto3" json:"ownerID,omitempty"`
}

func (x *AgentByIDReq) Reset() {
	*x = AgentByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleet_pb_fleet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentByIDReq) ProtoMessage() {}

func (x *AgentByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_fleet_pb_fleet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentByIDReq.ProtoReflect.Descriptor instead.
func (*AgentByIDReq) Descriptor() ([]byte, []int) {
	return file_fleet_pb_fleet_proto_rawDescGZIP(), []int{0}
}

func (x *AgentByIDReq) GetAgentID() string {
	if x != nil {
		return x.AgentID
	}
	return ""
}

func (x *AgentByIDReq) GetOwnerID() string {
	if x != nil {
		return x.OwnerID
	}
	return ""
}

type AgentRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Channel string `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *AgentRes) Reset() {
	*x = AgentRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleet_pb_fleet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentRes) ProtoMessage() {}

func (x *AgentRes) ProtoReflect() protoreflect.Message {
	mi := &file_fleet_pb_fleet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentRes.ProtoReflect.Descriptor instead.
func (*AgentRes) Descriptor() ([]byte, []int) {
	return file_fleet_pb_fleet_proto_rawDescGZIP(), []int{1}
}

func (x *AgentRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AgentRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AgentRes) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

var File_fleet_pb_fleet_proto protoreflect.FileDescriptor

var file_fleet_pb_fleet_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x22, 0x42, 0x0a,
	0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49,
	0x44, 0x22, 0x48, 0x0a, 0x08, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x32, 0x47, 0x0a, 0x0c, 0x46,
	0x6c, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x1a, 0x0f, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fleet_pb_fleet_proto_rawDescOnce sync.Once
	file_fleet_pb_fleet_proto_rawDescData = file_fleet_pb_fleet_proto_rawDesc
)

func file_fleet_pb_fleet_proto_rawDescGZIP() []byte {
	file_fleet_pb_fleet_proto_rawDescOnce.Do(func() {
		file_fleet_pb_fleet_proto_rawDescData = protoimpl.X.CompressGZIP(file_fleet_pb_fleet_proto_rawDescData)
	})
	return file_fleet_pb_fleet_proto_rawDescData
}

var file_fleet_pb_fleet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fleet_pb_fleet_proto_goTypes = []interface{}{
	(*AgentByIDReq)(nil), // 0: fleet.AgentByIDReq
	(*AgentRes)(nil),     // 1: fleet.AgentRes
}
var file_fleet_pb_fleet_proto_depIdxs = []int32{
	0, // 0: fleet.FleetService.RetrieveAgent:input_type -> fleet.AgentByIDReq
	1, // 1: fleet.FleetService.RetrieveAgent:output_type -> fleet.AgentRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fleet_pb_fleet_proto_init() }
func file_fleet_pb_fleet_proto_init() {
	if File_fleet_pb_fleet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fleet_pb_fleet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentByIDReq); i {
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
		file_fleet_pb_fleet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentRes); i {
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
			RawDescriptor: file_fleet_pb_fleet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fleet_pb_fleet_proto_goTypes,
		DependencyIndexes: file_fleet_pb_fleet_proto_depIdxs,
		MessageInfos:      file_fleet_pb_fleet_proto_msgTypes,
	}.Build()
	File_fleet_pb_fleet_proto = out.File
	file_fleet_pb_fleet_proto_rawDesc = nil
	file_fleet_pb_fleet_proto_goTypes = nil
	file_fleet_pb_fleet_proto_depIdxs = nil
}
