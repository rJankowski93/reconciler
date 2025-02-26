package server

import (
	"encoding/json"
	"github.com/kyma-incubator/reconciler/pkg/logger"
	"github.com/pkg/errors"
	"net/http"
)

func SendHTTPError(w http.ResponseWriter, httpCode int, resp interface{}) {
	log := logger.NewLogger(false)

	payload, err := json.Marshal(resp)
	if err != nil {
		err = errors.Wrap(err, "failed to encode HTTP error response to JSON")
	} else {
		w.WriteHeader(httpCode)
		w.Header().Set("content-type", "application/json")
		if _, err = w.Write(payload); err == nil {
			log.Warnf("Sending HTTP error response (httpCode %d): %s", httpCode, payload)
			return
		}
		err = errors.Wrap(err, "failed to write error payload to HTTP response writer")
	}

	log.Errorf("Failed to process HTTP error response (fallback to %d - internal server error): %v",
		http.StatusInternalServerError, err)
	w.WriteHeader(http.StatusInternalServerError)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
