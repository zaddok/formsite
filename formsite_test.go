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
	fmt.Println("Fetch all results from form 45")
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
		fmt.Println(result.Id)
	}

	// Fetch results for a form
	fmt.Println("Fetch one page of results from 45, from a high ID that doesnt exist. Expect 0 results")
	results, err = api.GetResultsFrom("form45", 999999999, 5)
	if err != nil {
		t.Errorf("api.GetResults() failed: %v", err)
		return
	}
	if len(results) > 0 {
		t.Errorf("api.GetResults() failed: expecting no results")
		return
	}
	for _, result := range results {
		fmt.Println(result)
	}

	// Fetch results for a form
	fmt.Println("Check we can fetch only the first 5 items from a list")
	results, err = api.GetResultsFrom("form44", 0, 5)
	if err != nil {
		t.Errorf("api.GetResults() failed: %v", err)
		return
	}
	if len(results) != 5 {
		t.Errorf("api.GetResults() failed: expecting 5 results")
		return
	}
	for _, result := range results {
		fmt.Println(result.Id)
	}
	fmt.Println("Check we can fetch only the most recent 5 items. Starting from id", results[4].Id)
	results, err = api.GetResultsFrom("form44", results[4].Id, 5)
	if err != nil {
		t.Errorf("api.GetResults() failed: %v", err)
		return
	}
	if len(results) != 5 {
		t.Errorf("api.GetResults() failed: expecting 5 results")
		return
	}
	for _, result := range results {
		fmt.Println(result.Id + 1)
	}

}
