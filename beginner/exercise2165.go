package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()
	fmt.Printf("%d\n",len(line))
	if len(line) <140 {
		fmt.Print("TWEET\n")
	}else{
		fmt.Print("MUTE\n")
	}
}
