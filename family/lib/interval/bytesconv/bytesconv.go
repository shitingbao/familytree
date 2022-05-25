package bytesconv

import "unsafe"

func String2Bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct {
		string
		Cap int
	}{s, len(s)},
	))
}
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
