package single

import (
	"fmt"
	"sync"
)

var once sync.Once
var instance2 *single

func getInstance() *single {
	if instance2 == nil {
		once.Do(func() {
			fmt.Println("init instance2")
			instance2 = &single{}
		})
	}
	fmt.Println("get instance2 success")
	return instance2
}
