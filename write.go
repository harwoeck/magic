package magic

import (
	"fmt"
	"os"
)

func (m *Manager) writeBytes(buf []byte) {
	_, err := m.out.Write(buf)
	if err != nil {
		fmt.Printf("[manager] unable to write to output: %v\n", err)
		os.Exit(-1)
	}
}

func (m *Manager) Write(v interface{}) {
	m.WriteString(fmt.Sprintf("%v", v))
}

func (m *Manager) Writef(format string, a ...interface{}) {
	m.WriteString(fmt.Sprintf(format, a...))
}

func (m *Manager) WriteBool(v bool) {
	m.Write(v)
}

func (m *Manager) WriteInt8(v int8) {
	m.Write(v)
}

func (m *Manager) WriteInt16(v int16) {
	m.Write(v)
}

func (m *Manager) WriteInt32(v int32) {
	m.Write(v)
}

func (m *Manager) WriteInt64(v int64) {
	m.Write(v)
}

func (m *Manager) WriteInt(v int) {
	m.Write(v)
}

func (m *Manager) WriteByte(v byte) {
	m.Write(v)
}

func (m *Manager) WriteUint8(v uint8) {
	m.Write(v)
}

func (m *Manager) WriteUint16(v uint16) {
	m.Write(v)
}

func (m *Manager) WriteUint32(v uint32) {
	m.Write(v)
}

func (m *Manager) WriteUint64(v uint64) {
	m.Write(v)
}

func (m *Manager) WriteUint(v uint) {
	m.Write(v)
}

func (m *Manager) WriteFloat32(v float32) {
	m.Write(v)
}

func (m *Manager) WriteFloat64(v float64) {
	m.Write(v)
}

func (m *Manager) WriteString(s string) {
	m.writeBytes([]byte(s))
}
