package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/PuerkitoBio/goquery"

)


var (

	downloadCmd = &cobra.Command{
		Use: "download",
		Short: "Download images",
		Long: "Iterates over a page and finds all images for download.",
		Run: func(cmd *cobra.Command, args []string) {
			crawlImages()
		},
	}

)


var images = map[string]bool{}

var (
	counter 				int64
  pages           int64
)



func init() {

	downloadCmd.Flags().StringVarP(&fFile, "file", "f", DEFAULT_FILE_NAME,
    "Filename to read URLs")

} // init


func createChapter(page string) string {

	if len(page) == 0 {
		panic(nil)
	}

	u := getUrl(page)

	str := filepath.Base(u.Path)

  title := strings.Split(str, CHAPTER)

	dir := filepath.Join(title[0], title[len(title) - 1])

	os.MkdirAll(dir, 0644)

	return dir
	
} // createChapter


func download(location string, dir string) {

	res, err := http.Get(location)

	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	u := getUrl(location)

	name := filepath.Base(u.Path)
	
	f, err := os.Create(filepath.Join(dir, name))	

	if err != nil {
		log.Println(err)
	} else {

		b, err := io.Copy(f, res.Body)

		if err != nil {
			log.Println(err)
		}

		counter = counter + b

		pages = pages + 1

	}

} // download


func parseImages(page string) {

	res, err := http.Get(page)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	dir := createChapter(page)

	color.Green("Directory created: " + dir)

	doc, err := goquery.NewDocumentFromReader(res.Body)

  if err != nil {
		log.Println(err)
  } else {

    doc.Find(TAG_IMG).Each(func(index int, item *goquery.Selection) {

      l, _ := item.Attr(ATTR_DATA_SRC)

			color.Green("Downloading: " + l)

			download(l, dir)
  
    })

	}

} // parseImages


func crawlImages() {

	buf, err := ioutil.ReadFile(fFile)

	if err != nil {
		log.Println(err)
	}

	links := strings.Split(string(buf), NEWLINE)

	start := time.Now()

	for _, l := range links {

		
		if len(l) > 0 && l[0] != COMMENT {
			
			color.Blue("Parsing: " + l)
			parseImages(l)

		}

	}

	color.Blue("\nDownload summary:")
	color.Blue("\tTotal bytes: %d\n", counter)
	color.Blue("\tTotal pages: %d\n", pages)
  color.Blue(fmt.Sprintf("\tDuration: %.0fs", time.Since(start).Seconds()))

} // crawlImages
