## Callbacks

I want to show how to do callbacks in Golang although many say that callbacks are not idiomatic in Golang

[Callbacks](https://en.wikipedia.org/wiki/Callback_(computer_programming))

> a callback is any executable code that is passed as an argument to other code, which is expected to call back (execute) the argument at a given time.

There are two types of callbacks, differing in how they control data flow at runtime: 

1. Blocking callbacks (also known as synchronous callbacks or just callbacks) 
2. Deferred callbacks (also known as asynchronous callbacks). 

* Blocking callbacks are invoked before a function returns
* Deferred callbacks may be invoked after a function returns. 
    * Deferred callbacks are often used in the context of I/O operations or event handling,
    * Deferred callbacks are called by interrupts or by a different thread in case of multiple threads.
* Due to their nature, blocking callbacks can work without interrupts or multiple threads, meaning that blocking callbacks are not commonly used for synchronization or delegating work to another thread.

```go
package workshop

func average(numbers []float32, cb func(float32) float32) float32 {
	var sum float32
	for _, val := range numbers {
		sum += cb(val)
	}
	return sum / float32(len(numbers))
}

func callbacks() {
	nums := average([]float32{1.0, 2.0, 3.0, 4.0, 5.0}, func(n float32) float32 {
		return n
	})
	assert(nums == 3.0)
}
```

Notice here that the average function takes 2 parameters `numbers` which is a list of float32 and a function
Also notice that the name cb is used for the cb function and all it does in the average function is return `n`

[Callbacks in Golang](https://play.golang.org/p/YeGAaivQMq)