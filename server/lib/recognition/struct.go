package recognition

const (
	RecognitionAPIEndpoint = "http://recognition:8000"
	Recognition = "recognition"
)

type Paifu struct {
	Zikaze         int   `json:"zikaze"`
	Bakaze         int   `json:"bakaze"`
	Turn           int   `json:"turn"`
	SyantenType    int   `json:"syanten_type"`
	DoraIndicators []int `json:"dora_indicators"`
	Flag           int   `json:"flag"`
	Handtiles      []int `json:"hand_tiles"`
	MeldedBlocks   []int `json:"melded_blocks"`
}