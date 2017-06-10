package json_manipulation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	var servicePlans ServicePlans
	plans := servicePlans.readFile()
	for _, resources := range plans.Resources {
		guid := resources.Metadata.GUID
		var expectedGUID = "6fecf53b-7553-4cb3-b97e-930f9c4e3385"
		assert.Equal(t, guid, expectedGUID)
	}
}
