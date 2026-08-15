package kernel

type ConvertRequest struct {
	Category string  `json:"category"`
	Value    float64 `json:"value"`
	UnitFrom string  `json:"from"`
	UnitTo   string  `json:"to"`
}
