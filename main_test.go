//+build unit

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tableTest = [1]struct {
	message            string
	expectedStatusCode int
}{
	{
		message:            "Status Code are not 200",
		expectedStatusCode: 200,
	},
}

//TestIndex Test the Index function
func TestIndex(t *testing.T) {
	//Arrange
	assert := assert.New(t)
	request, _ := http.NewRequest("GET", "/", nil)

	for _, tt := range tableTest {
		w := httptest.NewRecorder()

		//Act
		Index(w, request)

		//Assert
		assert.Equal(tt.expectedStatusCode, w.Code, tt.message)
	}
}
