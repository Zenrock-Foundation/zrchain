// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zrchain/identity/keyring.proto

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

// defines the keyring format
type Keyring struct {
	Address     string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Creator     string   `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Admins      []string `protobuf:"bytes,4,rep,name=admins,proto3" json:"admins,omitempty"`
	Parties     []string `protobuf:"bytes,5,rep,name=parties,proto3" json:"parties,omitempty"`
	// The MPC threshold, i.e. the number of parties required to submit signed txs
	// in order for a request to be fulfilled
	PartyThreshold uint32 `protobuf:"varint,6,opt,name=party_threshold,json=partyThreshold,proto3" json:"party_threshold,omitempty"`
	// deprecated
	KeyReqFee uint64 `protobuf:"varint,7,opt,name=key_req_fee,json=keyReqFee,proto3" json:"key_req_fee,omitempty"`
	// deprecated
	SigReqFee    uint64       `protobuf:"varint,8,opt,name=sig_req_fee,json=sigReqFee,proto3" json:"sig_req_fee,omitempty"`
	IsActive     bool         `protobuf:"varint,9,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	DelegateFees bool         `protobuf:"varint,10,opt,name=delegate_fees,json=delegateFees,proto3" json:"delegate_fees,omitempty"`
	Fees         *KeyringFees `protobuf:"bytes,11,opt,name=fees,proto3" json:"fees,omitempty"`
}

func (m *Keyring) Reset()         { *m = Keyring{} }
func (m *Keyring) String() string { return proto.CompactTextString(m) }
func (*Keyring) ProtoMessage()    {}
func (*Keyring) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1bb3c09d479ee27, []int{0}
}
func (m *Keyring) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Keyring) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Keyring.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Keyring) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keyring.Merge(m, src)
}
func (m *Keyring) XXX_Size() int {
	return m.Size()
}
func (m *Keyring) XXX_DiscardUnknown() {
	xxx_messageInfo_Keyring.DiscardUnknown(m)
}

var xxx_messageInfo_Keyring proto.InternalMessageInfo

func (m *Keyring) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Keyring) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Keyring) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Keyring) GetAdmins() []string {
	if m != nil {
		return m.Admins
	}
	return nil
}

func (m *Keyring) GetParties() []string {
	if m != nil {
		return m.Parties
	}
	return nil
}

func (m *Keyring) GetPartyThreshold() uint32 {
	if m != nil {
		return m.PartyThreshold
	}
	return 0
}

func (m *Keyring) GetKeyReqFee() uint64 {
	if m != nil {
		return m.KeyReqFee
	}
	return 0
}

func (m *Keyring) GetSigReqFee() uint64 {
	if m != nil {
		return m.SigReqFee
	}
	return 0
}

func (m *Keyring) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *Keyring) GetDelegateFees() bool {
	if m != nil {
		return m.DelegateFees
	}
	return false
}

func (m *Keyring) GetFees() *KeyringFees {
	if m != nil {
		return m.Fees
	}
	return nil
}

type KeyringFee struct {
	RockAmount uint64 `protobuf:"varint,1,opt,name=rock_amount,json=rockAmount,proto3" json:"rock_amount,omitempty"`
	UsdAmount  uint64 `protobuf:"varint,2,opt,name=usd_amount,json=usdAmount,proto3" json:"usd_amount,omitempty"`
}

func (m *KeyringFee) Reset()         { *m = KeyringFee{} }
func (m *KeyringFee) String() string { return proto.CompactTextString(m) }
func (*KeyringFee) ProtoMessage()    {}
func (*KeyringFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1bb3c09d479ee27, []int{1}
}
func (m *KeyringFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyringFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyringFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyringFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyringFee.Merge(m, src)
}
func (m *KeyringFee) XXX_Size() int {
	return m.Size()
}
func (m *KeyringFee) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyringFee.DiscardUnknown(m)
}

var xxx_messageInfo_KeyringFee proto.InternalMessageInfo

func (m *KeyringFee) GetRockAmount() uint64 {
	if m != nil {
		return m.RockAmount
	}
	return 0
}

func (m *KeyringFee) GetUsdAmount() uint64 {
	if m != nil {
		return m.UsdAmount
	}
	return 0
}

