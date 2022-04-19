package model

import "encoding/json"

type (
	Link struct {
		ID      string `json:"correlation_id"`
		URL     string `json:"original_url"`
		Deleted bool   `json:"-"`
	}
	ShortLink struct {
		ID    string `json:"correlation_id"`
		Short string `json:"original_url"`
	}
	Links      map[string]Link
	ShortLinks map[string]ShortLink
	UserLink   struct {
		ShortURL    string `json:"short_url"`
		OriginalURL string `json:"original_url"`
	}
	Result struct {
		Result string `json:"result"`
	}
)

func (links Links) MarshalJSON() ([]byte, error) {
	var linksPrepared []UserLink
	for k, v := range links {
		linksPrepared = append(linksPrepared, UserLink{
			ShortURL:    k,
			OriginalURL: v.URL,
		})
	}
	return json.Marshal(linksPrepared)
}
