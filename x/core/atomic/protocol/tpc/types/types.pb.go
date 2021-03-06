// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cross/core/atomic/tpc/types.proto

package types

import (
	fmt "fmt"
	types1 "github.com/datachainlab/cross/x/core/atomic/types"
	github_com_datachainlab_cross_x_core_tx_types "github.com/datachainlab/cross/x/core/tx/types"
	types "github.com/datachainlab/cross/x/core/tx/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CommitStatus int32

const (
	COMMIT_STATUS_UNKNOWN CommitStatus = 0
	COMMIT_STATUS_OK      CommitStatus = 1
	COMMIT_STATUS_FAILED  CommitStatus = 2
)

var CommitStatus_name = map[int32]string{
	0: "COMMIT_STATUS_UNKNOWN",
	1: "COMMIT_STATUS_OK",
	2: "COMMIT_STATUS_FAILED",
}

var CommitStatus_value = map[string]int32{
	"COMMIT_STATUS_UNKNOWN": 0,
	"COMMIT_STATUS_OK":      1,
	"COMMIT_STATUS_FAILED":  2,
}

func (x CommitStatus) String() string {
	return proto.EnumName(CommitStatus_name, int32(x))
}

func (CommitStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0e47c0d3dab52b5, []int{0}
}

type PacketDataPrepare struct {
	TxId    github_com_datachainlab_cross_x_core_tx_types.TxID    `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3,casttype=github.com/datachainlab/cross/x/core/tx/types.TxID" json:"tx_id,omitempty"`
	TxIndex github_com_datachainlab_cross_x_core_tx_types.TxIndex `protobuf:"varint,2,opt,name=tx_index,json=txIndex,proto3,casttype=github.com/datachainlab/cross/x/core/tx/types.TxIndex" json:"tx_index,omitempty"`
	Tx      types.ResolvedContractTransaction                     `protobuf:"bytes,3,opt,name=tx,proto3" json:"tx"`
}

func (m *PacketDataPrepare) Reset()         { *m = PacketDataPrepare{} }
func (m *PacketDataPrepare) String() string { return proto.CompactTextString(m) }
func (*PacketDataPrepare) ProtoMessage()    {}
func (*PacketDataPrepare) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0e47c0d3dab52b5, []int{0}
}
func (m *PacketDataPrepare) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketDataPrepare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketDataPrepare.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketDataPrepare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketDataPrepare.Merge(m, src)
}
func (m *PacketDataPrepare) XXX_Size() int {
	return m.Size()
}
func (m *PacketDataPrepare) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketDataPrepare.DiscardUnknown(m)
}

var xxx_messageInfo_PacketDataPrepare proto.InternalMessageInfo

type PacketAcknowledgementPrepare struct {
	Result types1.PrepareResult `protobuf:"varint,1,opt,name=result,proto3,enum=cross.core.atomic.PrepareResult" json:"result,omitempty"`
}

func (m *PacketAcknowledgementPrepare) Reset()         { *m = PacketAcknowledgementPrepare{} }
func (m *PacketAcknowledgementPrepare) String() string { return proto.CompactTextString(m) }
func (*PacketAcknowledgementPrepare) ProtoMessage()    {}
func (*PacketAcknowledgementPrepare) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0e47c0d3dab52b5, []int{1}
}
func (m *PacketAcknowledgementPrepare) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketAcknowledgementPrepare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketAcknowledgementPrepare.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketAcknowledgementPrepare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketAcknowledgementPrepare.Merge(m, src)
}
func (m *PacketAcknowledgementPrepare) XXX_Size() int {
	return m.Size()
}
func (m *PacketAcknowledgementPrepare) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketAcknowledgementPrepare.DiscardUnknown(m)
}

var xxx_messageInfo_PacketAcknowledgementPrepare proto.InternalMessageInfo

type PacketDataCommit struct {
	TxId          github_com_datachainlab_cross_x_core_tx_types.TxID    `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3,casttype=github.com/datachainlab/cross/x/core/tx/types.TxID" json:"tx_id,omitempty"`
	TxIndex       github_com_datachainlab_cross_x_core_tx_types.TxIndex `protobuf:"varint,2,opt,name=tx_index,json=txIndex,proto3,casttype=github.com/datachainlab/cross/x/core/tx/types.TxIndex" json:"tx_index,omitempty"`
	IsCommittable bool                                                  `protobuf:"varint,3,opt,name=is_committable,json=isCommittable,proto3" json:"is_committable,omitempty"`
}

