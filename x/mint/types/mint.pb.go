// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zrchain/mint/v1beta1/mint.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// Minter represents the minting state.
type Minter struct {
	// current annual inflation rate
	Inflation cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=inflation,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation"`
	// current annual expected provisions
	AnnualProvisions cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=annual_provisions,json=annualProvisions,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"annual_provisions"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_008fb58650beebd1, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

// Params defines the parameters for the x/mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// maximum annual change in inflation rate
	InflationRateChange cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=inflation_rate_change,json=inflationRateChange,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation_rate_change"`
	// maximum inflation rate
	InflationMax cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=inflation_max,json=inflationMax,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation_max"`
	// minimum inflation rate
	InflationMin cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=inflation_min,json=inflationMin,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation_min"`
	// goal of percent bonded atoms
	GoalBonded cosmossdk_io_math.LegacyDec `protobuf:"bytes,5,opt,name=goal_bonded,json=goalBonded,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"goal_bonded"`
	// expected blocks per year
	BlocksPerYear uint64 `protobuf:"varint,6,opt,name=blocks_per_year,json=blocksPerYear,proto3" json:"blocks_per_year,omitempty"`
	// goal of percent staking yield
	StakingYield cosmossdk_io_math.LegacyDec `protobuf:"bytes,7,opt,name=staking_yield,json=stakingYield,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"staking_yield"`
	// rewards burn rate
	BurnRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,8,opt,name=burn_rate,json=burnRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"burn_rate"`
	// rewards protocol wallet rate
	ProtocolWalletRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,9,opt,name=protocol_wallet_rate,json=protocolWalletRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"protocol_wallet_rate"`
	// address of protocol wallet
	ProtocolWalletAddress string `protobuf:"bytes,10,opt,name=protocol_wallet_address,json=protocolWalletAddress,proto3" json:"protocol_wallet_address,omitempty"`
	// additional staking rewards rate
	AdditionalStakingRewards cosmossdk_io_math.LegacyDec `protobuf:"bytes,11,opt,name=additional_staking_rewards,json=additionalStakingRewards,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"additional_staking_rewards"`
	// additional MPC rewards rate
	AdditionalMpcRewards cosmossdk_io_math.LegacyDec `protobuf:"bytes,12,opt,name=additional_mpc_rewards,json=additionalMpcRewards,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"additional_mpc_rewards"`
	// additional burn rate
	AdditionalBurnRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,13,opt,name=additional_burn_rate,json=additionalBurnRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"additional_burn_rate"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_008fb58650beebd1, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetBlocksPerYear() uint64 {
	if m != nil {
		return m.BlocksPerYear
	}
	return 0
}

func (m *Params) GetProtocolWalletAddress() string {
	if m != nil {
		return m.ProtocolWalletAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Minter)(nil), "zrchain.mint.v1beta1.Minter")
	proto.RegisterType((*Params)(nil), "zrchain.mint.v1beta1.Params")
}

func init() { proto.RegisterFile("zrchain/mint/v1beta1/mint.proto", fileDescriptor_008fb58650beebd1) }

