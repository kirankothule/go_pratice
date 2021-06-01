package main

import "fmt"

func main() {
	var s string

	s = "hi there" // string literals
	fmt.Println(s)
	s = `hi there` // Raw string literals
	fmt.Println(s) // output will be same

	s = "<html>\n\t<body>hi there</body>\n<html>"
	fmt.Println(s) // go compiler will process the esqape charaters
	/*<html>
	        <body>hi there</body>
		<html>*/

	s = `<html>\n\t<body>hi there</body>\n<html>`
	fmt.Println(s) // Go compiler will not process the string as its a raw string
	/*<html>\n\t<body>hi there</body>\n<html>*/
}
