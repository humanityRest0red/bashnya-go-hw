package myunique

import (
	"bufio"
	"flag"
	"fmt"
	"omap"
	"os"
	"strings"
)

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

	uniq_lines := uniqueLines(&options, input)
	Output(options, uniq_lines, output)

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
		return fmt.Errorf("use only one of c, d, u")
	}

	return nil
}

func myStrEq(s1, s2 string, o *Options) bool {
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

func uniqueLines(options *Options, file *os.File) omap.OMap[string, uint] {
	data := omap.New[string, uint]()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		exists := false
		for it := range data.Iter() {
			if exists = myStrEq(it.Key, line, options); exists {
				data.Add(it.Key, it.Value+1)
				break
			}
		}

		if !exists {
			data.Add(line, 1)
		}
	}

	return *data
}

func Output(options Options, lines omap.OMap[string, uint], file *os.File) {
	for it := range lines.Iter() {
		line := it.Key
		count := it.Value
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
