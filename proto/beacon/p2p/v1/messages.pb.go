// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/beacon/p2p/v1/messages.proto

package ethereum_beacon_p2p_v1

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Status struct {
	HeadForkVersion      []byte   `protobuf:"bytes,1,opt,name=head_fork_version,json=headForkVersion,proto3" json:"head_fork_version,omitempty" ssz-size:"4"`
	FinalizedRoot        []byte   `protobuf:"bytes,2,opt,name=finalized_root,json=finalizedRoot,proto3" json:"finalized_root,omitempty" ssz-size:"32"`
	FinalizedEpoch       uint64   `protobuf:"varint,3,opt,name=finalized_epoch,json=finalizedEpoch,proto3" json:"finalized_epoch,omitempty"`
	HeadRoot             []byte   `protobuf:"bytes,4,opt,name=head_root,json=headRoot,proto3" json:"head_root,omitempty" ssz-size:"32"`
	HeadSlot             uint64   `protobuf:"varint,5,opt,name=head_slot,json=headSlot,proto3" json:"head_slot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1d590cda035b632, []int{0}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Status.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return m.Size()
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetHeadForkVersion() []byte {
	if m != nil {
		return m.HeadForkVersion
	}
	return nil
}

func (m *Status) GetFinalizedRoot() []byte {
	if m != nil {
		return m.FinalizedRoot
	}
	return nil
}

func (m *Status) GetFinalizedEpoch() uint64 {
	if m != nil {
		return m.FinalizedEpoch
	}
	return 0
}

func (m *Status) GetHeadRoot() []byte {
	if m != nil {
		return m.HeadRoot
	}
	return nil
}

func (m *Status) GetHeadSlot() uint64 {
	if m != nil {
		return m.HeadSlot
	}
	return 0
}

type BeaconBlocksByRangeRequest struct {
	StartSlot            uint64   `protobuf:"varint,1,opt,name=start_slot,json=startSlot,proto3" json:"start_slot,omitempty"`
	Count                uint64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Step                 uint64   `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeaconBlocksByRangeRequest) Reset()         { *m = BeaconBlocksByRangeRequest{} }
func (m *BeaconBlocksByRangeRequest) String() string { return proto.CompactTextString(m) }
func (*BeaconBlocksByRangeRequest) ProtoMessage()    {}
func (*BeaconBlocksByRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1d590cda035b632, []int{1}
}
func (m *BeaconBlocksByRangeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BeaconBlocksByRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BeaconBlocksByRangeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BeaconBlocksByRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconBlocksByRangeRequest.Merge(m, src)
}
func (m *BeaconBlocksByRangeRequest) XXX_Size() int {
	return m.Size()
}
func (m *BeaconBlocksByRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconBlocksByRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconBlocksByRangeRequest proto.InternalMessageInfo

func (m *BeaconBlocksByRangeRequest) GetStartSlot() uint64 {
	if m != nil {
		return m.StartSlot
	}
	return 0
}

func (m *BeaconBlocksByRangeRequest) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *BeaconBlocksByRangeRequest) GetStep() uint64 {
	if m != nil {
		return m.Step
	}
	return 0
}

func init() {
	proto.RegisterType((*Status)(nil), "ethereum.beacon.p2p.v1.Status")
	proto.RegisterType((*BeaconBlocksByRangeRequest)(nil), "ethereum.beacon.p2p.v1.BeaconBlocksByRangeRequest")
}

func init() {
	proto.RegisterFile("proto/beacon/p2p/v1/messages.proto", fileDescriptor_a1d590cda035b632)
}

