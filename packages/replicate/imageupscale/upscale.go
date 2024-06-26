package imageupscale

import (
	"github.com/kingmariano/omnicron/config"
	rep "github.com/kingmariano/omnicron/packages/replicate"
	"github.com/kingmariano/omnicron/utils"
	"net/http"
)

func ImageUpscale(w http.ResponseWriter, r *http.Request, cfg *config.ApiConfig) {
	ctx := r.Context()
	model := r.URL.Query().Get("model")
	if model == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "image upscale model query parameter is required")
		return
	}
	repImageUpscaleModel, err := rep.GetModelByName(model, rep.ImageUpscaleModels)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "model not found")
		return
	}
	predictionFunc, ok := rep.ImageUpscaleGenModels[*repImageUpscaleModel]
	if !ok {
		utils.RespondWithError(w, http.StatusNotFound, "model not found")
		return
	}
	predictionInput, err := processImageUpscaleModelInput(repImageUpscaleModel, ctx, r, cfg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ImageUpscalePrediction, err := predictionFunc(ctx, cfg.ReplicateAPIKey, repImageUpscaleModel.Version, predictionInput, nil, false)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, ImageUpscalePrediction)
}
