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

func (api *NewsApi) Init(url, key string) {
	api.url = url
	api.key = key
	api.client = &http.Client{}
}

func (api *NewsApi) FetchSources() (*ApiSourcesResponse, error) {
	prefix := "newsApi.NewsApi.FetchSources"

	var sourceResponse ApiSourcesResponse
	err := api.getResponse("sources", nil, &sourceResponse)
	if err != nil {
		return nil, switchAndGetErr(prefix+" (Api get response): ", err)
	}

	return &sourceResponse, nil
}

func (api *NewsApi) FetchArticles(sourceIds []string, pageNum, pageSize int) (*ApiArticlesResponse, error) {
	prefix := "newsApi.NewsApi.FetchArticles"

	params := make(map[string]string)
	params["sources"] = strings.Join(sourceIds, ",")
	params["page"] = strconv.Itoa(pageNum)
	params["pageSize"] = strconv.Itoa(pageSize)

	var articleResponse ApiArticlesResponse
	err := api.getResponse("everything", params, &articleResponse)
	if err != nil {
		return nil, switchAndGetErr(prefix+" (Api get response): ", err)
	}

	return &articleResponse, nil
}

func switchAndGetErr(prefix string, err error) error {
	switch err.(type) {
	case ApiError:
		return err
	default:
		return errors.New(prefix + err.Error())
	}
}

// TODO: handle errors better
// if NOT 200 OK, then return a ApiError
func (api *NewsApi) getResponse(endpoint string, params map[string]string, toResponse interface{}) error {
	prefix := "newsApi.NewsApi.get"

	req, err := http.NewRequest("GET", api.url+"/"+endpoint, nil)
	if err != nil {
		return errors.New(prefix + " (http request): " + err.Error())
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
		return errors.New(prefix + " (client do): " + err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(prefix + " (read): " + err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		var apiErr ApiError
		err = json.Unmarshal(bodyBytes, &apiErr)
		if err != nil {
			return errors.New(prefix + " (unmarshal api error): " + err.Error())
		}

		return apiErr
	}

	err = json.Unmarshal(bodyBytes, toResponse)
	if err != nil {
		return errors.New(prefix + " (unmarshal api response): " + err.Error())
	}

	return nil
}
