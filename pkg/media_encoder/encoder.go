package media_encoder

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// Encoder interface
type Encoder interface {
	EncodeToHLS() error
}

// MediaEncoderToHLS struct
type MediaEncoderToHLS struct {
	InputFile          string
	OutputFile         string
	Preset             string
	ConstantRateFactor string
	HLSTime            string
	HLSListSize        string
	HLSWrap            string
}

// NewMediaEncoderToHLS function
func (m *MediaEncoderToHLS) EncodeToHLS() error {
	// FFmpeg command
	cmd := exec.Command("ffmpeg", "-i", m.InputFile, "-c:v", "libx264", "-preset", m.Preset, "-crf", m.ConstantRateFactor, "-hls_time", m.HLSTime, "-hls_list_size", m.HLSListSize, "-hls_wrap", m.HLSWrap, m.OutputFile)

	// Redirect output and error streams
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// Run the command
	err := cmd.Run()
	if err != nil {
		return err
	}

	// Check if the output contains error information and return an error with the error inforamtion if it does
	if strings.Contains(out.String(), "error") {
		return fmt.Errorf("ffmpeg command failed: %s", out.String())
	}

	log.Println("Conversion completed successfully.")
	return nil
}
