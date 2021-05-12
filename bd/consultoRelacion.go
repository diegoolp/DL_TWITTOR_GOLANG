package bd

import (
	"context"
	"time"

	"github.com/diegoolp/DL_TWITTOR_GOLANG/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoRelacion consulta la existencia de una relación en la BD*/
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return false, err
	}

	return true, nil
}