type KeyringFees struct {
	Signature *KeyringFee `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Key       *KeyringFee `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *KeyringFees) Reset()         { *m = KeyringFees{} }
func (m *KeyringFees) String() string { return proto.CompactTextString(m) }
func (*KeyringFees) ProtoMessage()    {}
func (*KeyringFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1bb3c09d479ee27, []int{2}
}
func (m *KeyringFees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyringFees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyringFees.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyringFees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyringFees.Merge(m, src)
}
func (m *KeyringFees) XXX_Size() int {
	return m.Size()
}
func (m *KeyringFees) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyringFees.DiscardUnknown(m)
}

var xxx_messageInfo_KeyringFees proto.InternalMessageInfo

func (m *KeyringFees) GetSignature() *KeyringFee {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *KeyringFees) GetKey() *KeyringFee {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*Keyring)(nil), "zrchain.identity.Keyring")
	proto.RegisterType((*KeyringFee)(nil), "zrchain.identity.KeyringFee")
	proto.RegisterType((*KeyringFees)(nil), "zrchain.identity.KeyringFees")
}

func init() { proto.RegisterFile("zrchain/identity/keyring.proto", fileDescriptor_a1bb3c09d479ee27) }

var fileDescriptor_a1bb3c09d479ee27 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x31, 0x8f, 0xd3, 0x30,
	0x18, 0x86, 0x9b, 0xb6, 0xb4, 0xcd, 0x17, 0x0e, 0x90, 0x07, 0x64, 0x09, 0x2e, 0x44, 0x65, 0x20,
	0x0b, 0x89, 0x38, 0xc4, 0x00, 0xdb, 0x31, 0x74, 0x81, 0x29, 0x30, 0xdd, 0x12, 0xf9, 0xe2, 0x8f,
	0xd4, 0xca, 0x35, 0xee, 0xd9, 0xce, 0x89, 0xf0, 0x2b, 0xf8, 0x59, 0x8c, 0x37, 0x32, 0xa2, 0xf6,
	0x07, 0xf0, 0x17, 0x90, 0x9d, 0x84, 0x56, 0x0c, 0xdc, 0x96, 0xf7, 0x7d, 0x5e, 0xdb, 0xf9, 0x3e,
	0xbd, 0x10, 0x7e, 0x53, 0xc5, 0x9a, 0x89, 0x3a, 0x15, 0x1c, 0x6b, 0x23, 0x4c, 0x9b, 0x56, 0xd8,
	0x2a, 0x51, 0x97, 0xc9, 0x56, 0x49, 0x23, 0xc9, 0xa3, 0x9e, 0x27, 0x03, 0x5f, 0xfe, 0x1e, 0xc3,
	0xfc, 0x43, 0x97, 0x21, 0x14, 0xe6, 0x8c, 0x73, 0x85, 0x5a, 0x53, 0x2f, 0xf2, 0x62, 0x3f, 0x1b,
	0xa4, 0x25, 0x85, 0x42, 0x66, 0xa4, 0xa2, 0xe3, 0x8e, 0xf4, 0x92, 0x44, 0x10, 0x70, 0xd4, 0x85,
	0x12, 0x5b, 0x23, 0x64, 0x4d, 0x27, 0x8e, 0x1e, 0x5b, 0xe4, 0x31, 0xcc, 0x18, 0xdf, 0x88, 0x5a,
	0xd3, 0x69, 0x34, 0x89, 0xfd, 0xac, 0x57, 0xf6, 0xce, 0x2d, 0x53, 0x46, 0xa0, 0xa6, 0xf7, 0x1c,
	0x18, 0x24, 0x79, 0x01, 0x0f, 0xed, 0x67, 0x9b, 0x9b, 0xb5, 0x42, 0xbd, 0x96, 0x57, 0x9c, 0xce,
	0x22, 0x2f, 0x3e, 0xc9, 0x1e, 0x38, 0xfb, 0xf3, 0xe0, 0x92, 0x10, 0x82, 0x0a, 0xdb, 0x5c, 0xe1,
	0x75, 0xfe, 0x05, 0x91, 0xce, 0x23, 0x2f, 0x9e, 0x66, 0x7e, 0x85, 0x6d, 0x86, 0xd7, 0x2b, 0x44,
	0xcb, 0xb5, 0x28, 0xff, 0xf2, 0x45, 0xc7, 0xb5, 0x28, 0x7b, 0xfe, 0x04, 0x7c, 0xa1, 0x73, 0x56,
	0x18, 0x71, 0x83, 0xd4, 0x8f, 0xbc, 0x78, 0x91, 0x2d, 0x84, 0x3e, 0x77, 0x9a, 0x3c, 0x87, 0x13,
	0x8e, 0x57, 0x58, 0x32, 0x83, 0xf6, 0xb4, 0xa6, 0xe0, 0x02, 0xf7, 0x07, 0x73, 0x85, 0xa8, 0xc9,
	0x2b, 0x98, 0x3a, 0x16, 0x44, 0x5e, 0x1c, 0x9c, 0x9d, 0x26, 0xff, 0xee, 0x37, 0xe9, 0x77, 0x6b,
	0xc3, 0x99, 0x8b, 0x2e, 0x3f, 0x02, 0x1c, 0x4c, 0xf2, 0x0c, 0x02, 0x25, 0x8b, 0x2a, 0x67, 0x1b,
	0xd9, 0xd4, 0xc6, 0xed, 0x7d, 0x9a, 0x81, 0xb5, 0xce, 0x9d, 0x43, 0x4e, 0x01, 0x1a, 0xcd, 0x07,
	0x3e, 0xee, 0x46, 0x68, 0x34, 0xef, 0xf0, 0xb2, 0x85, 0xe0, 0xe8, 0x09, 0xf2, 0x0e, 0xec, 0x78,
	0x35, 0x33, 0x8d, 0x42, 0x77, 0x59, 0x70, 0xf6, 0xf4, 0x7f, 0x3f, 0x95, 0x1d, 0xe2, 0x24, 0x81,
	0x49, 0x85, 0xad, 0x7b, 0xe2, 0xae, 0x53, 0x36, 0xf8, 0xfe, 0xd3, 0x8f, 0x5d, 0xe8, 0xdd, 0xee,
	0x42, 0xef, 0xd7, 0x2e, 0xf4, 0xbe, 0xef, 0xc3, 0xd1, 0xed, 0x3e, 0x1c, 0xfd, 0xdc, 0x87, 0xa3,
	0x8b, 0xb7, 0xa5, 0x30, 0xeb, 0xe6, 0x32, 0x29, 0xe4, 0x26, 0xbd, 0xc0, 0xda, 0x4e, 0xf3, 0x72,
	0x25, 0x9b, 0x9a, 0x33, 0xdb, 0x88, 0x74, 0x28, 0xe9, 0xcd, 0x9b, 0xf4, 0xeb, 0xa1, 0xa9, 0xa6,
	0xdd, 0xa2, 0xbe, 0x9c, 0xb9, 0xa2, 0xbe, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x14, 0x8c,
	0x64, 0xca, 0x02, 0x00, 0x00,
}

