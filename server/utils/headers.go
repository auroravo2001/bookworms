package utils

import "net/http"

func HandleCors(w *http.ResponseWriter, allowedMethods string) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", allowedMethods)
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
	(*w).Header().Set("Access-Control-Max-Age", "3600")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}
