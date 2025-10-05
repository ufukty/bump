package labels

import (
	"regexp"
)

var extractor = regexp.MustCompile(`v([0-9]+)\.([0-9]+)\.([0-9]+).*`)

type Labels [4]int

func Parse(verstr string) (Labels, error) {

}
