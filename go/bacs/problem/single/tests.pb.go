// Code generated by protoc-gen-go.
// source: bacs/problem/single/tests.proto
// DO NOT EDIT!

package single

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TestQuery_Wildcard_Flag int32

const (
	TestQuery_Wildcard_IGNORE_CASE TestQuery_Wildcard_Flag = 0
)

var TestQuery_Wildcard_Flag_name = map[int32]string{
	0: "IGNORE_CASE",
}
var TestQuery_Wildcard_Flag_value = map[string]int32{
	"IGNORE_CASE": 0,
}

func (x TestQuery_Wildcard_Flag) String() string {
	return proto.EnumName(TestQuery_Wildcard_Flag_name, int32(x))
}

type TestQuery_Regex_Flag int32

const (
	TestQuery_Regex_IGNORE_CASE TestQuery_Regex_Flag = 0
)

var TestQuery_Regex_Flag_name = map[int32]string{
	0: "IGNORE_CASE",
}
var TestQuery_Regex_Flag_value = map[string]int32{
	"IGNORE_CASE": 0,
}

func (x TestQuery_Regex_Flag) String() string {
	return proto.EnumName(TestQuery_Regex_Flag_name, int32(x))
}

type TestSequence_Order int32

const (
	TestSequence_IDENTITY        TestSequence_Order = 0
	TestSequence_NUMERIC         TestSequence_Order = 1
	TestSequence_LEXICOGRAPHICAL TestSequence_Order = 2
)

var TestSequence_Order_name = map[int32]string{
	0: "IDENTITY",
	1: "NUMERIC",
	2: "LEXICOGRAPHICAL",
}
var TestSequence_Order_value = map[string]int32{
	"IDENTITY":        0,
	"NUMERIC":         1,
	"LEXICOGRAPHICAL": 2,
}

func (x TestSequence_Order) String() string {
	return proto.EnumName(TestSequence_Order_name, int32(x))
}

type TestSequence_ContinueCondition int32

const (
	TestSequence_ALWAYS   TestSequence_ContinueCondition = 0
	TestSequence_WHILE_OK TestSequence_ContinueCondition = 1
)

var TestSequence_ContinueCondition_name = map[int32]string{
	0: "ALWAYS",
	1: "WHILE_OK",
}
var TestSequence_ContinueCondition_value = map[string]int32{
	"ALWAYS":   0,
	"WHILE_OK": 1,
}

func (x TestSequence_ContinueCondition) String() string {
	return proto.EnumName(TestSequence_ContinueCondition_name, int32(x))
}

type TestQuery struct {
	// Types that are valid to be assigned to Query:
	//	*TestQuery_Id
	//	*TestQuery_Wildcard_
	//	*TestQuery_Regex_
	Query isTestQuery_Query `protobuf_oneof:"query"`
}

func (m *TestQuery) Reset()         { *m = TestQuery{} }
func (m *TestQuery) String() string { return proto.CompactTextString(m) }
func (*TestQuery) ProtoMessage()    {}

type isTestQuery_Query interface {
	isTestQuery_Query()
}

type TestQuery_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,oneof"`
}
type TestQuery_Wildcard_ struct {
	Wildcard *TestQuery_Wildcard `protobuf:"bytes,2,opt,name=wildcard,oneof"`
}
type TestQuery_Regex_ struct {
	Regex *TestQuery_Regex `protobuf:"bytes,3,opt,name=regex,oneof"`
}

func (*TestQuery_Id) isTestQuery_Query()        {}
func (*TestQuery_Wildcard_) isTestQuery_Query() {}
func (*TestQuery_Regex_) isTestQuery_Query()    {}

func (m *TestQuery) GetQuery() isTestQuery_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *TestQuery) GetId() string {
	if x, ok := m.GetQuery().(*TestQuery_Id); ok {
		return x.Id
	}
	return ""
}

