package cuslog

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const (
	flag = log.LstdFlags | log.Lshortfile
)

var (
	todate   = time.Now().Format("2006-01-02")
	filesArr = [3]string{"info", "err", "warning"}
)

func init() {
	for i := 0; i < len(filesArr); i++ {
		if err := os.MkdirAll("./log/"+filesArr[i], 0775); err != nil {
			log.Panic("Can't create files :", err)
		}
	}

}
func Info(v ...interface{}) {

	file, err := openLogFile("./log/info/" + fmt.Sprint(todate) + ".log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	infoLogger := log.New(os.Stdout, "INFO :", flag)
	infoLogger.SetOutput(file)
	infoLogger.Println(v...)
}
func Err(v ...interface{}) {

	file, err := openLogFile("./log/err/" + fmt.Sprint(todate) + ".log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	infoLogger := log.New(os.Stdout, "ERROR :", flag)
	infoLogger.SetOutput(file)
	infoLogger.Println(v...)
}
func Warn(v ...interface{}) {

	file, err := openLogFile("./log/warn/" + fmt.Sprint(todate) + ".log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	infoLogger := log.New(os.Stdout, "WARNING :", flag)
	infoLogger.SetOutput(file)
	infoLogger.Println(v...)
}
func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
func ReadFile(path string) []string {
	var container []string
	// open file
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		container = append(container, scanner.Text())
		return container
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}
func GetListFile(originPath string) []string {
	var container []string
	files, err := ioutil.ReadDir("./" + originPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		container = append(container, f.Name())
	}
	return container
}
