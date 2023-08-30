package resource

import "embed"

const (
	PathMusic = "audio/macan-asphalt-8.wav"
)

//go:embed audio/*.wav sprites/*.png
var FS embed.FS
