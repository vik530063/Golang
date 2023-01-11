package main

import "fmt"

func enqueueOps[T comparable](queue []T, element T) []T {
	queue = append(queue, element) //append to slice
	fmt.Println("Enqueue the element:", element)
	return queue
}

func dequeueOps[T comparable](queue []T) []T {
	element := queue[0] // return the first element
	fmt.Println("Dequeue the element :", element)
	return queue[1:]
}

func peakOps[T comparable](queue []T) T {
	element := queue[0] // return the first element
	return element
}

func isEmptyOps[T comparable](queue []T) bool {
	return len(queue) == 0
}

func main() {
	var queue []string // Make a queue of strings.

	queue = enqueueOps(queue, "Roshan")
	queue = enqueueOps(queue, "Dharmendar")
	queue = enqueueOps(queue, "Vivek")

	fmt.Println("Queue:", queue)

	peak := peakOps(queue)
	fmt.Println("peak:", peak)

	queue = dequeueOps(queue)
	fmt.Println("Queue:", queue)

	isEmpty := isEmptyOps(queue)
	fmt.Println("is empty:", isEmpty)
}
