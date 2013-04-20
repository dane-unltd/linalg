package matrix

import (
	//	"fmt"
	"github.com/dane-unltd/linalg/lapack"
	"github.com/kortschak/blas"
	"sync"
)

var ops matops
var mut sync.Mutex

type matops interface {
	blas.Float64
	lapack.Float64
}

func Register(o matops) {
	mut.Lock()
	if ops != nil {
		//		fmt.Println("Warning: multiple operator registrations")
	}
	ops = o
	mut.Unlock()
}
