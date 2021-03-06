// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cross/core/auth/packets.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	github_com_datachainlab_cross_x_core_account_types "github.com/datachainlab/cross/x/core/account/types"
	github_com_datachainlab_cross_x_core_tx_types "github.com/datachainlab/cross/x/core/tx/types"
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

type PacketDataIBCSignTx struct {
	TxID    github_com_datachainlab_cross_x_core_tx_types.TxID             `protobuf:"bytes,1,opt,name=txID,proto3,casttype=github.com/datachainlab/cross/x/core/tx/types.TxID" json:"txID,omitempty"`
	Signers []github_com_datachainlab_cross_x_core_account_types.AccountID `protobuf:"bytes,2,rep,name=signers,proto3,casttype=github.com/datachainlab/cross/x/core/account/types.AccountID" json:"signers,omitempty"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight types.Height `protobuf:"bytes,3,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
}

func (m *PacketDataIBCSignTx) Reset()         { *m = PacketDataIBCSignTx{} }
func (m *PacketDataIBCSignTx) String() string { return proto.CompactTextString(m) }
func (*PacketDataIBCSignTx) ProtoMessage()    {}
func (*PacketDataIBCSignTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf3907971b9c2156, []int{0}
}
func (m *PacketDataIBCSignTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketDataIBCSignTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketDataIBCSignTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketDataIBCSignTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketDataIBCSignTx.Merge(m, src)
}
func (m *PacketDataIBCSignTx) XXX_Size() int {
	return m.Size()
}
func (m *PacketDataIBCSignTx) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketDataIBCSignTx.DiscardUnknown(m)
}

var xxx_messageInfo_PacketDataIBCSignTx proto.InternalMessageInfo

type PacketAcknowledgementIBCSignTx struct {
	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *PacketAcknowledgementIBCSignTx) Reset()         { *m = PacketAcknowledgementIBCSignTx{} }
func (m *PacketAcknowledgementIBCSignTx) String() string { return proto.CompactTextString(m) }
func (*PacketAcknowledgementIBCSignTx) ProtoMessage()    {}
func (*PacketAcknowledgementIBCSignTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf3907971b9c2156, []int{1}
}
func (m *PacketAcknowledgementIBCSignTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketAcknowledgementIBCSignTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketAcknowledgementIBCSignTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketAcknowledgementIBCSignTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketAcknowledgementIBCSignTx.Merge(m, src)
}
func (m *PacketAcknowledgementIBCSignTx) XXX_Size() int {
	return m.Size()
}
func (m *PacketAcknowledgementIBCSignTx) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketAcknowledgementIBCSignTx.DiscardUnknown(m)
}

var xxx_messageInfo_PacketAcknowledgementIBCSignTx proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PacketDataIBCSignTx)(nil), "cross.core.auth.PacketDataIBCSignTx")
	proto.RegisterType((*PacketAcknowledgementIBCSignTx)(nil), "cross.core.auth.PacketAcknowledgementIBCSignTx")
}

func init() { proto.RegisterFile("cross/core/auth/packets.proto", fileDescriptor_bf3907971b9c2156) }

