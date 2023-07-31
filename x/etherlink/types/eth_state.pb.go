// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: etherlink/etherlink/eth_state.proto

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

type EthState struct {
	EthAddress     string `protobuf:"bytes,1,opt,name=eth_address,json=ethAddress,proto3" json:"eth_address,omitempty"`
	EthSlot        string `protobuf:"bytes,2,opt,name=eth_slot,json=ethSlot,proto3" json:"eth_slot,omitempty"`
	EthMerkleRoot  string `protobuf:"bytes,3,opt,name=eth_merkle_root,json=ethMerkleRoot,proto3" json:"eth_merkle_root,omitempty"`
	EthMerkleProof string `protobuf:"bytes,4,opt,name=eth_merkle_proof,json=ethMerkleProof,proto3" json:"eth_merkle_proof,omitempty"`
	EthBlockHeight string `protobuf:"bytes,5,opt,name=eth_block_height,json=ethBlockHeight,proto3" json:"eth_block_height,omitempty"`
	EthState       string `protobuf:"bytes,6,opt,name=eth_state,json=ethState,proto3" json:"eth_state,omitempty"`
}

func (m *EthState) Reset()         { *m = EthState{} }
func (m *EthState) String() string { return proto.CompactTextString(m) }
func (*EthState) ProtoMessage()    {}
func (*EthState) Descriptor() ([]byte, []int) {
	return fileDescriptor_904decc7374ff2af, []int{0}
}
func (m *EthState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthState.Merge(m, src)
}
func (m *EthState) XXX_Size() int {
	return m.Size()
}
func (m *EthState) XXX_DiscardUnknown() {
	xxx_messageInfo_EthState.DiscardUnknown(m)
}

var xxx_messageInfo_EthState proto.InternalMessageInfo

func (m *EthState) GetEthAddress() string {
	if m != nil {
		return m.EthAddress
	}
	return ""
}

func (m *EthState) GetEthSlot() string {
	if m != nil {
		return m.EthSlot
	}
	return ""
}

func (m *EthState) GetEthMerkleRoot() string {
	if m != nil {
		return m.EthMerkleRoot
	}
	return ""
}

func (m *EthState) GetEthMerkleProof() string {
	if m != nil {
		return m.EthMerkleProof
	}
	return ""
}

func (m *EthState) GetEthBlockHeight() string {
	if m != nil {
		return m.EthBlockHeight
	}
	return ""
}

func (m *EthState) GetEthState() string {
	if m != nil {
		return m.EthState
	}
	return ""
}

func init() {
	proto.RegisterType((*EthState)(nil), "etherlink.etherlink.EthState")
}

func init() {
	proto.RegisterFile("etherlink/etherlink/eth_state.proto", fileDescriptor_904decc7374ff2af)
}

var fileDescriptor_904decc7374ff2af = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x2d, 0xc9, 0x48,
	0x2d, 0xca, 0xc9, 0xcc, 0xcb, 0xd6, 0x47, 0x61, 0xc5, 0x17, 0x97, 0x24, 0x96, 0xa4, 0xea, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0xc3, 0xa5, 0xf4, 0xe0, 0x2c, 0xa5, 0xfb, 0x8c, 0x5c, 0x1c,
	0xae, 0x25, 0x19, 0xc1, 0x20, 0x75, 0x42, 0xf2, 0x5c, 0xdc, 0x20, 0x4d, 0x89, 0x29, 0x29, 0x45,
	0xa9, 0xc5, 0xc5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x5c, 0xa9, 0x25, 0x19, 0x8e, 0x10,
	0x11, 0x21, 0x49, 0x2e, 0x0e, 0xb0, 0xa9, 0x39, 0xf9, 0x25, 0x12, 0x4c, 0x60, 0x59, 0xf6, 0xd4,
	0x92, 0x8c, 0xe0, 0x9c, 0xfc, 0x12, 0x21, 0x35, 0x2e, 0x7e, 0x90, 0x54, 0x6e, 0x6a, 0x51, 0x76,
	0x4e, 0x6a, 0x7c, 0x51, 0x7e, 0x7e, 0x89, 0x04, 0x33, 0x58, 0x05, 0x6f, 0x6a, 0x49, 0x86, 0x2f,
	0x58, 0x34, 0x28, 0x3f, 0xbf, 0x44, 0x48, 0x83, 0x4b, 0x00, 0x49, 0x5d, 0x41, 0x51, 0x7e, 0x7e,
	0x9a, 0x04, 0x0b, 0x58, 0x21, 0x1f, 0x5c, 0x61, 0x00, 0x48, 0x14, 0xa6, 0x32, 0x29, 0x27, 0x3f,
	0x39, 0x3b, 0x3e, 0x23, 0x35, 0x33, 0x3d, 0xa3, 0x44, 0x82, 0x15, 0xae, 0xd2, 0x09, 0x24, 0xec,
	0x01, 0x16, 0x15, 0x92, 0xe6, 0xe2, 0x84, 0x7b, 0x56, 0x82, 0x0d, 0xac, 0x04, 0xe4, 0x4e, 0xb0,
	0xa7, 0x9c, 0x4c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x1a, 0x11,
	0x56, 0x15, 0x48, 0xe1, 0x56, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x0e, 0x34, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x8c, 0xc7, 0xdd, 0x5b, 0x01, 0x00, 0x00,
}

