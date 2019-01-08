package kehe

import (
	"bytes"
	//"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"bufio"
	"os"
	"path/filepath"
	//"strings"
	"time"
)

var (
	workingName = "./market6/output/active.csv"
	hasHeaders  = false
)

func ReadMaster(fileDate time.Time) (fileCT, recordCT int) {
	workingFile, err := os.Create(workingName)
	check(err)
	writeTo := bufio.NewWriter(workingFile)
	defer workingFile.Close()
	defer writeTo.Flush()
	filepath.Walk("C:\\workspace\\Go\\bin\\Market6", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Printf("nodir: %+v \n", info.Name())
			if info.Name() != "Market6" {return filepath.SkipDir
				}else {return nil
				}
		}
		fmt.Printf("f: %s\n",path)
		fRecordCT, body := parseFile(path, writeTo)
		fmt.Printf("%s", body)
		fileCT++
		recordCT += fRecordCT
		return nil
	})
	return
}
func parseFile(fPath string, outFile *bufio.Writer) (fRecordCT int) {
	cont, err := ioutil.ReadFile(fPath)
	check(err)
	reader := bufio.NewReader(bytes.NewReader(cont))
	obody, ispfx, err := reader.ReadLine()
	for err != io.EOF{
		//if(ispfx){panic("ispfx true")}
		check(err)
		//process line, in this case cut off headers
		if !headerRead{ //has not seen header this file
			headerRead = bytes.Contains(obody, []byte("textbox")) //current line matches header pattern
			if !headerRead || hasheaders{ //hasnt read the header from this file yet  OR header written already
				continue
			}
		} 
			_, err2 := outFile.Write(obody)
			check(err2)
			_, err2 = outFile.WriteRune('\x0a')
			check(err2)
			fRecordCT++
		

		//body = fmt.Sprintf("%s", obody) //[]byte -> string
		


		obody, ispfx, err = reader.ReadLine()
	}
	return fRecordCT
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