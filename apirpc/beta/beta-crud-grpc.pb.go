// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: beta-crud-grpc.proto

package betarpc

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

type BetaUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Status      string `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *BetaUpdateRequest) Reset() {
	*x = BetaUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beta_crud_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BetaUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BetaUpdateRequest) ProtoMessage() {}

func (x *BetaUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beta_crud_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BetaUpdateRequest.ProtoReflect.Descriptor instead.
func (*BetaUpdateRequest) Descriptor() ([]byte, []int) {
	return file_beta_crud_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *BetaUpdateRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *BetaUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BetaUpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BetaUpdateRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type BetaUpdateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *BetaUpdateReply) Reset() {
	*x = BetaUpdateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beta_crud_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BetaUpdateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BetaUpdateReply) ProtoMessage() {}

func (x *BetaUpdateReply) ProtoReflect() protoreflect.Message {
	mi := &file_beta_crud_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BetaUpdateReply.ProtoReflect.Descriptor instead.
func (*BetaUpdateReply) Descriptor() ([]byte, []int) {
	return file_beta_crud_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *BetaUpdateReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type BetaGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityID int64 `protobuf:"varint,1,opt,name=entityID,proto3" json:"entityID,omitempty"`
}

func (x *BetaGetRequest) Reset() {
	*x = BetaGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beta_crud_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BetaGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BetaGetRequest) ProtoMessage() {}

func (x *BetaGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beta_crud_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BetaGetRequest.ProtoReflect.Descriptor instead.
func (*BetaGetRequest) Descriptor() ([]byte, []int) {
	return file_beta_crud_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *BetaGetRequest) GetEntityID() int64 {
	if x != nil {
		return x.EntityID
	}
	return 0
}

type BetaGetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Status      string `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *BetaGetReply) Reset() {
	*x = BetaGetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beta_crud_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BetaGetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BetaGetReply) ProtoMessage() {}

func (x *BetaGetReply) ProtoReflect() protoreflect.Message {
	mi := &file_beta_crud_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BetaGetReply.ProtoReflect.Descriptor instead.
func (*BetaGetReply) Descriptor() ([]byte, []int) {
	return file_beta_crud_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *BetaGetReply) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *BetaGetReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BetaGetReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BetaGetReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_beta_crud_grpc_proto protoreflect.FileDescriptor

var file_beta_crud_grpc_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x65, 0x74, 0x61, 0x2d, 0x63, 0x72, 0x75, 0x64, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x71, 0x0a, 0x11,
	0x42, 0x65, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x29, 0x0a, 0x0f, 0x42, 0x65, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2c, 0x0a, 0x0e, 0x42, 0x65,
	0x74, 0x61, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x22, 0x6c, 0x0a, 0x0c, 0x42, 0x65, 0x74, 0x61,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x9a, 0x01, 0x0a, 0x0b, 0x42, 0x65, 0x74, 0x61, 0x43,
	0x52, 0x55, 0x44, 0x52, 0x50, 0x43, 0x12, 0x40, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x65, 0x74,
	0x61, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x42, 0x65, 0x74, 0x61, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x65, 0x74, 0x61, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x65, 0x74, 0x61, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x42, 0x65, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x63, 0x72, 0x75,
	0x64, 0x2f, 0x62, 0x65, 0x74, 0x61, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_beta_crud_grpc_proto_rawDescOnce sync.Once
	file_beta_crud_grpc_proto_rawDescData = file_beta_crud_grpc_proto_rawDesc
)

func file_beta_crud_grpc_proto_rawDescGZIP() []byte {
	file_beta_crud_grpc_proto_rawDescOnce.Do(func() {
		file_beta_crud_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_beta_crud_grpc_proto_rawDescData)
	})
	return file_beta_crud_grpc_proto_rawDescData
}

var file_beta_crud_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_beta_crud_grpc_proto_goTypes = []interface{}{
	(*BetaUpdateRequest)(nil), // 0: main.BetaUpdateRequest
	(*BetaUpdateReply)(nil),   // 1: main.BetaUpdateReply
	(*BetaGetRequest)(nil),    // 2: main.BetaGetRequest
	(*BetaGetReply)(nil),      // 3: main.BetaGetReply
}
var file_beta_crud_grpc_proto_depIdxs = []int32{
	2, // 0: main.BetaCRUDRPC.GetBetaInformation:input_type -> main.BetaGetRequest
	0, // 1: main.BetaCRUDRPC.UpdateBetaInformation:input_type -> main.BetaUpdateRequest
	3, // 2: main.BetaCRUDRPC.GetBetaInformation:output_type -> main.BetaGetReply
	1, // 3: main.BetaCRUDRPC.UpdateBetaInformation:output_type -> main.BetaUpdateReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_beta_crud_grpc_proto_init() }
func file_beta_crud_grpc_proto_init() {
	if File_beta_crud_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_beta_crud_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BetaUpdateRequest); i {
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
		file_beta_crud_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BetaUpdateReply); i {
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
		file_beta_crud_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BetaGetRequest); i {
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
		file_beta_crud_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BetaGetReply); i {
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
			RawDescriptor: file_beta_crud_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_beta_crud_grpc_proto_goTypes,
		DependencyIndexes: file_beta_crud_grpc_proto_depIdxs,
		MessageInfos:      file_beta_crud_grpc_proto_msgTypes,
	}.Build()
	File_beta_crud_grpc_proto = out.File
	file_beta_crud_grpc_proto_rawDesc = nil
	file_beta_crud_grpc_proto_goTypes = nil
	file_beta_crud_grpc_proto_depIdxs = nil
}
