package charset

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
	a := asciize.NewAsciizer(
		asciize.Width(50),
		asciize.Charset([]byte("|<.>?=+*&@±^")),
	)
	s, err := a.Asciize(img)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
