package cmd

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/PuerkitoBio/goquery"

)


var (

	imageCmd = &cobra.Command{
		Use: "image",
		Short: "Download images",
		Long: "Iterates over a page and finds all images for download.",
		Run: func(cmd *cobra.Command, args []string) {
			crawlImages()
		},
	}

)


var images = map[string]bool{}


func init() {

	imageCmd.Flags().StringVarP(&fFile, "file", "f", DEFAULT_FILE_NAME,
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
		io.Copy(f, res.Body)
	}

} // download


func parseImages(page string) {

	res, err := http.Get(page)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	dir := createChapter(page)

	doc, err := goquery.NewDocumentFromReader(res.Body)

  if err != nil {
		log.Println(err)
  } else {

    doc.Find(TAG_IMG).Each(func(index int, item *goquery.Selection) {

      l, _ := item.Attr(ATTR_DATA_SRC)

			download(l, dir)
			//u := getUrl(l)

			//images[l] = true
			//images[fmt.Sprintf("%s://%s%s", page.Scheme, page.Host, u.Path)] = true
  
    })

	}

} // parseImages


func readFileLinks() {

	buf, err := ioutil.ReadFile(fFile)

	if err != nil {
		log.Println(err)
	}

	links := strings.Split(string(buf), NEWLINE)

	for _, l := range links {
		parseImages(l)
	}

} // readFileLinks


func crawlImages() {

	readFileLinks()

} // crawlImages