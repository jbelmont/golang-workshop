## Practice

1. Create a folder called practice
2. Create a file and name it whatever you like
3. Create a Test file and name it whatever you like but append `_test.go` to the end of the name

Exercise 1:

Write a function that computes the standard deviation using a closure and formats the number using 2 decimal places.
Write a test function to test the newly created function.

[Exercise Solution](https://play.golang.org/p/LeQLXHUrmW)

Exercise 2:

Write a function that reads the `service_plans.json` file and formats the json into Golang types

```go
type ServicePlans struct {
	TotalResults int         `json:"total_results"`
	TotalPages   int         `json:"total_pages"`
	PrevURL      interface{} `json:"prev_url"`
	NextURL      interface{} `json:"next_url"`
	Resources    []struct {
		Metadata struct {
			GUID      string    `json:"guid"`
			URL       string    `json:"url"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"metadata"`
		Entity struct {
			Name                string      `json:"name"`
			Free                bool        `json:"free"`
			Description         string      `json:"description"`
			ServiceGUID         string      `json:"service_guid"`
			Extra               interface{} `json:"extra"`
			UniqueID            string      `json:"unique_id"`
			Public              bool        `json:"public"`
			Active              bool        `json:"active"`
			Bindable            bool        `json:"bindable"`
			ServiceURL          string      `json:"service_url"`
			ServiceInstancesURL string      `json:"service_instances_url"`
		} `json:"entity"`
	} `json:"resources"`
}
```

Please use the following `struct` to help you along

Also write a test file to test this.

Steps to see solution

1. `cd json_manipulation`
2. Open `json_manipulation.go` and `json_manipulation_test.go`
3. Run `go test`

Exercise 3:

Write a recursive function that sums up the values of an integer along with a test.

[Exercise Solution](https://play.golang.org/p/Uu8TGQ2qDM)

Exercise 4:

Write a function using callbacks that filters a list of ints

[Exercise Solution](https://play.golang.org/p/cRmZByNXMR)

Exercise 5:

Write a function in its own package that reverses a string and write a test for it.

Steps to see solution:

1. `cd reverse`
2. Open `reverse.go` and `reverse_test.go`
3. Run `go test`

Exercise 6:

Write a function that queries Github Api for users and finds the repos with the most stargazers

Steps to see solution:

1. `cd flatten`
2. Open `flatten.go` and `flatten_test.go`
3. Run `go test`

[Go Playground Solution](https://play.golang.org/p/9HwlKyIcmS)
