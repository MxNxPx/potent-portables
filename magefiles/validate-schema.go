package main

import (
    "fmt"
    "github.com/xeipuuv/gojsonschema"
	//"io"
    "io/ioutil"
	//"net/http"
	//"os"
    "log"
    "encoding/json"
    "sigs.k8s.io/yaml"
)

var schemaFile = "zarf.schema.json"
var schemaUrl = "https://raw.githubusercontent.com/defenseunicorns/zarf/v0.29.2/zarf.schema.json"
var zarfConfigPath = "../app/zarf.yaml"

func main() {

    //Download zarf schame file, save to disk
    //   err := DownloadFile(schemaFile, schemaUrl)
    
    //Read in the zarf config yaml file
    zarfConfigFile, err := ioutil.ReadFile(zarfConfigPath)
    if err != nil {
        log.Fatalf("Unable to read the zarf schema: %v", err)
    }
    var zarfConfigYaml map[string]interface{}
    err = yaml.Unmarshal(zarfConfigFile, &zarfConfigYaml)
    //Read in the downloaded zarf schema file as byte
    // zarfSchemaRaw, err := os.Open(schemaFile)
    // if err != nil {
    //     log.Fatalf("Unable to read in zarf schema file: %v", err)
    // }
    // defer zarfSchemaRaw.Close()

    //Declare placeholder var for the json
    // var zarfSchema map[string]interface{}
    //Store file contents to var as JSON
    
    // byteValue, _ := ioutil.ReadAll(zarfSchemaRaw)
    // json.Unmarshal([]byte(byteValue), &zarfSchema)


    //Convert the loaded schema to json
    //fmt.Println(zarfConfigYaml)
    // for key, value := range zarfConfigYaml {
	// 	fmt.Printf("%s: %v\n", key, value)
	// }
    zarfConfigJson, err := json.Marshal(zarfConfigYaml)
    //fmt.Printf("JSON Data:\n%s\n", zarfConfigJson)
    // zarfConfig, err := yaml.YAMLToJSON(fmt.Sprint(zarfConfigYaml))
    // if err != nil {
    //     log.Fatalf("Unable to convert the zarf schema to json: %v", err)
    // }
    
    // var zarfConfig map[string]interface{}
    // json.Unmarshal([]byte(zarfConfigJson), &zarfConfig)

    // Load the zarf schema json into memory
    schemaLoader := gojsonschema.NewReferenceLoader(schemaUrl)
    //schemaLoader := gojsonschema.NewGoLoader(zarfSchema)
    //Load the converted zarf config into jsonschema
    documentLoader := gojsonschema.NewGoLoader(zarfConfigJson)

    //fmt.Println(json.Marshal(zarfConfig))
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


// func DownloadFile(filepath string, url string) error {

// 	// Get the data
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Create the file
// 	out, err := os.Create(filepath)
// 	if err != nil {
// 		return err
// 	}
// 	defer out.Close()

// 	// Write the body to file
// 	_, err = io.Copy(out, resp.Body)
// 	return err
// }

