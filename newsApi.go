package newsApi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jforcode/DeepError"
	"github.com/jforcode/Util"
)

const (
	maxSourcesInARequest int = 20
	defaultPageSize      int = 20
	maxPageSize          int = 100
)

type NewsApi struct {
	url    string
	key    string
	client *http.Client
}

func New(url, key string) *NewsApi {
	return &NewsApi{
		url,
		key,
		&http.Client{},
	}
}

type FetchSourcesParams struct {
	Category Category
	Language Language
	Country  Country
}

func (params *FetchSourcesParams) Validate() error {
	return nil
}

func (params *FetchSourcesParams) GetRequestParamsMap() (map[string]string, error) {
	fnName := "newsApi.FetchSourcesParams.GetRequestParamsMap"

	err := params.Validate()
	if err != nil {
		return nil, deepError.New(fnName, "validate", err)
	}

	reqParams := make(map[string]string)
	if params.Category != "" {
		reqParams["category"] = string(params.Category)
	}
	if params.Language != "" {
		reqParams["language"] = string(params.Language)
	}
	if params.Country != "" {
		reqParams["country"] = string(params.Country)
	}

	return reqParams, nil
}

func (api *NewsApi) FetchSources(params *FetchSourcesParams) (*ApiSourcesResponse, error) {
	fnName := "newsApi.NewsApi.FetchSources"

	reqParams, err := params.GetRequestParamsMap()
	if err != nil {
		return nil, deepError.New(fnName, "getting params map", err)
	}

	var sourceResponse ApiSourcesResponse
	err = api.getResponse("sources", reqParams, &sourceResponse)
	if err != nil {
		return nil, deepError.New(fnName, "getting sources response", err)
	}

	return &sourceResponse, nil
}

type FetchEverythingParams struct {
	Q              string
	Sources        []string
	Domains        []string
	ExcludeDomains []string
	From           time.Time
	To             time.Time
	Language       Language
	SortBy         SortBy
	PageSize       int
	Page           int
}

func (params *FetchEverythingParams) Validate() error {
	fnName := "newsApi.FetchEverythingParams.Validate"

	if len(params.Sources) > maxSourcesInARequest {
		return deepError.DeepErr{
			Function: fnName,
			Action:   "validating sources",
			Message:  "Maximum length of sources should be " + string(maxSourcesInARequest),
		}
	}

	if params.PageSize > maxPageSize {
		return deepError.DeepErr{
			Function: fnName,
			Action:   "validating page size",
			Message:  "Max allowed page size is " + string(maxPageSize),
		}
	}

	return nil
}

func (params *FetchEverythingParams) GetRequestParamsMap() (map[string]string, error) {
	fnName := "newsApi.FetchEvFetchEverythingParams.GetRequestParamsMap"

	err := params.Validate()
	if err != nil {
		return nil, deepError.New(fnName, "validate", err)
	}

	reqParams := make(map[string]string)
	if params.Q != "" {
		reqParams["q"] = params.Q
	}
	if !util.Array.IsEmptyStringArray(params.Sources) {
		reqParams["sources"] = strings.Join(params.Sources, ",")
	}
	if !util.Array.IsEmptyStringArray(params.Domains) {
		reqParams["domains"] = strings.Join(params.Domains, ",")
	}
	if !util.Array.IsEmptyStringArray(params.ExcludeDomains) {
		reqParams["excludeDomains"] = strings.Join(params.ExcludeDomains, ",")
	}
	if !params.From.IsZero() {
		reqParams["from"] = params.From.UTC().Format(time.RFC3339)
	}
	if !params.To.IsZero() {
		reqParams["to"] = params.To.UTC().Format(time.RFC3339)
	}
	if params.Language != "" {
		reqParams["language"] = string(params.Language)
	}
	if params.SortBy != "" {
		reqParams["sortBy"] = string(params.SortBy)
	}
	if params.PageSize != 0 {
		reqParams["pageSize"] = strconv.Itoa(params.PageSize)
	}
	if params.Page != 0 {
		reqParams["page"] = strconv.Itoa(params.Page)
	}

	return reqParams, nil
}

