package expr

import (
	"reflect"
	"unsafe"
)

func bytesToStr(data []byte) string {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	shdr := reflect.StringHeader{Data: h.Data, Len: h.Len}
	return *(*string)(unsafe.Pointer(&shdr))
}
