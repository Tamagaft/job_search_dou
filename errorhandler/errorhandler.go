package errorhandler

import "log"

func HandlePanicError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
