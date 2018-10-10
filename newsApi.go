package newsApi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jforcode/DeepError"
)

type NewsApi struct {
	url    string
	key    string
	client *http.Client
}

func (api *NewsApi) Init(url, key string) {
	api.url = url
	api.key = key
	api.client = &http.Client{}
}

type FetchSourcesParams struct {
	categories Category
	languages  Language
	countries  Country
}

// TODO
func (params *FetchSourcesParams) Validate() error {
	return nil
}

// TODO
func (api *NewsApi) FetchSources() (*ApiSourcesResponse, error) {
	fnName := "newsApi.NewsApi.FetchSources"

	var sourceResponse ApiSourcesResponse
	err := api.getResponse("sources", nil, &sourceResponse)
	if err != nil {
		return nil, deepError.New(fnName, "getting sources response", err)
	}

	return &sourceResponse, nil
}

type FetchEverythingParams struct {
	q              string
	sources        []string
	domains        []string
	excludeDomains []string
	from           time.Time
	to             time.Time
	language       Language
	sortBy         SortBy
	pageSize       int
	page           int
}

// TODO
func (params *FetchEverythingParams) Validate() error {
	return nil
}

// TODO
func (api *NewsApi) FetchEverything(sourceIds []string, pageNum, pageSize int) (*ApiArticlesResponse, error) {
	fnName := "newsApi.NewsApi.FetchEverything"

	params := make(map[string]string)
	params["sources"] = strings.Join(sourceIds, ",")
	params["page"] = strconv.Itoa(pageNum)
	params["pageSize"] = strconv.Itoa(pageSize)

	var articleResponse ApiArticlesResponse
	err := api.getResponse("everything", params, &articleResponse)
	if err != nil {
		return nil, deepError.New(fnName, "getting articles response", err)
	}

	return &articleResponse, nil
}

type FetchTopHeadlinesParams struct {
	country  Country
	category Category
	sources  []string
	q        string
	pageSize int
	page     int
}

// TODO
func (params *FetchTopHeadlinesParams) Validate() error {
	return nil
}

// TODO
func (api *NewsApi) FetchTopHeadlines(sourceIds []string, pageNum, pageSize int) (*ApiArticlesResponse, error) {
	fnName := "newsApi.NewsApi.FetchTopHeadlines"

	params := make(map[string]string)
	params["sources"] = strings.Join(sourceIds, ",")
	params["page"] = strconv.Itoa(pageNum)
	params["pageSize"] = strconv.Itoa(pageSize)

	var articleResponse ApiArticlesResponse
	err := api.getResponse("everything", params, &articleResponse)
	if err != nil {
		return nil, deepError.New(fnName, "getting articles response", err)
	}

	return &articleResponse, nil
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
