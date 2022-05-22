# gb

gb (golden boy) is a command line tool that aides in the download of media such as images, videos, and audio files.  there are essentially two services that it provides, the ability to crawl a site for links and the ability to download sets of images from those links.

since every site does things differently in terms of organizing content, this is supposed to be a generic tool that allows for a high degree of customization.

## dependencies

* golang 1.17.x+
* run on linux and windows
* github.com/PuerkitoBio/goquery (Excellent for parsing html pages)
* github.com/spf13/cobra (Industry standard for creating command line tools)

## setup

* `go build`
* `gb help`

## Usage

1. `gb crawl [URL] -p chapters # crawls URL for chapters, stores list in .chapters file`
  1. by default all links are stored to `.chapters` in the current directory, this is a human readable file, you can make changes and remove entries
  1. use the `-e` flag for URLs to exclude from crawling, this should be a comma delimited list
1. `gb image # crawls previously defined links for image files and downloads these`\
  1. this actually creates directories based on the URL path
