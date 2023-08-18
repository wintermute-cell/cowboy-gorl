package logging

import (
	"errors"
	"fmt"
	"log"
	"os"
)


// FIXME: wie wird das ganze jetzt "static", damit ich es in jeder .go benutzen
// kann? keine ahnung im Moment....

type Log struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

/*
	<log_path> + log.txt = c
*/
func (logger *Log) Init(log_path string) {

	file_path := log_path + "log.txt"
	fmt.Println("---filepath for logfile: ", file_path) // TODO: remove

	// filepath exists
	if _, err := os.Stat(file_path); err == nil {
		// do nothing...
	// filepath does not exist  
	} else if errors.Is(err, os.ErrNotExist) {
		// creates the directory (and subdirectories)
		if err := os.MkdirAll(log_path, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
	}


	file, err := os.OpenFile(file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger.infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (logger *Log) Info(text string) {
	logger.infoLogger.Printf(text)
}

func (logger *Log) Warning(text string) {
	logger.warningLogger.Printf(text)
}

func (logger *Log) Error(text string) {
	logger.errorLogger.Printf(text)
	// FIXME: lieber ErrorLogger.Fatal benutzen?
}
