package models

type UrlShortRequest struct {
	Url string `json:"url"`
}

type UrlShortResponse struct {
	ShortUrl string `json:"short_url"`
}

type GenerateQRCodeRequest struct {
	Url string `json:"url"`
}
