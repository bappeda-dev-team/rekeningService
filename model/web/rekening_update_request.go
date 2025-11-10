package web

type RekeningUpdateRequest struct {
	Id           int    `json:"id" validate:"required"`
	KodeRekening string `json:"kode_rekening" validate:"required"`
	NamaRekening string `json:"nama_rekening" validate:"required"`
	Tahun        string `json:"tahun" validate:"required"`
}
