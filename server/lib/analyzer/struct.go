package analyzer

const (
	AnalyzerAPIEndpoint = "http://analyzer:8888"
)

type Analysis struct {
	Success bool `json:"success"`
	Request struct {
		Zikaze         int           `json:"zikaze"`
		Bakaze         int           `json:"bakaze"`
		Turn           int           `json:"turn"`
		SyantenType    int           `json:"syanten_type"`
		DoraIndicators []int         `json:"dora_indicators"`
		Flag           int           `json:"flag"`
		HandTiles      []int         `json:"hand_tiles"`
		MeldedBlocks   []interface{} `json:"melded_blocks"`
	} `json:"request"`
	Response struct {
		ResultType int `json:"result_type"`
		Syanten    int `json:"syanten"`
		Time       int `json:"time"`
		Candidates []struct {
			Tile          int  `json:"tile"`
			SyantenDown   bool `json:"syanten_down"`
			RequiredTiles []struct {
				Tile  int `json:"tile"`
				Count int `json:"count"`
			} `json:"required_tiles"`
			ExpValues   []float64 `json:"exp_values"`
			WinProbs    []float64 `json:"win_probs"`
			TenpaiProbs []float64 `json:"tenpai_probs"`
		} `json:"candidates"`
	} `json:"response"`
}