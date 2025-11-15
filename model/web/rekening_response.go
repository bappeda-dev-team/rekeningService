package web

type RekeningResponse struct {
	Id           int     `json:"id"`
	KodeRekening string  `json:"kode_rekening"`
	NamaRekening *string `json:"nama_rekening"`
	Tahun        string  `json:"tahun"`
}
