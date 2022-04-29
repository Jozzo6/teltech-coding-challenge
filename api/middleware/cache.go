package middleware

import (
	"coding_challenge/models"
	"encoding/json"
	"net/http"
	"time"
)

var cachedRequests map[string]Cache = make(map[string]Cache)

type CacheHandler struct {
	handler http.Handler
}

func (c *CacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if val, ok := cachedRequests[r.URL.String()]; ok {
		if ok {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			val.LastUsed = time.Now()
			cachedRequests[r.URL.String()] = val
			json.NewEncoder(w).Encode(val.Result)
			return
		}
	}
	c.handler.ServeHTTP(w, r)
}

func NewCache(handlerToWrap http.Handler) *CacheHandler {
	StartTimer()
	return &CacheHandler{handlerToWrap}
}

func SaveRequest(url string, result models.Result) {
	result.Cached = true
	cachedRequests[url] = Cache{url, result, time.Now()}
}

func StartTimer() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				for key, element := range cachedRequests {
					lastUsed := element.LastUsed
					if lastUsed.Add(1 * time.Minute).Before(time.Now()) {
						delete(cachedRequests, key)
					}
				}
			}
		}
	}()
}

type Cache struct {
	Url      string
	Result   models.Result
	LastUsed time.Time
}
