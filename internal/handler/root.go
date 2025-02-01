package handler

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, StatusSuccess, "patients-golang-api is up and running !", nil, nil)
}

func InvalidPathHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, StatusError, "Invalid endpoint or URL path. Please check the documentation.", nil, nil)
}
