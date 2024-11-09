package main

// https://mholt.github.io/json-to-go/

type SpriteSheet struct {
	Comment string            `json:"_comment"`
	Frames  map[string]Sprite `json:"frames"`
	Meta    Meta              `json:"meta"`
}
type Frame struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}
type SpriteSourceSize struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}
type SourceSize struct {
	W int `json:"w"`
	H int `json:"h"`
}
type Pivot struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
type Sprite struct {
	Frame            Frame            `json:"frame"`
	Rotated          bool             `json:"rotated"`
	Trimmed          bool             `json:"trimmed"`
	SpriteSourceSize SpriteSourceSize `json:"spriteSourceSize"`
	SourceSize       SourceSize       `json:"sourceSize"`
	Pivot            Pivot            `json:"pivot"`
}
type Size struct {
	W int `json:"w"`
	H int `json:"h"`
}
type Meta struct {
	App         string `json:"app"`
	Version     string `json:"version"`
	Image       string `json:"image"`
	Format      string `json:"format"`
	Size        Size   `json:"size"`
	Scale       string `json:"scale"`
	Smartupdate string `json:"smartupdate"`
}
