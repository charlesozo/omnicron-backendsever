package generatemusic

import (
	"github.com/kingmariano/omnicron/config"
	rep "github.com/kingmariano/omnicron/packages/replicate"
	"github.com/kingmariano/omnicron/utils"
	"net/http"
)

func MusicGen(w http.ResponseWriter, r *http.Request, cfg *config.ApiConfig) {
	ctx := r.Context()
	model := r.URL.Query().Get("model")
	if model == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "music generation model query parameter is required")
		return
	}
	repMusicModel, err := rep.GetModelByName(model, rep.MusicModels)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "model not found")
		return
	}
	predictionFunc, ok := rep.MusicGenModels[*repMusicModel]
	if !ok {
		utils.RespondWithError(w, http.StatusNotFound, "model not found")
		return
	}
	predictionInput, err := processMusicModelInput(repMusicModel, ctx, r, cfg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	MusicGenPrediction, err := predictionFunc(ctx, cfg.ReplicateAPIKey, repMusicModel.Version, predictionInput, nil, false)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, MusicGenPrediction)

}
