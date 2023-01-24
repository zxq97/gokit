package concurrent

import (
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	Go(func() {
		panic(1)
	})
	<-time.After(time.Second)
}

func TestF(t *testing.T) {
	f := func() {
		panic(1)
	}
	F(f)
}
