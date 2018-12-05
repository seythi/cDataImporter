package market6

import (
	//"bytes"
	//"encoding/csv"
	//"errors"
	"fmt"
	//"io"
	//"io/ioutil"
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	workingName = "./market6/output/active.csv"
	hasHeaders  = false
)

func ReadMaster(fileDate time.Time) (fileCT, recordCT int) {
	workingFile, err := os.Create(workingName)
	writeTo := bufio.NewWriter(workingFile)
	check(err)
	defer working.Close()
	filepath.Walk("./market6", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() /*&& info.Name() == subDirToSkip*/ {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		fmt.Println(path)
		fRecordCT := parseFile(path)
		fileCT++
		recordCT += fRecordCT
		return nil
	})
	return
}
func parseFile(fPath string) (fRecordCT int, body string) {
	cont, err := ioutil.ReadFile("fPath")
	check(err)
	reader := NewReader(cont)
}

func check(e error) error {
	if e == io.EOF {
		return e
	}
	if e != nil {
		panic(e)
	}
	return errors.New("this should never ever return what is happening")
}
