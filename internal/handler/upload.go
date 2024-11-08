package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/karthikeyaspace/cdn-go/internal/service"
	"github.com/rs/xid"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	indexFile, _, err := r.FormFile("index.html")
	if err != nil {
		http.Error(w, "Missing index.html file", http.StatusBadRequest)
		return
	}

	defer indexFile.Close()

	styleFile, _, err := r.FormFile("style.css")
	if err != nil {
		http.Error(w, "Missing style.css file", http.StatusBadRequest)
		return
	}

	defer styleFile.Close()

	key := xid.New().String()[:10]
	indexFileKey := fmt.Sprintf("%v/index.html", key)
	styleFileKey := fmt.Sprintf("%v/style.css", key)

	err = service.UploadToS3(indexFileKey, indexFile)
	if err != nil {
		http.Error(w, "Failed to upload index.html", http.StatusInternalServerError)
		return
	}

	err = service.UploadToS3(styleFileKey, styleFile)
	if err != nil {
		http.Error(w, "Failed to upload style.css", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Context-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"success": true, "key": key})
}
