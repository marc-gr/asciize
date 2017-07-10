package main

import (
	"image"
	"os"

	"log"
	"net/url"

	"net/http"

	"github.com/marc-gr/asciize"

	"flag"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// BuildVersion is the version of the binary. It is automatically filled by the build step.
var BuildVersion = "dev"

func main() {
	src, opts := parseFlags()
	a := asciize.NewAsciizer(opts...)

	var m image.Image
	// We check if it is a URL
	u, err := url.ParseRequestURI(src)
	if err == nil && u.Host != "" {
		m, err = imageFromURL(src)
	} else {
		m, err = imageFromFile(src)
	}

	if err != nil {
		log.Fatal(err)
	}

	ascii, err := a.Asciize(m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ascii)
}

func imageFromURL(src string) (image.Image, error) {
	r, err := http.Get(src)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	m, _, err := image.Decode(r.Body)
	return m, err
}

func imageFromFile(src string) (image.Image, error) {
	f, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}

	m, _, err := image.Decode(f)
	return m, err
}

func parseFlags() (src string, opts []asciize.Option) {
	flag.StringVar(&src, "src", "", "Define the source of the image. Can be a local file or a URL")

	f := flag.String("f", "text", "Output format. Can be \"text\" or \"html\"")
	w := flag.Uint("w", 100, "Target width")
	cs := flag.String("cs", string(asciize.DefaultCharset()), "Used to define a custom charset")
	c := flag.Bool("c", false, "If set to true the output will be colored (ANSI or HTML)")
	i := flag.Bool("i", false, "If set to true the charset will be reversed. Can improve results for some images")
	v := flag.Bool("v", false, "Shows the build version")

	flag.Parse()

	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *v {
		fmt.Printf("version: %s\n", BuildVersion)
		os.Exit(0)
	}

	if src == "" {
		fmt.Printf("\n\tmissing parameter: src parameter is required\n\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	switch asciize.OutputFormat(*f) {
	case asciize.FormatHTML:
		opts = append(opts, asciize.Format(asciize.FormatHTML))
	case "", asciize.FormatText:
		opts = append(opts, asciize.Format(asciize.FormatText))
	default:
		fmt.Printf("\n\tinvalid parameter: unknown format %q\n\n", *f)
		flag.PrintDefaults()
		os.Exit(1)
	}

	opts = append(
		opts,
		asciize.Width(*w),
		asciize.Charset([]byte(*cs)),
		asciize.Colored(*c),
		asciize.InvertCharset(*i),
	)

	return
}
