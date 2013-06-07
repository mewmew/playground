package page

import (
	"bytes"
	"fmt"
	"os"

	"code.google.com/p/go.net/html"
	"github.com/mewkiz/pkg/htmlutil"
	"github.com/mewkiz/pkg/httputil"
)

// download downloads the page, locates the relevant HTML nodes based on the CSS
// selector and hashes the content.
func (p *Page) download() (err error) {
	buf, err := httputil.Get(p.RawUrl)
	if err != nil {
		return err
	}
	doc, err := html.Parse(bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	p.digest.Reset()
	if p.sel != nil {
		// Create a hash based on all nodes that match the CSS selector.
		nodes := p.sel.MatchAll(doc)
		if len(nodes) < 1 {
			return fmt.Errorf("page.download: No nodes matches the CSS selector ('%s'). Page size: %d.", p.RawSel, len(buf))
		}
		for _, node := range nodes {
			htmlutil.Render(os.Stdout, node)
			fmt.Println()
			htmlutil.Render(p.digest, node)
		}
	} else {
		// Create a hash of the entire HTML page.
		htmlutil.Render(os.Stdout, doc)
		fmt.Println()
		htmlutil.Render(p.digest, doc)
	}
	return nil
}

func init() {
	httputil.SetClient(httputil.InsecureClient)
}
