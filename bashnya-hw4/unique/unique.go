package myunique

import (
	"bufio"
	"flag"
	"fmt"
	"omap"
	"os"
	"strings"
)

var ErrCDUFlags = fmt.Errorf("use only one of c, d, u")

type Options struct {
	c     bool
	d     bool
	u     bool
	i     bool
	f     uint
	s     uint
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

	input, err = choose(options.ifile, os.Stdin)
	if err != nil {
		return err
	}
	defer func() {
		if input != os.Stdin {
			input.Close()
		}
	}()

	output, err = choose(options.ofile, os.Stdout)
	if err != nil {
		return err
	}
	defer func() {
		if output != os.Stdout {
			output.Close()
		}
	}()

	lines := readLines(input)
	result := uniqueLines(lines, options)
	writeLines(result, output)

	return nil
}

func choose(file_name string, default_value *os.File) (*os.File, error) {
	if file_name == "" {
		return default_value, nil
	}
	if default_value == os.Stdin {
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
		return ErrCDUFlags
	}

	return nil
}

func readLines(file *os.File) []string {
	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func uniqueLines(lines []string, options Options) []string {
	omap_result := omap.New[string, uint]()

	for _, line := range lines {
		exists := false
		for it := range omap_result.Iter() {
			if exists = myStrEq(it.Key, line, options); exists {
				omap_result.Add(it.Key, it.Value+1)
				break
			}
		}

		if !exists {
			omap_result.Add(line, 1)
		}
	}

	result := omapToStr(*omap_result, options)

	return result
}

func omapToStr(lines omap.OMap[string, uint], options Options) []string {
	result := make([]string, 0)

	for it := range lines.Iter() {
		line := it.Key
		count := it.Value
		if options.d {
			if count > 1 {
				result = append(result, line)
			}
		} else if options.u {
			if count == 1 {
				result = append(result, line)
			}
		} else {
			if options.c {
				line = fmt.Sprintf("%v %s", count, line)
			}
			result = append(result, line)
		}
	}

	return result
}

func myStrEq(s1, s2 string, o Options) bool {
	n := int(o.s)
	if len(s1) < n || len(s2) < n {
		return false
	}

	for i := uint(0); i < o.f; i++ {
		s1 = s1[strings.Index(s1, " ")+1:]
		s2 = s2[strings.Index(s2, " ")+1:]
	}

	if o.i {
		return strings.EqualFold(s1[n:], s2[n:])
	}

	return s1[n:] == s2[n:]
}

func writeLines(lines []string, file *os.File) {
	for _, line := range lines {
		fmt.Fprintln(file, line)
	}
}
