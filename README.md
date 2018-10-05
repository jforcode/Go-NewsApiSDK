This package is kind of an sdk for NewsAPI (http://newsapi.org).

# Features

- fetch sources
  - fetches sources from {{ baseApiUrl }}/sources
  - no parameters
- fetch articles
  - fetches articles from {{ baseApi }}/everything
  - fetches based on sources
  - supports pagination
- daily refresh of articles
  - parameters: config, chArticles, chNumTransactionsUpdated, chError
    - config: as mentioned in daily refresh config
    - chArticles: channel of type []\*ApiArticle. Used to communicate the articles fetched from NewsAPI api.
    - chNumTransactionsUpdated: channel of type int. Used to communicate the number of transactions done by the routine.
    - chError: channel of type error. Used to communicate any errors in the routine.

# Configurable

## Root config
baseApiUrl
- (string, optional, default: http://newsapi.org/v2)

apiKey: client key for your account.
- (string, required)

## Daily Refresh config
remainingRequests
- number of requests remaining for the day
- (int, required)

sourceIds
- articles will be fetched for these source Ids in a round-robin batched fashion.
- ([]string, required)

sourcesBatchSize
- articles are fetched for sources in batch. This parameter specifies number of sources in one batch.
- (int, optional:20, [1, 20])

startPageNum
- (int, optional:1, [1, ])

pageSize
- articles are paginated. pageSize specifies number of articles in one request.
- (int, optional:10, [1, 100])

lastMomentMinutes
- number of minutes to subtract from the end of day. At this threshold time, fetching would stop for the day. // TODO: better name
- (int, optional:30, [1, ])

sleepSeconds
- after fetching 1 page for all the sources provided (in batch), this parameter specifies how many seconds to sleep for
- (int, optional:60, [1, ])

# TODO

- support all options (q, domains, excludeDomains etc.) for fetching articles
- support all options (category, language etc.) for fetching sources
- add fetch of top headlines
- add a command channel (in) for the refresher to handle command: exit. Can use in case of errors.
