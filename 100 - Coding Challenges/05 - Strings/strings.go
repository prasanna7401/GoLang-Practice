package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	// Basic String Printing
	var name = "Prasanna Aravindan"
	country := "Canada"
	fmt.Printf("Your Name: %s\nCountry: %s\n", name, country)
	fmt.Println(strings.Repeat("#", 20))
	fmt.Println("He says: \"Hello!\"")
	fmt.Println(`C:\Users\a.txt`)
	// END

	// Runes
	b := '中'                                       // => Note the single quote
	fmt.Printf("Rune: %c and Unicode: %U\n", b, b) // => Rune: 中 and Unicode: U+4E2D

	r := 'ă'
	fmt.Printf("Type: %T\n", r) // => int32

	// END

	word1, word2 := "m", "am"
	newWord := word1 + word2 + string(r)
	fmt.Println(newWord) // => mamă

	s1 := "țară means country in Romanian"

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c ", s1[i]) // => (Wrong) È  a r Ä    m e a n s   c o u n t r y   i n   R o m a n i a n
	}

	fmt.Println("\n" + strings.Repeat("#", 20))
	// Decoding a string rune by rune manually
	for i := 0; i < len(s1); {
		r, size := utf8.DecodeRuneInString(s1[i:]) // it returns the rune in string in variable r
		//and its size in bytes in variable size

		// printing out each rune
		fmt.Printf("%c", r) // -> țară means country in Romanian
		i += size           // incrementing i by the size of the rune in bytes
	}

	fmt.Println("\n" + strings.Repeat("#", 20))

	// END

	// Bytes in a string
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v ", s1[i]) // => 200 155 97 114 196 131 32 109 101 97 110 115 32 99 111 117 110 116 114 121 32 105 110 32 82 111 109 97 110 105 97 110
	}

	fmt.Println("\n" + strings.Repeat("#", 20))

	for i, r := range s1 {
		fmt.Printf("byte index: %d -> rune: %c\n", i, r)
	}

	// END

	s2 := "Go is cool!"

	x := s2[0]
	fmt.Println(x)

	// ERROR -> cannot assign to s1[0]
	// s1[0] = 'x'

	// printing the number of runes in the string
	// fmt.Println(len(s1)) // => INCORRECT LENGTH DUE TO RUNES
	fmt.Println(utf8.RuneCountInString(s2))

	// END

	// String Operations
	s := "你好 Go!"
	fmt.Println(strings.Fields(s)) // => [你好 Go!]

	byte_s := []byte(s)
	fmt.Println(byte_s) // => [228 189 160 229 165 189 32 71 111 33]

	rune_s := []rune(s)
	fmt.Println(rune_s) // => [20320 22909 32 71 111 33]

}
