package main

import (
	/*
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
*/

  "github.com/stephenhu/gb/cmd"
	//"github.com/PuerkitoBio/goquery"

)

/*
var links = map[string]bool{}

var (
	page	= flag.String("page", "", "root download page")
	query	= flag.String("query", "", "query string")
)

func download(link string) {

	u, err := url.Parse(link)

	if err != nil {
		log.Println(err)
	} else {

		filename := path.Base(u.Path)

		dirs := strings.Split(path.Dir(u.Path), "/")

		dirname := dirs[len(dirs)-1]

		log.Println(filename)
		log.Println(dirname)

		_, err := os.Stat(dirname)

		if err != nil {

			if os.IsNotExist(err) {
			
				err := os.Mkdir(dirname, 0755)
	
				if err != nil {
					log.Println(err)
				}
	
			} else {
				log.Println(err)
			}
	
		} else {

			res, err := http.Get(link)

			if err != nil {
				log.Println(err)
			} else {

				defer res.Body.Close()

				f, err := os.Create(path.Join(dirname, filename))

				if err != nil {
					log.Println(err)
				} else {
					io.Copy(f, res.Body)
				}

			}
	
		}


	}

	
} // download


func getImages(search string) {

	for link := range links {

		res, err := http.Get(link)
		
		if err != nil {
			log.Println(err)
			continue
		} else {

			defer res.Body.Close()

			doc, err := goquery.NewDocumentFromReader(res.Body)

			if err != nil {
				log.Println(err)
			} else {
		
				doc.Find("img").Each(func(index int, item *goquery.Selection) {
		
					image, _ := item.Attr("src")
		
					if strings.Contains(image, search) {
						download(image)
					}
					
				})
		
			}

		}

	}
	
} // getImages



func main() {

	flag.Parse()
	
	if *page == "" {
		log.Fatal("Error: empty page")
	}

	if *query == "" {
		log.Fatal("Error: empty query string")
	}



	getLinks(res.Body, *query)
	log.Println(links)
	getImages(*query)

} // main

*/

func main() {

	cmd.Execute()

} // main
