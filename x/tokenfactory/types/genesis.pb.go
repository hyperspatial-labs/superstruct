// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/tokenfactory/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the tokenfactory module's genesis state.
type GenesisState struct {
	// params defines the paramaters of the module.
	Params        Params         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	FactoryDenoms []GenesisDenom `protobuf:"bytes,2,rep,name=factory_denoms,json=factoryDenoms,proto3" json:"factory_denoms" yaml:"factory_denoms"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5749c3f71850298b, []int{0}
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

func (m *GenesisState) GetFactoryDenoms() []GenesisDenom {
	if m != nil {
		return m.FactoryDenoms
	}
	return nil
}

// GenesisDenom defines a tokenfactory denom that is defined within genesis
// state. The structure contains DenomAuthorityMetadata which defines the
// denom's admin.
type GenesisDenom struct {
	Denom             string                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
	AuthorityMetadata DenomAuthorityMetadata `protobuf:"bytes,2,opt,name=authority_metadata,json=authorityMetadata,proto3" json:"authority_metadata" yaml:"authority_metadata"`
}

func (m *GenesisDenom) Reset()         { *m = GenesisDenom{} }
func (m *GenesisDenom) String() string { return proto.CompactTextString(m) }
func (*GenesisDenom) ProtoMessage()    {}
func (*GenesisDenom) Descriptor() ([]byte, []int) {
	return fileDescriptor_5749c3f71850298b, []int{1}
}
func (m *GenesisDenom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisDenom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisDenom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisDenom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisDenom.Merge(m, src)
}
func (m *GenesisDenom) XXX_Size() int {
	return m.Size()
}
func (m *GenesisDenom) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisDenom.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisDenom proto.InternalMessageInfo

func (m *GenesisDenom) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *GenesisDenom) GetAuthorityMetadata() DenomAuthorityMetadata {
	if m != nil {
		return m.AuthorityMetadata
	}
	return DenomAuthorityMetadata{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "osmosis.tokenfactory.v1beta1.GenesisState")
	proto.RegisterType((*GenesisDenom)(nil), "osmosis.tokenfactory.v1beta1.GenesisDenom")
}

func init() {
	proto.RegisterFile("osmosis/tokenfactory/v1beta1/genesis.proto", fileDescriptor_5749c3f71850298b)
}

var fileDescriptor_5749c3f71850298b = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x33, 0xfd, 0xfb, 0x17, 0x4c, 0xab, 0x68, 0x50, 0xa8, 0x45, 0x93, 0x1a, 0x44, 0x6a,
	0xc1, 0x84, 0xd6, 0xae, 0x0a, 0x2e, 0x0c, 0x05, 0x57, 0x82, 0xc4, 0x85, 0xe0, 0xa6, 0x4c, 0xda,
	0x31, 0x0d, 0x36, 0x9d, 0x90, 0xb9, 0x11, 0xf3, 0x02, 0xae, 0x7d, 0x04, 0x1f, 0x46, 0xa1, 0xcb,
	0x2e, 0x5d, 0x15, 0x69, 0x37, 0xae, 0xfb, 0x04, 0x92, 0xc9, 0xa0, 0xd6, 0x42, 0x76, 0xc9, 0x9d,
	0xef, 0x9c, 0x39, 0x77, 0x8e, 0x5c, 0xa7, 0xcc, 0xa7, 0xcc, 0x63, 0x26, 0xd0, 0x7b, 0x32, 0xba,
	0xc3, 0x3d, 0xa0, 0x61, 0x6c, 0x3e, 0x34, 0x1c, 0x02, 0xb8, 0x61, 0xba, 0x64, 0x44, 0x98, 0xc7,
	0x8c, 0x20, 0xa4, 0x40, 0x95, 0x3d, 0xc1, 0x1a, 0xbf, 0x59, 0x43, 0xb0, 0x95, 0x6d, 0x97, 0xba,
	0x94, 0x83, 0x66, 0xf2, 0x95, 0x6a, 0x2a, 0xad, 0x4c, 0x7f, 0x1c, 0xc1, 0x80, 0x86, 0x1e, 0xc4,
	0x97, 0x04, 0x70, 0x1f, 0x03, 0x16, 0xaa, 0xe3, 0x4c, 0x55, 0x80, 0x43, 0xec, 0x8b, 0x50, 0xfa,
	0x2b, 0x92, 0x4b, 0x17, 0x69, 0xcc, 0x6b, 0xc0, 0x40, 0x14, 0x4b, 0x2e, 0xa4, 0x40, 0x19, 0x55,
	0x51, 0xad, 0xd8, 0x3c, 0x34, 0xb2, 0x62, 0x1b, 0x57, 0x9c, 0xb5, 0xf2, 0xe3, 0xa9, 0x26, 0xd9,
	0x42, 0xa9, 0x04, 0xf2, 0x86, 0xe0, 0xba, 0x7d, 0x32, 0xa2, 0x3e, 0x2b, 0xe7, 0xaa, 0xff, 0x6a,
	0xc5, 0x66, 0x3d, 0xdb, 0x4b, 0xe4, 0xe8, 0x24, 0x12, 0x6b, 0x3f, 0x71, 0x5c, 0x4c, 0xb5, 0x9d,
	0x18, 0xfb, 0xc3, 0xb6, 0xbe, 0xec, 0xa7, 0xdb, 0xeb, 0x62, 0xd0, 0x49, 0xff, 0xdf, 0x7e, 0xd6,
	0xe0, 0x13, 0xe5, 0x48, 0xfe, 0xcf, 0x51, 0xbe, 0xc5, 0x9a, 0xb5, 0xb9, 0x98, 0x6a, 0xa5, 0xd4,
	0x89, 0x8f, 0x75, 0x3b, 0x3d, 0x56, 0x9e, 0x90, 0xac, 0x7c, 0x3f, 0x63, 0xd7, 0x17, 0xef, 0x58,
	0xce, 0xf1, 0xdd, 0x5b, 0xd9, 0x79, 0xf9, 0x4d, 0xe7, 0x7f, 0x3b, 0xb0, 0x0e, 0x44, 0xf2, 0xdd,
	0xf4, 0xbe, 0x55, 0x77, 0xdd, 0xde, 0x5a, 0x69, 0xae, 0x9d, 0xff, 0x7c, 0xd1, 0x90, 0x75, 0x33,
	0x9e, 0xa9, 0x68, 0x32, 0x53, 0xd1, 0xc7, 0x4c, 0x45, 0xcf, 0x73, 0x55, 0x9a, 0xcc, 0x55, 0xe9,
	0x7d, 0xae, 0x4a, 0xb7, 0x67, 0xae, 0x07, 0x83, 0xc8, 0x31, 0x7a, 0xd4, 0x37, 0x07, 0x71, 0x40,
	0x42, 0x16, 0x60, 0xf0, 0xf0, 0xf0, 0x64, 0x88, 0x1d, 0x66, 0xb2, 0x28, 0x99, 0x40, 0x18, 0xf5,
	0xc0, 0x7c, 0x5c, 0xae, 0x1d, 0xe2, 0x80, 0x30, 0xa7, 0xc0, 0xeb, 0x3e, 0xfd, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0x72, 0xed, 0xac, 0xd3, 0xb1, 0x02, 0x00, 0x00,
}

func (this *GenesisDenom) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenesisDenom)
	if !ok {
		that2, ok := that.(GenesisDenom)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	if !this.AuthorityMetadata.Equal(&that1.AuthorityMetadata) {
		return false
	}
	return true
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
	if len(m.FactoryDenoms) > 0 {
		for iNdEx := len(m.FactoryDenoms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FactoryDenoms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
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

func (m *GenesisDenom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisDenom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisDenom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.AuthorityMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Denom)))
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
	if len(m.FactoryDenoms) > 0 {
		for _, e := range m.FactoryDenoms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisDenom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.AuthorityMetadata.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field FactoryDenoms", wireType)
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
			m.FactoryDenoms = append(m.FactoryDenoms, GenesisDenom{})
			if err := m.FactoryDenoms[len(m.FactoryDenoms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *GenesisDenom) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisDenom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisDenom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorityMetadata", wireType)
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
			if err := m.AuthorityMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
