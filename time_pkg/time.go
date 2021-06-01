package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch hr := t.Hour(); true {
	case hr > 6 && hr < 12:
		fmt.Println("GM")
	case hr > 12 && hr < 17:
		fmt.Println("Good afternoon")
	case hr > 17 && hr < 20:
		fmt.Println("Good evening")
	case hr > 20 && hr < 6:
		fmt.Println("GN")
	}

}
