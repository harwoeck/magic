package magic

import (
	"bufio"
	"reflect"
	"strings"
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
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"string slice, 0 len, 0 cap", fields{"42"}, args{reflect.ValueOf([]string{})}, 42},
		{"string slice, 0 len, 42 cap", fields{""}, args{reflect.ValueOf(make([]string, 0, 42))}, 42},
		{"string slice, 42 len, 42 cap", fields{""}, args{reflect.ValueOf(make([]string, 42))}, 42},
		{"string array, 53 len", fields{""}, args{reflect.ValueOf([53]string{})}, 53},

		{"int slice, 0 len, 0 cap", fields{"42"}, args{reflect.ValueOf([]int{})}, 42},
		{"int slice, 0 len, 42 cap", fields{""}, args{reflect.ValueOf(make([]int, 0, 42))}, 42},
		{"int slice, 42 len, 42 cap", fields{""}, args{reflect.ValueOf(make([]int, 42))}, 42},
		{"int array, 53 len", fields{""}, args{reflect.ValueOf([53]int{})}, 53},

		{"float slice, 0 len, 0 cap", fields{"42"}, args{reflect.ValueOf([]float32{})}, 42},
		{"float slice, 0 len, 42 cap", fields{""}, args{reflect.ValueOf(make([]float32, 0, 42))}, 42},
		{"float slice, 42 len, 42 cap", fields{""}, args{reflect.ValueOf(make([]float32, 42))}, 42},
		{"float array, 53 len", fields{""}, args{reflect.ValueOf([53]float32{})}, 53},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager(bufio.NewScanner(strings.NewReader(tt.fields.input)))

			got := m.vLen(tt.args.v)
			assert.Equal(t, tt.want, got)
		})
	}
}
