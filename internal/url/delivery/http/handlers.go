package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"url-shortener/config"
	"url-shortener/internal/models"
	sUrl "url-shortener/internal/url"
)

type urlHandler struct {
	cfg        *config.Config
	urlUseCase sUrl.UrlUseCase
}

func NewAuthHandlers(cfg *config.Config, urlUseCase sUrl.UrlUseCase) sUrl.UrlHandler {
	return &urlHandler{
		cfg:        cfg,
		urlUseCase: urlUseCase,
	}
}

// RedirectToOriginalUrl implements url.UrlHandler.
func (u *urlHandler) RedirectToOriginalUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("RedirectToOriginalUrl")
		ctx := context.Background()

		urlToken := r.URL.Path[3:]
		urlValue, err := u.urlUseCase.GetUrl(ctx, urlToken)
		if err != nil {
			log.Printf("Error getting URL: %v", err)
			http.Error(w, "Error getting URL", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, urlValue, http.StatusSeeOther)
	}
}

// CreateShortUrl implements url.UrlHandler.
func (u *urlHandler) CreateShortUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("CreateShortUrl")
		ctx := context.Background()

		var urlShortRequest models.UrlShortRequest
		err := json.NewDecoder(r.Body).Decode(&urlShortRequest)
		if err != nil {
			log.Printf("Error decoding request: %v", err)
			http.Error(w, "Error decoding request", http.StatusBadRequest)
			return
		}

		result, err := url.QueryUnescape(urlShortRequest.Url)
		if err != nil {
			log.Printf("Error unescaping URL: %v", err)
			http.Error(w, "Error unescaping URL", http.StatusBadRequest)
			return
		}

		result = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(result, "")

		urlToken := u.urlUseCase.GenerateUrlToken(ctx, result)
		shortUrl := u.urlUseCase.GenerateShortUrl(ctx, urlToken)

		err = u.urlUseCase.SaveUrl(ctx, urlToken, urlShortRequest.Url)
		if err != nil {
			log.Printf("Error saving URL: %v", err)
			http.Error(w, "Error saving URL", http.StatusInternalServerError)
			return
		}

		response := models.UrlShortResponse{
			ShortUrl: shortUrl,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusOK)
	}
}
