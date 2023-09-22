# learnings

## pros & cons

* If you don't know Go or Mage, ChatGPT is able to help with reasonable accuracy
* Dependencies are flexible
  * If not specified in a Target that contains other Targets, runs in parallel
  * If supplied in a Target that contains other Targets, then will [sequence](https://magefile.org/dependencies/#example-dependencies)
* Easy to split up the Magefile code into multiple files and it stitches it together
* Will not output the Target step stage by default without adding verbosity
  * `mage -v <TARGET>`
  * `export MAGEFILE_VERBOSE=1`

## useful info & links

* Mage license [Apache 2](https://github.com/magefile/mage/blob/master/LICENSE)
* Mage env vars: https://magefile.org/environment/
* Hugo uses Mage: https://github.com/gohugoio/hugo/blob/master/magefile.go
* Dagger uses Mage??: https://github.com/dagger/dagger/blob/main/internal/mage/magefile.go
* Awesome Mage repo: https://github.com/magefile/awesome-mage
* Best to use the ./magefiles directory to store all the code (bootstrap via: `mage -d magefiles -w . -init`)
* The Mage mascot - **Gary**...

![Gary, aka the Mage](https://magefile.org/images/gary.svg)
