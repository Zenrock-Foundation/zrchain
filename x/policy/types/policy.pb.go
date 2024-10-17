// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zrchain/policy/policy.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

// Policy defines a policy that can be used to validate a transaction
type Policy struct {
	Creator string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64     `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name    string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Policy  *types.Any `protobuf:"bytes,4,opt,name=policy,proto3" json:"policy,omitempty"`
	Btl     uint64     `protobuf:"varint,5,opt,name=btl,proto3" json:"btl,omitempty"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_53e400df02d23a81, []int{0}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return m.Size()
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Policy) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Policy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policy) GetPolicy() *types.Any {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *Policy) GetBtl() uint64 {
	if m != nil {
		return m.Btl
	}
	return 0
}

// BoolparserPolicy defines a policy that can be used to validate a transaction
type BoolparserPolicy struct {
	// Definition of the policy, eg.
	// "t1 + t2 + t3 > 1"
	Definition   string               `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
	Participants []*PolicyParticipant `protobuf:"bytes,2,rep,name=participants,proto3" json:"participants,omitempty"`
}

func (m *BoolparserPolicy) Reset()         { *m = BoolparserPolicy{} }
func (m *BoolparserPolicy) String() string { return proto.CompactTextString(m) }
func (*BoolparserPolicy) ProtoMessage()    {}
func (*BoolparserPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_53e400df02d23a81, []int{1}
}
func (m *BoolparserPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BoolparserPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BoolparserPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BoolparserPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolparserPolicy.Merge(m, src)
}
func (m *BoolparserPolicy) XXX_Size() int {
	return m.Size()
}
func (m *BoolparserPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolparserPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_BoolparserPolicy proto.InternalMessageInfo

func (m *BoolparserPolicy) GetDefinition() string {
	if m != nil {
		return m.Definition
	}
	return ""
}

func (m *BoolparserPolicy) GetParticipants() []*PolicyParticipant {
	if m != nil {
		return m.Participants
	}
	return nil
}

// PolicyParticipant defines a participant in a policy
type PolicyParticipant struct {
	Abbreviation string `protobuf:"bytes,1,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	Address      string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *PolicyParticipant) Reset()         { *m = PolicyParticipant{} }
func (m *PolicyParticipant) String() string { return proto.CompactTextString(m) }
func (*PolicyParticipant) ProtoMessage()    {}
func (*PolicyParticipant) Descriptor() ([]byte, []int) {
	return fileDescriptor_53e400df02d23a81, []int{2}
}
func (m *PolicyParticipant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PolicyParticipant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PolicyParticipant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PolicyParticipant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyParticipant.Merge(m, src)
}
func (m *PolicyParticipant) XXX_Size() int {
	return m.Size()
}
func (m *PolicyParticipant) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyParticipant.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyParticipant proto.InternalMessageInfo

func (m *PolicyParticipant) GetAbbreviation() string {
	if m != nil {
		return m.Abbreviation
	}
	return ""
}

func (m *PolicyParticipant) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Policy)(nil), "zrchain.policy.Policy")
	proto.RegisterType((*BoolparserPolicy)(nil), "zrchain.policy.BoolparserPolicy")
	proto.RegisterType((*PolicyParticipant)(nil), "zrchain.policy.PolicyParticipant")
}

func init() { proto.RegisterFile("zrchain/policy/policy.proto", fileDescriptor_53e400df02d23a81) }

