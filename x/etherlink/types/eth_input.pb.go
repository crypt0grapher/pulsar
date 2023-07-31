// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: etherlink/etherlink/eth_input.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type EthInput struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	EthAddress string `protobuf:"bytes,2,opt,name=eth_address,json=ethAddress,proto3" json:"eth_address,omitempty"`
	EthSlot    string `protobuf:"bytes,3,opt,name=eth_slot,json=ethSlot,proto3" json:"eth_slot,omitempty"`
}

func (m *EthInput) Reset()         { *m = EthInput{} }
func (m *EthInput) String() string { return proto.CompactTextString(m) }
func (*EthInput) ProtoMessage()    {}
func (*EthInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb7ed35340dc9b7d, []int{0}
}
func (m *EthInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthInput.Merge(m, src)
}
func (m *EthInput) XXX_Size() int {
	return m.Size()
}
func (m *EthInput) XXX_DiscardUnknown() {
	xxx_messageInfo_EthInput.DiscardUnknown(m)
}

var xxx_messageInfo_EthInput proto.InternalMessageInfo

func (m *EthInput) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *EthInput) GetEthAddress() string {
	if m != nil {
		return m.EthAddress
	}
	return ""
}

func (m *EthInput) GetEthSlot() string {
	if m != nil {
		return m.EthSlot
	}
	return ""
}

func init() {
	proto.RegisterType((*EthInput)(nil), "etherlink.etherlink.EthInput")
}

func init() {
	proto.RegisterFile("etherlink/etherlink/eth_input.proto", fileDescriptor_fb7ed35340dc9b7d)
}

var fileDescriptor_fb7ed35340dc9b7d = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x2d, 0xc9, 0x48,
	0x2d, 0xca, 0xc9, 0xcc, 0xcb, 0xd6, 0x47, 0x61, 0xc5, 0x67, 0xe6, 0x15, 0x94, 0x96, 0xe8, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0xc3, 0xa5, 0xf4, 0xe0, 0x2c, 0xa5, 0x04, 0x2e, 0x0e, 0xd7,
	0x92, 0x0c, 0x4f, 0x90, 0x32, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xa2, 0xd4, 0xc4, 0x92, 0xfc, 0x22,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x48, 0x9e, 0x8b, 0x1b, 0x64, 0x5a, 0x62,
	0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1, 0x04, 0x13, 0x58, 0x96, 0x2b, 0xb5, 0x24, 0xc3, 0x11, 0x22,
	0x22, 0x24, 0xc9, 0xc5, 0x01, 0x52, 0x50, 0x9c, 0x93, 0x5f, 0x22, 0xc1, 0x0c, 0xd1, 0x9b, 0x5a,
	0x92, 0x11, 0x9c, 0x93, 0x5f, 0xe2, 0x64, 0x7a, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72,
	0x0c, 0x51, 0xd2, 0x08, 0xb7, 0x56, 0x20, 0xb9, 0xbb, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d,
	0xec, 0x68, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0x60, 0x06, 0x69, 0xdb, 0x00, 0x00,
	0x00,
}

func (m *EthInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthInput) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthInput) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EthSlot) > 0 {
		i -= len(m.EthSlot)
		copy(dAtA[i:], m.EthSlot)
		i = encodeVarintEthInput(dAtA, i, uint64(len(m.EthSlot)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EthAddress) > 0 {
		i -= len(m.EthAddress)
		copy(dAtA[i:], m.EthAddress)
		i = encodeVarintEthInput(dAtA, i, uint64(len(m.EthAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEthInput(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEthInput(dAtA []byte, offset int, v uint64) int {
	offset -= sovEthInput(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EthInput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEthInput(uint64(l))
	}
	l = len(m.EthAddress)
	if l > 0 {
		n += 1 + l + sovEthInput(uint64(l))
	}
	l = len(m.EthSlot)
	if l > 0 {
		n += 1 + l + sovEthInput(uint64(l))
	}
	return n
}

func sovEthInput(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEthInput(x uint64) (n int) {
	return sovEthInput(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthInput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEthInput
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
			return fmt.Errorf("proto: EthInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthInput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthInput
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
				return ErrInvalidLengthEthInput
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthInput
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
				return ErrInvalidLengthEthInput
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthSlot", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthInput
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
				return ErrInvalidLengthEthInput
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthSlot = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEthInput(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEthInput
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
func skipEthInput(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEthInput
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
					return 0, ErrIntOverflowEthInput
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
					return 0, ErrIntOverflowEthInput
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
				return 0, ErrInvalidLengthEthInput
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEthInput
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEthInput
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEthInput        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEthInput          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEthInput = fmt.Errorf("proto: unexpected end of group")
)
