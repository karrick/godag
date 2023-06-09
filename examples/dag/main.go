package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/karrick/godag"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	dag := godag.New()

	for scanner.Scan() {
		line := scanner.Text()

		index := strings.Index(line, "#")
		if index >= 0 {
			line = line[:index]
		}
		if line == "" {
			continue
		}
		fields := strings.Fields(line)

		dag.Insert(fields[0], fields[1:])
	}

	err := scanner.Err()
	bailWhenError(err, 1)

	ordered, err := dag.Order()
	bailWhenError(err, 2)

	fmt.Println(strings.Join(ordered, "\n"))
}

func bailWhenError(err error, code int) {
	if err == nil {
		return
	}
	program := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "%s: %s\n", program, err)
	os.Exit(code)
}
