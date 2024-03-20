package tags

import "strings"

// heavily inspired by arrayFlags from
// https://stackoverflow.com/questions/28322997/how-to-get-a-list-of-values-into-a-flag-in-golang

type CommaSeparatedStrings []string

func (i *CommaSeparatedStrings) String() string {
	return strings.Join(*i, ",")
}

func (i *CommaSeparatedStrings) Set(value string) error {
	if value == "" {
		return nil
	}
	parts := strings.Split(value, ",")
	*i = append(*i, parts...)
	return nil
}
