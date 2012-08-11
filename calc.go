package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print(">")
	r := bufio.NewReaderSize(os.Stdin, 16)

	for {
		cmd := ""
		for {

			sl, pre, err := r.ReadLine()
			if err != nil {
				fmt.Errorf("readline", err)
				return
			}

			cmd += string(sl)

			if !pre {
				break
			}

		}

		if cmd == "quit" {
			break
		} else {
			Parse(cmd)
			fmt.Print(">")
		}
	}
}
