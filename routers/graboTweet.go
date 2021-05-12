package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/diegoolp/DL_TWITTOR_GOLANG/bd"
	"github.com/diegoolp/DL_TWITTOR_GOLANG/models"
)

/*GraboTweet permite grabar el tweet en la base de datos*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, intente nuevamente "+err.Error(), 400)
		return
	}

	registro := models.GraboTweet{
		UserId:  IDUsario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, intente nuevamente "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}