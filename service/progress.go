package service

import (
	"fmt"
	"os"

	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

// NewProgressBar
//
//	@param maxValue
//	@param message
//	@return *progressbar.ProgressBar
func NewProgressBar(maxValue int, message string) *progressbar.ProgressBar {
	return progressbar.NewOptions(maxValue,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(20),
		progressbar.OptionSetDescription(message),
		progressbar.OptionShowDescriptionAtLineEnd(),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
}
