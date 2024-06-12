package videodownloader

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/charlesozo/omnicron-backendsever/golang-server/config"
	"github.com/charlesozo/omnicron-backendsever/golang-server/storage"
	"github.com/charlesozo/omnicron-backendsever/golang-server/utils"
	"net/http"
	"os"
	"path/filepath"
)

const outputPath = "./downloadedvideo/" //The output path where the downloaded video will be temporarily stored
type DownloadParams struct {
	URL        string `json:"url"`
	Resolution string `json:"resolution"`
}
type Responseparams struct {
	Response string `json:"response"`
}

func Download(w http.ResponseWriter, r *http.Request, cfg *config.ApiConfig) {
	ctx := context.Background()

	decode := json.NewDecoder(r.Body)
	params := DownloadParams{}
	err := decode.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error unmarshalling json, %v", err))
		return
	}
	stream := HandleStreamResolution(params.Resolution)
	outputName := "youtube"
	err = DownloadVideoData(params.URL, outputName, outputPath, stream)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	files, err := filepath.Glob(outputPath + outputName + ".*")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	videoPath := files[0]
	urlLink, err := storage.HandleFileUpload(ctx, videoPath, cfg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	//remove the video file after uploading
	err = os.Remove(videoPath)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, Responseparams{Response: urlLink})
}