var fileDescriptor_bf3907971b9c2156 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x6e, 0xd4, 0x40,
	0x10, 0xc6, 0xed, 0xe4, 0x14, 0xa4, 0x25, 0xe1, 0x8f, 0xf9, 0x23, 0xeb, 0x44, 0xd6, 0x27, 0x57,
	0x57, 0x79, 0x95, 0x20, 0x51, 0x44, 0x08, 0x11, 0xe3, 0x02, 0x23, 0x21, 0x21, 0x73, 0x55, 0x1a,
	0x58, 0x6f, 0x56, 0xf6, 0x2a, 0xb6, 0xd7, 0xf2, 0x8e, 0x83, 0xf3, 0x16, 0x3c, 0x02, 0x8f, 0x73,
	0x65, 0x4a, 0x2a, 0x0b, 0xee, 0x1a, 0x4a, 0x94, 0x32, 0x15, 0xf2, 0xae, 0x8f, 0x28, 0x5d, 0x2a,
	0xcf, 0xcc, 0xf7, 0xcd, 0xcf, 0xf6, 0xa7, 0x41, 0xfb, 0xac, 0x91, 0x4a, 0x11, 0x26, 0x1b, 0x4e,
	0x68, 0x0b, 0x39, 0xa9, 0x29, 0x3b, 0xe3, 0xa0, 0x82, 0xba, 0x91, 0x20, 0x9d, 0x87, 0x5a, 0x0e,
	0x06, 0x39, 0x18, 0xe4, 0xe9, 0xd3, 0x4c, 0x66, 0x52, 0x6b, 0x64, 0xa8, 0x8c, 0x6d, 0xea, 0x89,
	0x94, 0x19, 0x06, 0x2b, 0x04, 0xaf, 0x80, 0x9c, 0x1f, 0x8c, 0x95, 0x31, 0xf8, 0x7f, 0xb7, 0xd0,
	0x93, 0x4f, 0x9a, 0x1c, 0x51, 0xa0, 0x71, 0xf8, 0xee, 0xb3, 0xc8, 0xaa, 0x45, 0xe7, 0x7c, 0x40,
	0x13, 0xe8, 0xe2, 0xc8, 0xb5, 0x67, 0xf6, 0x7c, 0x37, 0x7c, 0x75, 0xdd, 0x7b, 0x87, 0x99, 0x80,
	0xbc, 0x4d, 0x03, 0x26, 0x4b, 0x72, 0x4a, 0x81, 0xb2, 0x9c, 0x8a, 0xaa, 0xa0, 0x29, 0x31, 0x1f,
	0xda, 0x99, 0xd7, 0x40, 0x47, 0xe0, 0xa2, 0xe6, 0x2a, 0x58, 0x74, 0x71, 0x94, 0x68, 0x86, 0x73,
	0x82, 0xee, 0x29, 0x91, 0x55, 0xbc, 0x51, 0xee, 0xd6, 0x6c, 0x7b, 0xbe, 0x1b, 0xbe, 0xbd, 0xee,
	0xbd, 0xd7, 0x77, 0xc2, 0x51, 0xc6, 0x64, 0x5b, 0xc1, 0xc8, 0x3c, 0x36, 0x5d, 0x1c, 0x25, 0x1b,
	0xa0, 0xf3, 0x15, 0x3d, 0x00, 0x51, 0x72, 0xd9, 0xc2, 0x97, 0x9c, 0x8b, 0x2c, 0x07, 0x77, 0x7b,
	0x66, 0xcf, 0xef, 0x1f, 0x4e, 0x03, 0x91, 0x32, 0x13, 0xcf, 0xf8, 0xbf, 0xe7, 0x07, 0xc1, 0x7b,
	0xed, 0x08, 0xf7, 0x97, 0xbd, 0x67, 0x5d, 0xf5, 0xde, 0xb3, 0x0b, 0x5a, 0x16, 0x47, 0xfe, 0xed,
	0x7d, 0x3f, 0xd9, 0x1b, 0x07, 0xc6, 0xed, 0xc4, 0xe8, 0xf1, 0xc6, 0x31, 0x3c, 0x15, 0xd0, 0xb2,
	0x76, 0x27, 0x33, 0x7b, 0x3e, 0x09, 0x5f, 0x5c, 0xf5, 0x9e, 0x7b, 0x1b, 0xf2, 0xdf, 0xe2, 0x27,
	0x8f, 0xc6, 0xd9, 0x62, 0x33, 0x3a, 0x9a, 0xfc, 0xf9, 0xe1, 0x59, 0xfe, 0x1b, 0x84, 0x4d, 0xe2,
	0xc7, 0xec, 0xac, 0x92, 0xdf, 0x0a, 0x7e, 0x9a, 0xf1, 0x92, 0x57, 0x70, 0x13, 0xfe, 0x73, 0xb4,
	0xa3, 0x80, 0x42, 0xab, 0x74, 0xfc, 0x7b, 0xc9, 0xd8, 0x99, 0xfd, 0xf0, 0xe3, 0xf2, 0x37, 0xb6,
	0x96, 0x2b, 0x6c, 0x5f, 0xae, 0xb0, 0xfd, 0x6b, 0x85, 0xed, 0xef, 0x6b, 0x6c, 0x5d, 0xae, 0xb1,
	0xf5, 0x73, 0x8d, 0xad, 0x13, 0x72, 0xb7, 0x5c, 0x87, 0x8b, 0xd2, 0xa1, 0xa6, 0x3b, 0xfa, 0x10,
	0x5e, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x87, 0xd3, 0x04, 0xcd, 0x71, 0x02, 0x00, 0x00,
}

func (m *PacketDataIBCSignTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketDataIBCSignTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketDataIBCSignTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintPackets(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPackets(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintPackets(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.TxID) > 0 {
		i -= len(m.TxID)
		copy(dAtA[i:], m.TxID)
		i = encodeVarintPackets(dAtA, i, uint64(len(m.TxID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PacketAcknowledgementIBCSignTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketAcknowledgementIBCSignTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketAcknowledgementIBCSignTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintPackets(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPackets(dAtA []byte, offset int, v uint64) int {
	offset -= sovPackets(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PacketDataIBCSignTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxID)
	if l > 0 {
		n += 1 + l + sovPackets(uint64(l))
	}
	if len(m.Signers) > 0 {
		for _, b := range m.Signers {
			l = len(b)
			n += 1 + l + sovPackets(uint64(l))
		}
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovPackets(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovPackets(uint64(m.TimeoutTimestamp))
	}
	return n
}

func (m *PacketAcknowledgementIBCSignTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovPackets(uint64(m.Status))
	}
	return n
}

func sovPackets(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPackets(x uint64) (n int) {
	return sovPackets(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PacketDataIBCSignTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPackets
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
			return fmt.Errorf("proto: PacketDataIBCSignTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketDataIBCSignTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
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
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxID = append(m.TxID[:0], dAtA[iNdEx:postIndex]...)
			if m.TxID == nil {
				m.TxID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
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
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, make([]byte, postIndex-iNdEx))
			copy(m.Signers[len(m.Signers)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
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
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPackets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPackets
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPackets
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
func (m *PacketAcknowledgementIBCSignTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPackets
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
			return fmt.Errorf("proto: PacketAcknowledgementIBCSignTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketAcknowledgementIBCSignTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPackets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPackets
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPackets
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
func skipPackets(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPackets
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
					return 0, ErrIntOverflowPackets
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
					return 0, ErrIntOverflowPackets
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
				return 0, ErrInvalidLengthPackets
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPackets
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPackets
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPackets        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPackets          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPackets = fmt.Errorf("proto: unexpected end of group")
)
