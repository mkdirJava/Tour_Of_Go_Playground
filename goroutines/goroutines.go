package goroutines

func BasicGoRoutine() {
	go func() {
		println("I am running")
	}()
}

// Goroutines are like Java Thread construct or the Future, ie it returns stuff
func GoRoutineExample() int {
	sumFunc := func(items []int, returnChan chan int) {
		runningValue := 0
		for _, item := range items {
			runningValue = runningValue + item
		}
		returnChan <- runningValue
	}

	channel := make(chan int)

	go func() {
		sumFunc([]int{1, 2, 3, 4}, channel)
	}()

	go func() {
		sumFunc([]int{5, 6, 7, 8}, channel)
	}()
	return <-channel + <-channel
}

func populateChannel(channel chan int) {
	for value := range 100 {
		channel <- value
	}
	close(channel)
}

func InterestingSelect() int {

	runningNumberChannel := make(chan int)
	quit := make(chan int)
	runningNumber := 0
	go func() {
		for number := range 10 {
			runningNumberChannel <- number
		}
		quit <- 0
	}()

	for {
		select {
		case <-quit:
			return runningNumber
		case number := <-runningNumberChannel:
			runningNumber = runningNumber + number
		}

	}
}
