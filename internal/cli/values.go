package cli

import "strings"

type Put []string

func (i *Put) String() string {
	return strings.Join(*i, " ")
}

func (i *Put) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *Put) Type() string {
	return "string"
}
