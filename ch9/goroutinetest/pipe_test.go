package main

import (
	"testing"
)

func TestPipeline(t *testing.T) {
	//in, out := pipeline()
	//for i := 0; i < b.N; i++ {
	//	in <- 1
	//	<-out
	//}
	//close(in)

	pipeline(700000)
}
