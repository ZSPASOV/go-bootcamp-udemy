package main

func main() {
	// a string is just a series of byte
	// "hey" can be represented as the following byte slice []byte{104, 101, 121}
	// they are interchangeably convertible
	// you can represent a character as rune literal like so: 'h' -> 104, 'e' -> 101, 'y' -> 121,  104, 101 and 121 are unicode code points, aka Runes
	// a rune literal can represent: byte(uint8), rune(int32) or any other integer type
	// 104 is a decimal number, so its base is 10, it is the most popular numeric system
	// however computers work with binary data, so runes can be represented with other number system
	// when a number starts with the 0b prefix, go thinks it is a binary number 0b0110_1000 is 104 in binary - see below
	// in Binary base 2: the same rune is 2^6 = 64, 2^6 = 32, 2^3 = 8, 64 + 32 + 8 = 104
	// hexaecimal is base 16, if a number starts with the 0x prefix, go thinks it is a hexadecimal number so 0x68 -> 16 ^1 *6 = 96, 16^0 * 8 = 8, 96 + 8 = 104
}
