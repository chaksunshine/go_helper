package str

import "regexp"

var (
	regexpAZHaving, _ = regexp.Compile("[^\\w+]")
)
