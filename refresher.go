package newsApi

import (
	"errors"
	"strings"
	"time"

	"github.com/jforcode/DeepError"
)

type RefresherConfig struct {
	RemainingRequests int
	SourceIds         []string
	SourcesBatchSize  int
	StartPageNum      int
	PageSize          int
	LastMomentMinutes int
	SleepSeconds      int
}

func (config *RefresherConfig) Validate() error {
	fnName := "newsApi.RefresherConfig.Validate"

	errs := make([]string, 0)
	if config.SourceIds == nil || len(config.SourceIds) == 0 {
		errs = append(errs, "Invalid Source IDs")
	}

	if config.SourcesBatchSize == 0 {
		config.SourcesBatchSize = 20
	} else if config.SourcesBatchSize < 1 || config.SourcesBatchSize > 20 {
		errs = append(errs, "Invalid sources batch size")
	}

	if config.StartPageNum == 0 {
		config.StartPageNum = 1
	} else if config.StartPageNum < 1 {
		errs = append(errs, "Invalid start page number")
	}

	if config.PageSize == 0 {
		config.PageSize = 10
	} else if config.PageSize < 1 || config.PageSize > 100 {
		errs = append(errs, "Invalid page size")
	}

	if config.LastMomentMinutes == 0 {
		config.LastMomentMinutes = 30
	} else if config.LastMomentMinutes < 1 {
		errs = append(errs, "Invalid last moment minutes")
	}

	if config.SleepSeconds == 0 {
		config.SleepSeconds = 60
	} else if config.SleepSeconds < 1 {
		errs = append(errs, "Invalid sleep seconds")
	}

	errsI := make([]interface{}, len(errs))
	for ind, err := range errs {
		errsI[ind] = err
	}

	if len(errs) == 0 {
		return nil
	} else {
		return deepError.DeepErr{
			Function: fnName,
			Action:   "validating",
			Params:   errsI,
			Cause:    errors.New(strings.Join(errs, "|")),
		}
	}
}

type Refresher struct {
	newsApi *NewsApi
}

func (refr *Refresher) Init(newsApi *NewsApi) {
	refr.newsApi = newsApi
}

func (refr *Refresher) DailyRefresh(config *RefresherConfig, chArticles chan []*ApiArticle, chNumRequestsUpdated chan int, chError chan error) {
	prefix := "newsApi.Refresher.FetchArticles"
	closeChannels := func() {
		close(chArticles)
		close(chNumRequestsUpdated)
		close(chError)
	}

	err := config.Validate()
	if err != nil {
		chError <- err
		closeChannels()
		return
	}

	remainingRequests := config.RemainingRequests
	pageNum := config.StartPageNum
	lenSourceIds := len(config.SourceIds)
	today := time.Now().UTC()
	lastMoment := time.Date(today.Year(), today.Month(), today.Day()+1, 0, -config.LastMomentMinutes, 0, 0, today.Location())

	for {
		startIndex := 0
		for startIndex < lenSourceIds {
			if time.Now().UTC().After(lastMoment) || remainingRequests <= 0 {
				closeChannels()
				return
			}

			remainingRequests--
			chNumRequestsUpdated <- 1

			endIndex := minInt(startIndex+config.SourcesBatchSize, lenSourceIds)
			batchSources := config.SourceIds[startIndex:endIndex]

			articlesResponse, err := refr.newsApi.FetchEverything(&FetchEverythingParams{
				Sources:  batchSources,
				Page:     pageNum,
				PageSize: config.PageSize,
			})

			if err != nil {
				chError <- deepError.DeepErr{
					Function: prefix,
					Action:   "fetching articles",
					Cause:    err,
				}
			} else {
				chArticles <- articlesResponse.Articles
			}

			startIndex += config.SourcesBatchSize
		}

		pageNum++
		time.Sleep(time.Duration(config.SleepSeconds) * time.Second)
	}

}
