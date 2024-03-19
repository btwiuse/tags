package tags

import "strings"

// heavily inspired by arrayFlags from
// https://stackoverflow.com/questions/28322997/how-to-get-a-list-of-values-into-a-flag-in-golang

type SpaceSeparatedStrings []string

func (i *SpaceSeparatedStrings) String() string {
	return strings.Join(*i, " ")
}

func (i *SpaceSeparatedStrings) Set(value string) error {
	if value == "" {
		return nil
	}
	parts := strings.Split(value, " ")
	*i = append(*i, parts...)
	return nil
}
