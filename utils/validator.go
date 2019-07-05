package utils

import (
	"sync"
)

func CheckErr(err error) {
	if err != nil {
		WriteCritLog(err.Error())
		panic(err)
	}
}

func CheckErrWithWaitGroup(err error, group *sync.WaitGroup) {
	if err != nil {
		group.Done()
		WriteCritLog(err.Error())
		panic(err)
	}
}
