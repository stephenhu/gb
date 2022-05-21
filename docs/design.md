# design

first there needs to be a generic crawler which can get a set of links which ultimately contain the images that need to be downloaded, each site could have different ways of listing which requires a bit of pattern matching.  the output will be a csv file which will
feed the image downloader.

`gb crawl --pattern chapter`

after the links are gained, there needs to be a generic way to collect image links.