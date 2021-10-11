package ddwrapper

import (
	"reflect"
	"unsafe"
)

type Tags []string

func (t Tags) Pair() []string {
	// array needs to be even in len
	if len(t)&1 == 1 {
		return t
	}

	tag := make([]string, len(t)/2)
	for i := 0; i < len(t)/2; i++ {
		var b = make([]byte, 0, len(t[i*2])+len(t[i*2+1])+1)
		b = append(b, t[i*2]...)
		b = append(b, ':')
		b = append(b, t[i*2+1]...)
		tag[i] = bytesToString(b)
	}
	return tag
}
func bytesToString(bytes []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}))
}
