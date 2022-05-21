package cmd

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)


const (
	TAG_A										= "a"
	TAG_IMG                 = "img"
)


const (
	ATTR_HREF               = "href"
	ATTR_SRC                = "src"
)


const (
  DEFAULT_FILE_NAME				= ".download"
	NEWLINE                 = "\n"
)


const (
	ERR_EMPTY_URL						= "Must provide a URL"
)


var (
	
	fPattern				string
	fFile						string

	crawlCmd = &cobra.Command{
		Use: "crawl [URL]",
		Short: "Crawler for a site",
		Long: "Crawler iterates over a site for page links",
		Args: cobra.ExactArgs(1),
		ValidArgs: []string{"URL"},
		Run: func(cmd *cobra.Command, args []string) {
			crawl(args[0])
		},
	}

)


var links = map[string]bool{}
var depth = map[string]bool{}


func init() {

	crawlCmd.Flags().StringVarP(&fPattern, "pattern", "p", "",
	  "URL search pattern")
	crawlCmd.Flags().StringVarP(&fFile, "file", "f", DEFAULT_FILE_NAME,
    "Filename to save URLs")

} // init


func getDomain(s string) string {

	u, err := url.Parse(s)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println(u)
	return u.Host

} // getDomain


func save() {

	f, err := os.Create(fFile)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	for k, _ := range links {

		_, err := f.WriteString(k + NEWLINE)

		if err != nil {
			
			log.Println(err)
			panic(err)

		}

	}

} // save


func crawl(location string) {

	res, err := http.Get(location)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	parseTags(res.Body, getDomain(location), TAG_A, ATTR_HREF)

	log.Println(links)
	log.Println(depth)

	save()

} // crawl


func parseTags(body io.Reader, location string, tag string, attr string) {

	doc, err := goquery.NewDocumentFromReader(body)

  if err != nil {
    log.Println(err)
  } else {

    doc.Find(tag).Each(func(index int, item *goquery.Selection) {

      l, _ := item.Attr(attr)

			if strings.Contains(l, fPattern) {
				links[l] = true
			} else if strings.Contains(l, location) {
				depth[l] = true
			}
  
    })

	}
	
} // parseTags
