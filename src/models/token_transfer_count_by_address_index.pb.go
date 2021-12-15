// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: token_transfer_count_by_address_index.proto

package models

import (
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
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

type TokenTransferCountByAddressIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionHash string `protobuf:"bytes,1,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash"`
	LogIndex        uint64 `protobuf:"varint,2,opt,name=log_index,json=logIndex,proto3" json:"log_index"`
	Address         string `protobuf:"bytes,3,opt,name=address,proto3" json:"address"`
	BlockNumber     uint64 `protobuf:"varint,4,opt,name=block_number,json=blockNumber,proto3" json:"block_number"`
}

func (x *TokenTransferCountByAddressIndex) Reset() {
	*x = TokenTransferCountByAddressIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_transfer_count_by_address_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenTransferCountByAddressIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenTransferCountByAddressIndex) ProtoMessage() {}

func (x *TokenTransferCountByAddressIndex) ProtoReflect() protoreflect.Message {
	mi := &file_token_transfer_count_by_address_index_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenTransferCountByAddressIndex.ProtoReflect.Descriptor instead.
func (*TokenTransferCountByAddressIndex) Descriptor() ([]byte, []int) {
	return file_token_transfer_count_by_address_index_proto_rawDescGZIP(), []int{0}
}

func (x *TokenTransferCountByAddressIndex) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *TokenTransferCountByAddressIndex) GetLogIndex() uint64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *TokenTransferCountByAddressIndex) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TokenTransferCountByAddressIndex) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

var File_token_transfer_count_by_address_index_proto protoreflect.FileDescriptor

var file_token_transfer_count_by_address_index_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8d, 0x02, 0x0a, 0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x33, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x28, 0x01, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x25, 0x0a, 0x09, 0x6c,
	0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08,
	0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x22, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x28, 0x01, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x61, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x3e, 0xba, 0xb9,
	0x19, 0x3a, 0x0a, 0x38, 0x52, 0x36, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x78, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0b, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_transfer_count_by_address_index_proto_rawDescOnce sync.Once
	file_token_transfer_count_by_address_index_proto_rawDescData = file_token_transfer_count_by_address_index_proto_rawDesc
)

func file_token_transfer_count_by_address_index_proto_rawDescGZIP() []byte {
	file_token_transfer_count_by_address_index_proto_rawDescOnce.Do(func() {
		file_token_transfer_count_by_address_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_transfer_count_by_address_index_proto_rawDescData)
	})
	return file_token_transfer_count_by_address_index_proto_rawDescData
}

var file_token_transfer_count_by_address_index_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_token_transfer_count_by_address_index_proto_goTypes = []interface{}{
	(*TokenTransferCountByAddressIndex)(nil), // 0: models.TokenTransferCountByAddressIndex
}
var file_token_transfer_count_by_address_index_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_token_transfer_count_by_address_index_proto_init() }
func file_token_transfer_count_by_address_index_proto_init() {
	if File_token_transfer_count_by_address_index_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_token_transfer_count_by_address_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenTransferCountByAddressIndex); i {
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
			RawDescriptor: file_token_transfer_count_by_address_index_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_token_transfer_count_by_address_index_proto_goTypes,
		DependencyIndexes: file_token_transfer_count_by_address_index_proto_depIdxs,
		MessageInfos:      file_token_transfer_count_by_address_index_proto_msgTypes,
	}.Build()
	File_token_transfer_count_by_address_index_proto = out.File
	file_token_transfer_count_by_address_index_proto_rawDesc = nil
	file_token_transfer_count_by_address_index_proto_goTypes = nil
	file_token_transfer_count_by_address_index_proto_depIdxs = nil
}
