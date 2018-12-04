// Collection of shared methods used by other solution files
package shared

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// GetInputPath takes a day integer and returns the path to that days
// input.txt file.
func GetInputPath(d int) string {
	p := ""
	if d < 10 {
		p = fmt.Sprintf("./day_0%d/input.txt", d)
	} else {
		p = fmt.Sprintf("./day_%d/input.txt", d)
	}
	return p
}

// ReadInputLines takes an abs path to an input.txt file and returns the
// lines in a []string.
func ReadInputLines(path string) []string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}