func (m *TestQuery) GetWildcard() *TestQuery_Wildcard {
	if x, ok := m.GetQuery().(*TestQuery_Wildcard_); ok {
		return x.Wildcard
	}
	return nil
}

func (m *TestQuery) GetRegex() *TestQuery_Regex {
	if x, ok := m.GetQuery().(*TestQuery_Regex_); ok {
		return x.Regex
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TestQuery) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _TestQuery_OneofMarshaler, _TestQuery_OneofUnmarshaler, []interface{}{
		(*TestQuery_Id)(nil),
		(*TestQuery_Wildcard_)(nil),
		(*TestQuery_Regex_)(nil),
	}
}

func _TestQuery_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TestQuery)
	// query
	switch x := m.Query.(type) {
	case *TestQuery_Id:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Id)
	case *TestQuery_Wildcard_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Wildcard); err != nil {
			return err
		}
	case *TestQuery_Regex_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Regex); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TestQuery.Query has unexpected type %T", x)
	}
	return nil
}

func _TestQuery_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TestQuery)
	switch tag {
	case 1: // query.id
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Query = &TestQuery_Id{x}
		return true, err
	case 2: // query.wildcard
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TestQuery_Wildcard)
		err := b.DecodeMessage(msg)
		m.Query = &TestQuery_Wildcard_{msg}
		return true, err
	case 3: // query.regex
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TestQuery_Regex)
		err := b.DecodeMessage(msg)
		m.Query = &TestQuery_Regex_{msg}
		return true, err
	default:
		return false, nil
	}
}

type TestQuery_Wildcard struct {
	Flag  []TestQuery_Wildcard_Flag `protobuf:"varint,1,rep,name=flag,enum=bacs.problem.single.TestQuery_Wildcard_Flag" json:"flag,omitempty"`
	Value string                    `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *TestQuery_Wildcard) Reset()         { *m = TestQuery_Wildcard{} }
func (m *TestQuery_Wildcard) String() string { return proto.CompactTextString(m) }
func (*TestQuery_Wildcard) ProtoMessage()    {}

type TestQuery_Regex struct {
	Flag  []TestQuery_Regex_Flag `protobuf:"varint,1,rep,name=flag,enum=bacs.problem.single.TestQuery_Regex_Flag" json:"flag,omitempty"`
	Value string                 `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *TestQuery_Regex) Reset()         { *m = TestQuery_Regex{} }
func (m *TestQuery_Regex) String() string { return proto.CompactTextString(m) }
func (*TestQuery_Regex) ProtoMessage()    {}

type TestSequence struct {
	Query             []*TestQuery                   `protobuf:"bytes,1,rep,name=query" json:"query,omitempty"`
	Order             TestSequence_Order             `protobuf:"varint,2,opt,name=order,enum=bacs.problem.single.TestSequence_Order" json:"order,omitempty"`
	ContinueCondition TestSequence_ContinueCondition `protobuf:"varint,3,opt,name=continue_condition,enum=bacs.problem.single.TestSequence_ContinueCondition" json:"continue_condition,omitempty"`
}

func (m *TestSequence) Reset()         { *m = TestSequence{} }
func (m *TestSequence) String() string { return proto.CompactTextString(m) }
func (*TestSequence) ProtoMessage()    {}

func (m *TestSequence) GetQuery() []*TestQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

func init() {
	proto.RegisterEnum("bacs.problem.single.TestQuery_Wildcard_Flag", TestQuery_Wildcard_Flag_name, TestQuery_Wildcard_Flag_value)
	proto.RegisterEnum("bacs.problem.single.TestQuery_Regex_Flag", TestQuery_Regex_Flag_name, TestQuery_Regex_Flag_value)
	proto.RegisterEnum("bacs.problem.single.TestSequence_Order", TestSequence_Order_name, TestSequence_Order_value)
	proto.RegisterEnum("bacs.problem.single.TestSequence_ContinueCondition", TestSequence_ContinueCondition_name, TestSequence_ContinueCondition_value)
}
