package main

import (
	"fmt"
	"os"
)

func main(){
	for _, arg := range os.Args[1:] {
		fmt.Print(arg)
		fmt.Print(" ")
	}
}
