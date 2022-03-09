package model

import "encoding/json"

type Links map[string]Link

type ShortLinks map[string]ShortLink

type UserLink struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

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

type Link struct {
	ID      string `json:"correlation_id"`
	URL     string `json:"original_url"`
	Deleted bool   `json:"-"`
}

type ShortLink struct {
	ID    string `json:"correlation_id"`
	Short string `json:"original_url"`
}
