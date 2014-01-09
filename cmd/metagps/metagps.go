// metagps outputs the GPS coordinates of provided images.
package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	gps "github.com/kurtcc/goexifgps"
	"github.com/rwcarlsen/goexif/exif"
)

func main() {
	flag.Parse()
	for _, imgPath := range flag.Args() {
		geoCoord, err := getCoordinate(imgPath)
		if err != nil {
			continue
		}
		fmt.Println("path:", imgPath)
		fmt.Println("coord:", geoCoord)
		fmt.Println()
	}
}

// getCoordinate returns the GPS coordinate of an image. The information is
// stored in the image's EXIF data.
func getCoordinate(imgPath string) (coord *gps.GeoFields, err error) {
	fr, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}
	defer fr.Close()
	x, err := exif.Decode(fr)
	if err != nil {
		return nil, fmt.Errorf("exif.Decode: failed for %q; %s.", imgPath, err)
	}
	coord, err = gps.GetGPS(x)
	if err != nil {
		return nil, fmt.Errorf("gps.GetGPS: failed for %q; %s.", imgPath, err)
	}
	if math.IsNaN(float64(coord.Lat)) || math.IsNaN(float64(coord.Long)) {
		return nil, fmt.Errorf("getCoordinate: failed for %q; unable to locate lat and long in EXIF data.", imgPath)
	}
	return coord, nil
}
