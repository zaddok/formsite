# Formsite Golang API

How to use the formsite Go API:

        api := formsite.NewFormsiteApi("https://fs22.formsite.com/api/users/myusername/forms", "myapikey")
        
	// List all available forms
        forms, err := api.GetForms()
        if err != nil {
                return err
        }
        for _, form := range forms {
                fmt.Println(form)
        }
        
	// Get the first page of data from a form
        results, err := api.GetResults("form22", 1)
        if err != nil {
                return err
        }
        for _, result := range results {
                fmt.Println(result)
        }
        
	// Get the next 5 form submissions, starting at particular response id
	lastFetchedId := 104992
        results, err := api.GetResults("form22", lastFetchedId, 5)
        if err != nil {
                return err
        }
        for _, result := range results {
                fmt.Println(result)
        }

## Related links

 - Formsite API Documentation: https://support.formsite.com/hc/en-us/articles/360000288594-API
 - To find your API key and URL:
    - go to a form you have created
    - click on "Settings" for that form,
    - click on "Integrations"
    - click on "Formsite API"
