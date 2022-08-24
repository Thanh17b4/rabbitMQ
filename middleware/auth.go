package middleware

import (
	"net/http"
)

func RequestToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//token := r.Header.Get("Authorization")
		//tokenArray := strings.Split(token, " ")
		//if len(tokenArray) != 2 {
		//	responses.Error(w, r, http.StatusUnauthorized, nil, "token invalid")
		//	return
		//}
		//realToken := tokenArray[1]
		//if ok := model.VerifyToken(realToken); !ok {
		//	responses.Error(w, r, http.StatusUnauthorized, nil, "token invalid")
		//	return
		//}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
