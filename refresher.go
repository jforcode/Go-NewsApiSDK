package newsApi

import (
	"time"
)

type Refresher struct {
	newsApi *NewsApi
}

func (refr *Refresher) Init(newsApi *NewsApi) {
	refr.newsApi = newsApi
}

func (refr *Refresher) DailyRefresh(config *RefresherConfig, chArticles chan []*ApiArticle, chNumTransactionsUpdated chan int, chError chan error) {
	// prefix := "newsApi.Refresher.FetchArticles"
	closeChannels := func() {
		close(chArticles)
		close(chNumTransactionsUpdated)
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
			chNumTransactionsUpdated <- 1

			endIndex := minInt(startIndex+config.SourcesBatchSize, lenSourceIds)
			batchSources := config.SourceIds[startIndex:endIndex]
			articlesResponse, err := refr.newsApi.FetchArticles(batchSources, pageNum, config.PageSize)

			if err != nil {
				chError <- err // TODO: better error format
			} else {
				chArticles <- articlesResponse.Articles
			}

			startIndex += config.SourcesBatchSize
		}

		pageNum++
		time.Sleep(time.Duration(config.SleepSeconds) * time.Second)
	}

}
