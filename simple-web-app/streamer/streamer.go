package streamer

import (
	"fmt"
	"github.com/tsawler/toolbox"
	"path"
	"strings"
)

// get back when once the encode is finished
type ProcessingMessage struct {
	ID          int
	Successfull bool
	Message     string
	OutputFiile string
}

// hold the unit of work thatworker pool will perform
type VideoProcessingJob struct {
	Video Video
}

// ADAPTER
// will encode or simulate encoding  videos
type Processor struct {
	Engine Encoder
}

type Video struct {
	ID           int
	InputFile    string
	OutputDir    string
	EncodingType string
	NotifyChan   chan ProcessingMessage
	Options      *VideoOptions
	Encoder      Processor
}

type VideoOptions struct {
	RenameOutput    bool
	SegmentDuration int
	MaxRate1080p    string
	MaxRate720p     string
	MaxRate480p     string
}

// Factory function
func (vd *VideoDispatcher) NewVideo(id int, inputFile, outputDir, encodingType string, notifyChan chan ProcessingMessage, options *VideoOptions) *Video {
	if options == nil {
		options = &VideoOptions{}
	}

	fmt.Println("NewVideo(): new video created: ", id, inputFile)

	return &Video{
		ID:           id,
		InputFile:    inputFile,
		OutputDir:    outputDir,
		EncodingType: encodingType,
		NotifyChan:   notifyChan,
		Encoder:      vd.Processor,
		Options:      options,
	}
}

// all pushes to notify channel resides inside this function
func (v *Video) encode() {
	// call the encoder
	var fileName string

	switch v.EncodingType {
	case "mp4":
		// encode video
		fmt.Println("v.encode(): about to encode video to mp4 ", v.ID)
		name, err := v.encodeToMP4()
		if err != nil {
			// send info to the  notifyChann
			v.sendToNotifyChan(false, "",
				fmt.Sprintf("encoding failed for %d: %s", v.ID, err.Error()))
			return
		}

		fileName = fmt.Sprintf("%s.mp4", name)

	case "hls":
		// encode video
		fmt.Println("v.encode(): about to encode video to mp4 ", v.ID)
		name, err := v.encodeToHLS()
		if err != nil {
			// send info to the  notifyChann
			v.sendToNotifyChan(false, "",
				fmt.Sprintf("encoding failed for %d: %s", v.ID, err.Error()))
			return
		}

		fileName = fmt.Sprintf("%s.m3u8", name)

	default:
		fmt.Println("v.encode(): err trying to encode video", v.ID)
		v.sendToNotifyChan(false, "",
			fmt.Sprintf("error processing for %d: unknown or invalid encoding type", v.ID))
		return
	}

	// send info to the notifyChann
	fmt.Println("v.encode(): sending success message for video id", v.ID, "to notify channel")
	v.sendToNotifyChan(true, fileName,
		fmt.Sprintf("video id %d processed and saved as %s", v.ID, fmt.Sprintf("%s/%s", v.OutputDir, fileName)))
}

func (v *Video) encodeToMP4() (string, error) {
	baseFileName := "/output.mp4"
	//baseFileName := ""

	fmt.Println("v.encodeToMP4(): about to try to encode video id to mp4 ", v.ID)
	if !v.Options.RenameOutput {
		b := path.Base(v.InputFile)
		baseFileName = strings.TrimSuffix(b, path.Ext(b))
	} else {
		//baseFileName = path.Base(baseFileName)

		// TODO generate random file name
		var t toolbox.Tools
		baseFileName = t.RandomString(10)
	}

	// encode video
	err := v.Encoder.Engine.EncodeToMP4(v, baseFileName)
	if err != nil {
		return "", err
	}

	fmt.Println("v.encodeToMP4(): video successfully encoded to mp4 ", v.ID)
	return baseFileName, nil
}

func (v *Video) encodeToHLS() (string, error) {
	baseFileName := "/output.m3u8"
	//baseFileName := ""

	fmt.Println("v.encodeToHLS(): about to try to encode video id to hls ", v.ID)

	if !v.Options.RenameOutput {
		b := path.Base(v.InputFile)
		baseFileName = strings.TrimSuffix(b, path.Ext(b))
	} else {
		// TODO generate random file name
		baseFileName = path.Base(baseFileName)
	}

	// encode video
	err := v.Encoder.Engine.EncodeToHLS(v, baseFileName)
	if err != nil {
		return "", err
	}

	fmt.Println("v.encodeToHLS(): video successfully encoded to hls ", v.ID)
	return baseFileName, nil
}

// sends to notify channel pushes a message down the notify channel
func (v *Video) sendToNotifyChan(success bool, fileName, message string) {
	fmt.Println("v.sendToNotifyChan(): sending message to notify channel for video id", v.ID)
	v.NotifyChan <- ProcessingMessage{
		ID:          v.ID,
		Successfull: success,
		Message:     message,
		OutputFiile: fileName,
	}
}

// New creates and returns a new worker pool
func New(jobQueue chan VideoProcessingJob, maxWorkers int) *VideoDispatcher {
	fmt.Println("New(): creating new worker pool")
	workerPool := make(chan chan VideoProcessingJob, maxWorkers)

	// TODO implement processor logic
	var e VideoEncoder
	p := Processor{
		Engine: &e,
	}

	return &VideoDispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		WorkerPool: workerPool,
		Processor:  p,
	}

}
