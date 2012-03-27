// Package str implements simple functions to manipulate strings, extending the
// strings package.
package str

import "strings"

// IndexAfter returns the index directly after the first instance of sep in s,
// or -1 if sep is not present in s.
func IndexAfter(s, sep string) (pos int) {
   pos = strings.Index(s, sep)
   if pos == -1 {
      return pos
   }
   pos += len(sep)
   return pos
}
