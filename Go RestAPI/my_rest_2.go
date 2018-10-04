if r.Method != "POST" {
	http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
	return
}