func (m *EthState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EthState) > 0 {
		i -= len(m.EthState)
		copy(dAtA[i:], m.EthState)
		i = encodeVarintEthState(dAtA, i, uint64(len(m.EthState)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.EthBlockHeight) > 0 {
		i -= len(m.EthBlockHeight)
		copy(dAtA[i:], m.EthBlockHeight)
		i = encodeVarintEthState(dAtA, i, uint64(len(m.EthBlockHeight)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.EthMerkleProof) > 0 {
		i -= len(m.EthMerkleProof)
		copy(dAtA[i:], m.EthMerkleProof)
		i = encodeVarintEthState(dAtA, i, uint64(len(m.EthMerkleProof)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.EthMerkleRoot) > 0 {
		i -= len(m.EthMerkleRoot)
		copy(dAtA[i:], m.EthMerkleRoot)
		i = encodeVarintEthState(dAtA, i, uint64(len(m.EthMerkleRoot)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EthSlot) > 0 {
		i -= len(m.EthSlot)
		copy(dAtA[i:], m.EthSlot)
		i = encodeVarintEthState(dAtA, i, uint64(len(m.EthSlot)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.EthAddress) > 0 {
		i -= len(m.EthAddress)
		copy(dAtA[i:], m.EthAddress)
		i = encodeVarintEthState(dAtA, i, uint64(len(m.EthAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEthState(dAtA []byte, offset int, v uint64) int {
	offset -= sovEthState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EthState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EthAddress)
	if l > 0 {
		n += 1 + l + sovEthState(uint64(l))
	}
	l = len(m.EthSlot)
	if l > 0 {
		n += 1 + l + sovEthState(uint64(l))
	}
	l = len(m.EthMerkleRoot)
	if l > 0 {
		n += 1 + l + sovEthState(uint64(l))
	}
	l = len(m.EthMerkleProof)
	if l > 0 {
		n += 1 + l + sovEthState(uint64(l))
	}
	l = len(m.EthBlockHeight)
	if l > 0 {
		n += 1 + l + sovEthState(uint64(l))
	}
	l = len(m.EthState)
	if l > 0 {
		n += 1 + l + sovEthState(uint64(l))
	}
	return n
}

func sovEthState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEthState(x uint64) (n int) {
	return sovEthState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEthState
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
			return fmt.Errorf("proto: EthState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthState
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
				return ErrInvalidLengthEthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthSlot", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthState
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
				return ErrInvalidLengthEthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthSlot = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthMerkleRoot", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthState
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
				return ErrInvalidLengthEthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthMerkleRoot = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthMerkleProof", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthState
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
				return ErrInvalidLengthEthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthMerkleProof = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthBlockHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthState
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
				return ErrInvalidLengthEthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthBlockHeight = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthState", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthState
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
				return ErrInvalidLengthEthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthState = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEthState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEthState
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
func skipEthState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEthState
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
					return 0, ErrIntOverflowEthState
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
					return 0, ErrIntOverflowEthState
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
				return 0, ErrInvalidLengthEthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEthState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEthState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEthState = fmt.Errorf("proto: unexpected end of group")
)
