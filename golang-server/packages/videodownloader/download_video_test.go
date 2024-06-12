package videodownloader

import (
	"os"
	"testing"
)

func TestDownloadVideoData(t *testing.T) {
	type args struct {
		url        string
		outputName string
		outputPath string
		resolution string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Download 1080p video",
			args: args{
				url:        "https://www.youtube.com/watch?v=3ODP6tTpjqA",
				outputName: "test_video_1080p",
				outputPath: "./test_videos",
				resolution: "1080p",
			},
		},
		{
			name: "Download 720p video",
			args: args{
				url:        "https://www.youtube.com/watch?v=pPtK5KTDcXM",
				outputName: "test_video_720p",
				outputPath: "./test_videos",
				resolution: "720p",
			},
		},
		{
			name: "Download 480p video",
			args: args{
				url:        "https://www.youtube.com/watch?v=3ODP6tTpjqA",
				outputName: "test_video_480p",
				outputPath: "./test_videos",
				resolution: "",
			},
		},
		{
			name: "Download 360p video",
			args: args{
				url:        "https://www.youtube.com/watch?v=xN1-2p06Urc",
				outputName: "test_video_360p",
				outputPath: "./test_videos",
				resolution: "360p",
			},
		},
		{
			name: "Download 240p video",
			args: args{
				url:        "https://youtu.be/ZT0yQgUIZho",
				outputName: "test_video_240p",
				outputPath: "./test_videos",
				resolution: "360p",
			},
		},
	}
	outputDir := "./test_videos"

	// Create the output directory before running tests
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create output directory: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stream := HandleStreamResolution(tt.args.resolution)
			err := DownloadVideoData(tt.args.url, tt.args.outputName, tt.args.outputPath, stream)
			if err != nil {
				t.Errorf("DownloadVideoData() error = %v", err)
				return
			}
			err = deleteContents(tt.args.outputPath)
			if err != nil {
				t.Errorf("Failed to cleanup: %v", err)
			}
		})

	}
	// Cleanup after all tests
	err = deleteContents(outputDir)
	if err != nil {
		t.Errorf("Failed to cleanup output directory: %v", err)
	}

	err = os.Remove(outputDir)
	if err != nil {
		t.Errorf("Failed to remove output directory: %v", err)
	}

}
