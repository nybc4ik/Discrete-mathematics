package main

import (
	"fmt"
	"unicode/utf8"
)

func convertChar(char rune) []byte {
	switch {
	case char>>7 == 0:
		return []byte{byte(char)}
	case char>>11 == 0:
		var a, b byte
		a = byte(char >> 6)
		b = byte(char & 63)
		return []byte{a | 192, b | 128}
	case char>>16 == 0:
		var a, b, c byte
		a = byte(char >> 12)
		b = byte(char & 4032 >> 6)
		c = byte(char & 63)
		return []byte{a | 224, b | 128, c | 128}
	default:
		var a, b, c, d byte
		a = byte(char >> 18)
		b = byte(char & 258048 >> 12)
		c = byte(char & 4032 >> 6)
		d = byte(char & 63)
		return []byte{a | 240, b | 128, c | 128, d | 128}
	}
}
func encode(utf32 []rune) []byte {
	var res []byte
	for _, char := range utf32 {
		res = append(res, convertChar(char)...)
	}
	return res
}

func convertBytes(bytes []byte) rune {
	switch len(bytes) {
	case 4:
		return rune(bytes[0])&7<<18 + rune(bytes[1])&63<<12 + rune(bytes[2])&63<<6 + rune(bytes[3])&63
	case 3:
		return (rune(bytes[0])&15)<<12 + (rune(bytes[1])&63)<<6 + rune(bytes[2])&63
	case 2:
		return rune(bytes[0]&31)<<6 + rune(bytes[1])&63
	case 1:
		return rune(bytes[0])
	default:
		panic("wrong bytes")
	}
}
func decode(utf8 []byte) []rune {
	var decoded []rune
	for len(utf8) > 0 {
		switch a := utf8[0] >> 4; {
		case a == 15:
			decoded = append(decoded, convertBytes(utf8[:4]))
			utf8 = utf8[4:]
			continue
		case a == 14:
			decoded = append(decoded, convertBytes(utf8[:3]))
			utf8 = utf8[3:]
			continue
		case a >= 12:
			decoded = append(decoded, convertBytes(utf8[:2]))
			utf8 = utf8[2:]
			continue
		default:
			decoded = append(decoded, convertBytes(utf8[:1]))
			utf8 = utf8[1:]
			continue
		}
	}
	return decoded
}

func encodeLibrary(utf32 []rune) []byte {
	var encoded []byte
	for _, char := range utf32 {
		bytes := make([]byte, 4)
		utf8.EncodeRune(bytes, char)
		for _, char := range bytes {
			if char != 0 {
				encoded = append(encoded, char)
			}
		}

	}
	return encoded
}
func decodeLibrary(utf8array []byte) []rune {
	var decoded []rune
	for len(utf8array) > 0 {
		r, size := utf8.DecodeRune(utf8array)
		decoded = append(decoded, r)
		utf8array = utf8array[size:]
	}
	return decoded
}

func main() {
	raw := []rune("12345")
	encoded := encode(raw)
	decoded := decode(encoded)
	encodedLibrary := encodeLibrary(raw)
	decodedLibrary := decodeLibrary(encodedLibrary)
	fmt.Printf("Raw: %v\nImplemented:\nEncoded: %v\nDecoded: %v\n\nLibrary:\nEncoded: %v\nDecoded: %v", raw, encoded, decoded, encodedLibrary, decodedLibrary)
}
