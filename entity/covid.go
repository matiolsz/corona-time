package entity

type Covid struct {
	Date                     string `json:"date"`
	States                   string `json:"states"`
	Positive                 int64  `json:"positive"`
	Negative                 int64  `json:"negative"`
	Pending                  string `json:"pending"`
	HospitalizedCurrently    string `json:"hospitalizedCurrently"`
	HospitalizedCumulative   string `json:"hospitalizedCumulative"`
	InIcuCurrently           string `json:"inIcuCurrently"`
	InIcuCumulative          string `json:"inIcuCumulative"`
	OnVentilatorCurrently    string `json:"onVentilatorCurrently"`
	OnVentilatorCumulative   string `json:"onVentilatorCumulative"`
	DateChecked              string `json:"dateChecked"`
	Death                    string `json:"death"`
	Hospitalized             int64  `json:"hospitalized"`
	TotalTestResults         string `json:"totalTestResults"`
	LastModified             string `json:"lastModified"`
	Recovered                string `json:"recovered"`
	Total                    string `json:"total"`
	PosNeg                   string `json:"posNeg"`
	DeathIncrease            string `json:"deathIncrease"`
	HospitalizedIncrease     string `json:"hospitalizedIncrease"`
	NegativeIncrease         string `json:"negativeIncrease"`
	PositiveIncrease         string `json:"positiveIncrease"`
	TotalTestResultsIncrease string `json:"totalTestResultsIncrease"`
	Hash                     string `json:"hash"`
}
