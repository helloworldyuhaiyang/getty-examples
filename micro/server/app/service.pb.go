// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

/*
	Package main is a generated protocol buffer package.

	It is generated from these files:
		service.proto

	It has these top-level messages:
		TestReq
		TestRsp
		AddReq
		AddRsp
		ErrReq
		ErrRsp
		EventReq
*/
package main

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TestReq struct {
	A string `protobuf:"bytes,1,opt,name=A" json:"A"`
	B string `protobuf:"bytes,2,opt,name=B" json:"B"`
	C string `protobuf:"bytes,3,opt,name=C" json:"C"`
}

func (m *TestReq) Reset()                    { *m = TestReq{} }
func (*TestReq) ProtoMessage()               {}
func (*TestReq) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{0} }

type TestRsp struct {
	A string `protobuf:"bytes,1,opt,name=A" json:"A"`
}

func (m *TestRsp) Reset()                    { *m = TestRsp{} }
func (*TestRsp) ProtoMessage()               {}
func (*TestRsp) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{1} }

type AddReq struct {
	A int32 `protobuf:"varint,1,opt,name=A" json:"A"`
	B int32 `protobuf:"varint,2,opt,name=B" json:"B"`
}

func (m *AddReq) Reset()                    { *m = AddReq{} }
func (*AddReq) ProtoMessage()               {}
func (*AddReq) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{2} }

type AddRsp struct {
	Sum int32 `protobuf:"varint,1,opt,name=Sum" json:"Sum"`
}

func (m *AddRsp) Reset()                    { *m = AddRsp{} }
func (*AddRsp) ProtoMessage()               {}
func (*AddRsp) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{3} }

type ErrReq struct {
	A int32 `protobuf:"varint,1,opt,name=A" json:"A"`
}

func (m *ErrReq) Reset()                    { *m = ErrReq{} }
func (*ErrReq) ProtoMessage()               {}
func (*ErrReq) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{4} }

type ErrRsp struct {
	A int32 `protobuf:"varint,1,opt,name=A" json:"A"`
}

func (m *ErrRsp) Reset()                    { *m = ErrRsp{} }
func (*ErrRsp) ProtoMessage()               {}
func (*ErrRsp) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{5} }

type EventReq struct {
	A string `protobuf:"bytes,1,opt,name=A" json:"A"`
}

func (m *EventReq) Reset()                    { *m = EventReq{} }
func (*EventReq) ProtoMessage()               {}
func (*EventReq) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{6} }

