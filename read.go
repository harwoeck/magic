package magic

import (
	"reflect"
	"strconv"
)

// Read tries to magically return a fully populated objet of the same type as
// you whished. It does this by allocating a new object, taking input src from
// the current cursor position and writing it to the object's memory address.
// This works for primitive types, structs, slices, arrays and combinations of
// them (e.g. slice of struct). Furthermore recursive resolution is supported.
// Therefore you can, for example, create slice of structs with arrays of other
// structs.
//
// As this method purly relies on runtime information (e.g. reflection) it will
// take longer in execution than all the statically typed functions, but it is
// a way more powerful tool.
//
// Rules: if an array or slice is encountered Read will first check if this
// collection is already initialized. If so, the len or cap of the collection
// will be taken as the number of following elements to read from the
// SrcProvider. If the collection is not initialized Read will consider the
// collection as dynamic and first read an integer from the SrcProvider, that
// defines at runtime who many items the collection will contain.
//
//     Read([2]string{}) // static -> load 2 string elements from SrcProvider
//     Read(make([]string, 2)) // static -> load 2 string elements from SrcProvider
//     Read([]string{}) // dynamic -> load 1 int as counter and then counter string from SrcProvider
func (m *Manager) Read(wish interface{}) interface{} {
	t := reflect.TypeOf(wish)
	v := m.newV(t)
	m.pop(t, v)
	return m.getI(t, v)
}

// ReadBool reads a single element from the SrcProvider and tries to interpret
// it as bool.
func (m *Manager) ReadBool() bool {
	b, err := strconv.ParseBool(m.cursor.next())
	if err != nil {
		panic(err)
	}

	return b
}

func (m *Manager) readInt(bitSize int) int64 {
	i, err := strconv.ParseInt(m.cursor.next(), 10, bitSize)
	if err != nil {
		panic(err)
	}

	return i
}

func (m *Manager) readUint(bitSize int) uint64 {
	i, err := strconv.ParseUint(m.cursor.next(), 10, bitSize)
	if err != nil {
		panic(err)
	}

	return i
}

func (m *Manager) readFloat(bitSize int) float64 {
	f, err := strconv.ParseFloat(m.cursor.next(), bitSize)
	if err != nil {
		panic(err)
	}

	return f
}

// ReadInt8 reads a single element from the SrcProvider and tries to interpret
// it as int8.
func (m *Manager) ReadInt8() int8 {
	return int8(m.readInt(8))
}

// ReadInt16 reads a single element from the SrcProvider and tries to interpret
// it as int16.
func (m *Manager) ReadInt16() int16 {
	return int16(m.readInt(16))
}

// ReadInt32 reads a single element from the SrcProvider and tries to interpret
// it as int32.
func (m *Manager) ReadInt32() int32 {
	return int32(m.readInt(32))
}

// ReadInt64 reads a single element from the SrcProvider and tries to interpret
// it as int64.
func (m *Manager) ReadInt64() int64 {
	return m.readInt(64)
}

// ReadInt reads a single element from the SrcProvider and tries to interpret
// it as int.
func (m *Manager) ReadInt() int {
	return int(m.readInt(0))
}

// ReadByte reads a single element from the SrcProvider and tries to interpret
// it as byte.
func (m *Manager) ReadByte() byte {
	return byte(m.readUint(8))
}

// ReadUint8 reads a single element from the SrcProvider and tries to interpret
// it as uint8.
func (m *Manager) ReadUint8() uint8 {
	return uint8(m.readUint(8))
}

// ReadUint16 reads a single element from the SrcProvider and tries to interpret
// it as uint16.
func (m *Manager) ReadUint16() uint16 {
	return uint16(m.readUint(16))
}

// ReadUint32 reads a single element from the SrcProvider and tries to interpret
// it as uint16.
func (m *Manager) ReadUint32() uint32 {
	return uint32(m.readUint(32))
}

// ReadUint64 reads a single element from the SrcProvider and tries to interpret
// it as uint64.
func (m *Manager) ReadUint64() uint64 {
	return m.readUint(64)
}

// ReadUint reads a single element from the SrcProvider and tries to interpret
// it as uint.
func (m *Manager) ReadUint() uint {
	return uint(m.readUint(0))
}

// ReadFloat32 reads a single element from the SrcProvider and tries to interpret
// it as uint64.
func (m *Manager) ReadFloat32() float32 {
	return float32(m.readFloat(32))
}

// ReadFloat64 reads a single element from the SrcProvider and tries to interpret
// it as uint64.
func (m *Manager) ReadFloat64() float64 {
	return m.readFloat(64)
}

// ReadString reads a single element from the SrcProvider and tries to interpret
// it as string.
func (m *Manager) ReadString() string {
	return m.cursor.next()
}
