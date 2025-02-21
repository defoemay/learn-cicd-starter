package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	req, err := http.NewRequest("GET", "http://1.1.1.1/", nil)
	if err != nil {
		t.Fatalf("Error creating the request")
	}

	_, err = GetAPIKey(req.Header)
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("Expected ErrNoAuthHeaderIncluded on request with empty header")
	}
}

// func TestGetAPIKey(t *testing.T) {
// 	tests := []struct {
// 		key       string
// 		value     string
// 		expect    string
// 		expectErr string
// 	}{
// 		{
// 			expectErr: "no authorization header",
// 		},
// 		{
// 			key:       "Authorization",
// 			expectErr: "no authorization header",
// 		},
// 		{
// 			key:       "Authorization",
// 			value:     "-",
// 			expectErr: "malformed authorization header",
// 		},
// 		{
// 			key:       "Authorization",
// 			value:     "Bearer xxxxxx",
// 			expectErr: "malformed authorization header",
// 		},
// 		{
// 			key:       "Authorization",
// 			value:     "ApiKey xxxxxx",
// 			expect:    "xxxxxx",
// 			expectErr: "not expecting an error",
// 		},
// 	}

// 	for i, test := range tests {
// 		t.Run(fmt.Sprintf("TestGetAPIKey Case #%v:", i), func(t *testing.T) {
// 			header := http.Header{}
// 			header.Add(test.key, test.value)

// 			output, err := GetAPIKey(header)
// 			if err != nil {
// 				if strings.Contains(err.Error(), test.expectErr) {
// 					return
// 				}
// 				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
// 				return
// 			}

// 			if output != test.expect {
// 				t.Errorf("Unexpected: TestGetAPIKey:%s", output)
// 				return
// 			}
// 		})
// 	}
// }
