package main

func pipeline(n int) {
	mid := make([]chan int, n)
	for i := 0; i < n; i++ {
		mid[i] = make(chan int)

	}
	head, tail := mid[0], mid[n-1]
	for i := 0; i < n; i++ {
		go func(i int) {
			mid[i+1] <- <-mid[i]

		}(i)

	}
	head <- 1
	<-tail

}

func main() {

	pipeline(10)
}