var fileDescriptor_a1d590cda035b632 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0xb3, 0x5a, 0x88, 0x6c, 0x40, 0x64, 0x63, 0x4c, 0x83, 0x11, 0x48, 0x2f, 0x72, 0xa1,
	0x0d, 0xe0, 0xc1, 0x18, 0x4f, 0x4d, 0xf4, 0x01, 0x4a, 0xe2, 0x95, 0xb4, 0x65, 0x68, 0x1b, 0x4a,
	0xa7, 0x76, 0xb7, 0x24, 0xf2, 0x84, 0x1e, 0x7d, 0x02, 0x62, 0x78, 0x04, 0x0e, 0x9e, 0x4d, 0x67,
	0x89, 0x9c, 0xbc, 0xed, 0xcc, 0x7c, 0xff, 0x97, 0xd9, 0xe1, 0x56, 0x5e, 0xa0, 0x42, 0x27, 0x00,
	0x3f, 0xc4, 0xcc, 0xc9, 0x27, 0xb9, 0xb3, 0x19, 0x3b, 0x6b, 0x90, 0xd2, 0x8f, 0x40, 0xda, 0x34,
	0x14, 0x37, 0xa0, 0x62, 0x28, 0xa0, 0x5c, 0xdb, 0x1a, 0xb3, 0xf3, 0x49, 0x6e, 0x6f, 0xc6, 0xdd,
	0x51, 0x94, 0xa8, 0xb8, 0x0c, 0xec, 0x10, 0xd7, 0x4e, 0x84, 0x11, 0x3a, 0x84, 0x07, 0xe5, 0x92,
	0x2a, 0x2d, 0xae, 0x5e, 0x5a, 0x63, 0xfd, 0x30, 0x5e, 0x9f, 0x29, 0x5f, 0x95, 0x52, 0x3c, 0xf3,
	0x4e, 0x0c, 0xfe, 0x62, 0xbe, 0xc4, 0x62, 0x35, 0xdf, 0x40, 0x21, 0x13, 0xcc, 0x4c, 0x36, 0x60,
	0xc3, 0xa6, 0x7b, 0x75, 0xd8, 0xf5, 0x9b, 0x52, 0x6e, 0x47, 0x32, 0xd9, 0xc2, 0x93, 0xf5, 0x60,
	0x79, 0xed, 0x0a, 0x7d, 0xc5, 0x62, 0xf5, 0xa6, 0x41, 0xf1, 0xc8, 0x2f, 0x97, 0x49, 0xe6, 0xa7,
	0xc9, 0x16, 0x16, 0xf3, 0x02, 0x51, 0x99, 0x67, 0x14, 0xed, 0x1c, 0x76, 0xfd, 0xd6, 0x29, 0x3a,
	0x9d, 0x58, 0x5e, 0xeb, 0x0f, 0xf4, 0x10, 0x95, 0xb8, 0xe7, 0xed, 0x53, 0x12, 0x72, 0x0c, 0x63,
	0xf3, 0x7c, 0xc0, 0x86, 0x86, 0x77, 0x12, 0xbe, 0x54, 0x5d, 0x61, 0xf3, 0x06, 0x2d, 0x48, 0x76,
	0xe3, 0x3f, 0xfb, 0x45, 0xc5, 0x90, 0xf8, 0xf6, 0xc8, 0xcb, 0x14, 0x95, 0x59, 0x23, 0x25, 0x0d,
	0x67, 0x29, 0x2a, 0x0b, 0x78, 0xd7, 0xa5, 0xc3, 0xb9, 0x29, 0x86, 0x2b, 0xe9, 0x7e, 0x78, 0x7e,
	0x16, 0x81, 0x07, 0xef, 0x25, 0x48, 0x25, 0xee, 0x38, 0x97, 0xca, 0x2f, 0x94, 0xce, 0x32, 0xca,
	0x36, 0xa8, 0x53, 0x85, 0xc5, 0x35, 0xaf, 0x85, 0x58, 0x66, 0xfa, 0x8f, 0x86, 0xa7, 0x0b, 0x21,
	0xb8, 0x21, 0x15, 0xe4, 0xc7, 0xed, 0xe9, 0xed, 0x36, 0x3f, 0xf7, 0x3d, 0xf6, 0xb5, 0xef, 0xb1,
	0xef, 0x7d, 0x8f, 0x05, 0x75, 0x3a, 0xfa, 0xf4, 0x37, 0x00, 0x00, 0xff, 0xff, 0x60, 0xbe, 0x44,
	0x15, 0xe1, 0x01, 0x00, 0x00,
}

func (m *Status) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Status) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Status) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.HeadSlot != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.HeadSlot))
		i--
		dAtA[i] = 0x28
	}
	if len(m.HeadRoot) > 0 {
		i -= len(m.HeadRoot)
		copy(dAtA[i:], m.HeadRoot)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.HeadRoot)))
		i--
		dAtA[i] = 0x22
	}
	if m.FinalizedEpoch != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.FinalizedEpoch))
		i--
		dAtA[i] = 0x18
	}
	if len(m.FinalizedRoot) > 0 {
		i -= len(m.FinalizedRoot)
		copy(dAtA[i:], m.FinalizedRoot)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.FinalizedRoot)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.HeadForkVersion) > 0 {
		i -= len(m.HeadForkVersion)
		copy(dAtA[i:], m.HeadForkVersion)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.HeadForkVersion)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BeaconBlocksByRangeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BeaconBlocksByRangeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BeaconBlocksByRangeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Step != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.Step))
		i--
		dAtA[i] = 0x18
	}
	if m.Count != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	if m.StartSlot != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.StartSlot))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessages(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessages(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Status) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HeadForkVersion)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.FinalizedRoot)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.FinalizedEpoch != 0 {
		n += 1 + sovMessages(uint64(m.FinalizedEpoch))
	}
	l = len(m.HeadRoot)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.HeadSlot != 0 {
		n += 1 + sovMessages(uint64(m.HeadSlot))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BeaconBlocksByRangeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartSlot != 0 {
		n += 1 + sovMessages(uint64(m.StartSlot))
	}
	if m.Count != 0 {
		n += 1 + sovMessages(uint64(m.Count))
	}
	if m.Step != 0 {
		n += 1 + sovMessages(uint64(m.Step))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMessages(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Status) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: Status: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Status: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeadForkVersion", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeadForkVersion = append(m.HeadForkVersion[:0], dAtA[iNdEx:postIndex]...)
			if m.HeadForkVersion == nil {
				m.HeadForkVersion = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalizedRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FinalizedRoot = append(m.FinalizedRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.FinalizedRoot == nil {
				m.FinalizedRoot = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalizedEpoch", wireType)
			}
			m.FinalizedEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinalizedEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeadRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeadRoot = append(m.HeadRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.HeadRoot == nil {
				m.HeadRoot = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeadSlot", wireType)
			}
			m.HeadSlot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeadSlot |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BeaconBlocksByRangeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: BeaconBlocksByRangeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BeaconBlocksByRangeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartSlot", wireType)
			}
			m.StartSlot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartSlot |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			m.Step = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Step |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
				return 0, ErrInvalidLengthMessages
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessages
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessages
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessages        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessages          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessages = fmt.Errorf("proto: unexpected end of group")
)
