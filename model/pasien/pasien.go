package pasien



type Pasien struct {
	Id           int          `json:"id"`
	Namalengkap  string       `json:"nama_lengkap"`
	Nik          string       `json:"nik"`
	Jeniskelamin string       `json:"jenis_kelamin"`
	Tempatlahir  string       `json:"tempat_lahir"`
	Tanggallahir string       `json:"tanggal_lahir"`
	Alamat       string       `json:"alamat"`
	Nohp         string       `json:"no_hp"`
	Reservasi    string       `json:"reservasi"`
	IdJadwal     string       `json:"id_jadwal"`
	TglReservasi string       `json:"tgl_reservasi"`
}