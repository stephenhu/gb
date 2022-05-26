package cmd

import (
	//"io"
	//"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/signintech/gopdf"
	"github.com/spf13/cobra"

)


var (

	fDir 					string
	fSubDir       bool

	generateCmd = &cobra.Command{
		Use: "generate",
		Short: "Generate pdf from images",
		Long: "Creates pdf with multiple jpg images.",
		Run: func(cmd *cobra.Command, args []string) {
			generatePdf()
		},
	}

)


func init() {

	generateCmd.Flags().StringVarP(&fFile, "file", "f", DEFAULT_PDF_NAME,
    "Filename for pdf file")
	generateCmd.Flags().StringVarP(&fDir, "dir", "d", DEFAULT_DIR,
    "Filename for pdf file")
	generateCmd.Flags().BoolVarP(&fSubDir, "subdir", "s", true,
	  "Recurse over sub-directories")

} // init


func createChapte2r(page string) string {

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


func generateBook(dir string) {

	files, err := os.ReadDir(filepath.Join(fDir, dir))

	if err != nil {
		log.Println(err)
	}

	pdf := gopdf.GoPdf{}

	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
	})

	for _, f := range files {

		file := filepath.Join(fDir, f.Name())

		if !f.IsDir() && filepath.Ext(file) == EXT_JPG {
			log.Println(f.Name())
		}

	}

} // generateBook


func generatePdf() {

	_, err := os.Stat(fDir)

	if err != nil || os.IsNotExist(err) {
		log.Fatal(err)
	} else {

		files, err := os.ReadDir(fDir)

		if err != nil {
			log.Fatal(err)
		} else {

			if fSubDir {

				for _, file := range files {
				
					if fSubDir && file.IsDir() {
						generateBook(file.Name())
					}

				}
	
			} else {
				generateBook(fDir)
			}

/*				buf, err := ioutil.ReadFile(file.Name())

				if err != nil {
					log.Println(err)
				}*/

		}

	}


} // generatePdf
