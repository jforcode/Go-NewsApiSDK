package newsApi

type Category string

const (
	CAT_BUSINESS      Category = "business"
	CAT_ENTERTAINMENT Category = "entertainment"
	CAT_GENERAL       Category = "general"
	CAT_HEALTH        Category = "health"
	CAT_SCIENCE       Category = "science"
	CAT_SPORTS        Category = "sports"
	CAT_TECHNOLOGY    Category = "technology"
)

type Language string

const (
	LANG_ARABIC        Language = "ar"
	LANG_AR            Language = "ar"
	LANG_GERMAN        Language = "de"
	LANG_DE            Language = "de"
	LANG_ENGLISH       Language = "en"
	LANG_EN            Language = "en"
	LANG_SPANISH       Language = "es"
	LANG_ES            Language = "es"
	LANG_FRENCH        Language = "fr"
	LANG_FR            Language = "fr"
	LANG_HEBREW        Language = "he"
	LANG_HE            Language = "he"
	LANG_ITALIAN       Language = "it"
	LANG_IT            Language = "it"
	LANG_DUTCH         Language = "nl"
	LANG_NL            Language = "nl"
	LANG_NORWEGIAN     Language = "no"
	LANG_NO            Language = "no"
	LANG_PORTUGUESE    Language = "pt"
	LANG_PT            Language = "pt"
	LANG_RUSSIAN       Language = "ru"
	LANG_RU            Language = "ru"
	LANG_NORTHERN_SAMI Language = "se"
	LANG_SE            Language = "se"
	LANG_UD            Language = "ud" // TODO: which country/full-form
	LANG_CHINESE       Language = "zh"
	LANG_ZH            Language = "zh"
)

type Country string

const () // TODO

type SortBy string

const (
	SORT_BY_RELEVANCY    SortBy = "relevancy"
	SORT_BY_POPULARITY   SortBy = "popularity"
	SORT_BY_PUBLISHED_AT SortBy = "publishedAt"
)
