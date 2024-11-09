package handler

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karthikeyaspace/cdn-go/internal/service"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	indexFile := fmt.Sprintf("%v/index.html", key)
	styleFile := fmt.Sprintf("%v/style.css", key)

	indexBytes, err := service.GetFilesFromS3(indexFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error getting file from S3: %v", err)
		return
	}

	styleBytes, err := service.GetFilesFromS3(styleFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error getting file from S3: %v", err)
		return
	}

	responseBytes := bytes.Replace(indexBytes, []byte("</head>"), []byte(fmt.Sprintf("<style>%s</style></head>", styleBytes)), 1)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)
}
