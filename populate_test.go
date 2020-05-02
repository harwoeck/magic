package magic

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManager_vLen(t *testing.T) {
	type fields struct {
		input string
	}
	type args struct {
		v reflect.Value
	}
	type want struct {
		len          int
		dynamicInfer bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{"string slice, 0 len, 0 cap", fields{"42"}, args{reflect.ValueOf([]string{})}, want{42, true}},
		{"string slice, 0 len, 42 cap", fields{""}, args{reflect.ValueOf(make([]string, 0, 42))}, want{42, false}},
		{"string slice, 42 len, 42 cap", fields{""}, args{reflect.ValueOf(make([]string, 42))}, want{42, false}},
		{"string array, 53 len", fields{""}, args{reflect.ValueOf([53]string{})}, want{53, false}},

		{"int slice, 0 len, 0 cap", fields{"42"}, args{reflect.ValueOf([]int{})}, want{42, true}},
		{"int slice, 0 len, 42 cap", fields{""}, args{reflect.ValueOf(make([]int, 0, 42))}, want{42, false}},
		{"int slice, 42 len, 42 cap", fields{""}, args{reflect.ValueOf(make([]int, 42))}, want{42, false}},
		{"int array, 53 len", fields{""}, args{reflect.ValueOf([53]int{})}, want{53, false}},

		{"float slice, 0 len, 0 cap", fields{"42"}, args{reflect.ValueOf([]float32{})}, want{42, true}},
		{"float slice, 0 len, 42 cap", fields{""}, args{reflect.ValueOf(make([]float32, 0, 42))}, want{42, false}},
		{"float slice, 42 len, 42 cap", fields{""}, args{reflect.ValueOf(make([]float32, 42))}, want{42, false}},
		{"float array, 53 len", fields{""}, args{reflect.ValueOf([53]float32{})}, want{53, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager(FromString(tt.fields.input), NewStringSplitDecoder(" "), nil)

			len, dynamicInfer := m.vLen(tt.args.v)
			assert.Equal(t, tt.want.len, len)
			assert.Equal(t, tt.want.dynamicInfer, dynamicInfer)
		})
	}
}
