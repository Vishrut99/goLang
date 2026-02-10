package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var lc, wc, cc int
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		s := scan.Text()
		wc += len(strings.Fields(s))
		cc += len(s)
		lc++
		
	}
	os.WriteFile("a.txt", []byte("yes"), 0744) // but now file content changed

	fmt.Printf("Words:%d Characters:%d line:%d", wc, cc, lc)
}