func init() {
	proto.RegisterType((*TestReq)(nil), "main.TestReq")
	proto.RegisterType((*TestRsp)(nil), "main.TestRsp")
	proto.RegisterType((*AddReq)(nil), "main.AddReq")
	proto.RegisterType((*AddRsp)(nil), "main.AddRsp")
	proto.RegisterType((*ErrReq)(nil), "main.ErrReq")
	proto.RegisterType((*ErrRsp)(nil), "main.ErrRsp")
	proto.RegisterType((*EventReq)(nil), "main.EventReq")
}
func (this *TestReq) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*TestReq)
	if !ok {
		that2, ok := that.(TestReq)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *TestReq")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *TestReq but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *TestReq but is not nil && this == nil")
	}
	if this.A != that1.A {
		return fmt.Errorf("A this(%v) Not Equal that(%v)", this.A, that1.A)
	}
	if this.B != that1.B {
		return fmt.Errorf("B this(%v) Not Equal that(%v)", this.B, that1.B)
	}
	if this.C != that1.C {
		return fmt.Errorf("C this(%v) Not Equal that(%v)", this.C, that1.C)
	}
	return nil
}
func (this *TestReq) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*TestReq)
	if !ok {
		that2, ok := that.(TestReq)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.A != that1.A {
		return false
	}
	if this.B != that1.B {
		return false
	}
	if this.C != that1.C {
		return false
	}
	return true
}
func (this *TestRsp) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*TestRsp)
	if !ok {
		that2, ok := that.(TestRsp)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *TestRsp")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *TestRsp but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *TestRsp but is not nil && this == nil")
	}
	if this.A != that1.A {
		return fmt.Errorf("A this(%v) Not Equal that(%v)", this.A, that1.A)
	}
	return nil
}
func (this *TestRsp) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*TestRsp)
	if !ok {
		that2, ok := that.(TestRsp)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.A != that1.A {
		return false
	}
	return true
}
func (this *AddReq) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*AddReq)
	if !ok {
		that2, ok := that.(AddReq)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *AddReq")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *AddReq but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *AddReq but is not nil && this == nil")
	}
	if this.A != that1.A {
		return fmt.Errorf("A this(%v) Not Equal that(%v)", this.A, that1.A)
	}
	if this.B != that1.B {
		return fmt.Errorf("B this(%v) Not Equal that(%v)", this.B, that1.B)
	}
	return nil
}
func (this *AddReq) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*AddReq)
	if !ok {
		that2, ok := that.(AddReq)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.A != that1.A {
		return false
	}
	if this.B != that1.B {
		return false
	}
	return true
}
func (this *AddRsp) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*AddRsp)
	if !ok {
		that2, ok := that.(AddRsp)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *AddRsp")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *AddRsp but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *AddRsp but is not nil && this == nil")
	}
	if this.Sum != that1.Sum {
		return fmt.Errorf("Sum this(%v) Not Equal that(%v)", this.Sum, that1.Sum)
	}
	return nil
}
func (this *AddRsp) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*AddRsp)
	if !ok {
		that2, ok := that.(AddRsp)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Sum != that1.Sum {
		return false
	}
	return true
}
func (this *ErrReq) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ErrReq)
	if !ok {
		that2, ok := that.(ErrReq)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ErrReq")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ErrReq but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ErrReq but is not nil && this == nil")
	}
	if this.A != that1.A {
		return fmt.Errorf("A this(%v) Not Equal that(%v)", this.A, that1.A)
	}
	return nil
}
func (this *ErrReq) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ErrReq)
	if !ok {
		that2, ok := that.(ErrReq)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.A != that1.A {
		return false
	}
	return true
}
func (this *ErrRsp) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ErrRsp)
	if !ok {
		that2, ok := that.(ErrRsp)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ErrRsp")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ErrRsp but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ErrRsp but is not nil && this == nil")
	}
	if this.A != that1.A {
		return fmt.Errorf("A this(%v) Not Equal that(%v)", this.A, that1.A)
	}
	return nil
}
func (this *ErrRsp) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ErrRsp)
	if !ok {
		that2, ok := that.(ErrRsp)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.A != that1.A {
		return false
	}
	return true
}
func (this *EventReq) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*EventReq)
	if !ok {
		that2, ok := that.(EventReq)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *EventReq")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *EventReq but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *EventReq but is not nil && this == nil")
	}
	if this.A != that1.A {
		return fmt.Errorf("A this(%v) Not Equal that(%v)", this.A, that1.A)
	}
	return nil
}
func (this *EventReq) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*EventReq)
	if !ok {
		that2, ok := that.(EventReq)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.A != that1.A {
		return false
	}
	return true
}
func (this *TestReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&main.TestReq{")
	s = append(s, "A: "+fmt.Sprintf("%#v", this.A)+",\n")
	s = append(s, "B: "+fmt.Sprintf("%#v", this.B)+",\n")
	s = append(s, "C: "+fmt.Sprintf("%#v", this.C)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TestRsp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&main.TestRsp{")
	s = append(s, "A: "+fmt.Sprintf("%#v", this.A)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AddReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&main.AddReq{")
	s = append(s, "A: "+fmt.Sprintf("%#v", this.A)+",\n")
	s = append(s, "B: "+fmt.Sprintf("%#v", this.B)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AddRsp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&main.AddRsp{")
	s = append(s, "Sum: "+fmt.Sprintf("%#v", this.Sum)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ErrReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&main.ErrReq{")
	s = append(s, "A: "+fmt.Sprintf("%#v", this.A)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ErrRsp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&main.ErrRsp{")
	s = append(s, "A: "+fmt.Sprintf("%#v", this.A)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *EventReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&main.EventReq{")
	s = append(s, "A: "+fmt.Sprintf("%#v", this.A)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringService(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *TestReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintService(dAtA, i, uint64(len(m.A)))
	i += copy(dAtA[i:], m.A)
	dAtA[i] = 0x12
	i++
	i = encodeVarintService(dAtA, i, uint64(len(m.B)))
	i += copy(dAtA[i:], m.B)
	dAtA[i] = 0x1a
	i++
	i = encodeVarintService(dAtA, i, uint64(len(m.C)))
	i += copy(dAtA[i:], m.C)
	return i, nil
}

func (m *TestRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintService(dAtA, i, uint64(len(m.A)))
	i += copy(dAtA[i:], m.A)
	return i, nil
}

func (m *AddReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintService(dAtA, i, uint64(m.A))
	dAtA[i] = 0x10
	i++
	i = encodeVarintService(dAtA, i, uint64(m.B))
	return i, nil
}

func (m *AddRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintService(dAtA, i, uint64(m.Sum))
	return i, nil
}

func (m *ErrReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintService(dAtA, i, uint64(m.A))
	return i, nil
}

func (m *ErrRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintService(dAtA, i, uint64(m.A))
	return i, nil
}

func (m *EventReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintService(dAtA, i, uint64(len(m.A)))
	i += copy(dAtA[i:], m.A)
	return i, nil
}

func encodeFixed64Service(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Service(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TestReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.A)
	n += 1 + l + sovService(uint64(l))
	l = len(m.B)
	n += 1 + l + sovService(uint64(l))
	l = len(m.C)
	n += 1 + l + sovService(uint64(l))
	return n
}

func (m *TestRsp) Size() (n int) {
	var l int
	_ = l
	l = len(m.A)
	n += 1 + l + sovService(uint64(l))
	return n
}

func (m *AddReq) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovService(uint64(m.A))
	n += 1 + sovService(uint64(m.B))
	return n
}

func (m *AddRsp) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovService(uint64(m.Sum))
	return n
}

func (m *ErrReq) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovService(uint64(m.A))
	return n
}

func (m *ErrRsp) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovService(uint64(m.A))
	return n
}

func (m *EventReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.A)
	n += 1 + l + sovService(uint64(l))
	return n
}

func sovService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TestReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TestReq{`,
		`A:` + fmt.Sprintf("%v", this.A) + `,`,
		`B:` + fmt.Sprintf("%v", this.B) + `,`,
		`C:` + fmt.Sprintf("%v", this.C) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TestRsp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TestRsp{`,
		`A:` + fmt.Sprintf("%v", this.A) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AddReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AddReq{`,
		`A:` + fmt.Sprintf("%v", this.A) + `,`,
		`B:` + fmt.Sprintf("%v", this.B) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AddRsp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AddRsp{`,
		`Sum:` + fmt.Sprintf("%v", this.Sum) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ErrReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ErrReq{`,
		`A:` + fmt.Sprintf("%v", this.A) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ErrRsp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ErrRsp{`,
		`A:` + fmt.Sprintf("%v", this.A) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EventReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EventReq{`,
		`A:` + fmt.Sprintf("%v", this.A) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TestReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: TestReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.A = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.B = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.C = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
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
func (m *TestRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: TestRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.A = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
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
func (m *AddReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: AddReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			m.A = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.A |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			m.B = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.B |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
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
func (m *AddRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: AddRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
			}
			m.Sum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sum |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
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
func (m *ErrReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: ErrReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			m.A = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.A |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
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
func (m *ErrRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: ErrRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			m.A = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.A |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
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
func (m *EventReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: EventReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.A = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
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
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowService
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
				next, err := skipService(dAtA[start:])
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
	ErrInvalidLengthService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("service.proto", fileDescriptorService) }

var fileDescriptorService = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x93,
	0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf,
	0xd7, 0x07, 0x4b, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa4, 0xe4, 0xca,
	0xc5, 0x1e, 0x92, 0x5a, 0x5c, 0x12, 0x94, 0x5a, 0x28, 0x24, 0xc4, 0xc5, 0xe8, 0x28, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0xe9, 0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x10, 0xa3, 0x23, 0x48, 0xcc, 0x49,
	0x82, 0x09, 0x59, 0xcc, 0x09, 0x24, 0xe6, 0x2c, 0xc1, 0x8c, 0x2c, 0xe6, 0xac, 0x24, 0x0b, 0x35,
	0xa6, 0xb8, 0x00, 0x9b, 0x31, 0x4a, 0x06, 0x5c, 0x6c, 0x8e, 0x29, 0x29, 0x28, 0x96, 0xb0, 0x62,
	0xb1, 0x04, 0x2e, 0xe6, 0xa4, 0xa4, 0x00, 0xd1, 0x51, 0x5c, 0x20, 0x24, 0xc6, 0xc5, 0x1c, 0x5c,
	0x9a, 0x8b, 0xa2, 0x07, 0x24, 0xa0, 0x24, 0xc3, 0xc5, 0xe6, 0x5a, 0x54, 0x84, 0xc3, 0x4c, 0x98,
	0x2c, 0xb2, 0x7b, 0x90, 0x64, 0xe5, 0xb8, 0x38, 0x5c, 0xcb, 0x52, 0xf3, 0x70, 0x79, 0xdb, 0xc9,
	0xe4, 0xc4, 0x43, 0x39, 0x86, 0x0b, 0x0f, 0xe5, 0x18, 0x6e, 0x3c, 0x94, 0x63, 0x78, 0xf0, 0x50,
	0x8e, 0xf1, 0xc3, 0x43, 0x39, 0xc6, 0x86, 0x47, 0x72, 0x8c, 0x2b, 0x1e, 0xc9, 0x31, 0x9e, 0x78,
	0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x2f, 0x1e, 0xc9, 0x31, 0x7c,
	0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0x03, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x4b, 0x53,
	0x36, 0x90, 0x01, 0x00, 0x00,
}
