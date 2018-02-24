//+build integration

package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestIndex Integration Test for Index "/"
func TestIndex(t *testing.T) {
	//Arrange

	//Act
	response, _ := http.Get("http://localhost:8080/")

	//Assert
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
