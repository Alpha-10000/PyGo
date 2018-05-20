package main

import "fmt"
import "os"

func main () {
	for _, fileName := range os.Args[1:] {
		fmt.Println(fileName)
		interp(fileName)
	}
}
