package uniq

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Options struct {
	c bool
	d bool
	u bool
	i bool
	// f      uint
	// s      uint
	// input  bool
	// output bool
}

func Run() {
	options, err := GetOptions()
	if err == nil {
		lines := UniqLines()
		Output(options, lines)
	} else {
		fmt.Print(err)
	}
}

func GetOptions() (Options, error) {
	options := Options{}

	flag.BoolVar(&options.c, "c", false, "")
	flag.BoolVar(&options.d, "d", false, "")
	flag.BoolVar(&options.u, "u", false, "")
	flag.BoolVar(&options.i, "i", false, "")

	flag.Parse()

	count := 0
	if options.c {
		count++
	}
	if options.d {
		count++
	}
	if options.u {
		count++
	}

	if count > 1 {
		return options, fmt.Errorf("usage c d u")
	}

	return options, nil
}

func UniqLines() map[string]uint {
	unique := make(map[string]uint)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		unique[line]++
	}

	return unique
}

func Output(options Options, lines map[string]uint) {
	for line, num := range lines {
		if options.d {
			if num > 1 {
				fmt.Println(line)
			}
		} else if options.u {
			if num == 1 {
				fmt.Println(line)
			}
		} else {
			if options.c {
				fmt.Printf("%v ", num)
			}
			fmt.Println(line)
		}
	}
}
