package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

)


var (

	fPrefix						string

	fixCmd = &cobra.Command{
		Use: "fix [PATH]",
		Short: "fix site",
		Long: "Fix PDF filenames",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			fixNames(args[0])

		},
	}

)


func init() {

	fixCmd.Flags().StringVarP(&fPrefix, "prefix", "p", "",
	  "Filename prefix filter pattern")

} // init


func fixNames(dir string) {

	files, err := os.ReadDir(dir)

	if err != nil {
		log.Println(err)
	} else {

		for _, f := range files {

			if !f.IsDir() {

				if filepath.Ext(f.Name()) == EXT_PDF {

					c := strings.TrimPrefix(strings.TrimSuffix(f.Name(),
					  filepath.Ext(f.Name())), fPrefix)

					switch(len(c)) {
					case 1:
						
						err := os.Rename(filepath.Join(dir, f.Name()),
						  filepath.Join(dir, fmt.Sprintf("%s000%s.pdf", fPrefix, c)))

						if err != nil {
							log.Println(err)
						}

					case 2:
						
						err := os.Rename(filepath.Join(dir, f.Name()),
						  filepath.Join(fmt.Sprintf("%s00%s.pdf", fPrefix, c)))

						if err != nil {
							log.Println(err)
						}

					case 3:

						err := os.Rename(filepath.Join(dir, f.Name()),
						  filepath.Join(dir, fmt.Sprintf("%s0%s.pdf", fPrefix, c)))

						if err != nil {
							log.Println(err)
						}

					default:
						fmt.Println("Unknown chapter length, skipping...")
					}

				}

			}

		}

	}

} // fixNames
