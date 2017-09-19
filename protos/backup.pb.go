// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: backup.proto

/*
	Package protos is a generated protocol buffer package.

	It is generated from these files:
		backup.proto
		manifest.proto

	It has these top-level messages:
		BackupItem
		BackupStatus
		BackupStatusItem
		ManifestChangeSet
		ManifestChange
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// BackupItem holds the new information for a specific key: whether it has a value, and what that
// value is.  Also the Version value for that item.
type BackupItem struct {
	Key      []byte `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Counter  uint64 `protobuf:"varint,2,opt,name=Counter,proto3" json:"Counter,omitempty"`
	HasValue bool   `protobuf:"varint,3,opt,name=HasValue,proto3" json:"HasValue,omitempty"`
	// UserMeta always holds a value 0 <= ... < 256.
	UserMeta uint32 `protobuf:"varint,4,opt,name=UserMeta,proto3" json:"UserMeta,omitempty"`
	Value    []byte `protobuf:"bytes,5,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *BackupItem) Reset()                    { *m = BackupItem{} }
func (m *BackupItem) String() string            { return proto.CompactTextString(m) }
func (*BackupItem) ProtoMessage()               {}
func (*BackupItem) Descriptor() ([]byte, []int) { return fileDescriptorBackup, []int{0} }

func (m *BackupItem) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *BackupItem) GetCounter() uint64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *BackupItem) GetHasValue() bool {
	if m != nil {
		return m.HasValue
	}
	return false
}

func (m *BackupItem) GetUserMeta() uint32 {
	if m != nil {
		return m.UserMeta
	}
	return 0
}

func (m *BackupItem) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// BackupStatus describes the state of a backup store -- what backups are made, their max cas
// values, other info
type BackupStatus struct {
	Backups []*BackupStatusItem `protobuf:"bytes,1,rep,name=Backups" json:"Backups,omitempty"`
}

func (m *BackupStatus) Reset()                    { *m = BackupStatus{} }
func (m *BackupStatus) String() string            { return proto.CompactTextString(m) }
func (*BackupStatus) ProtoMessage()               {}
func (*BackupStatus) Descriptor() ([]byte, []int) { return fileDescriptorBackup, []int{1} }

func (m *BackupStatus) GetBackups() []*BackupStatusItem {
	if m != nil {
		return m.Backups
	}
	return nil
}

// BackupStatusItem describes a single backup.  (It might identify the original backup or an
// incremental backup after that.)  Counter is the last generated value of the Badger store's
// incrementing operation counter as of the backup.
type BackupStatusItem struct {
	BackupID uint64 `protobuf:"varint,1,opt,name=BackupID,proto3" json:"BackupID,omitempty"`
	Counter  uint64 `protobuf:"varint,2,opt,name=Counter,proto3" json:"Counter,omitempty"`
}

func (m *BackupStatusItem) Reset()                    { *m = BackupStatusItem{} }
func (m *BackupStatusItem) String() string            { return proto.CompactTextString(m) }
func (*BackupStatusItem) ProtoMessage()               {}
func (*BackupStatusItem) Descriptor() ([]byte, []int) { return fileDescriptorBackup, []int{2} }

func (m *BackupStatusItem) GetBackupID() uint64 {
	if m != nil {
		return m.BackupID
	}
	return 0
}

func (m *BackupStatusItem) GetCounter() uint64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func init() {
	proto.RegisterType((*BackupItem)(nil), "protos.BackupItem")
	proto.RegisterType((*BackupStatus)(nil), "protos.BackupStatus")
	proto.RegisterType((*BackupStatusItem)(nil), "protos.BackupStatusItem")
}
func (m *BackupItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupItem) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBackup(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if m.Counter != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.Counter))
	}
	if m.HasValue {
		dAtA[i] = 0x18
		i++
		if m.HasValue {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.UserMeta != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.UserMeta))
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintBackup(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func (m *BackupStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Backups) > 0 {
		for _, msg := range m.Backups {
			dAtA[i] = 0xa
			i++
			i = encodeVarintBackup(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *BackupStatusItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupStatusItem) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BackupID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.BackupID))
	}
	if m.Counter != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.Counter))
	}
	return i, nil
}

func encodeFixed64Backup(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Backup(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintBackup(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BackupItem) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovBackup(uint64(l))
	}
	if m.Counter != 0 {
		n += 1 + sovBackup(uint64(m.Counter))
	}
	if m.HasValue {
		n += 2
	}
	if m.UserMeta != 0 {
		n += 1 + sovBackup(uint64(m.UserMeta))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovBackup(uint64(l))
	}
	return n
}

func (m *BackupStatus) Size() (n int) {
	var l int
	_ = l
	if len(m.Backups) > 0 {
		for _, e := range m.Backups {
			l = e.Size()
			n += 1 + l + sovBackup(uint64(l))
		}
	}
	return n
}

func (m *BackupStatusItem) Size() (n int) {
	var l int
	_ = l
	if m.BackupID != 0 {
		n += 1 + sovBackup(uint64(m.BackupID))
	}
	if m.Counter != 0 {
		n += 1 + sovBackup(uint64(m.Counter))
	}
	return n
}

func sovBackup(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBackup(x uint64) (n int) {
	return sovBackup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BackupItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
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
			return fmt.Errorf("proto: BackupItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			m.Counter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Counter |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasValue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HasValue = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserMeta", wireType)
			}
			m.UserMeta = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserMeta |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
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
func (m *BackupStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
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
			return fmt.Errorf("proto: BackupStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Backups", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Backups = append(m.Backups, &BackupStatusItem{})
			if err := m.Backups[len(m.Backups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
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
func (m *BackupStatusItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
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
			return fmt.Errorf("proto: BackupStatusItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupStatusItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BackupID", wireType)
			}
			m.BackupID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BackupID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			m.Counter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Counter |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
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
func skipBackup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBackup
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
					return 0, ErrIntOverflowBackup
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
					return 0, ErrIntOverflowBackup
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
				return 0, ErrInvalidLengthBackup
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBackup
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
				next, err := skipBackup(dAtA[start:])
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
	ErrInvalidLengthBackup = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBackup   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("backup.proto", fileDescriptorBackup) }

var fileDescriptorBackup = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x4a, 0x4c, 0xce,
	0x2e, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0x6d, 0x8c,
	0x5c, 0x5c, 0x4e, 0x60, 0x09, 0xcf, 0x92, 0xd4, 0x5c, 0x21, 0x01, 0x2e, 0x66, 0xef, 0xd4, 0x4a,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x10, 0x53, 0x48, 0x82, 0x8b, 0xdd, 0x39, 0xbf, 0x34,
	0xaf, 0x24, 0xb5, 0x48, 0x82, 0x49, 0x81, 0x51, 0x83, 0x25, 0x08, 0xc6, 0x15, 0x92, 0xe2, 0xe2,
	0xf0, 0x48, 0x2c, 0x0e, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x08, 0x82,
	0xf3, 0x41, 0x72, 0xa1, 0xc5, 0xa9, 0x45, 0xbe, 0xa9, 0x25, 0x89, 0x12, 0x2c, 0x0a, 0x8c, 0x1a,
	0xbc, 0x41, 0x70, 0xbe, 0x90, 0x08, 0x17, 0x2b, 0x44, 0x13, 0x2b, 0xd8, 0x16, 0x08, 0x47, 0xc9,
	0x89, 0x8b, 0x07, 0xe2, 0x8e, 0xe0, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0x21, 0x23, 0x2e, 0x76, 0x08,
	0xbf, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x02, 0xe2, 0xf2, 0x62, 0x3d, 0x64, 0x65,
	0x20, 0x47, 0x07, 0xc1, 0x14, 0x2a, 0x79, 0x70, 0x09, 0xa0, 0x4b, 0x82, 0x5c, 0x02, 0xf5, 0x9f,
	0x0b, 0xd8, 0x5b, 0x2c, 0x41, 0x70, 0x3e, 0x6e, 0xbf, 0x39, 0x09, 0x9c, 0x78, 0x24, 0xc7, 0x78,
	0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0x41, 0x02, 0xcc,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x44, 0xaa, 0x84, 0x89, 0x47, 0x01, 0x00, 0x00,
}
