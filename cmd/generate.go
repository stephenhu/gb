package cmd

import (
	"image"
	//"io"
	//"io/ioutil
	"log"
	"os"
	"path/filepath"
	//"sort"
	"strconv"
	"strings"

	_ "image/jpeg"

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
	generateCmd.Flags().BoolVarP(&fSubDir, "subdir", "s", false,
	  "Recurse over sub-directories")

} // init


func sortFiles(dir string) []string {

	var sorted []string

	files, err := os.ReadDir(dir)

	if err != nil {
		
		log.Println(err)
		return nil

	} else {

		for i := 0; i < len(files); i++ {

			for _, f := range files {

				n := strings.Replace(filepath.Base(f.Name()), filepath.Ext(f.Name()), "", 1)

				c, err := strconv.Atoi(n)

				if err != nil {
				
					log.Println(err)
					return nil
		
				} else {

					if i == c {
						sorted = append(sorted, f.Name())						
					}
				}

			}

		}

/*
		for _, f := range files {

			n := strings.Replace(filepath.Base(f.Name()), filepath.Ext(f.Name()), "", 1)

			log.Println(n)
		
			i, err := strconv.Atoi(n)
	
			if err != nil {
				
				log.Println(err)
				return nil
	
			} else {
	
				if i % 10 == 0 {
					
				}
	
			}
			
		}
*/

		return sorted

	}

} // sortFiles


func generateBook(dir string) {

	files := sortFiles(dir)	

	pdf := gopdf.GoPdf{}

	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
	})

	for _, f := range files {

		file := filepath.Join(fDir, f)

		if filepath.Ext(file) == EXT_JPG {
			
			pdf.AddPage()

			fn, err := os.Open(file)

			if err != nil {
				log.Println(err)
			} else {

				x, _, err := image.Decode(fn)

				if err != nil {
					log.Println(err)
				} else {
					
					bounds := x.Bounds()

					pdf.Image(file, 0, 0, &gopdf.Rect{W: float64(bounds.Max.X/2), H: float64(bounds.Max.Y/2)})

				}
	
			}

			//pdf.Image(file, 0, 0, &gopdf.Rect{W: 572, H: 800})
			//pdf.Image(file, 0, 0, nil)

		}

	}

	pdf.WritePdf(fFile)
	
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
