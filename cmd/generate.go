package cmd

import (
	//"io"
	//"io/ioutil
	"log"
	"os"
	"path/filepath"
	"strconv"
	//"strings"

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


func sortFiles(dir string) []string {

	//var sorted []string

	files, err := os.ReadDir(dir)

	if err != nil {
		
		log.Println(err)
		return nil

	} else {

		for _, f := range files {

			n := filepath.Base(f.Name())

			log.Println(n)
		
			i, err := strconv.Atoi(n)
	
			if err != nil {
				
				log.Println(err)
				return nil
	
			} else {
	
				if i % 10 == 0 {
	
				}
	
				return nil
	
			}
			
		}

		return nil

	}

} // sortFiles


func generateBook(dir string) {

	files := sortFiles(filepath.Join(fDir, dir))
	

	pdf := gopdf.GoPdf{}

	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
	})

	for _, f := range files {

		file := filepath.Join(fDir, f)

		if filepath.Ext(file) == EXT_JPG {
			log.Println(f)
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
