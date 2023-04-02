package main

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"net/http/httptest"

	"github.com/stretchr/testify/assert" // add Testify package
)

func Test_main(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "should get HTTP status 200 when asking for an article",
			route:        "/article",
			expectedCode: 200,
		},
		// Second test case
		{
			description:  "should get HTTP status 404, when route does not exist",
			route:        "/not-found",
			expectedCode: 404,
		},
	}

	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest("POST", test.route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}

}