var fileDescriptor_008fb58650beebd1 = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x41, 0x4f, 0xd4, 0x4e,
	0x18, 0xc6, 0xb7, 0xff, 0x3f, 0xae, 0xec, 0x00, 0x51, 0xc6, 0x45, 0x2b, 0xc6, 0x85, 0x70, 0x30,
	0x04, 0xc3, 0x56, 0x62, 0xe4, 0xe0, 0xcd, 0x95, 0x78, 0x92, 0x40, 0xca, 0x81, 0x80, 0x89, 0xcd,
	0xdb, 0xe9, 0xd8, 0x8e, 0xdb, 0xce, 0x6c, 0x66, 0x66, 0x81, 0xf5, 0x23, 0x78, 0xf2, 0x33, 0x78,
	0xf2, 0xc8, 0xc1, 0x8b, 0xdf, 0x80, 0x23, 0xf1, 0x64, 0x3c, 0x10, 0xc3, 0x1e, 0xf8, 0x1a, 0xa6,
	0x33, 0xdd, 0x2d, 0x72, 0x52, 0xeb, 0xa5, 0xe9, 0xbc, 0xf3, 0xf6, 0xf7, 0x3c, 0x33, 0x4f, 0x67,
	0xd0, 0xc2, 0x3b, 0x49, 0x12, 0x60, 0xdc, 0xcb, 0x18, 0xd7, 0xde, 0xc1, 0x5a, 0x48, 0x35, 0xac,
	0x99, 0x41, 0xbb, 0x27, 0x85, 0x16, 0xb8, 0x59, 0x34, 0xb4, 0x4d, 0xad, 0x68, 0x98, 0x6f, 0xc6,
	0x22, 0x16, 0xa6, 0xc1, 0xcb, 0xdf, 0x6c, 0xef, 0xfc, 0x5d, 0x22, 0x54, 0x26, 0x54, 0x60, 0x27,
	0xec, 0xa0, 0x98, 0x9a, 0x85, 0x8c, 0x71, 0xe1, 0x99, 0xa7, 0x2d, 0x2d, 0x7d, 0x71, 0x50, 0x7d,
	0x93, 0x71, 0x4d, 0x25, 0xde, 0x42, 0x0d, 0xc6, 0xdf, 0xa4, 0xa0, 0x99, 0xe0, 0xae, 0xb3, 0xe8,
	0x2c, 0x37, 0x3a, 0x6b, 0x27, 0x67, 0x0b, 0xb5, 0xef, 0x67, 0x0b, 0xf7, 0x2c, 0x46, 0x45, 0xdd,
	0x36, 0x13, 0x5e, 0x06, 0x3a, 0x69, 0xbf, 0xa4, 0x31, 0x90, 0xc1, 0x06, 0x25, 0x5f, 0x3f, 0xaf,
	0xa2, 0x42, 0x65, 0x83, 0x12, 0xbf, 0x64, 0xe0, 0xd7, 0x68, 0x16, 0x38, 0xef, 0x43, 0x9a, 0x7b,
	0x39, 0x60, 0x8a, 0x09, 0xae, 0xdc, 0xff, 0xfe, 0x16, 0x7c, 0xd3, 0xb2, 0xb6, 0xc7, 0xa8, 0xa5,
	0x8f, 0x0d, 0x54, 0xdf, 0x06, 0x09, 0x99, 0xc2, 0xf7, 0x11, 0xca, 0xb7, 0x26, 0x88, 0x28, 0x17,
	0x99, 0x35, 0xef, 0x37, 0xf2, 0xca, 0x46, 0x5e, 0xc0, 0x6f, 0xd1, 0xdc, 0xd8, 0x56, 0x20, 0x41,
	0xd3, 0x80, 0x24, 0xc0, 0x63, 0x5a, 0xb8, 0x59, 0xff, 0x63, 0x37, 0x9f, 0x2e, 0x8e, 0x57, 0x1c,
	0xff, 0xd6, 0x18, 0xea, 0x83, 0xa6, 0xcf, 0x0d, 0x12, 0xbf, 0x42, 0x33, 0xa5, 0x56, 0x06, 0x47,
	0xee, 0xff, 0x95, 0x34, 0xa6, 0xc7, 0xb0, 0x4d, 0x38, 0xba, 0x02, 0x67, 0xdc, 0x9d, 0xf8, 0x57,
	0x70, 0xc6, 0xf1, 0x2e, 0x9a, 0x8a, 0x05, 0xa4, 0x41, 0x28, 0x78, 0x44, 0x23, 0xf7, 0x5a, 0x25,
	0x34, 0xca, 0x51, 0x1d, 0x43, 0xc2, 0x0f, 0xd0, 0x8d, 0x30, 0x15, 0xa4, 0xab, 0x82, 0x1e, 0x95,
	0xc1, 0x80, 0x82, 0x74, 0xeb, 0x8b, 0xce, 0xf2, 0x84, 0x3f, 0x63, 0xcb, 0xdb, 0x54, 0xee, 0x51,
	0x90, 0xf9, 0xea, 0x94, 0x86, 0x2e, 0xe3, 0x71, 0x30, 0x60, 0x34, 0x8d, 0xdc, 0xeb, 0xd5, 0x56,
	0x57, 0xc0, 0xf6, 0x72, 0x16, 0xde, 0x41, 0x8d, 0xb0, 0x2f, 0x6d, 0xfc, 0xee, 0x64, 0x25, 0xf0,
	0x64, 0x0e, 0xca, 0x23, 0xc7, 0x09, 0x6a, 0x9a, 0x73, 0x44, 0x44, 0x1a, 0x1c, 0x42, 0x9a, 0x52,
	0x6d, 0xf9, 0x8d, 0x4a, 0x7c, 0x3c, 0x62, 0xee, 0x1a, 0xa4, 0x51, 0x5a, 0x47, 0x77, 0xae, 0x2a,
	0x41, 0x14, 0x49, 0xaa, 0x94, 0x8b, 0xcc, 0xef, 0x3e, 0xf7, 0xeb, 0x47, 0xcf, 0xec, 0x24, 0xd6,
	0x68, 0x1e, 0xa2, 0x88, 0xe5, 0x19, 0x43, 0x1a, 0x8c, 0xb6, 0x57, 0xd2, 0x43, 0x90, 0x91, 0x72,
	0xa7, 0x2a, 0xf9, 0x74, 0x4b, 0xf2, 0x8e, 0x05, 0xfb, 0x96, 0x8b, 0x53, 0x74, 0xfb, 0x92, 0x6a,
	0xd6, 0x23, 0x63, 0xc5, 0xe9, 0x4a, 0x8a, 0xcd, 0x92, 0xba, 0xd9, 0x23, 0x23, 0xb5, 0x04, 0x5d,
	0xaa, 0x07, 0x65, 0xca, 0x33, 0xd5, 0x52, 0x28, 0x99, 0x9d, 0x22, 0xef, 0xa7, 0x8f, 0xde, 0x5f,
	0x1c, 0xaf, 0x3c, 0xdc, 0xa7, 0x5c, 0x0a, 0xd2, 0x5d, 0x7d, 0x21, 0xfa, 0x3c, 0x32, 0xc7, 0xc7,
	0x1b, 0xdd, 0xe0, 0x07, 0xeb, 0xde, 0x91, 0xbd, 0xc6, 0xed, 0xcd, 0xd4, 0xd9, 0x3a, 0x39, 0x6f,
	0x39, 0xa7, 0xe7, 0x2d, 0xe7, 0xc7, 0x79, 0xcb, 0xf9, 0x30, 0x6c, 0xd5, 0x4e, 0x87, 0xad, 0xda,
	0xb7, 0x61, 0xab, 0xb6, 0xff, 0x24, 0x66, 0x3a, 0xe9, 0x87, 0x6d, 0x22, 0x32, 0xef, 0xb7, 0x88,
	0x7a, 0xd0, 0xa3, 0x2a, 0xac, 0x9b, 0x9c, 0x1f, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x53, 0x1e,
	0x5d, 0xe7, 0x35, 0x06, 0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AnnualProvisions.Size()
		i -= size
		if _, err := m.AnnualProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AdditionalBurnRate.Size()
		i -= size
		if _, err := m.AdditionalBurnRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size := m.AdditionalMpcRewards.Size()
		i -= size
		if _, err := m.AdditionalMpcRewards.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.AdditionalStakingRewards.Size()
		i -= size
		if _, err := m.AdditionalStakingRewards.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	if len(m.ProtocolWalletAddress) > 0 {
		i -= len(m.ProtocolWalletAddress)
		copy(dAtA[i:], m.ProtocolWalletAddress)
		i = encodeVarintMint(dAtA, i, uint64(len(m.ProtocolWalletAddress)))
		i--
		dAtA[i] = 0x52
	}
	{
		size := m.ProtocolWalletRate.Size()
		i -= size
		if _, err := m.ProtocolWalletRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.BurnRate.Size()
		i -= size
		if _, err := m.BurnRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.StakingYield.Size()
		i -= size
		if _, err := m.StakingYield.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.BlocksPerYear != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.BlocksPerYear))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.GoalBonded.Size()
		i -= size
		if _, err := m.GoalBonded.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.InflationMin.Size()
		i -= size
		if _, err := m.InflationMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.InflationMax.Size()
		i -= size
		if _, err := m.InflationMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.InflationRateChange.Size()
		i -= size
		if _, err := m.InflationRateChange.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflation.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.AnnualProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = m.InflationRateChange.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.InflationMax.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.InflationMin.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.GoalBonded.Size()
	n += 1 + l + sovMint(uint64(l))
	if m.BlocksPerYear != 0 {
		n += 1 + sovMint(uint64(m.BlocksPerYear))
	}
	l = m.StakingYield.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.BurnRate.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.ProtocolWalletRate.Size()
	n += 1 + l + sovMint(uint64(l))
	l = len(m.ProtocolWalletAddress)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = m.AdditionalStakingRewards.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.AdditionalMpcRewards.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.AdditionalBurnRate.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnnualProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRateChange", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationRateChange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoalBonded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GoalBonded.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksPerYear", wireType)
			}
			m.BlocksPerYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksPerYear |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingYield", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StakingYield.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BurnRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolWalletRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ProtocolWalletRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolWalletAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProtocolWalletAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalStakingRewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AdditionalStakingRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalMpcRewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AdditionalMpcRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalBurnRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AdditionalBurnRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
