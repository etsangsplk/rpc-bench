// Code generated by protoc-gen-gogo.
// source: wire.proto
// DO NOT EDIT!

package rpcbench

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CompressionType int32

const (
	CompressionType_NONE CompressionType = 0
)

var CompressionType_name = map[int32]string{
	0: "NONE",
}
var CompressionType_value = map[string]int32{
	"NONE": 0,
}

func (x CompressionType) String() string {
	return proto.EnumName(CompressionType_name, int32(x))
}
func (CompressionType) EnumDescriptor() ([]byte, []int) { return fileDescriptorWire, []int{0} }

type RequestHeader struct {
	Id               uint64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Method           string          `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	MethodId         int32           `protobuf:"varint,3,opt,name=method_id,json=methodId,proto3" json:"method_id,omitempty"`
	Compression      CompressionType `protobuf:"varint,4,opt,name=compression,proto3,enum=rpcbench.CompressionType" json:"compression,omitempty"`
	UncompressedSize uint32          `protobuf:"varint,5,opt,name=uncompressed_size,json=uncompressedSize,proto3" json:"uncompressed_size,omitempty"`
}

func (m *RequestHeader) Reset()                    { *m = RequestHeader{} }
func (m *RequestHeader) String() string            { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()               {}
func (*RequestHeader) Descriptor() ([]byte, []int) { return fileDescriptorWire, []int{0} }

func (m *RequestHeader) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RequestHeader) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *RequestHeader) GetMethodId() int32 {
	if m != nil {
		return m.MethodId
	}
	return 0
}

func (m *RequestHeader) GetCompression() CompressionType {
	if m != nil {
		return m.Compression
	}
	return CompressionType_NONE
}

func (m *RequestHeader) GetUncompressedSize() uint32 {
	if m != nil {
		return m.UncompressedSize
	}
	return 0
}

type ResponseHeader struct {
	Id               uint64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Method           string          `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Error            string          `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Compression      CompressionType `protobuf:"varint,4,opt,name=compression,proto3,enum=rpcbench.CompressionType" json:"compression,omitempty"`
	UncompressedSize uint32          `protobuf:"varint,5,opt,name=uncompressed_size,json=uncompressedSize,proto3" json:"uncompressed_size,omitempty"`
}

func (m *ResponseHeader) Reset()                    { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string            { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()               {}
func (*ResponseHeader) Descriptor() ([]byte, []int) { return fileDescriptorWire, []int{1} }

func (m *ResponseHeader) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ResponseHeader) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ResponseHeader) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ResponseHeader) GetCompression() CompressionType {
	if m != nil {
		return m.Compression
	}
	return CompressionType_NONE
}

func (m *ResponseHeader) GetUncompressedSize() uint32 {
	if m != nil {
		return m.UncompressedSize
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestHeader)(nil), "rpcbench.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "rpcbench.ResponseHeader")
	proto.RegisterEnum("rpcbench.CompressionType", CompressionType_name, CompressionType_value)
}
func (m *RequestHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWire(dAtA, i, uint64(m.Id))
	}
	if len(m.Method) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWire(dAtA, i, uint64(len(m.Method)))
		i += copy(dAtA[i:], m.Method)
	}
	if m.MethodId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintWire(dAtA, i, uint64(m.MethodId))
	}
	if m.Compression != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintWire(dAtA, i, uint64(m.Compression))
	}
	if m.UncompressedSize != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintWire(dAtA, i, uint64(m.UncompressedSize))
	}
	return i, nil
}

func (m *ResponseHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWire(dAtA, i, uint64(m.Id))
	}
	if len(m.Method) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWire(dAtA, i, uint64(len(m.Method)))
		i += copy(dAtA[i:], m.Method)
	}
	if len(m.Error) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintWire(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	if m.Compression != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintWire(dAtA, i, uint64(m.Compression))
	}
	if m.UncompressedSize != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintWire(dAtA, i, uint64(m.UncompressedSize))
	}
	return i, nil
}

func encodeFixed64Wire(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Wire(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintWire(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RequestHeader) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovWire(uint64(m.Id))
	}
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovWire(uint64(l))
	}
	if m.MethodId != 0 {
		n += 1 + sovWire(uint64(m.MethodId))
	}
	if m.Compression != 0 {
		n += 1 + sovWire(uint64(m.Compression))
	}
	if m.UncompressedSize != 0 {
		n += 1 + sovWire(uint64(m.UncompressedSize))
	}
	return n
}

func (m *ResponseHeader) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovWire(uint64(m.Id))
	}
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovWire(uint64(l))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovWire(uint64(l))
	}
	if m.Compression != 0 {
		n += 1 + sovWire(uint64(m.Compression))
	}
	if m.UncompressedSize != 0 {
		n += 1 + sovWire(uint64(m.UncompressedSize))
	}
	return n
}

func sovWire(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWire(x uint64) (n int) {
	return sovWire(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWire
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RequestHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MethodId", wireType)
			}
			m.MethodId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MethodId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compression", wireType)
			}
			m.Compression = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Compression |= (CompressionType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UncompressedSize", wireType)
			}
			m.UncompressedSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UncompressedSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWire(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
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
func (m *ResponseHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWire
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResponseHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compression", wireType)
			}
			m.Compression = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Compression |= (CompressionType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UncompressedSize", wireType)
			}
			m.UncompressedSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UncompressedSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWire(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
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
func skipWire(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWire
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
					return 0, ErrIntOverflowWire
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWire
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthWire
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWire
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipWire(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthWire = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWire   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("wire.proto", fileDescriptorWire) }

var fileDescriptorWire = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xcf, 0x2c, 0x4a,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x2a, 0x48, 0x4e, 0x4a, 0xcd, 0x4b, 0xce,
	0x50, 0xda, 0xc7, 0xc8, 0xc5, 0x1b, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0xe2, 0x91, 0x9a, 0x98,
	0x92, 0x5a, 0x24, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0xc4,
	0x94, 0x99, 0x22, 0x24, 0xc6, 0xc5, 0x96, 0x9b, 0x5a, 0x92, 0x91, 0x9f, 0x22, 0xc1, 0xa4, 0xc0,
	0xa8, 0xc1, 0x19, 0x04, 0xe5, 0x09, 0x49, 0x73, 0x71, 0x42, 0x58, 0xf1, 0x99, 0x29, 0x12, 0xcc,
	0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x1c, 0x10, 0x01, 0xcf, 0x14, 0x21, 0x6b, 0x2e, 0xee, 0xe4, 0xfc,
	0xdc, 0x82, 0xa2, 0xd4, 0xe2, 0xe2, 0xcc, 0xfc, 0x3c, 0x09, 0x16, 0x05, 0x46, 0x0d, 0x3e, 0x23,
	0x49, 0x3d, 0x98, 0xb5, 0x7a, 0xce, 0x08, 0xc9, 0x90, 0xca, 0x82, 0xd4, 0x20, 0x64, 0xd5, 0x42,
	0xda, 0x5c, 0x82, 0xa5, 0x79, 0x30, 0x81, 0xd4, 0x94, 0xf8, 0xe2, 0xcc, 0xaa, 0x54, 0x09, 0x56,
	0x05, 0x46, 0x0d, 0xde, 0x20, 0x01, 0x64, 0x89, 0xe0, 0xcc, 0xaa, 0x54, 0xa5, 0x1d, 0x8c, 0x5c,
	0x7c, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x24, 0xfa, 0x40, 0x84, 0x8b, 0x35, 0xb5,
	0xa8, 0x28, 0xbf, 0x08, 0xec, 0x7a, 0xce, 0x20, 0x08, 0x87, 0x7e, 0x4e, 0xd7, 0x92, 0xe6, 0xe2,
	0x47, 0x33, 0x4c, 0x88, 0x83, 0x8b, 0xc5, 0xcf, 0xdf, 0xcf, 0x55, 0x80, 0xc1, 0x49, 0xe0, 0xc4,
	0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0x21,
	0x89, 0x0d, 0x1c, 0x77, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0xd9, 0x18, 0xf8, 0xc9,
	0x01, 0x00, 0x00,
}
