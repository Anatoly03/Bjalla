// API Testing for Guild Roles
// https://pocketbase.io/docs/go-testing/

package api_test

import (
	"net/http"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

// TestBjallaBotExists tests that the bjalla bot user exists in the
// database and can be authenticated.
func TestBjallaBotExists(t *testing.T) {
	app := setupTestApp(t)
	t.Cleanup(func() {
		app.Cleanup()
	})

	token := getAuthToken(app, "pbbotbjalla0000")
	testAppFactory := func(t testing.TB) *tests.TestApp {
		return app
	}

	scenarios := []tests.ApiScenario{
		{
			Name:   "the bot user bjalla should exist and be able to authenticate",
			Method: http.MethodGet,
			URL:    "/api/collections/users/records/pbbotbjalla0000",
			Headers: map[string]string{
				"Authorization": token,
			},
			DisableTestAppCleanup: true,
			TestAppFactory:        testAppFactory,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"name":"Bjalla"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
