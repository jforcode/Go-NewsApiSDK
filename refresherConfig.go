package newsApi

import (
	"errors"
	"strings"

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
