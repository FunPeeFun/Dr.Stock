package fman

import (
	"../buffer"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)


func Init() {
	readJson("/Users/neyzoter/code/golang/Dr.Stock/auto-delivery/buffer/config.json")
	printJsonConfig()
}

func writeJson()  {

}

func readJson(path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	buffer.StockConfigLock.Lock()
	err = json.Unmarshal(byteValue, &buffer.StockConfig)
	buffer.StockConfigLock.Unlock()
	if err != nil {
		log.Fatal(err)
	}
}

func printJsonConfig() {
	res, err := json.MarshalIndent(buffer.StockConfig, "", "  ")
	if err == nil {
		fmt.Printf("%s\n", res)
	} else {
		fmt.Println("marshal error")
	}
}