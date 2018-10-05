package newsApi

import "errors"

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
	if config.SourceIds == nil || len(config.SourceIds) == 0 {
		return errors.New("Invalid source ids")
	}

	if config.SourcesBatchSize == 0 {
		config.SourcesBatchSize = 20
	} else if config.SourcesBatchSize < 1 || config.SourcesBatchSize > 20 {
		return errors.New("Invalid sources batch size")
	}

	if config.StartPageNum == 0 {
		config.StartPageNum = 1
	} else if config.StartPageNum < 1 {
		return errors.New("Invalid start page number")
	}

	if config.PageSize == 0 {
		config.PageSize = 10
	} else if config.PageSize < 1 || config.PageSize > 100 {
		return errors.New("Invalid page size")
	}

	if config.LastMomentMinutes == 0 {
		config.LastMomentMinutes = 30
	} else if config.LastMomentMinutes < 1 {
		return errors.New("Invalid last moment minutes")
	}

	if config.SleepSeconds == 0 {
		config.SleepSeconds = 60
	} else if config.SleepSeconds < 1 {
		return errors.New("Invalid sleep seconds")
	}

	return nil
}
