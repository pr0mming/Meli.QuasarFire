package aggregate

// Normal Top Secret

type TopSecretSatelliteRequest struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type TopSecretRequest struct {
	Satellites []TopSecretSatelliteRequest `json:"satellites"`
}

type TopSecretCoordsResponse struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type TopSecretResponse struct {
	Position TopSecretCoordsResponse `json:"position"`
	Message  string                  `json:"message"`
}

// Split Top Secret

type TopSecretSplitSatelliteRequest struct {
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}
