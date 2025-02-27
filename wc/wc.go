package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

func main() {
	var t_lc, t_wc, t_cc int

	for _, fname := range os.Args[1:] {
		var lc, wc, cc int

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc += len(s)
			lc ++
		}

		fmt.Printf("%7d %7d %7d %s\n",
				   lc, wc, cc, fname)
		file.Close()
		t_lc += lc
		t_wc += wc
		t_cc += cc
	}

	if len(os.Args) > 2 { // prg name is at 0 position
		fmt.Printf("%7d %7d %7d %s\n",
				   t_lc, t_wc, t_cc, "total")

	}
}
