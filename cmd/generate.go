package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/spf13/cobra"

)


var (

	fPdfFile      string
	fDir 					string
	fSubDir       bool
	fOut          string

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
	generateCmd.Flags().StringVarP(&fOut, "out", "o", DEFAULT_DIR,
    "Filename for pdf file")

} // init


func createBookName(dir string) string {

	fs := strings.ReplaceAll(dir, FORWARD_SLASH, "")

	name := fmt.Sprintf("%s%s", strings.ReplaceAll(fs, BACKSLASH, ""), EXT_PDF)

	if len(fOut) != 0 {
		return filepath.Join(fOut, name)
	} else {
		return name
	}

} // createBookName


func sortFiles(dir string) []string { 

	var sorted []string

	files, err := os.ReadDir(dir)

	if err != nil {
		
		log.Println(err)
		return nil

	} else {

		for i := 0; i < len(files); i++ {

			for _, f := range files {

				file := filepath.Join(dir, f.Name())

				if filepath.Ext(file) == EXT_JPG {

					n := strings.Replace(filepath.Base(f.Name()), filepath.Ext(f.Name()), "", 1)

					c, err := strconv.Atoi(n)

					if err != nil {
					
						log.Println(err)
						return nil
			
					} else {

						if i == c {
							sorted = append(sorted, file)						
						}

					}

				}

			}

		}

		return sorted

	}

} // sortFiles


func generateBook(dir string) {

	files := sortFiles(dir)	

	imp, _ := api.Import("form:A4, pos: c, s:1.0", pdfcpu.POINTS)

	name := createBookName(dir)

	color.Green("Creating pdf: " + name)

	api.ImportImagesFile(files, name, imp, nil)

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
