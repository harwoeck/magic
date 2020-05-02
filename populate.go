package magic

import (
	"reflect"
	"strings"
	"unsafe"
)

func (m *Manager) newV(t reflect.Type) reflect.Value {
	return reflect.New(t).Elem()
}

func (m *Manager) getI(v reflect.Value) interface{} {
	return v.Interface()
}

func (m *Manager) vLen(vg reflect.Value) (len int, dynamicInfer bool) {
	if vlen := vg.Len(); vlen != 0 {
		return vlen, false
	} else if vcap := vg.Cap(); vcap != 0 {
		return vcap, false
	} else {
		return m.ReadInt(), true
	}
}

func (m *Manager) pop(t reflect.Type, v reflect.Value, vg reflect.Value) {
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
		m.popColl(t, v, vg)
	case reflect.Ptr:
		ptrv := reflect.New(t.Elem())
		m.pop(t.Elem(), ptrv.Elem(), ptrv.Elem())
		v.Set(ptrv)
	case reflect.String:
		v.SetString(m.cursor.next())
	case reflect.Struct:
		m.popStruct(t, v, vg)
	default:
		panic("unsupported type: " + t.Name())
	}
}

func (m *Manager) popColl(t reflect.Type, v reflect.Value, vg reflect.Value) {
	vlen, dynamicInfer := m.vLen(vg)

	var s reflect.Value

	switch t.Kind() {
	case reflect.Array:
		s = reflect.New(reflect.ArrayOf(vlen, t.Elem())).Elem()
	case reflect.Slice:
		s = reflect.MakeSlice(t, vlen, vlen)
	}

	for idx := 0; idx < vlen; idx++ {
		vi := s.Index(idx)

		var vgi reflect.Value
		if !dynamicInfer {
			vgi = vg.Index(idx)
		} else {
			vgi = s.Index(idx)
		}

		m.pop(vi.Type(), vi, vgi)
	}

	v.Set(s)
}

func (m *Manager) popStruct(t reflect.Type, v reflect.Value, vg reflect.Value) {
	for idx := 0; idx < t.NumField(); idx++ {
		f := t.Field(idx)
		if strings.HasPrefix(f.Name, "_") {
			continue
		}

		ft := f.Type
		fvPtr := unsafe.Pointer(v.Field(idx).UnsafeAddr())
		fv := reflect.NewAt(ft, fvPtr).Elem()
		m.pop(ft, fv, vg.Field(idx))
	}
}
