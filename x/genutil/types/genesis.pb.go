// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zrchain/genutil/v1beta1/genesis.proto

package types

import (
	encoding_json "encoding/json"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the raw genesis transaction in JSON.
type GenesisState struct {
	// gen_txs defines the genesis transactions.
	GenTxs []encoding_json.RawMessage `protobuf:"bytes,1,rep,name=gen_txs,json=genTxs,proto3,casttype=encoding/json.RawMessage" json:"gentxs"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e58394b79961166, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetGenTxs() []encoding_json.RawMessage {
	if m != nil {
		return m.GenTxs
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "zrchain.genutil.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("zrchain/genutil/v1beta1/genesis.proto", fileDescriptor_5e58394b79961166)
}

var fileDescriptor_5e58394b79961166 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xad, 0x2a, 0x4a, 0xce,
	0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x4f, 0xcd, 0x2b, 0x2d, 0xc9, 0xcc, 0xd1, 0x2f, 0x33, 0x4c, 0x4a,
	0x2d, 0x49, 0x34, 0x04, 0xf1, 0x53, 0x8b, 0x33, 0x8b, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0xc4, 0xa1, 0xca, 0xf4, 0xa0, 0xca, 0xf4, 0xa0, 0xca, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1,
	0x6a, 0xf4, 0x41, 0x2c, 0x88, 0x72, 0x29, 0xc1, 0xc4, 0xdc, 0xcc, 0xbc, 0x7c, 0x7d, 0x30, 0x09,
	0x11, 0x52, 0x8a, 0xe7, 0xe2, 0x71, 0x87, 0x18, 0x19, 0x5c, 0x92, 0x58, 0x92, 0x2a, 0xe4, 0xcf,
	0xc5, 0x9e, 0x9e, 0x9a, 0x17, 0x5f, 0x52, 0x51, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0xe3, 0x64,
	0xf6, 0xea, 0x9e, 0x3c, 0x5b, 0x7a, 0x6a, 0x5e, 0x49, 0x45, 0xf1, 0xaf, 0x7b, 0xf2, 0x12, 0xa9,
	0x79, 0xc9, 0xf9, 0x29, 0x99, 0x79, 0xe9, 0xfa, 0x59, 0xc5, 0xf9, 0x79, 0x7a, 0x41, 0x89, 0xe5,
	0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x8b, 0x9e, 0x6f, 0xd0, 0x82, 0x2a, 0x5b, 0xf1, 0x7c,
	0x83, 0x16, 0x63, 0x10, 0x88, 0x13, 0x52, 0x51, 0xec, 0x14, 0x7c, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3,
	0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x96, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9,
	0xfa, 0x55, 0xa9, 0x79, 0x45, 0xf9, 0xc9, 0xd9, 0x39, 0x89, 0x49, 0xc5, 0x30, 0xb6, 0x3e, 0x2c,
	0x08, 0xca, 0x4c, 0xf4, 0x2b, 0xe0, 0xe1, 0x50, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x76,
	0xbc, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x61, 0xe8, 0x26, 0x25, 0x27, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GenTxs) > 0 {
		for iNdEx := len(m.GenTxs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.GenTxs[iNdEx])
			copy(dAtA[i:], m.GenTxs[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.GenTxs[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GenTxs) > 0 {
		for _, b := range m.GenTxs {
			l = len(b)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenTxs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenTxs = append(m.GenTxs, make([]byte, postIndex-iNdEx))
			copy(m.GenTxs[len(m.GenTxs)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
