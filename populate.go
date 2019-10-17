package magic

import (
	"reflect"
	"unsafe"
)

func (m *Manager) newV(t reflect.Type) reflect.Value {
	v := reflect.New(t)

	switch t.Kind() {
	/*case reflect.Struct:
	return v*/
	default:
		return v.Elem()
	}
}

func (m *Manager) getI(t reflect.Type, v reflect.Value) interface{} {
	switch t.Kind() {
	/*case reflect.Struct:
	return v.Elem().Interface()*/
	default:
		return v.Interface()
	}
}

func (m *Manager) vLen(v reflect.Value) int {
	len := 0
	if vlen := v.Len(); vlen != 0 {
		len = vlen
	} else if vcap := v.Cap(); vcap != 0 {
		len = vcap
	} else {
		len = m.ReadInt()
	}
	return len
}

func (m *Manager) pop(t reflect.Type, v reflect.Value) {
	switch t.Kind() {
	case reflect.Bool:
		v.SetBool(m.ReadBool())
	case reflect.Int8:
		v.SetInt(m.readInt(8))
	case reflect.Int16:
		v.SetInt(m.readInt(16))
	case reflect.Int:
		v.SetInt(m.readInt(0))
	case reflect.Int32:
		v.SetInt(m.readInt(32))
	case reflect.Int64:
		v.SetInt(m.readInt(64))
	case reflect.Uint8:
		v.SetUint(m.readUint(8))
	case reflect.Uint16:
		v.SetUint(m.readUint(16))
	case reflect.Uint:
		v.SetUint(m.readUint(0))
	case reflect.Uint32:
		v.SetUint(m.readUint(32))
	case reflect.Uint64:
		v.SetUint(m.readUint(64))
	case reflect.Float32:
		v.SetFloat(m.readFloat(32))
	case reflect.Float64:
		v.SetFloat(m.readFloat(64))
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		m.popColl(t, v)
	case reflect.Ptr:
		ptrv := reflect.New(t.Elem())
		m.pop(t.Elem(), ptrv.Elem())
		v.Set(ptrv)
	case reflect.String:
		v.SetString(m.cursor.next())
	case reflect.Struct:
		m.popStruct(t, v)
	default:
		panic("unsupported type: " + t.Name())
	}
}

func (m *Manager) popColl(t reflect.Type, v reflect.Value) {
	len := m.vLen(v)

	var s reflect.Value
	switch t.Kind() {
	case reflect.Array:
		s = reflect.New(reflect.ArrayOf(len, t.Elem())).Elem()
	case reflect.Slice:
		s = reflect.MakeSlice(t, len, len)
	}

	for idx := 0; idx < len; idx++ {
		vi := s.Index(idx)
		m.pop(vi.Type(), vi)
	}

	v.Set(s)
}

func (m *Manager) popStruct(t reflect.Type, v reflect.Value) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for idx := 0; idx < t.NumField(); idx++ {
		ft := t.Field(idx).Type
		fvPtr := unsafe.Pointer(v.Field(idx).UnsafeAddr())
		fv := reflect.NewAt(ft, fvPtr).Elem()
		m.pop(ft, fv)
	}
}