func (api *NewsApi) FetchEverything(params *FetchEverythingParams) (*ApiArticlesResponse, error) {
	fnName := "newsApi.NewsApi.FetchEverything"

	reqParamsMap, err := params.GetRequestParamsMap()
	if err != nil {
		return nil, deepError.New(fnName, "getting params map", err)
	}

	var articleResponse ApiArticlesResponse
	err = api.getResponse("everything", reqParamsMap, &articleResponse)
	if err != nil {
		return nil, deepError.New(fnName, "getting articles response", err)
	}

	return &articleResponse, nil
}

type FetchTopHeadlinesParams struct {
	Country  Country
	Category Category
	Sources  []string
	Q        string
	PageSize int
	Page     int
}

func (params *FetchTopHeadlinesParams) Validate() error {
	fnName := "newsApi.FetchTopHeadlinesParams.Validate"

	if !util.Array.IsEmptyStringArray(params.Sources) && (params.Country != "" || params.Category != "") {
		return deepError.DeepErr{
			Function: fnName,
			Action:   "validating sources, country and category",
			Message:  "Sources can't be mixed with Country or Category",
		}
	}

	if params.PageSize > maxPageSize {
		return deepError.DeepErr{
			Function: fnName,
			Action:   "validating page size",
			Message:  "Max allowed page size is " + string(maxPageSize),
		}
	}

	return nil
}

func (params *FetchTopHeadlinesParams) GetRequestParamsMap() (map[string]string, error) {
	fnName := "newsApi.FetchTopHeadlinesParams.GetRequestParamsMap"

	err := params.Validate()
	if err != nil {
		return nil, deepError.New(fnName, "validate", err)
	}

	reqParamsMap := make(map[string]string)
	if params.Country != "" {
		reqParamsMap["country"] = string(params.Country)
	}
	if params.Category != "" {
		reqParamsMap["category"] = string(params.Category)
	}
	if !util.Array.IsEmptyStringArray(params.Sources) {
		reqParamsMap["sources"] = strings.Join(params.Sources, ",")
	}
	if params.Q != "" {
		reqParamsMap["q"] = params.Q
	}
	if params.PageSize != 0 {
		reqParamsMap["pageSize"] = strconv.Itoa(params.PageSize)
	}
	if params.Page != 0 {
		reqParamsMap["page"] = strconv.Itoa(params.Page)
	}

	return reqParamsMap, nil
}

func (api *NewsApi) FetchTopHeadlines(params *FetchTopHeadlinesParams) (*ApiArticlesResponse, error) {
	fnName := "newsApi.NewsApi.FetchTopHeadlines"

	reqParams, err := params.GetRequestParamsMap()
	if err != nil {
		return nil, deepError.New(fnName, "validating request parameters", err)
	}

	var topHeadlinesResponse ApiArticlesResponse
	err = api.getResponse("top-headlines", reqParams, &topHeadlinesResponse)
	if err != nil {
		return nil, deepError.New(fnName, "getting top headlines response", err)
	}

	return &topHeadlinesResponse, nil
}

func (api *NewsApi) getResponse(endpoint string, params map[string]string, toResponse interface{}) error {
	fnName := "newsApi.NewsApi.get"

	req, err := http.NewRequest("GET", api.url+"/"+endpoint, nil)
	if err != nil {
		return deepError.New(fnName, "making http request", err)
	}

	req.Header.Add("X-Api-Key", api.key)

	if params != nil {
		q := req.URL.Query()
		for key, value := range params {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := api.client.Do(req)
	if err != nil {
		return deepError.New(fnName, "client do", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return deepError.New(fnName, "reading response", err)
	}

	if resp.StatusCode != http.StatusOK {
		var apiErr ApiError
		err = json.Unmarshal(bodyBytes, &apiErr)
		if err != nil {
			return deepError.New(fnName, "unmarshalling api error", err)
		}

		return apiErr
	}

	err = json.Unmarshal(bodyBytes, toResponse)
	if err != nil {
		return deepError.New(fnName, "unmarshalling api response", err)
	}

	return nil
}