func (m *Keyring) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Keyring) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Keyring) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fees != nil {
		{
			size, err := m.Fees.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintKeyring(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if m.DelegateFees {
		i--
		if m.DelegateFees {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.SigReqFee != 0 {
		i = encodeVarintKeyring(dAtA, i, uint64(m.SigReqFee))
		i--
		dAtA[i] = 0x40
	}
	if m.KeyReqFee != 0 {
		i = encodeVarintKeyring(dAtA, i, uint64(m.KeyReqFee))
		i--
		dAtA[i] = 0x38
	}
	if m.PartyThreshold != 0 {
		i = encodeVarintKeyring(dAtA, i, uint64(m.PartyThreshold))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Parties) > 0 {
		for iNdEx := len(m.Parties) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Parties[iNdEx])
			copy(dAtA[i:], m.Parties[iNdEx])
			i = encodeVarintKeyring(dAtA, i, uint64(len(m.Parties[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Admins) > 0 {
		for iNdEx := len(m.Admins) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Admins[iNdEx])
			copy(dAtA[i:], m.Admins[iNdEx])
			i = encodeVarintKeyring(dAtA, i, uint64(len(m.Admins[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintKeyring(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintKeyring(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintKeyring(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KeyringFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyringFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyringFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UsdAmount != 0 {
		i = encodeVarintKeyring(dAtA, i, uint64(m.UsdAmount))
		i--
		dAtA[i] = 0x10
	}
	if m.RockAmount != 0 {
		i = encodeVarintKeyring(dAtA, i, uint64(m.RockAmount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *KeyringFees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyringFees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyringFees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Key != nil {
		{
			size, err := m.Key.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintKeyring(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Signature != nil {
		{
			size, err := m.Signature.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintKeyring(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintKeyring(dAtA []byte, offset int, v uint64) int {
	offset -= sovKeyring(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Keyring) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovKeyring(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovKeyring(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovKeyring(uint64(l))
	}
	if len(m.Admins) > 0 {
		for _, s := range m.Admins {
			l = len(s)
			n += 1 + l + sovKeyring(uint64(l))
		}
	}
	if len(m.Parties) > 0 {
		for _, s := range m.Parties {
			l = len(s)
			n += 1 + l + sovKeyring(uint64(l))
		}
	}
	if m.PartyThreshold != 0 {
		n += 1 + sovKeyring(uint64(m.PartyThreshold))
	}
	if m.KeyReqFee != 0 {
		n += 1 + sovKeyring(uint64(m.KeyReqFee))
	}
	if m.SigReqFee != 0 {
		n += 1 + sovKeyring(uint64(m.SigReqFee))
	}
	if m.IsActive {
		n += 2
	}
	if m.DelegateFees {
		n += 2
	}
	if m.Fees != nil {
		l = m.Fees.Size()
		n += 1 + l + sovKeyring(uint64(l))
	}
	return n
}

func (m *KeyringFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RockAmount != 0 {
		n += 1 + sovKeyring(uint64(m.RockAmount))
	}
	if m.UsdAmount != 0 {
		n += 1 + sovKeyring(uint64(m.UsdAmount))
	}
	return n
}

func (m *KeyringFees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Signature != nil {
		l = m.Signature.Size()
		n += 1 + l + sovKeyring(uint64(l))
	}
	if m.Key != nil {
		l = m.Key.Size()
		n += 1 + l + sovKeyring(uint64(l))
	}
	return n
}

func sovKeyring(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKeyring(x uint64) (n int) {
	return sovKeyring(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Keyring) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeyring
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
			return fmt.Errorf("proto: Keyring: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Keyring: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyring
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
				return ErrInvalidLengthKeyring
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeyring
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyring
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
				return ErrInvalidLengthKeyring
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeyring
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyring
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
				return ErrInvalidLengthKeyring
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeyring
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admins", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyring
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
				return ErrInvalidLengthKeyring
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeyring
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admins = append(m.Admins, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyring
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
				return ErrInvalidLengthKeyring
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeyring
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parties = append(m.Parties, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartyThreshold", wireType)
			}
			m.PartyThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyring
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PartyThreshold |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyReqFee", wireType)
			}
			m.KeyReqFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyring
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyReqFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigReqFee", wireType)
			}
			m.SigReqFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyring
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigReqFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyring
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
			m.IsActive = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegateFees", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyring
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
			m.DelegateFees = bool(v != 0)
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyring
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
				return ErrInvalidLengthKeyring
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeyring
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fees == nil {
				m.Fees = &KeyringFees{}
			}
			if err := m.Fees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeyring(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeyring
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
func (m *KeyringFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeyring
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
			return fmt.Errorf("proto: KeyringFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyringFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RockAmount", wireType)
			}
			m.RockAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyring
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RockAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsdAmount", wireType)
			}
			m.UsdAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyring
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UsdAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKeyring(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeyring
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
func (m *KeyringFees) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeyring
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
			return fmt.Errorf("proto: KeyringFees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyringFees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyring
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
				return ErrInvalidLengthKeyring
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeyring
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Signature == nil {
				m.Signature = &KeyringFee{}
			}
			if err := m.Signature.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyring
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
				return ErrInvalidLengthKeyring
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeyring
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Key == nil {
				m.Key = &KeyringFee{}
			}
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeyring(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeyring
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
func skipKeyring(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeyring
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
					return 0, ErrIntOverflowKeyring
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
					return 0, ErrIntOverflowKeyring
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
				return 0, ErrInvalidLengthKeyring
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKeyring
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKeyring
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKeyring        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeyring          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKeyring = fmt.Errorf("proto: unexpected end of group")
)
