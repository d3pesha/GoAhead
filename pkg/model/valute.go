package model

type Valute struct {
	Vcurs   float64 `xml:"Vcurs" json:"curs"`
	VchCode string  `xml:"VchCode" json:"vchCode"`
}

type ValuteData struct {
	ValuteCursOnDate []Valute `xml:"diffgram>ValuteData>ValuteCursOnDate"`
}

type GetCursOnDateResult struct {
	ValuteData ValuteData `xml:"GetCursOnDateResult"`
}

type GetCursOnDateResponse struct {
	GetCursOnDateResult GetCursOnDateResult `xml:"GetCursOnDateResponse"`
}

type Envelope struct {
	Body GetCursOnDateResponse `xml:"Body"`
}
