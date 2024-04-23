package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/jmmalqui/terminal-renderer/internal/buffer"
	"github.com/jmmalqui/terminal-renderer/internal/color"
	"github.com/jmmalqui/terminal-renderer/internal/vec"
	"golang.org/x/term"
)

var clear map[string]func()

type Terminal struct {
	Size       vec.Vec2
	WorkBuffer buffer.PixelBuffer
	ShowBuffer buffer.PixelBuffer
}

func (t *Terminal) Flip() {
	var stringBuffer string
	{
		for idx, value := range t.WorkBuffer.Data {
			stringBuffer += value.ColorToString()
			if idx%t.Size.X == 0 {
				stringBuffer += "\n"
			}
		}
	}
	t.Clear()
	fmt.Println(stringBuffer)

}

func MakeTerminal() Terminal {
	width, height, err := term.GetSize(0)
	size :=
		vec.Vec2{X: 0, Y: 0}
	if err == nil {
		size.X = width
		size.Y = height
		fmt.Printf("Building terminal of size %d, %d\n", width, height)
	} else {
		panic("Something went wrong when getting the terminal size")
	}
	colorBuffer := make([]color.Color, width*height)
	for i := 0; i < width*height; i++ {
		colorBuffer[i] = color.Color{R: 0, G: 0, B: 0}
	}
	var pBuffer buffer.PixelBuffer
	{
		pBuffer.Data = colorBuffer
	}
	terminal := Terminal{Size: size, WorkBuffer: pBuffer, ShowBuffer: pBuffer}
	terminal.InitClear()
	return terminal
}

func (t *Terminal) InitClear() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (t *Terminal) Clear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Platform not supported")
	}
}

func (t *Terminal) DrawPixel(x int, y int, c color.Color) {
	t.WorkBuffer.Data[(y*t.Size.X)+x] = c
}

func (t *Terminal) Fill(c vec.Vec3) {
	for idx := range t.WorkBuffer.Data {
		var trueColor color.Color
		{
			trueColor = color.Vec3ToColor(c)

		}

		t.WorkBuffer.Data[idx] = trueColor
	}
}
