package main

import (
	//"bytes"
	//"encoding/csv"
	//"errors"
	//"fmt"
	//"io"
	//"io/ioutil"
	"os"
	//"path"
	//"path/filepath"
	//"strings"
	"github.com/seythi/cDataImporter/market6"
	"time"
)

var (
	fileDate time.Time = time.Now()
)

func main() {
	//starts off the individual datatypes
	//interfaces with GUI when the time comes
	initFolder()
	market6.ReadMaster(fileDate)
}

func initFolder() bool {
	//MARKET6 folder init
	if _, err := os.Stat("./Market6"); os.IsNotExist(err) {
		os.Mkdir("./Market6", os.ModePerm)
	}
	if _, err := os.Stat("./Market6/archive"); os.IsNotExist(err) {
		os.Mkdir("./Market6/archive", os.ModePerm)
	}
	if _, err := os.Stat("./Market6/output"); os.IsNotExist(err) {
		os.Mkdir("./Market6/output", os.ModePerm)
	}
	if _, err := os.Stat("./Market6/output/archive"); os.IsNotExist(err) {
		os.Mkdir("./Market6/output/archive", os.ModePerm)
	}
	//PETERSON folder init
	if _, err := os.Stat("./Peterson"); os.IsNotExist(err) {
		os.Mkdir("./Peterson", os.ModePerm)
	}
	if _, err := os.Stat("./Peterson/archive"); os.IsNotExist(err) {
		os.Mkdir("./Peterson/archive", os.ModePerm)
	}
	if _, err := os.Stat("./Peterson/output"); os.IsNotExist(err) {
		os.Mkdir("./Peterson/output", os.ModePerm)
	}
	if _, err := os.Stat("./Peterson/output/archive"); os.IsNotExist(err) {
		os.Mkdir("./Peterson/output/archive", os.ModePerm)
	}
	//KEHE folder init
	if _, err := os.Stat("./KeHe"); os.IsNotExist(err) {
		os.Mkdir("./KeHe", os.ModePerm)
	}
	if _, err := os.Stat("./KeHe/archive"); os.IsNotExist(err) {
		os.Mkdir("./KeHe/archive", os.ModePerm)
	}
	if _, err := os.Stat("./KeHe/output"); os.IsNotExist(err) {
		os.Mkdir("./KeHe/output", os.ModePerm)
	}
	if _, err := os.Stat("./KeHe/output/archive"); os.IsNotExist(err) {
		os.Mkdir("./KeHe/output/archive", os.ModePerm)
	}
	//WF folder init
	if _, err := os.Stat("./WholeFoods"); os.IsNotExist(err) {
		os.Mkdir("./WholeFoods", os.ModePerm)
	}
	if _, err := os.Stat("./WholeFoods/archive"); os.IsNotExist(err) {
		os.Mkdir("./WholeFoods/archive", os.ModePerm)
	}
	if _, err := os.Stat("./WholeFoods/output"); os.IsNotExist(err) {
		os.Mkdir("./WholeFoods/output", os.ModePerm)
	}
	if _, err := os.Stat("./WholeFoods/output/archive"); os.IsNotExist(err) {
		os.Mkdir("./WholeFoods/output/archive", os.ModePerm)
	}
	return true
}
