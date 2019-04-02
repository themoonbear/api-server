package utils

import b64 "encoding/base64"

//B64Decode b64 解密
func B64Decode(src string) string {
	dec, _ := b64.StdEncoding.DecodeString(src)
	return string(dec)
}

//B64Encode b64 加密
func B64Encode(src string) string {
	return b64.StdEncoding.EncodeToString([]byte(src))
}
