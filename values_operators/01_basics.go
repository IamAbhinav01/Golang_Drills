package main

import "fmt"

func main(){
	fmt.Println("Hi hello")
	fmt.Println(12)
	fmt.Println(12.5)
	fmt.Println('H','E') //Go code is valid and will compile, but the output may not be what you expect for 'H' and 'E'.
						// Why 72 69?

						// In Go, single quotes (' ') represent a rune (Unicode code point), not a string.

						// 'H' → Unicode value 72
						// 'E' → Unicode value 69

						// So fmt.Println('H', 'E') prints their numeric values.

fmt.Printf("%c\n %c ", 'H','E')
					}

// "" -> refers to string 
// '' -> charactrs are only allowed

