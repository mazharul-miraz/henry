package utility

import (
	"os"
	"io/ioutil"
	"strconv"
)

func DoesFileExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func GetCounter(counterpath string) (counter int){
	counter =0
	data, err := ioutil.ReadFile(counterpath)
	if (err == nil) {
		counter, err = strconv.Atoi(string(data))
		if err != nil {
			counter = 0
		}
	} else {
		counter = 0
	}
	return
}

func UpdateCounter(counterpath string, counter int){
	counter++
	ioutil.WriteFile(counterpath,[]byte(strconv.Itoa(counter)),0644)
}

func IntToString(i int) string{
	return strconv.Itoa(i)
}
