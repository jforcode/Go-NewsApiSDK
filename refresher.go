package newsApi

import (
	"time"

	"github.com/jforcode/DeepError"
)

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
			articlesResponse, err := refr.newsApi.FetchArticles(batchSources, pageNum, config.PageSize)

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
