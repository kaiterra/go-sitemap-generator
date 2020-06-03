package stm

import "fmt"

var poolBuffer = NewBufferPool()

// BuilderError provides interface for it can confirm the error in some difference.
type BuilderError interface {
	error
	FullError() bool
}

// Builder provides interface for adds some kind of url sitemap.
type Builder interface {
	XMLContent() []byte
	Content() []byte
	Add(interface{}) BuilderError
	Write()
}

// SitemapURL provides generated xml interface.
type SitemapURL interface {
	XML() []byte
}

// Attrs defines for xml attribute.
type Attrs []interface{}

// Attr defines for xml attribute.
type Attr map[string]string

// URL stores a single URL in a sitemap, along with all its properties.  It's really a
// map[string]interface{} where each item in the slice is a key-value pair (that is, it must have exactly 2 elements).
type URL [][]interface{}

// URLJoinBy takes the value of the 'key' property and sets it to the concatenated values
// of all the values of keys listed in 'joins'.
func (u URL) URLJoinBy(key string, joins ...string) URL {
	var values []string
	for _, k := range joins {
		var vals interface{}
		for _, v := range u {
			if v[0] == k {
				vals = v[1]
				break
			}
		}
		values = append(values, fmt.Sprint(vals))
	}
	var index int
	var v []interface{}
	for index, v = range u {
		if v[0] == key {
			break
		}
	}
	u[index][1] = URLJoin("", values...)
	return u
}

// BungURLJoinBy is an in-place version of URLJoinBy.  I wonder what Bung means.
func (u *URL) BungURLJoinBy(key string, joins ...string) {
	*u = u.URLJoinBy(key, joins...)
}

// Get returns the value at the specified key, if it exists.  This is like map[string]interface{}'s
// operator [].
func (u URL) Get(key string) (interface{}, bool) {
	for _, pairs := range u {
		if pairs[0] == key {
			return pairs[1], true
		}
	}
	return nil, false
}

// type News map[string]interface{}
