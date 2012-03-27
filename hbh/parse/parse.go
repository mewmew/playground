package parse

import "errors"
import "exp/html"
import "fmt"
import "strings"
import "unicode"

type Search struct {
   Tag   string
   Id    string
   Class string
}

func Html(text string, search *Search) (err error) {
   doc, err := html.Parse(strings.NewReader(text))
   if err != nil {
      return err
   }
   /*
    *    Dump all if no search Tag, Id or Class is set.
    */
   if len(search.Tag) < 1 && len(search.Id) < 1 && len(search.Class) < 1 {
      err = dump(doc, 0)
      if err != nil {
         return err
      }
      return nil
   }
   err = search.node(doc, 0)
   if err != nil {
      return err
   }
   return nil
}

func (search *Search) node(n *html.Node, pad int) (err error) {
   switch n.Type {
   case html.ElementNode:
      err = search.elem(n, pad)
      if err != nil {
         if err == ErrNodeHandledByDump {
            return nil
         }
         return err
      }
   }
   pad += 3
   for _, c := range n.Child {
      err = search.node(c, pad)
      if err != nil {
         return err
      }
   }
   return nil
}

var ErrNodeHandledByDump = errors.New("node handled by dump.")

func (search *Search) elem(n *html.Node, pad int) (err error) {
   /*
    *    If search.Tag is set check the tag name.
    */
   if len(search.Tag) > 0 {
      if n.Data != search.Tag {
         return nil
      }
   }
   var hasClass bool
   var hasId bool
   for _, attr := range n.Attr {
      switch attr.Key {
      case "class":
         /*
          *    If search.Class is set check the class.
          */
         if len(search.Class) > 0 {
            if search.Class != attr.Val {
               return nil
            }
         }
         hasClass = true
      case "id":
         /*
          *    If search.Id is set check the id.
          */
         if len(search.Id) > 0 {
            if search.Id != attr.Val {
               return nil
            }
         }
         hasId = true
      }
   }
   /*
    *    Tag doesn't contain a Class which is required by search.
    */
   if len(search.Class) > 1 && hasClass == false {
      return nil
   }
   /*
    *    Tag doesn't contain an Id which is required by search.
    */
   if len(search.Id) > 1 && hasId == false {
      return nil
   }
   err = dump(n, pad)
   if err != nil {
      return err
   }
   return ErrNodeHandledByDump
}

func dump(n *html.Node, pad int) (err error) {
   switch n.Type {
   case html.ElementNode:
      fmt.Printf("%*s<%s>\n", pad, "", n.Data)
   case html.CommentNode:
      fmt.Printf("%*s<!-- %s -->\n", pad, "", n.Data)
   case html.TextNode:
      if isOnlySpace(n.Data) {
         break
      }
      fmt.Printf("%*s%s", pad, "", n.Data)
      if n.Data[len(n.Data) - 1] != '\n' {
         fmt.Printf("\n")
      }
   }
   pad += 3
   for _, c := range n.Child {
      err = dump(c, pad)
      if err != nil {
         return err
      }
   }
   return nil
}

func isOnlySpace(s string) bool {
   for _, r := range s {
      if !unicode.IsSpace(r) {
         return false
      }
   }
   return true
}
