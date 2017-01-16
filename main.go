package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/rhinodavid/medianstream/mediankeeper"
)

func main() {
	k := mediankeeper.New()
	medianSum := 0

	flag.Parse()
	if len(flag.Args()) < 1 {
		panic("Enter the name of the file with the integer list")
	}
	f, err := os.Open(flag.Args()[0])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		x, err1 := strconv.Atoi(line)
		if err1 != nil {
			panic(err1)
		}
		k.Push(x)
		medianSum += k.Median()
	}
	fmt.Println("The sum of all the medians is:")
	fmt.Println(medianSum)
}
