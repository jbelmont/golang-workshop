## Practice

1. Create a folder called practice
2. Create a file and name it whatever you like
3. Create a Test file and name it whatever you like but append `_test.go` to the end of the name

Exercise 1:

Write a function that computes the standard deviation using a closure and formats the number using 2 decimal places.
Write a test function to test the newly created function.

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

Exercise 3:

Write a recursive function that sums up the values of an integer along with a test.