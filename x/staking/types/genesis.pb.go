// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/staking/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines all staking state that must be provided at genesis.
type GenesisState struct {
	Params               Params                                 `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	LastTotalPower       github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=last_total_power,json=lastTotalPower,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"last_total_power" yaml:"last_total_power"`
	LastValidatorPowers  []LastValidatorPower                   `protobuf:"bytes,3,rep,name=last_validator_powers,json=lastValidatorPowers,proto3" json:"last_validator_powers" yaml:"last_validator_powers"`
	Validators           []Validator                            `protobuf:"bytes,4,rep,name=validators,proto3" json:"validators"`
	Delegations          []Delegation                           `protobuf:"bytes,5,rep,name=delegations,proto3" json:"delegations"`
	UnbondingDelegations []UnbondingDelegation                  `protobuf:"bytes,6,rep,name=unbonding_delegations,json=unbondingDelegations,proto3" json:"unbonding_delegations" yaml:"unbonding_delegations"`
	Redelegations        []Redelegation                         `protobuf:"bytes,7,rep,name=redelegations,proto3" json:"redelegations"`
	Exported             bool                                   `protobuf:"varint,8,opt,name=exported,proto3" json:"exported,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_10c2ca02caf42801, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetLastValidatorPowers() []LastValidatorPower {
	if m != nil {
		return m.LastValidatorPowers
	}
	return nil
}

func (m *GenesisState) GetValidators() []Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *GenesisState) GetDelegations() []Delegation {
	if m != nil {
		return m.Delegations
	}
	return nil
}

func (m *GenesisState) GetUnbondingDelegations() []UnbondingDelegation {
	if m != nil {
		return m.UnbondingDelegations
	}
	return nil
}

func (m *GenesisState) GetRedelegations() []Redelegation {
	if m != nil {
		return m.Redelegations
	}
	return nil
}

func (m *GenesisState) GetExported() bool {
	if m != nil {
		return m.Exported
	}
	return false
}

// LastValidatorPower required for validator set update logic.
type LastValidatorPower struct {
	Address github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"address,omitempty"`
	Power   int64                                         `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
}

func (m *LastValidatorPower) Reset()         { *m = LastValidatorPower{} }
func (m *LastValidatorPower) String() string { return proto.CompactTextString(m) }
func (*LastValidatorPower) ProtoMessage()    {}
func (*LastValidatorPower) Descriptor() ([]byte, []int) {
	return fileDescriptor_10c2ca02caf42801, []int{1}
}
func (m *LastValidatorPower) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LastValidatorPower) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LastValidatorPower.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LastValidatorPower) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastValidatorPower.Merge(m, src)
}
func (m *LastValidatorPower) XXX_Size() int {
	return m.Size()
}
func (m *LastValidatorPower) XXX_DiscardUnknown() {
	xxx_messageInfo_LastValidatorPower.DiscardUnknown(m)
}

var xxx_messageInfo_LastValidatorPower proto.InternalMessageInfo

func (m *LastValidatorPower) GetAddress() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *LastValidatorPower) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmos.staking.GenesisState")
	proto.RegisterType((*LastValidatorPower)(nil), "cosmos.staking.LastValidatorPower")
}

func init() { proto.RegisterFile("cosmos/staking/genesis.proto", fileDescriptor_10c2ca02caf42801) }

var fileDescriptor_10c2ca02caf42801 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x1b, 0xba, 0x76, 0x95, 0x5b, 0x26, 0x64, 0x3a, 0x08, 0x55, 0x95, 0x54, 0x01, 0xa1,
	0x1e, 0x58, 0x22, 0x06, 0x27, 0x2e, 0x88, 0x08, 0x01, 0x13, 0x1c, 0x26, 0x03, 0x3b, 0x70, 0xa9,
	0xdc, 0xc5, 0x0a, 0xd1, 0xd2, 0x38, 0xca, 0xe7, 0xb2, 0x4d, 0x42, 0x3c, 0x03, 0xcf, 0xc0, 0xd3,
	0xec, 0xb8, 0x23, 0xe2, 0x10, 0xa1, 0xf6, 0x0d, 0x76, 0xe4, 0x84, 0x62, 0xa7, 0x21, 0x75, 0x11,
	0xe2, 0xe4, 0xda, 0xdf, 0xff, 0xff, 0xfb, 0x7f, 0x75, 0x3e, 0xa3, 0xe1, 0x31, 0x87, 0x19, 0x07,
	0x0f, 0x04, 0x3d, 0x89, 0x92, 0xd0, 0x0b, 0x59, 0xc2, 0x20, 0x02, 0x37, 0xcd, 0xb8, 0xe0, 0x78,
	0x47, 0x55, 0xdd, 0xb2, 0x3a, 0xe8, 0x87, 0x3c, 0xe4, 0xb2, 0xe4, 0x15, 0xbf, 0x94, 0x6a, 0xa0,
	0x33, 0xca, 0x55, 0x55, 0x9d, 0x6f, 0x2d, 0xd4, 0x7b, 0xa9, 0xa8, 0x6f, 0x05, 0x15, 0x0c, 0x3f,
	0x46, 0xed, 0x94, 0x66, 0x74, 0x06, 0xa6, 0x31, 0x32, 0xc6, 0xdd, 0xfd, 0x5b, 0xee, 0x7a, 0x8a,
	0x7b, 0x28, 0xab, 0xfe, 0xd6, 0x45, 0x6e, 0x37, 0x48, 0xa9, 0xc5, 0x80, 0x6e, 0xc4, 0x14, 0xc4,
	0x44, 0x70, 0x41, 0xe3, 0x49, 0xca, 0x4f, 0x59, 0x66, 0x5e, 0x1b, 0x19, 0xe3, 0x9e, 0x7f, 0x50,
	0xe8, 0x7e, 0xe4, 0xf6, 0xfd, 0x30, 0x12, 0x1f, 0xe7, 0x53, 0xf7, 0x98, 0xcf, 0xbc, 0xb2, 0x23,
	0xb5, 0xec, 0x41, 0x70, 0xe2, 0x89, 0xf3, 0x94, 0x81, 0x7b, 0x90, 0x88, 0xab, 0xdc, 0xbe, 0x7d,
	0x4e, 0x67, 0xf1, 0x13, 0x47, 0xe7, 0x39, 0x64, 0xa7, 0x38, 0x7a, 0x57, 0x9c, 0x1c, 0x16, 0x07,
	0xf8, 0x33, 0xda, 0x95, 0xa2, 0x4f, 0x34, 0x8e, 0x02, 0x2a, 0x78, 0xa6, 0x84, 0x60, 0x36, 0x47,
	0xcd, 0x71, 0x77, 0xdf, 0xd1, 0x3b, 0x7f, 0x43, 0x41, 0x1c, 0xad, 0xb4, 0x12, 0xe1, 0xdf, 0x2b,
	0xba, 0xbb, 0xca, 0xed, 0x61, 0x2d, 0x53, 0xc7, 0x39, 0xe4, 0x66, 0xbc, 0xe1, 0x04, 0xfc, 0x14,
	0xa1, 0x4a, 0x09, 0xe6, 0x96, 0x8c, 0xbc, 0xa3, 0x47, 0x56, 0xa6, 0xf2, 0xbe, 0x6a, 0x16, 0xec,
	0xa3, 0x6e, 0xc0, 0x62, 0x16, 0x52, 0x11, 0xf1, 0x04, 0xcc, 0x96, 0x24, 0x0c, 0x74, 0xc2, 0xf3,
	0x4a, 0x52, 0x22, 0xea, 0x26, 0xfc, 0x05, 0xed, 0xce, 0x93, 0x29, 0x4f, 0x82, 0x28, 0x09, 0x27,
	0x75, 0x5a, 0x5b, 0xd2, 0xee, 0xea, 0xb4, 0xf7, 0x2b, 0x71, 0x0d, 0xab, 0xdd, 0xc1, 0x5f, 0x79,
	0x0e, 0xe9, 0xcf, 0x37, 0xad, 0x80, 0x5f, 0xa1, 0xeb, 0x19, 0xab, 0xe7, 0x6e, 0xcb, 0xdc, 0xa1,
	0x9e, 0x4b, 0x6a, 0xa2, 0xf2, 0x7f, 0xac, 0x1b, 0xf1, 0x00, 0x75, 0xd8, 0x59, 0xca, 0x33, 0xc1,
	0x02, 0xb3, 0x33, 0x32, 0xc6, 0x1d, 0x52, 0xed, 0x9d, 0x53, 0x84, 0x37, 0xbf, 0x1d, 0x7e, 0x8d,
	0xb6, 0x69, 0x10, 0x64, 0x0c, 0xd4, 0xa8, 0xf6, 0xfc, 0x87, 0xbf, 0x72, 0x7b, 0xef, 0x3f, 0xc6,
	0xec, 0x88, 0xc6, 0xcf, 0x94, 0x91, 0xac, 0x08, 0xb8, 0x8f, 0x5a, 0x7f, 0xa6, 0xb6, 0x49, 0xd4,
	0xc6, 0x7f, 0x71, 0xb1, 0xb0, 0x8c, 0xcb, 0x85, 0x65, 0xfc, 0x5c, 0x58, 0xc6, 0xd7, 0xa5, 0xd5,
	0xb8, 0x5c, 0x5a, 0x8d, 0xef, 0x4b, 0xab, 0xf1, 0xe1, 0xc1, 0x3f, 0x73, 0xce, 0xaa, 0xd7, 0x26,
	0x13, 0xa7, 0x6d, 0xf9, 0xd8, 0x1e, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x07, 0xac, 0x0d, 0x57,
	0xd0, 0x03, 0x00, 0x00,
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
	if m.Exported {
		i--
		if m.Exported {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if len(m.Redelegations) > 0 {
		for iNdEx := len(m.Redelegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Redelegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.UnbondingDelegations) > 0 {
		for iNdEx := len(m.UnbondingDelegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UnbondingDelegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Delegations) > 0 {
		for iNdEx := len(m.Delegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Delegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.LastValidatorPowers) > 0 {
		for iNdEx := len(m.LastValidatorPowers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LastValidatorPowers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size := m.LastTotalPower.Size()
		i -= size
		if _, err := m.LastTotalPower.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *LastValidatorPower) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastValidatorPower) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LastValidatorPower) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Power != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.LastTotalPower.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.LastValidatorPowers) > 0 {
		for _, e := range m.LastValidatorPowers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Delegations) > 0 {
		for _, e := range m.Delegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UnbondingDelegations) > 0 {
		for _, e := range m.UnbondingDelegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Redelegations) > 0 {
		for _, e := range m.Redelegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Exported {
		n += 2
	}
	return n
}

func (m *LastValidatorPower) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovGenesis(uint64(m.Power))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTotalPower", wireType)
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
			if err := m.LastTotalPower.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastValidatorPowers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastValidatorPowers = append(m.LastValidatorPowers, LastValidatorPower{})
			if err := m.LastValidatorPowers[len(m.LastValidatorPowers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegations = append(m.Delegations, Delegation{})
			if err := m.Delegations[len(m.Delegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingDelegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnbondingDelegations = append(m.UnbondingDelegations, UnbondingDelegation{})
			if err := m.UnbondingDelegations[len(m.UnbondingDelegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Redelegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Redelegations = append(m.Redelegations, Redelegation{})
			if err := m.Redelegations[len(m.Redelegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exported", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.Exported = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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
func (m *LastValidatorPower) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: LastValidatorPower: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastValidatorPower: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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
