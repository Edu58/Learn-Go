package utils

import "log"


func FatalErrorHandler(err error)  {
	if err != nil {
		log.Fatalln(err)
	}
}