package main

import (
    "fmt"
    "github.com/xeipuuv/gojsonschema"
	  "io"
    "ioutil"
	  "net/http"
	  "os"
    "github.com/ghodss/yaml"
)

var schemaFile = "zarf.schema.json"
var schemaUrl = "https://raw.githubusercontent.com/defenseunicorns/zarf/v0.29.2/zarf.schema.json"

func main() {

    err := DownloadFile(schemaFile, schemaUrl)
    schemaLoader := gojsonschema.NewReferenceLoader(schemaFile)

    // TODO: get zarf.yaml and convert to json
    
    //Read in the zarf config yaml file
    


    //Load the converted zarf config into jsonschema
    documentLoader := gojsonschema.NewReferenceLoader("file:///home/me/document.json")

    result, err := gojsonschema.Validate(schemaLoader, documentLoader)
    if err != nil {
        panic(err.Error())
    }

    if result.Valid() {
        fmt.Printf("The document is valid\n")
    } else {
        fmt.Printf("The document is not valid. see errors :\n")
        for _, desc := range result.Errors() {
            fmt.Printf("- %s\n", desc)
        }
    }
}


func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

