# gb

gb (golden boy) is a command line tool that aides in the download of media such as images, videos, and audio files.  there are essentially three services that it provides, the ability to crawl a site for links, the ability to download sets of images from those links, and the ability
to convert sets of jpg images into a pdf file.

since every site does things differently in terms of organizing content, this is supposed to be a generic tool that allows for a high degree of customization.

## dependencies

* golang 1.17.x+
* https://github.com/PuerkitoBio/goquery (Excellent for parsing html pages)
* https://github.com/spf13/cobra (Industry standard for creating command line tools)

## platforms

* tested on linux and windows, probably works fine on mac os x


## setup

* `go build`
* `go install`
* `gb help`

## Usage

1. `gb crawl [URL] -p chapters # crawls URL for chapters, stores list in .chapters file`
   * by default all links are stored to `.chapters` in the current directory, this is a human readable file, you can make changes and remove entries, entries are separated by a newline
   * to comment out a line, use the pound sign (#) as the first character before the URL, lines that are commented out are ignored
   * use the `-e` flag to exclude URLs from being crawled, this should be a comma delimited list (no spaces or enclose comma delimited list in quotations to preserve spaces)
1. `gb download -f .chapters # crawls and downloads previously defined links for image files`
   * this actually creates directories based on the URL path and stores images respectively
1. `gb generate [DIR] -s -o tmp`
   * will use the provided directory path as the pdf filename, so you should use relative paths `i.e. gb generate gundam -s -o d:\manga`, do not use the following format: `gb generate c:\Users\manga\gundam -s -o d:\manga`
   * reads directory for sub-directories and generates pdf file from sub-directory images
   * `-s` walks sub-directories creating a pdf file per sub-directory, the default without `-s` will take the directory and combine jpg's into a single pdf file.
   * `-o` designates where to save the resultant pdf file
