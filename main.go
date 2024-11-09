package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

var (
	input  = flag.String("i", "", "image file")
	output = flag.String("o", "", "metadata file")
	rows   = flag.Int("rows", 1, "number of rows")
	cols   = flag.Int("columns", 1, "number of columns")
)

func main() {
	flag.Parse()
	reader, err := os.Open(*input)
	if err != nil {
		slog.Error("failed to open image", "error", err)
		return
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		slog.Error("failed to decode image", "error", err)
		return
	}

	slog.Info("image decoded", "width", img.Bounds().Dx(), "height", img.Bounds().Dy())

	sss := SpriteSourceSize{W: img.Bounds().Dx(), H: img.Bounds().Dy()}
	sheet := SpriteSheet{
		Comment: getCommandLine(),
		Frames:  make(map[string]Sprite),
		Meta: Meta{
			App:     "github.com/ubanita/spritesheetgen",
			Version: "1.0",
			Image:   filepath.Base(*input), // local name
			Format:  "RGBA8888",
			Size: Size{
				W: img.Bounds().Dx(),
				H: img.Bounds().Dy(),
			},
			Scale: "1",
		},
	}
	dx := img.Bounds().Dx() / *cols
	dy := img.Bounds().Dy() / *rows
	ss := SourceSize{W: dx, H: dy}
	count := 1
	offx := 0
	offy := 0
	for y := 0; y < *rows; y++ {
		offx = 0
		for x := 0; x < *cols; x++ {
			s := Sprite{
				SourceSize:       ss,
				SpriteSourceSize: sss,
				Frame: Frame{
					X: offx,
					Y: offy,
					W: dx,
					H: dy,
				},
			}
			sheet.Frames[fmt.Sprintf("frame_%d.png", count)] = s
			offx += dx
			count++
		}
		offy += dy
	}

	outFile := *output
	if outFile == "" {
		outFile = strings.TrimSuffix(*input, ".png") + ".json"
	}
	writer, err := os.Create(outFile)
	if err != nil {
		slog.Error("failed to create output file", "error", err)
		return
	}
	enc := json.NewEncoder(writer)
	enc.SetIndent("", "  ")
	err = enc.Encode(sheet)
	if err != nil {
		slog.Error("failed to encode json", "error", err)
		return
	}
	defer writer.Close()
}

func getCommandLine() string {
	return strings.Join(os.Args, " ")
}
