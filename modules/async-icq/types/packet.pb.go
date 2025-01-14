// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: icq/v1/packet.proto

package types

import (
	fmt "fmt"
	types "github.com/cometbft/cometbft/abci/types"
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

// InterchainQueryPacketData is comprised of raw query.
type InterchainQueryPacketData struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// optional memo
	Memo string `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (m *InterchainQueryPacketData) Reset()         { *m = InterchainQueryPacketData{} }
func (m *InterchainQueryPacketData) String() string { return proto.CompactTextString(m) }
func (*InterchainQueryPacketData) ProtoMessage()    {}
func (*InterchainQueryPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_13b1ec1d226ce757, []int{0}
}
func (m *InterchainQueryPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterchainQueryPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterchainQueryPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterchainQueryPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterchainQueryPacketData.Merge(m, src)
}
func (m *InterchainQueryPacketData) XXX_Size() int {
	return m.Size()
}
func (m *InterchainQueryPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_InterchainQueryPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_InterchainQueryPacketData proto.InternalMessageInfo

func (m *InterchainQueryPacketData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InterchainQueryPacketData) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

// InterchainQueryPacketAck is comprised of an ABCI query response with non-deterministic fields left empty (e.g. Codespace, Log, Info and ...).
type InterchainQueryPacketAck struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *InterchainQueryPacketAck) Reset()         { *m = InterchainQueryPacketAck{} }
func (m *InterchainQueryPacketAck) String() string { return proto.CompactTextString(m) }
func (*InterchainQueryPacketAck) ProtoMessage()    {}
func (*InterchainQueryPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_13b1ec1d226ce757, []int{1}
}
func (m *InterchainQueryPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterchainQueryPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterchainQueryPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterchainQueryPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterchainQueryPacketAck.Merge(m, src)
}
func (m *InterchainQueryPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *InterchainQueryPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_InterchainQueryPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_InterchainQueryPacketAck proto.InternalMessageInfo

func (m *InterchainQueryPacketAck) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// CosmosQuery contains a list of tendermint ABCI query requests. It should be used when sending queries to an SDK host chain.
type CosmosQuery struct {
	Requests []types.RequestQuery `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests"`
}

func (m *CosmosQuery) Reset()         { *m = CosmosQuery{} }
func (m *CosmosQuery) String() string { return proto.CompactTextString(m) }
func (*CosmosQuery) ProtoMessage()    {}
func (*CosmosQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_13b1ec1d226ce757, []int{2}
}
func (m *CosmosQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CosmosQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CosmosQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CosmosQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CosmosQuery.Merge(m, src)
}
func (m *CosmosQuery) XXX_Size() int {
	return m.Size()
}
func (m *CosmosQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_CosmosQuery.DiscardUnknown(m)
}

var xxx_messageInfo_CosmosQuery proto.InternalMessageInfo

func (m *CosmosQuery) GetRequests() []types.RequestQuery {
	if m != nil {
		return m.Requests
	}
	return nil
}

// CosmosResponse contains a list of tendermint ABCI query responses. It should be used when receiving responses from an SDK host chain.
type CosmosResponse struct {
	Responses []types.ResponseQuery `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses"`
}

