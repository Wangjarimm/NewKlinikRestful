package dokter

type Dokter struct {
	Id       int    `json:"id"`
	Nid      int    `json:"nid"`
	Nama     string `json:"nama"`
	Keahlian string `json:"keahlian"`
	Nohp     string `json:"no_hp"`
}