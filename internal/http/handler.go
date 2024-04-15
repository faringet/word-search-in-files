package http

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"word-search-in-files/pkg/searcher"
)

type SearchHandler struct {
	Searcher *searcher.Searcher
	logger   *zap.Logger
}

func NewSearchHandler(searcher *searcher.Searcher, logger *zap.Logger) *SearchHandler {
	return &SearchHandler{Searcher: searcher, logger: logger}
}

func (h *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling search request")

	word := r.URL.Query().Get("word")
	if word == "" {
		h.logger.Error("Missing 'word' parameter in request")
		http.Error(w, "Parameter 'word' is required", http.StatusBadRequest)
		return
	}

	files, err := h.Searcher.Search(word)
	if err != nil {
		h.logger.Error("Error searching for files", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Отправляем результаты в формате JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(files); err != nil {
		h.logger.Error("Error encoding JSON response", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	h.logger.Info("Search request handled")
}