var fileDescriptor_53e400df02d23a81 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x4f, 0xc2, 0x30,
	0x18, 0xc6, 0xe9, 0x40, 0x0c, 0x85, 0x10, 0x6c, 0x3c, 0x54, 0x4d, 0x9a, 0xb9, 0xd3, 0x0e, 0xa6,
	0x4b, 0xd0, 0x83, 0x57, 0x49, 0xbc, 0xe3, 0x8e, 0xde, 0xda, 0xad, 0x40, 0xe3, 0x68, 0x97, 0xae,
	0x10, 0xc7, 0x07, 0xf0, 0xec, 0xc7, 0xf2, 0xc8, 0xd1, 0xa3, 0x81, 0x2f, 0x62, 0xe8, 0xb6, 0x08,
	0xf1, 0xb4, 0xf7, 0xcf, 0x6f, 0xef, 0xf3, 0xf6, 0x79, 0xe1, 0xcd, 0xc6, 0x24, 0x0b, 0x26, 0x55,
	0x94, 0xeb, 0x4c, 0x26, 0x65, 0xfd, 0xa1, 0xb9, 0xd1, 0x56, 0xa3, 0x61, 0xdd, 0xa4, 0x55, 0xf5,
	0xfa, 0x6a, 0xae, 0xf5, 0x3c, 0x13, 0x91, 0xeb, 0xf2, 0xd5, 0x2c, 0x62, 0xaa, 0x46, 0x83, 0x0f,
	0x00, 0xbb, 0x53, 0x47, 0x21, 0x0c, 0xcf, 0x13, 0x23, 0x98, 0xd5, 0x06, 0x03, 0x1f, 0x84, 0xbd,
	0xb8, 0x49, 0xd1, 0x10, 0x7a, 0x32, 0xc5, 0x9e, 0x0f, 0xc2, 0x4e, 0xec, 0xc9, 0x14, 0x21, 0xd8,
	0x51, 0x6c, 0x29, 0x70, 0xdb, 0x61, 0x2e, 0x46, 0x77, 0xb0, 0x5b, 0xa9, 0xe1, 0x8e, 0x0f, 0xc2,
	0xfe, 0xf8, 0x92, 0x56, 0xa2, 0xb4, 0x11, 0xa5, 0x4f, 0xaa, 0x8c, 0x6b, 0x06, 0x8d, 0x60, 0x9b,
	0xdb, 0x0c, 0x9f, 0xb9, 0x91, 0x87, 0x30, 0x28, 0xe1, 0x68, 0xa2, 0x75, 0x96, 0x33, 0x53, 0x08,
	0x53, 0x6f, 0x44, 0x20, 0x4c, 0xc5, 0x4c, 0x2a, 0x69, 0xa5, 0x56, 0xf5, 0x52, 0x47, 0x15, 0xf4,
	0x0c, 0x07, 0x39, 0x33, 0x56, 0x26, 0x32, 0x67, 0xca, 0x16, 0xd8, 0xf3, 0xdb, 0x61, 0x7f, 0x7c,
	0x4b, 0x4f, 0x9f, 0x4f, 0xab, 0x69, 0xd3, 0x3f, 0x32, 0x3e, 0xf9, 0x2d, 0x78, 0x81, 0x17, 0xff,
	0x10, 0x14, 0xc0, 0x01, 0xe3, 0xdc, 0x88, 0xb5, 0x64, 0x47, 0xea, 0x27, 0xb5, 0x83, 0x63, 0x2c,
	0x4d, 0x8d, 0x28, 0x0a, 0x67, 0x4e, 0x2f, 0x6e, 0xd2, 0x49, 0xfc, 0xb5, 0x23, 0x60, 0xbb, 0x23,
	0xe0, 0x67, 0x47, 0xc0, 0xe7, 0x9e, 0xb4, 0xb6, 0x7b, 0xd2, 0xfa, 0xde, 0x93, 0xd6, 0xeb, 0xe3,
	0x5c, 0xda, 0xc5, 0x8a, 0xd3, 0x44, 0x2f, 0xa3, 0x8d, 0x50, 0x46, 0x27, 0x6f, 0x19, 0xe3, 0x45,
	0x13, 0x47, 0xcd, 0x5d, 0xd7, 0x0f, 0xd1, 0x7b, 0x73, 0x5c, 0x5b, 0xe6, 0xa2, 0xe0, 0x5d, 0xe7,
	0xe4, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x29, 0xf9, 0x76, 0xfb, 0x01, 0x00, 0x00,
}

func (m *Policy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Policy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Policy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Btl != 0 {
		i = encodeVarintPolicy(dAtA, i, uint64(m.Btl))
		i--
		dAtA[i] = 0x28
	}
	if m.Policy != nil {
		{
			size, err := m.Policy.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPolicy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintPolicy(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BoolparserPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BoolparserPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BoolparserPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Participants) > 0 {
		for iNdEx := len(m.Participants) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Participants[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPolicy(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Definition) > 0 {
		i -= len(m.Definition)
		copy(dAtA[i:], m.Definition)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Definition)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PolicyParticipant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PolicyParticipant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PolicyParticipant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Abbreviation) > 0 {
		i -= len(m.Abbreviation)
		copy(dAtA[i:], m.Abbreviation)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Abbreviation)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPolicy(dAtA []byte, offset int, v uint64) int {
	offset -= sovPolicy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Policy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovPolicy(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	if m.Policy != nil {
		l = m.Policy.Size()
		n += 1 + l + sovPolicy(uint64(l))
	}
	if m.Btl != 0 {
		n += 1 + sovPolicy(uint64(m.Btl))
	}
	return n
}

func (m *BoolparserPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Definition)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	if len(m.Participants) > 0 {
		for _, e := range m.Participants {
			l = e.Size()
			n += 1 + l + sovPolicy(uint64(l))
		}
	}
	return n
}

func (m *PolicyParticipant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Abbreviation)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	return n
}

func sovPolicy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPolicy(x uint64) (n int) {
	return sovPolicy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Policy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
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
			return fmt.Errorf("proto: Policy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Policy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Policy == nil {
				m.Policy = &types.Any{}
			}
			if err := m.Policy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Btl", wireType)
			}
			m.Btl = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Btl |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicy
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
func (m *BoolparserPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
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
			return fmt.Errorf("proto: BoolparserPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BoolparserPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Definition", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Definition = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participants", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participants = append(m.Participants, &PolicyParticipant{})
			if err := m.Participants[len(m.Participants)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicy
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
func (m *PolicyParticipant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
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
			return fmt.Errorf("proto: PolicyParticipant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PolicyParticipant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abbreviation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Abbreviation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicy
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
func skipPolicy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPolicy
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
					return 0, ErrIntOverflowPolicy
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
					return 0, ErrIntOverflowPolicy
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
				return 0, ErrInvalidLengthPolicy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPolicy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPolicy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPolicy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPolicy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPolicy = fmt.Errorf("proto: unexpected end of group")
)
