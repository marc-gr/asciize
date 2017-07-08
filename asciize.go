// package asciize provides methods and types to transform images to ASCII art.
package asciize

import (
	"fmt"
	"image"
	"image/color"

	"strings"

	"errors"

	"github.com/andybons/gogif"
	"github.com/aybabtme/rgbterm"
	"github.com/nfnt/resize"
)

type OutputFormat string

const (
	// FormatText represents plain text output format
	FormatText OutputFormat = "text"
	// FormatHTML represents HTML representation of the ASCII art
	FormatHTML OutputFormat = "html"

	defaultWidth   = 100
	defaultCharset = " .~:+=o*x^%#@$MW"
)

// Asciizer allows to transform any image into its
// ASCII art representation with a set of configurable options.
type Asciizer struct {
	width         uint
	outFmt        OutputFormat
	colored       bool
	charset       []byte
	invertCharset bool
}

// NewAsciizer initialize a new asciizer with the given options.
func NewAsciizer(opts ...Option) *Asciizer {
	a := &Asciizer{
		width:   defaultWidth,
		outFmt:  FormatText,
		charset: []byte(defaultCharset),
	}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// Asciize receives an image and transforms it to an ASCII art
// representation.
func (a *Asciizer) Asciize(m image.Image) (ascii string, err error) {
	if m == nil {
		return "", errors.New("no image provided")
	}

	mSize := m.Bounds().Size()

	w := a.width
	if w == 0 {
		w = uint(mSize.X)
	}
	scale := float64(w) / float64(mSize.X)
	h := uint(float64(mSize.Y) * scale / 2)

	rsM := resize.Resize(w, h, m, resize.NearestNeighbor)
	grayM := imageToGray(rsM, a.charset)
	for y := 0; y <= grayM.Bounds().Max.Y; y++ {
		for x := 0; x <= grayM.Bounds().Max.X; x++ {
			r, _, _, _ := grayM.At(x, y).RGBA()

			// TODO: find a better way of weighting the index
			i := (int(r) / len(a.charset)) % len(a.charset)

			if a.invertCharset {
				i = len(a.charset) - 1 - i
			}

			character := fmt.Sprintf("%c", a.charset[i])
			if a.colored {
				r, g, b, _ := rsM.At(x, y).RGBA()
				character = decorateCharacter(character, a.outFmt, uint8(r), uint8(g), uint8(b), a.colored)
			}
			ascii += character
		}

		switch a.outFmt {
		case FormatText:
			ascii += "\n"
		case FormatHTML:
			ascii += "<br \\>"
		}
	}

	return
}

// Option is a type defining functions that can change
// Asciizer behavior.
type Option func(*Asciizer)

// Format indicates what the format to output will be, can be
// "text" or "html".
func Format(f OutputFormat) Option {
	return func(a *Asciizer) {
		a.outFmt = f
	}
}

// Width sets the output target width.
func Width(w uint) Option {
	return func(a *Asciizer) {
		a.width = w
	}
}

// Charset defines what charset is being used to transform images.
func Charset(charset []byte) Option {
	return func(a *Asciizer) {
		a.charset = charset
	}
}

// Colored indicates the output will be colored in the
// specified format (ANSI for text or HTML).
func Colored(c bool) Option {
	return func(a *Asciizer) {
		a.colored = c
	}
}

// InvertCharset allows to enable or disable inverted charset.
// This can make the result clearer in some images.
func InvertCharset(i bool) Option {
	return func(a *Asciizer) {
		a.invertCharset = i
	}
}

// DefaultCharset returns the default charset " .~:+=o*x^%#@$MW".
func DefaultCharset() []byte {
	return []byte(defaultCharset)
}

// TODO: get rid of external dependencies to quantize and color

func imageToGray(m image.Image, cs []byte) image.Image {
	if _, ok := m.(*image.Gray); ok {
		return m // early return if image already in gray scale
	}

	// TODO: implement custom quantizer
	pm := image.NewPaletted(m.Bounds(), nil)
	quantizer := gogif.MedianCutQuantizer{NumColor: len(cs)}
	quantizer.Quantize(pm, pm.Bounds(), m, image.ZP)

	grayM := image.NewGray(m.Bounds())
	for y := 0; y < m.Bounds().Max.Y; y++ {
		for x := 0; x < pm.Bounds().Max.X; x++ {
			oldPixel := pm.At(x, y)
			pixel := color.GrayModel.Convert(oldPixel)
			grayM.Set(x, y, pixel)
		}
	}
	return grayM
}

func decorateCharacter(c string, f OutputFormat, r, g, b uint8, colored bool) string {
	// TODO: implement custom ANSI coloring
	switch f {
	case FormatText:
		if colored {
			return rgbterm.FgString(c, r, g, b)
		}
		return c
	case FormatHTML:
		c = strings.Replace(c, " ", "&nbsp;", -1)
		color := ""
		if colored {
			color = fmt.Sprintf("color: #%2x%2x%2x;", r, g, b)
		}
		return fmt.Sprintf("<span style=\"font-family: 'Lucida Console', Monaco, monospace; %s\">%s</span>", color, c)
	}
	return c
}
