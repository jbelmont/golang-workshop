package json_manipulation

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type JsonManipulator interface {
	readFile()
}

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

func (plans ServicePlans) readFile() ServicePlans {
	servicePlans, err := ioutil.ReadFile("service_plans.json")
	if err != nil {
		log.Fatal(err)
	}
	var services ServicePlans
	json.Unmarshal(servicePlans, &services)
	return services
}
