package utils

import "log"

//Check function for managing errors
func Check(e error, kind string) {
	if e != nil {
		log.Printf("%v :: %v\n", kind, e.Error())
	}
}
