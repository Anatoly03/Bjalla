package api_test

// TEST SOURCE FROM
// https://github.com/presentator/presentator/blob/master/api_users_test.go

import (
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

func setupTestApp(t testing.TB) *tests.TestApp {
	_, filePath, _, _ := runtime.Caller(0)
	dataDir := filepath.Clean(filepath.Join(filepath.Dir(filePath), "..", "pb_data"))

	testApp, err := tests.NewTestApp(dataDir)

	if err != nil {
		t.Fatal(err)
	}

	return testApp
}

// TestUsersCreate ensures everyone can create a new user record
func TestUsersCreate(t *testing.T) {
	t.Parallel()

	scenarios := []tests.ApiScenario{
		{
			Name:   "everyone should be able to create a user",
			Method: http.MethodPost,
			URL:    "/api/collections/users/records",
			Body: strings.NewReader(`{
				"name":            "test_new",
				"email":           "test_new@example.com",
				"password":        "1234567890",
				"passwordConfirm": "1234567890"
			}`),
			TestAppFactory: setupTestApp,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"name":"test_new"`,
				`"verified":false`,
			},
			ExpectedEvents: map[string]int{
				"*":                          0,
				"OnRecordCreateRequest":      1,
				"OnModelValidate":            1,
				"OnModelCreate":              1,
				"OnModelCreateExecute":       1,
				"OnModelAfterCreateSuccess":  1,
				"OnRecordEnrich":             1,
				"OnRecordValidate":           1,
				"OnRecordCreate":             1,
				"OnRecordCreateExecute":      1,
				"OnRecordAfterCreateSuccess": 1,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
