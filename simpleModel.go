package czml

// Simple Model describes a 3D model
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Model
type SimpleModel struct {
	Gltf             string  `json:"gltf"`
	Scale            float64 `json:"scale,omitempty"`
	MinimumPixelSize float64 `json:"minimumPixelSize,omitempty"`
}
