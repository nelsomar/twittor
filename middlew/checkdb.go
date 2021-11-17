package middlew

import (
	"net/http"

	"github.com/nelsomar/twittor/bd"
)

/* CheckDB, middleware check status of database */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Lost connection to database", 500)
		}
		next.ServeHTTP(w, r)
	}
}
