package magic

import (
	"bufio"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type rfields struct {
	input string
}
type rargs struct {
	wish interface{}
}
type rcases struct {
	name   string
	fields rfields
	args   rargs
	want   interface{}
}

var boolCases = []rcases{
	{"bool", rfields{"true"}, rargs{false}, true},
	{"bool", rfields{"TRUE"}, rargs{false}, true},
	{"bool", rfields{"True"}, rargs{false}, true},
	{"bool", rfields{"false"}, rargs{false}, false},
	{"bool", rfields{"FALSE"}, rargs{false}, false},
	{"bool", rfields{"False"}, rargs{false}, false},
	{"bool", rfields{"true"}, rargs{true}, true},
	{"bool", rfields{"TRUE"}, rargs{true}, true},
	{"bool", rfields{"True"}, rargs{true}, true},
	{"bool", rfields{"false"}, rargs{true}, false},
	{"bool", rfields{"FALSE"}, rargs{true}, false},
	{"bool", rfields{"False"}, rargs{true}, false},
}

var zeroInt8 int8
var int8Cases = []rcases{
	{"int8", rfields{"0"}, rargs{zeroInt8}, int8(0)},
	{"int8", rfields{"42"}, rargs{zeroInt8}, int8(42)},
	{"int8", rfields{strconv.Itoa(math.MaxInt8)}, rargs{zeroInt8}, int8(math.MaxInt8)},
	{"int8", rfields{"-53"}, rargs{zeroInt8}, int8(-53)},
	{"int8", rfields{strconv.Itoa(math.MinInt8)}, rargs{zeroInt8}, int8(math.MinInt8)},
	{"int8", rfields{"0"}, rargs{int8(100)}, int8(0)},
	{"int8", rfields{"42"}, rargs{int8(12)}, int8(42)},
	{"int8", rfields{strconv.Itoa(math.MaxInt8)}, rargs{int8(-50)}, int8(math.MaxInt8)},
	{"int8", rfields{"-53"}, rargs{int8(-40)}, int8(-53)},
	{"int8", rfields{strconv.Itoa(math.MinInt8)}, rargs{int8(117)}, int8(math.MinInt8)},
}

var zeroInt16 int16
var int16Cases = []rcases{
	{"int16", rfields{"0"}, rargs{zeroInt16}, int16(0)},
	{"int16", rfields{"42"}, rargs{zeroInt16}, int16(42)},
	{"int16", rfields{strconv.Itoa(math.MaxInt16)}, rargs{zeroInt16}, int16(math.MaxInt16)},
	{"int16", rfields{"-53"}, rargs{zeroInt16}, int16(-53)},
	{"int16", rfields{strconv.Itoa(math.MinInt16)}, rargs{zeroInt16}, int16(math.MinInt16)},
	{"int16", rfields{"0"}, rargs{int16(100)}, int16(0)},
	{"int16", rfields{"42"}, rargs{int16(12)}, int16(42)},
	{"int16", rfields{strconv.Itoa(math.MaxInt16)}, rargs{int16(-50)}, int16(math.MaxInt16)},
	{"int16", rfields{"-53"}, rargs{int16(-40)}, int16(-53)},
	{"int16", rfields{strconv.Itoa(math.MinInt16)}, rargs{int16(117)}, int16(math.MinInt16)},
}

var zeroInt32 int32
var int32Cases = []rcases{
	{"int32", rfields{"0"}, rargs{zeroInt32}, int32(0)},
	{"int32", rfields{"42"}, rargs{zeroInt32}, int32(42)},
	{"int32", rfields{strconv.Itoa(math.MaxInt32)}, rargs{zeroInt32}, int32(math.MaxInt32)},
	{"int32", rfields{"-53"}, rargs{zeroInt32}, int32(-53)},
	{"int32", rfields{strconv.Itoa(math.MinInt32)}, rargs{zeroInt32}, int32(math.MinInt32)},
	{"int32", rfields{"0"}, rargs{int32(100)}, int32(0)},
	{"int32", rfields{"42"}, rargs{int32(12)}, int32(42)},
	{"int32", rfields{strconv.Itoa(math.MaxInt32)}, rargs{int32(-50)}, int32(math.MaxInt32)},
	{"int32", rfields{"-53"}, rargs{int32(-40)}, int32(-53)},
	{"int32", rfields{strconv.Itoa(math.MinInt32)}, rargs{int32(117)}, int32(math.MinInt32)},
}

var zeroInt64 int64
var int64Cases = []rcases{
	{"int64", rfields{"0"}, rargs{zeroInt64}, int64(0)},
	{"int64", rfields{"42"}, rargs{zeroInt64}, int64(42)},
	{"int64", rfields{strconv.Itoa(math.MaxInt64)}, rargs{zeroInt64}, int64(math.MaxInt64)},
	{"int64", rfields{"-53"}, rargs{zeroInt64}, int64(-53)},
	{"int64", rfields{strconv.Itoa(math.MinInt64)}, rargs{zeroInt64}, int64(math.MinInt64)},
	{"int64", rfields{"0"}, rargs{int64(100)}, int64(0)},
	{"int64", rfields{"42"}, rargs{int64(12)}, int64(42)},
	{"int64", rfields{strconv.Itoa(math.MaxInt64)}, rargs{int64(-50)}, int64(math.MaxInt64)},
	{"int64", rfields{"-53"}, rargs{int64(-40)}, int64(-53)},
	{"int64", rfields{strconv.Itoa(math.MinInt64)}, rargs{int64(117)}, int64(math.MinInt64)},
}

var zeroInt int
var intCases = []rcases{
	{"int", rfields{"0"}, rargs{zeroInt}, int(0)},
	{"int", rfields{"42"}, rargs{zeroInt}, int(42)},
	{"int", rfields{strconv.Itoa(MaxInt)}, rargs{zeroInt}, MaxInt},
	{"int", rfields{"-53"}, rargs{zeroInt}, int(-53)},
	{"int", rfields{strconv.Itoa(MinInt)}, rargs{zeroInt}, MinInt},
	{"int", rfields{"0"}, rargs{int(100)}, int(0)},
	{"int", rfields{"42"}, rargs{int(12)}, int(42)},
	{"int", rfields{strconv.Itoa(MaxInt)}, rargs{int(-50)}, MaxInt},
	{"int", rfields{"-53"}, rargs{int(-40)}, int(-53)},
	{"int", rfields{strconv.Itoa(MinInt)}, rargs{int(117)}, MinInt},
}

var zeroByte byte
var byteCases = []rcases{
	{"byte", rfields{"0"}, rargs{zeroByte}, byte(0)},
	{"byte", rfields{"42"}, rargs{zeroByte}, byte(42)},
	{"byte", rfields{strconv.Itoa(math.MaxUint8)}, rargs{zeroByte}, byte(math.MaxUint8)},
	{"byte", rfields{"0"}, rargs{byte(100)}, byte(0)},
	{"byte", rfields{strconv.Itoa(math.MaxUint8)}, rargs{byte(100)}, byte(math.MaxUint8)},
}

var zeroUint8 uint8
var uint8Cases = []rcases{
	{"uint8", rfields{"0"}, rargs{zeroUint8}, uint8(0)},
	{"uint8", rfields{"42"}, rargs{zeroUint8}, uint8(42)},
	{"uint8", rfields{strconv.Itoa(math.MaxUint8)}, rargs{zeroUint8}, uint8(math.MaxUint8)},
	{"uint8", rfields{"0"}, rargs{uint8(100)}, uint8(0)},
	{"uint8", rfields{strconv.Itoa(math.MaxUint8)}, rargs{uint8(100)}, uint8(math.MaxUint8)},
}

var zeroUint16 uint16
var uint16Cases = []rcases{
	{"uint16", rfields{"0"}, rargs{zeroUint16}, uint16(0)},
	{"uint16", rfields{"42"}, rargs{zeroUint16}, uint16(42)},
	{"uint16", rfields{strconv.Itoa(math.MaxUint16)}, rargs{zeroUint16}, uint16(math.MaxUint16)},
	{"uint16", rfields{"0"}, rargs{uint16(100)}, uint16(0)},
	{"uint16", rfields{strconv.Itoa(math.MaxUint16)}, rargs{uint16(100)}, uint16(math.MaxUint16)},
}

var zeroUint32 uint32
var uint32Cases = []rcases{
	{"uint32", rfields{"0"}, rargs{zeroUint32}, uint32(0)},
	{"uint32", rfields{"42"}, rargs{zeroUint32}, uint32(42)},
	{"uint32", rfields{strconv.Itoa(math.MaxUint32)}, rargs{zeroUint32}, uint32(math.MaxUint32)},
	{"uint32", rfields{"0"}, rargs{uint32(100)}, uint32(0)},
	{"uint32", rfields{strconv.Itoa(math.MaxUint32)}, rargs{uint32(100)}, uint32(math.MaxUint32)},
}

var zeroUint64 uint64
var uint64Cases = []rcases{
	{"uint64", rfields{"0"}, rargs{zeroUint64}, uint64(0)},
	{"uint64", rfields{"42"}, rargs{zeroUint64}, uint64(42)},
	{"uint64", rfields{strconv.FormatUint(math.MaxUint64, 10)}, rargs{zeroUint64}, uint64(math.MaxUint64)},
	{"uint64", rfields{"0"}, rargs{uint64(100)}, uint64(0)},
	{"uint64", rfields{strconv.FormatUint(math.MaxUint64, 10)}, rargs{uint64(100)}, uint64(math.MaxUint64)},
}

var zeroUint uint
var uintCases = []rcases{
	{"uint", rfields{"0"}, rargs{zeroUint}, uint(0)},
	{"uint", rfields{"42"}, rargs{zeroUint}, uint(42)},
	{"uint", rfields{strconv.FormatUint(uint64(MaxUint), 10)}, rargs{zeroUint}, uint(MaxUint)},
	{"uint", rfields{"0"}, rargs{uint(100)}, uint(0)},
	{"uint", rfields{strconv.FormatUint(uint64(MaxUint), 10)}, rargs{uint(100)}, uint(MaxUint)},
}

var floatFmt byte = 'e'

var zeroFloat32 float32
var float32Cases = []rcases{
	{"float32", rfields{"0"}, rargs{zeroFloat32}, float32(0)},
	{"float32", rfields{"42.42"}, rargs{zeroFloat32}, float32(42.42)},
	{"float32", rfields{strconv.FormatFloat(math.MaxFloat32, floatFmt, -1, 32)}, rargs{zeroFloat32}, float32(math.MaxFloat32)},
	{"float32", rfields{"-53.53"}, rargs{zeroFloat32}, float32(-53.53)},
	{"float32", rfields{strconv.FormatFloat(math.SmallestNonzeroFloat32, floatFmt, -1, 32)}, rargs{zeroFloat32}, float32(math.SmallestNonzeroFloat32)},
	{"float32", rfields{"0"}, rargs{float32(100.001)}, float32(0)},
	{"float32", rfields{"42.42"}, rargs{float32(12.12)}, float32(42.42)},
	{"float32", rfields{strconv.FormatFloat(math.MaxFloat32, floatFmt, -1, 32)}, rargs{float32(-50.05)}, float32(math.MaxFloat32)},
	{"float32", rfields{"-53.53"}, rargs{float32(-40.04)}, float32(-53.53)},
	{"float32", rfields{strconv.FormatFloat(math.SmallestNonzeroFloat32, floatFmt, -1, 32)}, rargs{float32(117.711)}, float32(math.SmallestNonzeroFloat32)},
}

var zeroFloat64 float64
var float64Cases = []rcases{
	{"float64", rfields{"0"}, rargs{zeroFloat64}, float64(0)},
	{"float64", rfields{"42.42"}, rargs{zeroFloat64}, float64(42.42)},
	{"float64", rfields{strconv.FormatFloat(math.MaxFloat64, floatFmt, -1, 64)}, rargs{zeroFloat64}, float64(math.MaxFloat64)},
	{"float64", rfields{"-53.53"}, rargs{zeroFloat64}, float64(-53.53)},
	{"float64", rfields{strconv.FormatFloat(math.SmallestNonzeroFloat64, floatFmt, -1, 64)}, rargs{zeroFloat64}, float64(math.SmallestNonzeroFloat64)},
	{"float64", rfields{"0"}, rargs{float64(100.001)}, float64(0)},
	{"float64", rfields{"42.42"}, rargs{float64(12.12)}, float64(42.42)},
	{"float64", rfields{strconv.FormatFloat(math.MaxFloat64, floatFmt, -1, 64)}, rargs{float64(-50.05)}, float64(math.MaxFloat64)},
	{"float64", rfields{"-53.53"}, rargs{float64(-40.04)}, float64(-53.53)},
	{"float64", rfields{strconv.FormatFloat(math.SmallestNonzeroFloat64, floatFmt, -1, 64)}, rargs{float64(117.711)}, float64(math.SmallestNonzeroFloat64)},
}

var zeroString string
var stringCases = []rcases{
	{"string", rfields{"x"}, rargs{zeroString}, "x"},
	{"string", rfields{"X"}, rargs{zeroString}, "X"},
	{"string", rfields{"LoreMImpSuM"}, rargs{zeroString}, "LoreMImpSuM"},
	{"string", rfields{"x"}, rargs{"zeroString"}, "x"},
	{"string", rfields{"X"}, rargs{"x"}, "X"},
	{"string", rfields{"LoreMImpSuM"}, rargs{"y"}, "LoreMImpSuM"},
}

var zeroIntColl []int
var intCollCases = []rcases{
	{"intColl", rfields{"0"}, rargs{zeroIntColl}, []int{}},
	{"intColl", rfields{"5 2 17 13 1 0"}, rargs{zeroIntColl}, []int{2, 17, 13, 1, 0}},
	{"intColl", rfields{"5 2 17 13 1 0"}, rargs{make([]int, 0, 6)}, []int{2, 17, 13, 1, 0}},
	{"intColl", rfields{"5 2 17 13 1 0"}, rargs{make([]int, 6)}, []int{2, 17, 13, 1, 0}},
	{"intColl", rfields{"5 2 17 13 1 0"}, rargs{[6]int{}}, [6]int{5, 2, 17, 13, 1, 0}},
}

type child struct {
	childName string
}

type testStruct struct {
	version int
	name    string
	c0      child
	c1      *child
	c2      *[]*child
	c3      [2]bool
}

var testStructCases = []rcases{
	{"testStruct1", rfields{"100 intensiveTester baby0 baby1 2 baby2 baby3 false TRUE"}, rargs{testStruct{}}, testStruct{
		version: 100,
		name:    "intensiveTester",
		c0:      child{childName: "baby0"},
		c1:      &child{childName: "baby1"},
		c2:      &[]*child{&child{childName: "baby2"}, &child{childName: "baby3"}},
		c3:      [2]bool{false, true},
	}},
	{"testStruct2", rfields{"17 xyw babyChamp byb 3 ab1 ab2 ab3 true FALSE"}, rargs{&testStruct{}}, &testStruct{
		version: 17,
		name:    "xyw",
		c0:      child{childName: "babyChamp"},
		c1:      &child{childName: "byb"},
		c2:      &[]*child{&child{childName: "ab1"}, &child{childName: "ab2"}, &child{childName: "ab3"}},
		c3:      [2]bool{true, false},
	}},
	{"testStructSlice", rfields{
		"2\n" +
			"100 intensiveTester baby0 baby1 2 baby2 baby3 false TRUE\n" +
			"17 xyw babyChamp byb 3 ab1 ab2 ab3 true FALSE"}, rargs{[]*testStruct{}}, []*testStruct{
		&testStruct{
			version: 100,
			name:    "intensiveTester",
			c0:      child{childName: "baby0"},
			c1:      &child{childName: "baby1"},
			c2:      &[]*child{&child{childName: "baby2"}, &child{childName: "baby3"}},
			c3:      [2]bool{false, true},
		},
		&testStruct{
			version: 17,
			name:    "xyw",
			c0:      child{childName: "babyChamp"},
			c1:      &child{childName: "byb"},
			c2:      &[]*child{&child{childName: "ab1"}, &child{childName: "ab2"}, &child{childName: "ab3"}},
			c3:      [2]bool{true, false},
		},
	}},
}

func TestManager_Read(t *testing.T) {
	tests := []rcases{}

	tests = append(tests, boolCases...)

	tests = append(tests, int8Cases...)
	tests = append(tests, int16Cases...)
	tests = append(tests, int32Cases...)
	tests = append(tests, int64Cases...)
	tests = append(tests, intCases...)

	tests = append(tests, byteCases...)
	tests = append(tests, uint8Cases...)
	tests = append(tests, uint16Cases...)
	tests = append(tests, uint32Cases...)
	tests = append(tests, uint64Cases...)
	tests = append(tests, uintCases...)

	tests = append(tests, float32Cases...)
	tests = append(tests, float64Cases...)

	tests = append(tests, stringCases...)

	tests = append(tests, intCollCases...)

	tests = append(tests, testStructCases...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager(bufio.NewScanner(strings.NewReader(tt.fields.input)))
			assert.NotNil(t, m)

			got := m.Read(tt.args.wish)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestManager_ReadSpecific(t *testing.T) {
	tests := []struct {
		cases []rcases
		exec  func(*Manager) interface{}
	}{
		{boolCases, func(m *Manager) interface{} { return m.ReadBool() }},

		{int8Cases, func(m *Manager) interface{} { return m.ReadInt8() }},
		{int16Cases, func(m *Manager) interface{} { return m.ReadInt16() }},
		{int32Cases, func(m *Manager) interface{} { return m.ReadInt32() }},
		{int64Cases, func(m *Manager) interface{} { return m.ReadInt64() }},
		{intCases, func(m *Manager) interface{} { return m.ReadInt() }},

		{byteCases, func(m *Manager) interface{} { return m.ReadByte() }},
		{uint8Cases, func(m *Manager) interface{} { return m.ReadUint8() }},
		{uint16Cases, func(m *Manager) interface{} { return m.ReadUint16() }},
		{uint32Cases, func(m *Manager) interface{} { return m.ReadUint32() }},
		{uint64Cases, func(m *Manager) interface{} { return m.ReadUint64() }},
		{uintCases, func(m *Manager) interface{} { return m.ReadUint() }},

		{float32Cases, func(m *Manager) interface{} { return m.ReadFloat32() }},
		{float64Cases, func(m *Manager) interface{} { return m.ReadFloat64() }},

		{stringCases, func(m *Manager) interface{} { return m.ReadString() }},
	}

	for _, tt := range tests {
		for _, ttc := range tt.cases {
			t.Run(ttc.name, func(t *testing.T) {
				m := NewManager(bufio.NewScanner(strings.NewReader(ttc.fields.input)))
				assert.NotNil(t, m)

				got := tt.exec(m)
				assert.Equal(t, ttc.want, got)
			})
		}
	}
}

func TestManager_ReadPanic(t *testing.T) {
	m := NewManager(bufio.NewScanner(strings.NewReader("tRuE NOTanINT noUINT FloatNotHere")))

	assert.Panics(t, func() {
		m.Read(complex(1, 2))
	})
	assert.Panics(t, func() {
		m.Read(complex64(complex(3, 4)))
	})
	assert.Panics(t, func() {
		m.Read(make(chan int))
	})
	assert.Panics(t, func() {
		m.Read(func() {})
	})
	assert.Panics(t, func() {
		m.Read(make(map[string]string))
	})
	assert.Panics(t, func() {
		m.Read(unsafe.Pointer(reflect.ValueOf(0).UnsafeAddr()))
	})

	assert.Panics(t, func() {
		m.ReadBool()
	})
	assert.Panics(t, func() {
		m.ReadInt()
	})
	assert.Panics(t, func() {
		m.ReadUint()
	})
	assert.Panics(t, func() {
		m.ReadFloat32()
	})
}
