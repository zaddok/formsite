package formsite

import (
	"encoding/xml"
	"fmt"
)

type FormsiteApi struct {
	apiUrl string
	apiKey string
	fetch  LookupUrl
}

type Form struct {
	Id        int64  `xml:"id,attr"`
	Name      string `xml:"name"`
	Directory string `xml:"directory"`
}

func NewFormsiteApi(apiUrl, apiKey string) *FormsiteApi {
	return &FormsiteApi{
		apiUrl: apiUrl,
		apiKey: apiKey,
		fetch:  &DefaultLookupUrl{},
	}
}

func (fs *FormsiteApi) GetForms() ([]Form, error) {
	url := fmt.Sprintf("%s?fs_api_key=%s", fs.apiUrl, fs.apiKey)
	body, _, _, err := fs.fetch.GetUrl(url)
	if err != nil {
		return []Form{}, err
	}

	type ResponseForms struct {
		XMLName xml.Name `xml:"forms"`
		Forms   []Form   `xml:"form"`
	}

	type Response struct {
		XMLName xml.Name      `xml:"fs_response"`
		Forms   ResponseForms `xml:"forms"`
	}
	var result Response

	if err := xml.Unmarshal([]byte(body), &result); err != nil {
		return result.Forms.Forms, err
	}

	return result.Forms.Forms, nil
}

type Result struct {
	Id     int64
	Fields map[string][]string
	Metas  map[string]string
}

// GetResultsFrom returns any new form submissions after the last seen request id.
func (fs *FormsiteApi) GetResultsFrom(formName string, lastRequestId, limit int64) ([]*Result, error) {
	url := fmt.Sprintf("%s/%s/results?fs_api_key=%s&fs_limit=%d&fs_include_headings&fs_min_id=%d&fs_sort=result_id&fs_sort_direction=desc", fs.apiUrl, formName, fs.apiKey, limit, lastRequestId)
	return fs.fetchResults(url)
}

func (fs *FormsiteApi) GetResults(formName string, page int64) ([]*Result, error) {
	url := fmt.Sprintf("%s/%s/results?fs_api_key=%s&fs_limit=100&fs_page=%d&fs_include_headings", fs.apiUrl, formName, fs.apiKey, page)
	return fs.fetchResults(url)

}

func (fs *FormsiteApi) fetchResults(url string) ([]*Result, error) {
	body, _, _, err := fs.fetch.GetUrl(url)
	if err != nil {
		return []*Result{}, err
	}

	type ResponseResultMeta struct {
		XMLName xml.Name `xml:"meta"`
		Id      string   `xml:"id,attr"`
		Value   string   `xml:",chardata"`
	}
	type ResponseResultItem struct {
		XMLName xml.Name `xml:"item"`
		Id      string   `xml:"id,attr"`
		Type    string   `xml:"type,attr"`
		Index   string   `xml:"index,attr"`
		Value   []string `xml:"value"`
	}
	type ResponseHeading struct {
		XMLName xml.Name `xml:"heading"`
		Id      string   `xml:"for,attr"`
		Value   string   `xml:",chardata"`
	}

	type ResponseResult struct {
		XMLName xml.Name             `xml:"result"`
		Id      int64                `xml:"id,attr"`
		Metas   []ResponseResultMeta `xml:"metas>meta"`
		Items   []ResponseResultItem `xml:"items>item"`
	}

	type Response struct {
		XMLName  xml.Name          `xml:"fs_response"`
		Headings []ResponseHeading `xml:"headings>heading"`
		Results  []ResponseResult  `xml:"results>result"`
	}
	var result Response

	if err := xml.Unmarshal([]byte(body), &result); err != nil {
		return []*Result{}, err
	}

	h := map[string]string{}
	for _, heading := range result.Headings {
		h[heading.Id] = heading.Value
	}

	results := make([]*Result, 0)
	for _, result := range result.Results {
		r := &Result{
			Id:     result.Id,
			Fields: make(map[string][]string),
			Metas:  make(map[string]string),
		}
		for _, i := range result.Items {
			hn := i.Id
			if _, v := h[i.Id]; v {
				hn = h[i.Id] + "|" + i.Id
			}
			r.Fields[hn] = i.Value
		}
		for _, i := range result.Metas {
			//fmt.Println(i.Id, i.Value)
			r.Metas[i.Id] = i.Value
		}
		results = append(results, r)
	}

	return results, nil
}

func (fs *FormsiteApi) SetUrlFetcher(fetch LookupUrl) {
	fs.fetch = fetch
}
