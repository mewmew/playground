package str

import "strings"

func IndexAfter(s, sep string) (pos int) {
   pos = strings.Index(s, sep)
   if pos == -1 {
      return pos
   }
   pos += len(sep)
   return pos
}
