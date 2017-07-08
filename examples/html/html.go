package html

import (
	"fmt"
	"image/png"
	"os"

	"github.com/marc-gr/asciize"
)

func main() {
	f, _ := os.Open("./sample.png")
	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	a := asciize.NewAsciizer(asciize.Format(asciize.FormatHTML))
	s, err := a.Asciize(img)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
