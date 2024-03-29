package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
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
	ATTR_DATA_SRC           = "data-src"
	ATTR_SRC                = "src"
)


const (
	METHOD_GET							= "get"
)


const (
	USER_AGENT              = "User-Agent"
	MOZILLA_AGENT           = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 Safari/537.36"
	POSTMAN_AGENT           = "PostmanRuntime/7.28.4"
)


const (
	DEFAULT_DIR             = "."
	DEFAULT_FILE_NAME				= ".chapters"
	DEFAULT_PDF_NAME				= "book.pdf"
)


const (
	EXT_JPG									= ".jpg"
	EXT_PDF                 = ".pdf"
)


const (
	BACKSLASH               = "\\"
	CARRIAGE_RETURN         = "\r"
	CHAPTER                 = "-chapter-"
	DASH                    = "-"
	EMPTY_STR               = ""
	FORWARD_SLASH           = "/"
	NEWLINE                 = "\n"
)


const (
	COMMENT                 = '#'
)


const (
	ERR_EMPTY_URL						= "Error: Must provide a URL."
	ERR_EMPTY_DIR           = "Error: Must provide a directory name."
)


var (
	
	fPattern				string
	fFile						string
	fExclude        []string

	crawlCmd = &cobra.Command{
		Use: "crawl [URL]",
		Short: "Crawl site",
		Long: "Crawler iterates over a site for page links",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			loadExcludedMap()
			crawl(args[0])

		},
	}

)


var links 		= map[string]bool{}
var excluded	= map[string]bool{}


func init() {

	crawlCmd.Flags().StringVarP(&fPattern, "pattern", "p", "",
	  "URL search pattern")
	crawlCmd.Flags().StringVarP(&fFile, "file", "f", DEFAULT_FILE_NAME,
    "Filename to save URLs")
	crawlCmd.Flags().StringSliceVarP(&fExclude, "exclude", "e", []string{},
    "exclude URLs")

} // init


func loadExcludedMap() {

	for _, v := range fExclude {
		excluded[v] = true
	}

} // loadExcludedMap


func getUrl(s string) *url.URL {

	u, err := url.Parse(s)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	return u

} // getUrl


func exclude(l string) bool {

	_, ok := excluded[l]

	if ok {
		return true
	} else {
		return false
	}

} // exclude


func saveToFile() {

	f, err := os.Create(fFile)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	sortedLinks := []string{}

	for l, _ := range links {

		// TODO: optimize sort, right now it's O(N) * 2

		sortedLinks = append(sortedLinks, l)
		
	}

	sort.Strings(sortedLinks)

	for _, v := range sortedLinks {

		if !exclude(v) {

			_, err := f.WriteString(strings.TrimSpace(v) + NEWLINE)

			if err != nil {
				
				log.Println(err)
				panic(err)
	
			}
	
		}

	}

} // saveToFile	


func parseLinks(body io.Reader, page string) {

	doc, err := goquery.NewDocumentFromReader(body)

  if err != nil {
		log.Println(err)
  } else {

		p := getUrl(page)

    doc.Find(TAG_A).Each(func(index int, item *goquery.Selection) {
log.Println("got here")
      l, _ := item.Attr(ATTR_HREF)
log.Println(l)
			u := getUrl(l)

			if strings.Contains(l, fPattern) {
				links[fmt.Sprintf("%s://%s%s", p.Scheme, p.Host, u.Path)] = true
			}
  
    })

	}
	
} // parseLinks


func crawl(location string) {

	res, err := http.Get(location)

	log.Println(res)
	
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	parseLinks(res.Body, location)

	saveToFile()

} // crawl
