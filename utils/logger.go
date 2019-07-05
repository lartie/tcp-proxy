package utils

import (
	"log"
)

func WriteCritLog(s string) {
	log.Printf("[CRIT]: %s\n", s)
}

func WriteInfoLog(s string) {
	log.Printf("[INFO]: %s\n", s)
}
