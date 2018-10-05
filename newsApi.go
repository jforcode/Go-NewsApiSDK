package newsApi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type NewsApi struct {
	url    string
	key    string
	client *http.Client
}

func (api *NewsApi) Init(url, key string) error {
	api.url = url
	api.key = key
	api.client = &http.Client{}

	return nil
}

func (api *NewsApi) FetchSources() (*ApiSourcesResponse, error) {
	prefix := "newsApi.NewsApi.FetchSources"

	bodyBytes, err := api.get("sources", nil)
	if err != nil {
		return nil, errors.New(prefix + " (Api Get): " + err.Error())
	}

	var sourceResponse ApiSourcesResponse
	if err := json.Unmarshal(bodyBytes, &sourceResponse); err != nil {
		return nil, errors.New(prefix + " (json unmarshal): " + err.Error())
	}

	return &sourceResponse, nil
}

func (api *NewsApi) FetchArticles(sourceIds []string, pageNum, pageSize int) (*ApiArticlesResponse, error) {
	prefix := "newsApi.NewsApi.FetchArticles"

	params := make(map[string]string)
	params["sources"] = strings.Join(sourceIds, ",")
	params["page"] = strconv.Itoa(pageNum)
	params["pageSize"] = strconv.Itoa(pageSize)

	bodyBytes, err := api.get("everything", params)
	if err != nil {
		return nil, errors.New(prefix + " (Api get): " + err.Error())
	}

	var articleResponse ApiArticlesResponse
	if err := json.Unmarshal(bodyBytes, &articleResponse); err != nil {
		return nil, errors.New(prefix + " (json unmarshal): " + err.Error())
	}

	return &articleResponse, nil
}

// TODO: handle errors better
// if NOT 200 OK, then return a ApiError
func (api *NewsApi) get(endpoint string, params map[string]string) ([]byte, error) {
	prefix := "newsApi.NewsApi.get"

	req, err := http.NewRequest("GET", api.url+"/"+endpoint, nil)
	if err != nil {
		return nil, errors.New(prefix + " (http request): " + err.Error())
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
		return nil, errors.New(prefix + " (client do): " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(prefix + " (HTTP Status Error): " + resp.Status)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(prefix + " (read): " + err.Error())
	}

	return bodyBytes, nil
}
