package newsApi

import "testing"

func TestApiError(t *testing.T) {
	err := ApiError{status: "error", code: "apiKeyMissing", message: "Your API key is missing. Append this to the URL with the apiKey param, or use the x-api-key HTTP header."}
	expected := "apiKeyMissing - Your API key is missing. Append this to the URL with the apiKey param, or use the x-api-key HTTP header."

	if err.Error() != expected {
		t.FailNow()
	}
}
