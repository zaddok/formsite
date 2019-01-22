# Formsite Golang API

How to use the formsite Go API:

        api := formsite.NewFormsiteApi("https://fs22.formsite.com/api/users/myusername/forms", "myapikey")

        forms, err := api.GetForms()
        if err != nil {
                return err
        }
        for _, form := range forms {
                fmt.Println(form)
        }

        results, err := api.GetResults("form22")
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
