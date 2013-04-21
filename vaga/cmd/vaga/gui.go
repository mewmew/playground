package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/mewkiz/pkg/goutil"
	"github.com/mewkiz/pkg/imgutil"
	"github.com/mewmew/playground/vaga/grid"
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/init"
)

// gui plays a game of tic-tac-toe. The input is taken from the mouse.
func gui() (err error) {
	err = loadResources()
	if err != nil {
		return err
	}

	// Open a new window.
	width := imgGrid.Bounds().Dx()
	height := imgGrid.Bounds().Dy()
	win, err := wde.NewWindow(width, height)
	if err != nil {
		return err
	}
	win.SetTitle("vaga")
	win.Show()
	defer win.Close()

	g := grid.NewGrid()

	// Start the renderer.
	render := make(chan bool)
	go Render(win, g, render)

	// Render the grid.
	render <- true

	// Handle input.
	Input(win, g, render)

	return nil
}

// Input handles user input.
func Input(win wde.Window, g *grid.Grid, render chan bool) {
	// turnB keeps track of whose turn it is. When true, it's player B's turn.
	var turnB bool

	// When end is true the game has ended.
	var end bool

	// mde handles mouse down events.
	mde := func(e wde.MouseDownEvent) {
		if e.Which == wde.LeftButton {
			// Left click places a marker.

			if end {
				// Game has already ended.
				return
			}

			// Translate pixel location to grid coordinate.
			col, row, err := screenToCoord(g, e.Where.X, e.Where.Y)
			if err != nil {
				///log.Println(err)
				return
			}

			// Place marker.
			player := markA
			if turnB {
				player = markB
			}
			err = g.Place(col, row, player)
			if err != nil {
				///log.Println(err)
				return
			}
			turnB = !turnB
			render <- true

			// Check if the game has ended.
			end = true
			markWin := g.Check()
			if markWin != grid.MarkNone {
				fmt.Printf("===> %q wins :)\n", markWin)
				return
			}

		loop:
			for col := 0; col < g.Width(); col++ {
				for row := 0; row < g.Height(); row++ {
					if g[col][row] == grid.MarkNone {
						end = false
						break loop
					}
				}
			}
			if end {
				fmt.Println("===> no winner :(")
			}

		} else if e.Which == wde.MiddleButton {
			// Middle click clears the grid.
			g.Clear()
			turnB = false
			end = false
			render <- true
		}
	}

	for e := range win.EventChan() {
		switch v := e.(type) {
		case wde.CloseEvent:
			return
		case wde.MouseDownEvent:
			mde(v)
		}
	}
}

// Render renders the grid each time a value is sent on the render channel.
func Render(win wde.Window, g *grid.Grid, render chan bool) {
	screen := win.Screen()
	bounds := screen.Bounds()
	grey := image.NewUniform(color.RGBA{R: 0xF9, G: 0xF9, B: 0xF9}) /// ### todo ### A: 0xFF?
	for {
		<-render

		// Draw grid background.
		draw.Draw(screen, bounds, grey, image.ZP, draw.Src)
		draw.Draw(screen, bounds, imgGrid, image.ZP, draw.Over)

		// Draw grid markers.
		var img image.Image
		for col := 0; col < g.Width(); col++ {
			for row := 0; row < g.Height(); row++ {
				switch g[col][row] {
				case grid.MarkO:
					img = imgO
				case grid.MarkX:
					img = imgX
				default:
					continue
				}
				pt := getSlotPt(col, row, g[col][row])
				rect := image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy())
				draw.Draw(screen, rect.Add(pt), img, image.ZP, draw.Over)
			}
		}

		// Flush screen updates.
		win.FlushImage(bounds)
	}
}

// imgGrid is the grid background image.
var imgGrid image.Image

// imgO is the 'o' marker image.
var imgO image.Image

// imgX is the 'x' marker image.
var imgX image.Image

// loadResources loads the images required to render the grid and the markers.
func loadResources() (err error) {
	dataDir, err := goutil.SrcDir("github.com/mewmew/playground/vaga/data")
	if err != nil {
		return err
	}
	imgGrid, err = imgutil.ReadFile(dataDir + "/grid.png")
	if err != nil {
		return err
	}
	imgO, err = imgutil.ReadFile(dataDir + "/o.png")
	if err != nil {
		return err
	}
	imgX, err = imgutil.ReadFile(dataDir + "/x.png")
	if err != nil {
		return err
	}
	return nil
}

// getSlotPt returns the image.Point of the slot at which a marker can be drawn.
func getSlotPt(col, row int, mark grid.Mark) (pt image.Point) {
	// Note: a regular grid would simplify this code a lot.

	var dx int
	var dy int
	switch mark {
	case grid.MarkO:
		dx = 15
		dy = 7
	case grid.MarkX:
		dx = 7
		dy = 18
	}

	var pts = [][]image.Point{
		[]image.Point{
			image.Pt(0+dx, 0+dy),   // 0, 0
			image.Pt(0+dx, 136+dy), // 0, 1
			image.Pt(0+dx, 270+dy), // 0, 2
		},
		[]image.Point{
			image.Pt(165+dx, 0+dy),   // 1, 0
			image.Pt(165+dx, 136+dy), // 1, 1
			image.Pt(165+dx, 270+dy), // 1, 2
		},
		[]image.Point{
			image.Pt(323+dx, 0+dy),   // 2, 0
			image.Pt(323+dx, 136+dy), // 2, 1
			image.Pt(323+dx, 270+dy), // 2, 2
		},
	}

	return pts[col][row]
}

// screenToCoord translates the provided pixel location to grid coordinates.
func screenToCoord(g *grid.Grid, x, y int) (col, row int, err error) {
	// Note: a regular grid would simplify this code a lot.

	var area = [][]image.Rectangle{
		[]image.Rectangle{
			image.Rect(0, 0, 138, 119),   // 0, 0
			image.Rect(0, 133, 146, 243), // 0, 1
			image.Rect(0, 244, 157, 409), // 0, 2
		},
		[]image.Rectangle{
			image.Rect(154, 0, 309, 116),   // 1, 0
			image.Rect(165, 136, 299, 245), // 1, 1
			image.Rect(158, 261, 322, 409), // 1, 2
		},
		[]image.Rectangle{
			image.Rect(310, 0, 499, 128),   // 2, 0
			image.Rect(317, 129, 499, 258), // 2, 1
			image.Rect(323, 270, 499, 409), // 2, 2
		},
	}

	pt := image.Pt(x, y)
	for col = 0; col < g.Width(); col++ {
		for row = 0; row < g.Height(); row++ {
			if pt.In(area[col][row]) {
				return col, row, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("screenToCoord: invalid slot (x=%d, y=%d).", x, y)
}
