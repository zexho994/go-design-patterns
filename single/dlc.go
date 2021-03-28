package single

import (
	"fmt"
	"sync"
)

type single struct {
}

var lock = &sync.Mutex{}
var instance *single

func GetInstance() *single {
	if instance == nil {
		lock.Lock()
		fmt.Println("lock")
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("init instance")
			instance = &single{}
		}
		fmt.Println("un lock")
	}
	fmt.Println("get instance success")
	return instance
}
