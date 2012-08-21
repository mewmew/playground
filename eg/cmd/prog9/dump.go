package main

import "fmt"
import "io/ioutil"

const format = `<!DOCTYPE html>
<html>
   <head>
      <style>
         body
         {
            margin: 20px;
         }

         #game_grid
         {
            border: 1px solid #000000;
            background-color: #FFFFFF;
         }

         #game_grid td
         {
            border: 1px solid #EEEEEE;
            width: 40px;
            height: 40px;
            text-align: center;
         }
      </style>
   </head>
   <body>
      <table id='game_grid'>
%s
      </table>
   </body>
</html>

`

// dumpGrid dumps the grid to an HTML file.
func dumpGrid(grid string) (err error) {
	err = ioutil.WriteFile("grid.html", []byte(fmt.Sprintf(format, grid)), 0644)
	if err != nil {
		return err
	}
	return nil
}
