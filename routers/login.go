package routers

import (
	"encoding/json"
	"net/http"

	"github.com/diegoolp/DL_TWITTOR_GOLANG/bd"
	"github.com/diegoolp/DL_TWITTOR_GOLANG/jwt"
	"github.com/diegoolp/DL_TWITTOR_GOLANG/models"
)

/*Login realiza el login*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o Contraseña inválidos"+err.Error(), 400)
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)

	if !existe {
		http.Error(w, "Usuario y/o Contraseña inválidos", 400)
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrio un error al generar el Token correspondiente"+err.Error(), 400)
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	/*
		GRABACIÓN DE COOKIES
		expirationTime:=time.Now().Add(24*time.Hour)
		http.SetCookie(w,&http.Cookie{
			Name:"token",
			Value: jwtKey,
			Expires: expirationTime,
		})
	*/
}
