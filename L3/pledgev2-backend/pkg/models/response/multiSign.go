package response

type MultiSign struct {
	SpName           string   `json:"spName"`
	SpToken          string   `json:"spToken"`
	JpName           string   `json:"jpName"`
	JpToken          string   `json:"jpToken"`
	SpAddress        string   `json:"spAddress"`
	JpAddress        string   `json:"jpAddress"`
	SpHash           string   `json:"spHash"`
	JpHash           string   `json:"jpHash"`
	MultiSignAccount []string `json:"multiSignAccount"`
}
