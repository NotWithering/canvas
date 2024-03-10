// Canvas is a simple go package that provides functions to help you create and display canvases in the terminal.

// This is an implementation of functions and types that help you convert
// 2-dimensional arrays of color into strings using different types of
// subcharacter-layering to achieve the resolution the user desires.

// Copyright (c) 2016-2020: Jim Philip and licensed under the MIT license.
// For more information see https://github.com/notwithering/canvas/blob/master/LICENSE

package canvas

import (
	"fmt"
	"image"
	"image/color"
	"strings"
)

// Canvas is a two-dimensional array of colors.
type Canvas [][]color.Color

// TextHalfBlock renders the canvas using half block subcharacter-layering into a string.
func TextHalfBlock(canvas Canvas) string {
	var final string

	var x, y int
	var width, height = len(canvas), len(canvas[0])

	for y = 0; y < height; y += 2 {
		for x = 0; x < width; x++ {
			var top color.Color = canvas[x][y]
			var bot color.Color

			if y+1 < height {
				bot = canvas[x][y+1]
			} else {
				bot = color.Transparent
			}

			var topCol color.RGBA = color.RGBAModel.Convert(top).(color.RGBA)
			var botCol color.RGBA = color.RGBAModel.Convert(bot).(color.RGBA)

			if topCol.A != 0 && botCol.A != 0 {
				bg := fmt.Sprintf("48;2;%d;%d;%d", topCol.R, topCol.G, topCol.B)
				fg := fmt.Sprintf("38;2;%d;%d;%d", botCol.R, botCol.G, botCol.B)

				codes := fmt.Sprintf("\x1b[%s;%sm", bg, fg)

				final += fmt.Sprintf("%s▄", codes)
			} else if topCol.A != 0 && botCol.A == 0 {
				fg := fmt.Sprintf("38;2;%d;%d;%d", topCol.R, topCol.G, topCol.B)

				codes := fmt.Sprintf("\x1b[%sm", fg)

				final += fmt.Sprintf("%s▀", codes)
			} else if topCol.A == 0 && botCol.A != 0 {
				fg := fmt.Sprintf("38;2;%d;%d;%d", botCol.R, botCol.G, botCol.B)

				codes := fmt.Sprintf("\x1b[%sm", fg)

				final += fmt.Sprintf("%s▄", codes)
			} else if topCol.A == 0 && botCol.A == 0 {
				final += " "
			}
			final += "\x1b[0m"
		}
		final += "\r\n"
	}

	final = strings.TrimRight(final, "\r\n")
	return final
}

// ImageCanvas takes an image and converst it to a canvas
func ImageCanvas(img image.Image) Canvas {
	var canv = New(img.Bounds().Max.X, img.Bounds().Max.Y)

	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			canv[x][y] = color.RGBAModel.Convert(img.At(x, y))
		}
	}

	return canv
}

// Fill fills the entire Canvas with the specified color.
func Fill(canvas Canvas, color color.Color) Canvas {
	var x, y int
	var width, height = len(canvas), len(canvas[x])
	for y = 0; y < height; y++ {
		for x = 0; x < width; x++ {
			canvas[x][y] = color
		}
	}
	return canvas
}

// New creates a new Canvas with the specified width and height.
func New(width, height int) Canvas {
	canv := make(Canvas, width)
	for i := range canv {
		canv[i] = make([]color.Color, height)
	}
	return canv
}

// Clone takes the original canvas and returns the copy of it.
func Clone(original Canvas) Canvas {
	clone := make(Canvas, len(original))
	for i := range original {
		clone[i] = make([]color.Color, len(original[i]))
		copy(clone[i], original[i])
	}
	return clone
}

// Deprecated: Text renders the canvas using half block subcharacter-layering into a string.
// Please use TextHalfBlock instead.
func Text(canvas Canvas) string {
	return TextHalfBlock(canvas)
}

// Deprecated: ImageText renders the image using half block subcharacter-layering into a string.
// Please use ImageCanvas and TextHalfBlock instead.
func ImageText(img image.Image, width, height uint) string {
	return TextHalfBlock(ImageCanvas(img))
}
