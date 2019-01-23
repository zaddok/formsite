package formsite

import (
	"fmt"
	"os"
	"testing"
)

func TestFormsiteApi(t *testing.T) {

	// Initialize the api
	url := os.Getenv("API_URL")
	key := os.Getenv("API_KEY")
	if url == "" {
		t.Error("Please specify API_URL environment variable")
		return
	}
	if key == "" {
		t.Error("Please specify API_KEY environment variable")
		return
	}
	api := NewFormsiteApi(url, key)

	// Fetch a list of forms
	forms, err := api.GetForms()
	if err != nil {
		t.Errorf("api.GetForms() failed: %v", err)
	}
	if len(forms) == 0 {
		t.Errorf("api.GetForms() failed: expecting more than one form")
	}
	for _, form := range forms {
		fmt.Println(form)
	}

	// Fetch results for a form
	results, err := api.GetResults("form45", 1)
	if err != nil {
		t.Errorf("api.GetResults() failed: %v", err)
		return
	}
	if len(results) == 0 {
		t.Errorf("api.GetResults() failed: expecting more than one result")
		return
	}
	for _, result := range results {
		fmt.Println(result)
	}

}
