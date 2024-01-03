# Sublay [![MIT License](https://img.shields.io/badge/License-MIT-a10b31)](https://github.com/NotWithering/sublay/blob/master/LICENSE)

**Sublay** is a simple go package that provides functions to help you create and display canvases in the terminal.

### Contents
-  [Example](#example)

## Example
<img src="screenshot.png"/>

```go
package main

import (
	"fmt"
	"image/color"

	"github.com/NotWithering/sublay"
)

func main() {
	myCanvas := sublay.New(7, 3)
	myCanvas = sublay.Fill(myCanvas, color.Transparent)

	myCanvas[0][0] = color.RGBA{255, 85, 85, 255}
	myCanvas[0][1] = color.RGBA{255, 85, 85, 255}
	myCanvas[0][2] = color.RGBA{255, 85, 85, 255}
	myCanvas[1][1] = color.RGBA{255, 85, 85, 255}
	myCanvas[2][0] = color.RGBA{255, 85, 85, 255}
	myCanvas[2][1] = color.RGBA{255, 85, 85, 255}
	myCanvas[2][2] = color.RGBA{255, 85, 85, 255}

	myCanvas[4][0] = color.RGBA{255, 85, 85, 255}
	myCanvas[5][0] = color.RGBA{255, 85, 85, 255}
	myCanvas[6][0] = color.RGBA{255, 85, 85, 255}
	myCanvas[5][1] = color.RGBA{255, 85, 85, 255}
	myCanvas[4][2] = color.RGBA{255, 85, 85, 255}
	myCanvas[5][2] = color.RGBA{255, 85, 85, 255}
	myCanvas[6][2] = color.RGBA{255, 85, 85, 255}

	me := sublay.Text(myCanvas)
	fmt.Println(me)
}

```