func (m *PacketDataCommit) Reset()         { *m = PacketDataCommit{} }
func (m *PacketDataCommit) String() string { return proto.CompactTextString(m) }
func (*PacketDataCommit) ProtoMessage()    {}
func (*PacketDataCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0e47c0d3dab52b5, []int{2}
}
func (m *PacketDataCommit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketDataCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketDataCommit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketDataCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketDataCommit.Merge(m, src)
}
func (m *PacketDataCommit) XXX_Size() int {
	return m.Size()
}
func (m *PacketDataCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketDataCommit.DiscardUnknown(m)
}

var xxx_messageInfo_PacketDataCommit proto.InternalMessageInfo

type PacketAcknowledgementCommit struct {
	Status       CommitStatus `protobuf:"varint,1,opt,name=status,proto3,enum=cross.core.atomic.tpc.CommitStatus" json:"status,omitempty"`
	ErrorMessage string       `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (m *PacketAcknowledgementCommit) Reset()         { *m = PacketAcknowledgementCommit{} }
func (m *PacketAcknowledgementCommit) String() string { return proto.CompactTextString(m) }
func (*PacketAcknowledgementCommit) ProtoMessage()    {}
func (*PacketAcknowledgementCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0e47c0d3dab52b5, []int{3}
}
func (m *PacketAcknowledgementCommit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketAcknowledgementCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketAcknowledgementCommit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketAcknowledgementCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketAcknowledgementCommit.Merge(m, src)
}
func (m *PacketAcknowledgementCommit) XXX_Size() int {
	return m.Size()
}
func (m *PacketAcknowledgementCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketAcknowledgementCommit.DiscardUnknown(m)
}

var xxx_messageInfo_PacketAcknowledgementCommit proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("cross.core.atomic.tpc.CommitStatus", CommitStatus_name, CommitStatus_value)
	proto.RegisterType((*PacketDataPrepare)(nil), "cross.core.atomic.tpc.PacketDataPrepare")
	proto.RegisterType((*PacketAcknowledgementPrepare)(nil), "cross.core.atomic.tpc.PacketAcknowledgementPrepare")
	proto.RegisterType((*PacketDataCommit)(nil), "cross.core.atomic.tpc.PacketDataCommit")
	proto.RegisterType((*PacketAcknowledgementCommit)(nil), "cross.core.atomic.tpc.PacketAcknowledgementCommit")
}

func init() { proto.RegisterFile("cross/core/atomic/tpc/types.proto", fileDescriptor_e0e47c0d3dab52b5) }

var fileDescriptor_e0e47c0d3dab52b5 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0xc7, 0x93, 0x5a, 0xeb, 0x3a, 0xb6, 0x4b, 0x0d, 0x5d, 0xe8, 0x56, 0x4d, 0xeb, 0x2e, 0x42,
	0xd9, 0x43, 0x02, 0x15, 0xc5, 0x17, 0x10, 0xfb, 0xa2, 0x50, 0x6a, 0xdb, 0x25, 0xcd, 0x22, 0x88,
	0x18, 0xa6, 0x93, 0xa1, 0x1b, 0x36, 0xc9, 0x94, 0x99, 0xa7, 0x1a, 0x3f, 0x80, 0xe0, 0xd1, 0x8f,
	0x20, 0xf8, 0x65, 0x7a, 0xdc, 0xa3, 0xa7, 0xa2, 0xed, 0xc5, 0xcf, 0xb0, 0x07, 0x91, 0x4e, 0x22,
	0xed, 0x6a, 0x0f, 0x7a, 0xdc, 0xdb, 0xf0, 0xcc, 0x2f, 0xff, 0x99, 0xe7, 0x97, 0x79, 0xd0, 0x6d,
	0xc2, 0x99, 0x10, 0x26, 0x61, 0x9c, 0x9a, 0x18, 0x58, 0xe0, 0x11, 0x13, 0xc6, 0xc4, 0x84, 0xf7,
	0x63, 0x2a, 0x8c, 0x31, 0x67, 0xc0, 0xb4, 0x1d, 0x89, 0x18, 0x4b, 0xc4, 0x88, 0x11, 0x03, 0xc6,
	0xa4, 0x54, 0x18, 0xb1, 0x11, 0x93, 0x84, 0xb9, 0x5c, 0xc5, 0x70, 0x69, 0x77, 0x2d, 0x0f, 0xa2,
	0xf5, 0x9c, 0xd2, 0xad, 0x0d, 0x47, 0xad, 0xb6, 0xf7, 0x7e, 0xaa, 0xe8, 0xfa, 0x21, 0x26, 0x27,
	0x14, 0x5a, 0x18, 0xf0, 0x21, 0xa7, 0x63, 0xcc, 0xa9, 0xd6, 0x41, 0x97, 0x21, 0x72, 0x3c, 0xb7,
	0xa8, 0x56, 0xd4, 0x6a, 0xb6, 0x71, 0xff, 0x6c, 0x56, 0xae, 0x8d, 0x3c, 0x38, 0x9e, 0x0c, 0x0d,
	0xc2, 0x02, 0xd3, 0xc5, 0x80, 0xc9, 0x31, 0xf6, 0x42, 0x1f, 0x0f, 0xcd, 0x38, 0x3f, 0xfa, 0xe3,
	0x70, 0x3b, 0x6a, 0xb7, 0xac, 0x34, 0x44, 0x6d, 0x57, 0xb3, 0xd1, 0xd6, 0x32, 0x2c, 0x74, 0x69,
	0x54, 0x4c, 0x55, 0xd4, 0x6a, 0xae, 0xf1, 0xf0, 0x6c, 0x56, 0xbe, 0xf7, 0xdf, 0x79, 0xcb, 0x00,
	0xeb, 0x0a, 0xc4, 0x0b, 0xed, 0x29, 0x4a, 0x41, 0x54, 0xbc, 0x54, 0x51, 0xab, 0xd7, 0x6a, 0x07,
	0xc6, 0x9a, 0x2c, 0x88, 0x0c, 0x8b, 0x0a, 0xe6, 0xbf, 0xa5, 0x6e, 0x93, 0x85, 0xc0, 0x31, 0x01,
	0x9b, 0xe3, 0x50, 0x60, 0x02, 0x1e, 0x0b, 0x1b, 0xe9, 0xe9, 0xac, 0xac, 0x58, 0x29, 0x88, 0x1e,
	0xa5, 0x7f, 0x7c, 0x2e, 0x2b, 0x7b, 0x6f, 0xd0, 0xcd, 0xb8, 0xff, 0x3a, 0x39, 0x09, 0xd9, 0x3b,
	0x9f, 0xba, 0x23, 0x1a, 0xd0, 0x10, 0x7e, 0xab, 0x78, 0x80, 0x32, 0x9c, 0x8a, 0x89, 0x0f, 0xd2,
	0xc5, 0x76, 0xad, 0x62, 0xfc, 0xfd, 0x63, 0x12, 0xd6, 0x92, 0x9c, 0x95, 0xf0, 0x49, 0xfe, 0x5c,
	0x45, 0xf9, 0x95, 0xe0, 0x26, 0x0b, 0x02, 0x0f, 0x2e, 0x82, 0xdf, 0x3b, 0x68, 0xdb, 0x13, 0x0e,
	0x91, 0xf7, 0x05, 0x3c, 0xf4, 0xa9, 0x74, 0xbd, 0x65, 0xe5, 0x3c, 0xd1, 0x5c, 0x15, 0x93, 0x26,
	0x3f, 0xa8, 0xe8, 0xc6, 0x46, 0x8b, 0x49, 0xbf, 0x8f, 0x51, 0x46, 0x00, 0x86, 0x89, 0x48, 0x24,
	0xee, 0x1b, 0x1b, 0x5f, 0xb7, 0x11, 0xe3, 0x03, 0x89, 0x5a, 0xc9, 0x27, 0xda, 0x3e, 0xca, 0x51,
	0xce, 0x19, 0x77, 0x02, 0x2a, 0x04, 0x1e, 0x51, 0xd9, 0xe4, 0x55, 0x2b, 0x2b, 0x8b, 0xdd, 0xb8,
	0x16, 0xdf, 0xe3, 0xc0, 0x41, 0xd9, 0xf5, 0x08, 0x6d, 0x17, 0xed, 0x34, 0xfb, 0xdd, 0x6e, 0xdb,
	0x76, 0x06, 0x76, 0xdd, 0x3e, 0x1a, 0x38, 0x47, 0xbd, 0x4e, 0xaf, 0xff, 0xb2, 0x97, 0x57, 0xb4,
	0x02, 0xca, 0x9f, 0xdf, 0xea, 0x77, 0xf2, 0xaa, 0x56, 0x44, 0x85, 0xf3, 0xd5, 0xe7, 0xf5, 0xf6,
	0x8b, 0x67, 0xad, 0x7c, 0xaa, 0x94, 0xfe, 0xf8, 0x45, 0x57, 0x1a, 0xaf, 0xa7, 0xdf, 0x75, 0x65,
	0x3a, 0xd7, 0xd5, 0xd3, 0xb9, 0xae, 0x7e, 0x9b, 0xeb, 0xea, 0xa7, 0x85, 0xae, 0x9c, 0x2e, 0x74,
	0xe5, 0xeb, 0x42, 0x57, 0x5e, 0x3d, 0xf9, 0x27, 0xe7, 0xc9, 0x14, 0xca, 0xf9, 0x23, 0xcc, 0x5f,
	0x4d, 0xfe, 0x30, 0x23, 0x6b, 0x77, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x24, 0xb9, 0xaf,
	0x1f, 0x04, 0x00, 0x00,
}

func (m *PacketDataPrepare) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketDataPrepare) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketDataPrepare) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.TxIndex != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.TxIndex))
		i--
		dAtA[i] = 0x10
	}
	if len(m.TxId) > 0 {
		i -= len(m.TxId)
		copy(dAtA[i:], m.TxId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.TxId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PacketAcknowledgementPrepare) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketAcknowledgementPrepare) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketAcknowledgementPrepare) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Result != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PacketDataCommit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketDataCommit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketDataCommit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsCommittable {
		i--
		if m.IsCommittable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.TxIndex != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.TxIndex))
		i--
		dAtA[i] = 0x10
	}
	if len(m.TxId) > 0 {
		i -= len(m.TxId)
		copy(dAtA[i:], m.TxId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.TxId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PacketAcknowledgementCommit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketAcknowledgementCommit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketAcknowledgementCommit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ErrorMessage) > 0 {
		i -= len(m.ErrorMessage)
		copy(dAtA[i:], m.ErrorMessage)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ErrorMessage)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PacketDataPrepare) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.TxIndex != 0 {
		n += 1 + sovTypes(uint64(m.TxIndex))
	}
	l = m.Tx.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *PacketAcknowledgementPrepare) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != 0 {
		n += 1 + sovTypes(uint64(m.Result))
	}
	return n
}

func (m *PacketDataCommit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.TxIndex != 0 {
		n += 1 + sovTypes(uint64(m.TxIndex))
	}
	if m.IsCommittable {
		n += 2
	}
	return n
}

func (m *PacketAcknowledgementCommit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovTypes(uint64(m.Status))
	}
	l = len(m.ErrorMessage)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PacketDataPrepare) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PacketDataPrepare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketDataPrepare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxId = append(m.TxId[:0], dAtA[iNdEx:postIndex]...)
			if m.TxId == nil {
				m.TxId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxIndex", wireType)
			}
			m.TxIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxIndex |= github_com_datachainlab_cross_x_core_tx_types.TxIndex(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PacketAcknowledgementPrepare) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PacketAcknowledgementPrepare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketAcknowledgementPrepare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= types1.PrepareResult(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PacketDataCommit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PacketDataCommit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketDataCommit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxId = append(m.TxId[:0], dAtA[iNdEx:postIndex]...)
			if m.TxId == nil {
				m.TxId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxIndex", wireType)
			}
			m.TxIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxIndex |= github_com_datachainlab_cross_x_core_tx_types.TxIndex(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsCommittable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsCommittable = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PacketAcknowledgementCommit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PacketAcknowledgementCommit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketAcknowledgementCommit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= CommitStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
