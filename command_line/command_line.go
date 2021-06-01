package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("%#v", os.Args)

	for indx, value := range os.Args {
		fmt.Println("indx:", indx)
		fmt.Println("value:", value)
	}
}
