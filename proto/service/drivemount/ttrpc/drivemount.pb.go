// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.20.1
// source: drivemount.proto

package drivemount

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MountDriveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriveID         string   `protobuf:"bytes,1,opt,name=DriveID,proto3" json:"DriveID,omitempty"`
	DestinationPath string   `protobuf:"bytes,2,opt,name=DestinationPath,proto3" json:"DestinationPath,omitempty"`
	FilesytemType   string   `protobuf:"bytes,3,opt,name=FilesytemType,proto3" json:"FilesytemType,omitempty"`
	Options         []string `protobuf:"bytes,4,rep,name=Options,proto3" json:"Options,omitempty"`
}

func (x *MountDriveRequest) Reset() {
	*x = MountDriveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivemount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountDriveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountDriveRequest) ProtoMessage() {}

func (x *MountDriveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drivemount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountDriveRequest.ProtoReflect.Descriptor instead.
func (*MountDriveRequest) Descriptor() ([]byte, []int) {
	return file_drivemount_proto_rawDescGZIP(), []int{0}
}

func (x *MountDriveRequest) GetDriveID() string {
	if x != nil {
		return x.DriveID
	}
	return ""
}

func (x *MountDriveRequest) GetDestinationPath() string {
	if x != nil {
		return x.DestinationPath
	}
	return ""
}

func (x *MountDriveRequest) GetFilesytemType() string {
	if x != nil {
		return x.FilesytemType
	}
	return ""
}

func (x *MountDriveRequest) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

type UnmountDriveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriveID string `protobuf:"bytes,1,opt,name=DriveID,proto3" json:"DriveID,omitempty"`
}

func (x *UnmountDriveRequest) Reset() {
	*x = UnmountDriveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivemount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnmountDriveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnmountDriveRequest) ProtoMessage() {}

func (x *UnmountDriveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drivemount_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnmountDriveRequest.ProtoReflect.Descriptor instead.
func (*UnmountDriveRequest) Descriptor() ([]byte, []int) {
	return file_drivemount_proto_rawDescGZIP(), []int{1}
}

func (x *UnmountDriveRequest) GetDriveID() string {
	if x != nil {
		return x.DriveID
	}
	return ""
}

var File_drivemount_proto protoreflect.FileDescriptor

var file_drivemount_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x72, 0x69, 0x76, 0x65, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x97, 0x01, 0x0a, 0x11, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x72, 0x69, 0x76, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x44, 0x72, 0x69, 0x76, 0x65, 0x49, 0x44, 0x12,
	0x28, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x79, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x44, 0x72, 0x69, 0x76, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x44, 0x72, 0x69, 0x76, 0x65, 0x49, 0x44, 0x32, 0x86, 0x01, 0x0a, 0x0c, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0a, 0x4d,
	0x6f, 0x75, 0x6e, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x12, 0x12, 0x2e, 0x4d, 0x6f, 0x75, 0x6e,
	0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x0c, 0x55, 0x6e, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x12, 0x14, 0x2e, 0x55, 0x6e, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_drivemount_proto_rawDescOnce sync.Once
	file_drivemount_proto_rawDescData = file_drivemount_proto_rawDesc
)

func file_drivemount_proto_rawDescGZIP() []byte {
	file_drivemount_proto_rawDescOnce.Do(func() {
		file_drivemount_proto_rawDescData = protoimpl.X.CompressGZIP(file_drivemount_proto_rawDescData)
	})
	return file_drivemount_proto_rawDescData
}

var file_drivemount_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_drivemount_proto_goTypes = []interface{}{
	(*MountDriveRequest)(nil),   // 0: MountDriveRequest
	(*UnmountDriveRequest)(nil), // 1: UnmountDriveRequest
	(*emptypb.Empty)(nil),       // 2: google.protobuf.Empty
}
var file_drivemount_proto_depIdxs = []int32{
	0, // 0: DriveMounter.MountDrive:input_type -> MountDriveRequest
	1, // 1: DriveMounter.UnmountDrive:input_type -> UnmountDriveRequest
	2, // 2: DriveMounter.MountDrive:output_type -> google.protobuf.Empty
	2, // 3: DriveMounter.UnmountDrive:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_drivemount_proto_init() }
func file_drivemount_proto_init() {
	if File_drivemount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_drivemount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountDriveRequest); i {
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
		file_drivemount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnmountDriveRequest); i {
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
			RawDescriptor: file_drivemount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_drivemount_proto_goTypes,
		DependencyIndexes: file_drivemount_proto_depIdxs,
		MessageInfos:      file_drivemount_proto_msgTypes,
	}.Build()
	File_drivemount_proto = out.File
	file_drivemount_proto_rawDesc = nil
	file_drivemount_proto_goTypes = nil
	file_drivemount_proto_depIdxs = nil
}
