package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type Handler struct {
	store *Store    
}

func NewHandler(s *Store) *Handler { //Creating instance of handler
	return &Handler{store: s}
}

type request struct {          //consider this json obj or more better doing this for json encoding as this is stored in json format
	URL string `json:"url"`
}

//Generate Random shortLink (will use some other stronger encoding later)
func generateShortLink(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte,n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// shorten handler 
func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req request      //request instance
	json.NewDecoder(r.Body).Decode(&req)  //r is incoming Request , which is in form of json

	shortLink := generateShortLink(6)   //gen link of size 6chars alone (just arbitrary value)

	h.store.Save(shortLink,req.URL)

	json.NewEncoder(w).Encode(map[string]string{
		"shortURL":shortLink,
	})
}

// Get handler {slug} : any short link generated
func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[1:]

	if long,ok := h.store.GetLink(slug); ok {
		http.Redirect(w,r,long,http.StatusFound)
		return
	}

	http.NotFound(w,r)
}