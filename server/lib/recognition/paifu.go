package recognition

const (
	RecognitionAPIEndpoint = "http://recognition:8000"
	Recognition = "recognition"
)

type Paifu struct {
	Hands  []string `json:"hands"`
	Bakaze string   `json:"bakaze"`
	Jikaze string   `json:"jikaze"`
	Dora   []string `json:"dora"`
	Junme  int      `json:"junme"`
}