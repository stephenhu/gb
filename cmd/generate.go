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

	//"github.com/disintegration/imaging"
	"github.com/signintech/gopdf"
	"github.com/spf13/cobra"

)


var (

	fPdfFile      string
	fDir 					string
	fSubDir       bool

	generateCmd = &cobra.Command{
		Use: "generate [DIR]",
		Short: "Generate pdf from images",
		Long: "Creates pdf with multiple jpg images.",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			generatePdf(args[0])
		},
	}

)


func init() {

	generateCmd.Flags().StringVarP(&fPdfFile, "file", "f", DEFAULT_PDF_NAME,
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

		file := filepath.Join(dir, f)

		log.Println(file)
		if filepath.Ext(file) == EXT_JPG {
			
			pdf.AddPage()

			fn, err := os.Open(file)

			if err != nil {
				log.Println(err)
			} else {

				img, _, err := image.Decode(fn)

				if err != nil {
					log.Println(err)
				} else {

					ih, err := gopdf.ImageHolderByReader(fn)

					if err != nil {
						log.Println(err)
					} else {
						
						bounds := img.Bounds()
	
						if bounds.Max.X > bounds.Max.Y {
	
							err := pdf.ImageByHolderWithOptions(ih, gopdf.ImageOptions{
								DegreeAngle: 90,
								X: float64(bounds.Max.X),
								Y: float64(bounds.Max.Y),
								Rect: &gopdf.Rect{W: float64(bounds.Max.X), H: float64(bounds.Max.Y)},
							})
		
							if err != nil {
								log.Println(err)
							}
	
						} else {
	
							err := pdf.ImageByHolderWithOptions(ih, gopdf.ImageOptions{
								X: float64(bounds.Max.X),
								Y: float64(bounds.Max.Y),
								Rect: &gopdf.Rect{W: float64(bounds.Max.X), H: float64(bounds.Max.Y)},
							})
		
							if err != nil {
								log.Println(err)
							}
	
						}
	
						//pdf.Image(file, 0, 0, &gopdf.Rect{W: float64(bounds.Max.X),
							//H: float64(bounds.Max.Y)})
						
	
					}
	
				}
	
			}

			//pdf.Image(file, 0, 0, &gopdf.Rect{W: 572, H: 800})
			//pdf.Image(file, 0, 0, nil)

		}

	}

	pdf.WritePdf(fPdfFile)
	
} // generateBook


func generatePdf(dir string) {

	if len(dir) == 0 {
		log.Fatal(ERR_EMPTY_DIR)
	}
	_, err := os.Stat(dir)

	if err != nil || os.IsNotExist(err) {
		log.Fatal(err)
	} else {

		files, err := os.ReadDir(dir)

		if err != nil {
			log.Fatal(err)
		} else {

			if fSubDir {

				for _, file := range files {
				
					if fSubDir && file.IsDir() {
						generateBook(filepath.Join(dir, file.Name()))
					}

				}
	
			} else {
				generateBook(dir)
			}

		}

	}


} // generatePdf
