# Testing

* Go has a lightweight test framework composed of the go test command and the testing package.
* Write a test by creating a file with a name ending in _test.go
    * Write functions named `TestXXX` with signature func (t *testing.T)
* The test framework runs each such function
    * if the function calls a failure function such as t.Error or t.Fail, the test is considered to have failed.

```go
func TestAverage(t *testing.T) {
	numbers := []float64{
		1, 2, 3, 4, 5, 6, 7,
	}
	addEmUp := averageEmUp(numbers)
	if addEmUp != 4 {
		t.Error("Expected 4 but received ", addEmUp)
	}
}
```

Here is an example test function called `TestAverage` notice that it takes the signature `func (t *testing.T)`

Make sure to append `_test.go` to the end of the test so that go finds the test.

In Go you assert what should not occur which is a break from the `xunit` style of `assert.equal` 