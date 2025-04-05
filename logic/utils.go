package logic

import "strings"

func JoinWithSpace(slc []string) string {
	return strings.Join(slc, " ")
}

func updatePrefix(prefix *string, next string) {
	newPrefix := strings.Fields(*prefix)
	if len(newPrefix) == 1 {
		*prefix = next
	} else {
	 	*prefix = JoinWithSpace(newPrefix[1:])+" "+next
	}
	
}