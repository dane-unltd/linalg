package mat

import (
	//	"fmt"
	"github.com/gonum/blas"
	"sync"
)

var ops matops
var mut sync.Mutex

type matops interface {
	blas.Float64
}

func Register(o matops) {
	mut.Lock()
	if ops != nil {
		//		fmt.Println("Warning: multiple operator registrations")
	}
	ops = o
	mut.Unlock()
}
