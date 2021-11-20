package bing

import "time"

type Bing_res struct {
	Type            string `json:"_type"`
	Instrumentation struct {
		Type string `json:"_type"`
	} `json:"instrumentation"`
	ReadLink              string `json:"readLink"`
	WebSearchURL          string `json:"webSearchUrl"`
	TotalEstimatedMatches int    `json:"totalEstimatedMatches"`
	NextOffset            int    `json:"nextOffset"`
	Value                 []struct {
		WebSearchURL       string    `json:"webSearchUrl"`
		Name               string    `json:"name"`
		ThumbnailURL       string    `json:"thumbnailUrl"`
		DatePublished      time.Time `json:"datePublished"`
		ContentURL         string    `json:"contentUrl"`
		HostPageURL        string    `json:"hostPageUrl"`
		ContentSize        string    `json:"contentSize"`
		EncodingFormat     string    `json:"encodingFormat"`
		HostPageDisplayURL string    `json:"hostPageDisplayUrl"`
		Width              int       `json:"width"`
		Height             int       `json:"height"`
		Thumbnail          struct {
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"thumbnail"`
		ImageInsightsToken string `json:"imageInsightsToken"`
		InsightsMetadata   struct {
			RecipeSourcesCount      int `json:"recipeSourcesCount"`
			BestRepresentativeQuery struct {
				Text         string `json:"text"`
				DisplayText  string `json:"displayText"`
				WebSearchURL string `json:"webSearchUrl"`
			} `json:"bestRepresentativeQuery"`
			PagesIncludingCount int `json:"pagesIncludingCount"`
			AvailableSizesCount int `json:"availableSizesCount"`
		} `json:"insightsMetadata"`
		ImageID     string `json:"imageId"`
		AccentColor string `json:"accentColor"`
	} `json:"value"`
}