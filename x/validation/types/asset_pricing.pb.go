// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zrchain/validation/asset_pricing.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type AssetPrice struct {
	PriceUSD cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=priceUSD,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"priceUSD"`
}

func (m *AssetPrice) Reset()         { *m = AssetPrice{} }
func (m *AssetPrice) String() string { return proto.CompactTextString(m) }
func (*AssetPrice) ProtoMessage()    {}
func (*AssetPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_badfc8e144e6be45, []int{0}
}
func (m *AssetPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetPrice.Merge(m, src)
}
func (m *AssetPrice) XXX_Size() int {
	return m.Size()
}
func (m *AssetPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetPrice.DiscardUnknown(m)
}

var xxx_messageInfo_AssetPrice proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AssetPrice)(nil), "zrchain.validation.AssetPrice")
}

func init() {
	proto.RegisterFile("zrchain/validation/asset_pricing.proto", fileDescriptor_badfc8e144e6be45)
}

var fileDescriptor_badfc8e144e6be45 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xab, 0x2a, 0x4a, 0xce,
	0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x4f,
	0x2c, 0x2e, 0x4e, 0x2d, 0x89, 0x2f, 0x28, 0xca, 0x4c, 0xce, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x82, 0xaa, 0xd3, 0x43, 0xa8, 0x93, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0x4b, 0xeb, 0x83, 0x58, 0x10, 0x95, 0x52, 0x92, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xf1, 0x10,
	0x09, 0x08, 0x07, 0x22, 0xa5, 0x14, 0xcd, 0xc5, 0xe5, 0x08, 0x32, 0x3b, 0xa0, 0x28, 0x33, 0x39,
	0x55, 0xc8, 0x97, 0x8b, 0x03, 0x64, 0x47, 0x6a, 0x68, 0xb0, 0x8b, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0xa7, 0x93, 0xe1, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9, 0x4b, 0x43, 0x74, 0x15, 0xa7, 0x64,
	0xeb, 0x65, 0xe6, 0xeb, 0xe7, 0x26, 0x96, 0x64, 0xe8, 0xf9, 0xa4, 0xa6, 0x27, 0x26, 0x57, 0xba,
	0xa4, 0x26, 0x5f, 0xda, 0xa2, 0xcb, 0x05, 0x35, 0xd4, 0x25, 0x35, 0x39, 0x08, 0x6e, 0x84, 0x53,
	0xd8, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1,
	0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xd9, 0xa4, 0x67, 0x96, 0x64,
	0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x57, 0xa5, 0xe6, 0x15, 0xe5, 0x27, 0x67, 0xe7, 0x24,
	0x26, 0x15, 0xc3, 0xd8, 0xfa, 0xf0, 0x20, 0x30, 0xd1, 0xaf, 0x40, 0x0e, 0x87, 0x92, 0xca, 0x82,
	0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xdb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x0c, 0x30,
	0x10, 0x2a, 0x01, 0x00, 0x00,
}

func (m *AssetPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.PriceUSD.Size()
		i -= size
		if _, err := m.PriceUSD.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAssetPricing(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAssetPricing(dAtA []byte, offset int, v uint64) int {
	offset -= sovAssetPricing(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AssetPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PriceUSD.Size()
	n += 1 + l + sovAssetPricing(uint64(l))
	return n
}

func sovAssetPricing(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAssetPricing(x uint64) (n int) {
	return sovAssetPricing(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AssetPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAssetPricing
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
			return fmt.Errorf("proto: AssetPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceUSD", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetPricing
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
				return ErrInvalidLengthAssetPricing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAssetPricing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PriceUSD.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAssetPricing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAssetPricing
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
func skipAssetPricing(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAssetPricing
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
					return 0, ErrIntOverflowAssetPricing
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
					return 0, ErrIntOverflowAssetPricing
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
				return 0, ErrInvalidLengthAssetPricing
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAssetPricing
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAssetPricing
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAssetPricing        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAssetPricing          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAssetPricing = fmt.Errorf("proto: unexpected end of group")
)
