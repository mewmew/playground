// walls is a tool which downloads wallpapers from wallhaven.cc.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mewkiz/pkg/errutil"
	"github.com/mewmew/playground/archive/wallhaven"
)

func main() {
	// Parse command line parameters.
	var (
		// dir specifies the output directory.
		dir string
		// cats specifies the enabled wallpaper categories.
		cats wallhaven.Categories
		// purity specifies the enabled purity modes.
		purity wallhaven.Purity
		// res specifies the enabled screen resolutions.
		res wallhaven.Resolutions
		// ratios specifies the enabled aspect rations.
		ratios wallhaven.Ratios
		// sorting specifies the sorting method.
		sorting wallhaven.Sorting
		// order specifies the sorting order.
		order wallhaven.Order
		// page specifies the result page number.
		page wallhaven.Page
		// n specifies the largest number of wallpapers to download (0 =
		// infinite).
		n int
	)
	flag.StringVar(&dir, "o", "/tmp/walls", "Output directory")
	flag.Var(&cats, "cats", "Wallpaper categories (general,anime,people)")
	flag.Var(&purity, "purity", "Purity modes (SFW,sketchy,NSFW)")
	flag.Var(&res, "res", "Screen resolutions (1024x768,1280x800,1366x768,1280x960,1440x900,1600x900,1280x1024,1600x1200,1680x1050,1920x1080,1920x1200,2560x1440,2560x1600,3840x1080,5760x1080,3840x2160)")
	flag.Var(&ratios, "ratios", "Aspect ratios (4x3,5x4,16x9,16x10,21x9,32x9,48x9)")
	flag.Var(&sorting, "sorting", "Sorting method (relevance,random,date_added,views,favorites)")
	flag.Var(&order, "order", "Sorting order (asc,desc)")
	flag.Var(&page, "page", "Result page number")
	flag.IntVar(&n, "n", 0, "Download at most n wallpapers (0 = infinite)")
	flag.Parse()

	// Add command line options with non-zero values.
	var options []wallhaven.Option
	if cats != 0 {
		options = append(options, cats)
	}
	if purity != 0 {
		options = append(options, purity)
	}
	if res != 0 {
		options = append(options, res)
	}
	if ratios != 0 {
		options = append(options, ratios)
	}
	if sorting != 0 {
		options = append(options, sorting)
	}
	if order != 0 {
		options = append(options, order)
	}
	if page != 0 {
		options = append(options, page)
	}

	// Download wallpapers.
	query := strings.Join(flag.Args(), " ")
	if err := walls(query, dir, n, options); err != nil {
		log.Fatal(err)
	}
}

// walls downloads wallpapers based on the given search query, output dir, and
// search options. Download at most n wallpapers (0 = infinite).
func walls(query string, dir string, n int, options []wallhaven.Option) error {
	// Create output directory.
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errutil.Err(err)
	}

	// Download at most n wallpapers (0 = infinite).
	total := 0
	for {
		ids, err := wallhaven.Search(query, options...)
		if err != nil {
			return errutil.Err(err)
		}
		if len(ids) == 0 {
			return nil
		}
		for _, id := range ids {
			path, err := id.Download(dir)
			if err != nil {
				return errutil.Err(err)
			}
			fmt.Println(path)

			// Download at most n wallpapers.
			if total+1 == n {
				return nil
			}
			total++
		}
		// Turn to the next result page.
		options = nextPage(options)
	}
}

// nextPage returns the search options for querying the next page of search
// result.
func nextPage(options []wallhaven.Option) []wallhaven.Option {
	for i, option := range options {
		switch v := option.(type) {
		case wallhaven.Page:
			options[i] = v + 1
			return options
		case *wallhaven.Page:
			options[i] = *v + 1
			return options
		}
	}
	return append(options, wallhaven.Page(2))
}
