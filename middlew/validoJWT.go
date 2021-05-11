package middlew

import (
	"net/http"

	"github.com/diegoolp/DL_TWITTOR_GOLANG/routers"
)

/*ValidoJWT permite validar el JWT que nos viene en la petición*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Autorization"))
		if err != nil {
			http.Error(w, "Error en el Token !"+err.Error(), http.StatusBadRequest)
		}

		next.ServeHTTP(w, r)
	}
}
