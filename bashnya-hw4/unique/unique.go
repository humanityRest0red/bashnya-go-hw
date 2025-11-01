package myunique

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Options struct {
	c bool
	d bool
	u bool
	i bool
	f uint
	s uint
	ifile string
	ofile string
}

func Run() error {
	var err error
	var options Options
	var input, output *os.File

	options, err = getOptions()
	if err != nil {
		return err
	}

	input, err = chooce(options.ifile, os.Stdin)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err = chooce(options.ofile, os.Stdout)
	if err != nil {
		return err
	}
	defer output.Close()
	
	order, lines := uniqueLines(&options, input)
	Output(options, order, lines, output)

	return nil
}

func chooce(file_name string, default_value *os.File) (*os.File, error){
	if file_name == "" {
		return default_value, nil
	}
	if default_value == os.Stdin{
		return os.Open(file_name)
	}
	return os.Create(file_name)
	
}

func getOptions() (Options, error) {
	options := Options{}

	flag.BoolVar(&options.c, "c", false, "-c")
	flag.BoolVar(&options.d, "d", false, "-d")
	flag.BoolVar(&options.u, "u", false, "-u")
	flag.BoolVar(&options.i, "i", false, "-i")
	flag.UintVar(&options.f, "f", 0, "-f")
	flag.UintVar(&options.s, "s", 0, "-s")

	flag.Parse()

	args := flag.Args()

    if len(args) == 1 {
        options.ifile = args[0]
    } else if len(args) == 2 {
		options.ofile = args[1]
		options.ifile = args[0]
	}

	err := options.isColision()

	return options, err
}

func (options *Options) isColision() error {
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
		return fmt.Errorf("usage c d u")
	}

	return nil
}

func myStrEq(s1, s2 string, o *Options) bool {
	n := int(o.s)
	if len(s1) < n|| len(s2) < n {
		return false
	}

	for i := uint(0); i < o.f; i++ {
		s1 = s1[strings.Index(s1, " ")+1:]
		s2 = s2[strings.Index(s2, " ")+1:]
	}
	
	return o.i && strings.EqualFold(s1[n:], s2[n:]) || !o.i && s1[n:] == s2[n:]
}

func uniqueLines(options *Options, file *os.File) ([]string, map[string]uint) {
	unique := map[string]uint{}
	order := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// if options.f > 0 {

		// }

		exists := false
		for key, _ := range unique {
			if exists = myStrEq(key, line, options); exists {
				unique[key]++
				break
			}
		}

		if !exists {
			order = append(order, line)
			unique[line] = 1
		}


		// if options.i {
		// 	exists := false
		// 	for k, _ := range unique {
		// 		if strings.EqualFold(k, line){
		// 			exists = true
		// 			unique[k]++
		// 			break
		// 		}
		// 	}
		// 	if !exists {
		// 		order = append(order, line)
		// 		unique[line]++
		// 	}
		// } else {
		// 	if _, exists := unique[line]; !exists {
		// 		order = append(order, line)
		// 	}
		// 	unique[line]++
		// }
	}

	return order, unique
}

func Output(options Options, order []string, counts map[string]uint, file *os.File) {
	for _, line := range order {
		count := counts[line]
		if options.d {
			if count > 1 {
				fmt.Fprintln(file, line)
			}
		} else if options.u {
			if count == 1 {
				fmt.Fprintln(file, line)
			}
		} else {
			if options.c {
				fmt.Fprint(file, count, " ")
			}
			fmt.Fprintln(file, line)
		}
	}
}