func (m *CosmosResponse) Reset()         { *m = CosmosResponse{} }
func (m *CosmosResponse) String() string { return proto.CompactTextString(m) }
func (*CosmosResponse) ProtoMessage()    {}
func (*CosmosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_13b1ec1d226ce757, []int{3}
}
func (m *CosmosResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CosmosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CosmosResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CosmosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CosmosResponse.Merge(m, src)
}
func (m *CosmosResponse) XXX_Size() int {
	return m.Size()
}
func (m *CosmosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CosmosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CosmosResponse proto.InternalMessageInfo

func (m *CosmosResponse) GetResponses() []types.ResponseQuery {
	if m != nil {
		return m.Responses
	}
	return nil
}

func init() {
	proto.RegisterType((*InterchainQueryPacketData)(nil), "icq.v1.InterchainQueryPacketData")
	proto.RegisterType((*InterchainQueryPacketAck)(nil), "icq.v1.InterchainQueryPacketAck")
	proto.RegisterType((*CosmosQuery)(nil), "icq.v1.CosmosQuery")
	proto.RegisterType((*CosmosResponse)(nil), "icq.v1.CosmosResponse")
}

func init() { proto.RegisterFile("icq/v1/packet.proto", fileDescriptor_13b1ec1d226ce757) }

var fileDescriptor_13b1ec1d226ce757 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4e, 0x32, 0x31,
	0x14, 0xc5, 0xa7, 0xdf, 0x47, 0x88, 0x14, 0xe3, 0x62, 0x74, 0x31, 0x62, 0xac, 0x64, 0x56, 0x6c,
	0x68, 0x83, 0xc6, 0xb8, 0x34, 0x82, 0x1b, 0x37, 0xfe, 0x99, 0xb8, 0x72, 0xd7, 0x29, 0x37, 0xd0,
	0xe0, 0x4c, 0x4b, 0xdb, 0x21, 0xe1, 0x2d, 0x7c, 0x2c, 0x96, 0x2c, 0x5d, 0x19, 0x03, 0x2f, 0x62,
	0x68, 0x55, 0x5c, 0xb0, 0x3b, 0xb9, 0xf7, 0x9c, 0x5f, 0x6f, 0x4e, 0xf1, 0xa1, 0x14, 0x53, 0x36,
	0xeb, 0x31, 0xcd, 0xc5, 0x04, 0x1c, 0xd5, 0x46, 0x39, 0x15, 0xd7, 0xa5, 0x98, 0xd2, 0x59, 0xaf,
	0x75, 0x34, 0x52, 0x23, 0xe5, 0x47, 0x6c, 0xa3, 0xc2, 0xb6, 0x75, 0xe2, 0xa0, 0x1c, 0x82, 0x29,
	0x64, 0xe9, 0x18, 0xcf, 0x85, 0x64, 0x6e, 0xae, 0xc1, 0x86, 0x65, 0x3a, 0xc0, 0xc7, 0x77, 0xa5,
	0x03, 0x23, 0xc6, 0x5c, 0x96, 0x4f, 0x15, 0x98, 0xf9, 0xa3, 0x27, 0xdf, 0x72, 0xc7, 0xe3, 0x18,
	0xd7, 0x86, 0xdc, 0xf1, 0x04, 0xb5, 0x51, 0x67, 0x3f, 0xf3, 0x7a, 0x33, 0x2b, 0xa0, 0x50, 0xc9,
	0xbf, 0x36, 0xea, 0x34, 0x32, 0xaf, 0x53, 0x8a, 0x93, 0x9d, 0x90, 0x1b, 0x31, 0xd9, 0xc5, 0x48,
	0xef, 0x71, 0x73, 0xa0, 0x6c, 0xa1, 0xac, 0xf7, 0xc6, 0xd7, 0x78, 0xcf, 0xc0, 0xb4, 0x02, 0xeb,
	0x6c, 0x82, 0xda, 0xff, 0x3b, 0xcd, 0xf3, 0x53, 0xba, 0xbd, 0x99, 0x6e, 0x6e, 0xa6, 0x59, 0x30,
	0xf8, 0x40, 0xbf, 0xb6, 0xf8, 0x38, 0x8b, 0xb2, 0xdf, 0x50, 0xfa, 0x8c, 0x0f, 0x02, 0x2f, 0x03,
	0xab, 0x55, 0x69, 0x21, 0xee, 0xe3, 0x86, 0xf9, 0xd6, 0x3f, 0x4c, 0xb2, 0x83, 0x19, 0x1c, 0x7f,
	0xa1, 0xdb, 0x58, 0xff, 0x61, 0xb1, 0x22, 0x68, 0xb9, 0x22, 0xe8, 0x73, 0x45, 0xd0, 0xdb, 0x9a,
	0x44, 0xcb, 0x35, 0x89, 0xde, 0xd7, 0x24, 0x7a, 0xb9, 0x1c, 0x49, 0x37, 0xae, 0x72, 0x2a, 0x54,
	0xc1, 0x84, 0x7f, 0x98, 0xc9, 0x5c, 0x74, 0xb9, 0xd6, 0x96, 0x15, 0x6a, 0x58, 0xbd, 0x82, 0x65,
	0xdc, 0xce, 0x4b, 0xd1, 0xf5, 0xbf, 0x75, 0x15, 0x1a, 0xcf, 0xeb, 0xbe, 0xf2, 0x8b, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x02, 0x9a, 0x60, 0x5c, 0xc4, 0x01, 0x00, 0x00,
}

func (m *InterchainQueryPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterchainQueryPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterchainQueryPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InterchainQueryPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterchainQueryPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterchainQueryPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CosmosQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CosmosQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CosmosQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for iNdEx := len(m.Requests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPacket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CosmosResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CosmosResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CosmosResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Responses) > 0 {
		for iNdEx := len(m.Responses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Responses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPacket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InterchainQueryPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *InterchainQueryPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *CosmosQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for _, e := range m.Requests {
			l = e.Size()
			n += 1 + l + sovPacket(uint64(l))
		}
	}
	return n
}

func (m *CosmosResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Responses) > 0 {
		for _, e := range m.Responses {
			l = e.Size()
			n += 1 + l + sovPacket(uint64(l))
		}
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InterchainQueryPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: InterchainQueryPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterchainQueryPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *InterchainQueryPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: InterchainQueryPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterchainQueryPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *CosmosQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: CosmosQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CosmosQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requests = append(m.Requests, types.RequestQuery{})
			if err := m.Requests[len(m.Requests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *CosmosResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: CosmosResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CosmosResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Responses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Responses = append(m.Responses, types.ResponseQuery{})
			if err := m.Responses[len(m.Responses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
