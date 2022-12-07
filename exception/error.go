package exception

import "log"

func CheckError(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}
