package videodownloader

import (
	"encoding/json"
	"fmt"
	"github.com/kingmariano/omnicron/config"
	"github.com/kingmariano/omnicron/utils"
	"net/http"
)

type DownloadParams struct {
	URL        string `json:"url"`
	Resolution string `json:"resolution"`
}
type Responseparams struct {
	Response string `json:"response"`
}

// DownloadVideo handles the video download process.
// It accepts a POST request with JSON body containing URL and resolution.
// It downloads the video from the provided URL, stores it in a temporary directory,
// uploads the video to Cloudinary, and returns the Cloudinary URL of the uploaded video.
//
// Parameters:
//  w http.ResponseWriter: The response writer for the HTTP request.
//  r *http.Request: The HTTP request.
//  cfg *config.ApiConfig: The API configuration.
//
// Return values:
//  None.

func DownloadVideo(w http.ResponseWriter, r *http.Request, cfg *config.ApiConfig) {
	ctx := r.Context()
	decode := json.NewDecoder(r.Body)
	params := DownloadParams{}
	err := decode.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error unmarshalling json, %v", err))
		return
	}
	//creates a temporary file to store the downloaded video
	folderPath, err := utils.CreateUniqueFolder(utils.BasePath)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	videoPath, err := DownloadVideoData(params.URL, utils.OutputName, folderPath, params.Resolution)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//upload the video file to cloudinary and return the file URL
	urlLink, err := utils.HandleFileUpload(ctx, videoPath, cfg.CloudinaryUrl)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Remove the directory after uploading
	err = utils.DeleteFolder(folderPath)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, Responseparams{Response: urlLink})
}
