package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestBody(t *testing.T) {
	requestBody := RequestBody{Key: "item", Value: "test"}
	assert.Equal(t, requestBody.Key, "item", "Keys should be the same.")
	assert.Equal(t, requestBody.Value, "test", "Keys should be the same.")
	assert.NotEqual(t, requestBody.Key, "id", "item and id should not be equal")
}
