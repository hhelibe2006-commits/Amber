package value

import "strings"

type PathList []string

func (i *PathList) String() string {
	return strings.Join(*i, " ")
}

func (i *PathList) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *PathList) Type() string {
	return "string"
}
