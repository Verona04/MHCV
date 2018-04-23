package main

import "fmt"

func IterateOverASCIIStringLiteral(sl string) {
	for _, val := range sl {
		fmt.Printf("%X\t%s\t%b\n", val, string(val), val)
	}
}

func main() {
	var str2 string
	for i := 128; i <= 255; i++ {
		str2 += string(i)
	}

	IterateOverASCIIStringLiteral(str2)

	str := []byte{0x22, 0xe2, 0x82, 0xac, 0x20, 0xc3, 0xb7, 0x20, 0xc2, 0xbe, 0x20, 0x64, 0x6f, 0x6c, 0x6c, 0x61, 0x72, 0x22}
	extAsciiTxt := ExtendedASCIIText(string(str))
	fmt.Println(extAsciiTxt)
}

// Kode for Oppgave 2b
func ExtendedASCIIText(inputStr string) string {
	var resultStr string
	for _, val := range inputStr {
		if val >= 128 && val <= 255 {
			resultStr += string(val)
		}
	}

	return resultStr
}