package buffer

import "github.com/jmmalqui/terminal-renderer/internal/color"

type PixelBuffer struct {
	Data []color.Color
}

func (pb *PixelBuffer) Build() {
}
