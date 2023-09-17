package entity

type Nation struct {
	Nationality_id int        `gorm:"primaryKey;autoIncrement;type:int" json:"nationality_id"`
	Negara         string     `gorm:"type:varchar(50)" json:"negara"`
	Kode           string     `gorm:"type:char(2)" json:"kode"`
	Customer       []Customer `json:"customer" gorm:"foreignKey:nationality_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Customer struct {
	Keluarga        []Family_list `json:"keluarga" gorm:"foreignKey:Cst_id;constraint:constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Cst_id          int           `gorm:"primaryKey;autoIncrement;type:int" json:"-"`
	Nationality_id  int           `gorm:"type:int" json:"-"`
	Nama            string        `gorm:"type:char(50)" json:"nama"`
	TanggalLahir    int           `gorm:"type:char(50)" json:"tanggal_lahir"`
	Telepon         string        `gorm:"type:varchar(20)" json:"telepon"`
	Kewarganegaraan string        `gorm:"type:varchar(50)" json:"kewarganegaraan"`
	Email           string        `gorm:"type:varchar(50)" json:"email"`
}

type Family_list struct {
	Fl_id        int    `gorm:"primaryKey;autoIncrement;type:int" json:"-"`
	Cst_id       uint   `gorm:"type:int" json:"-"`
	Hubungan     string `gorm:"varchar(50)" json:"hubungan"`
	Nama         string `gorm:"varchar(50)" json:"nama"`
	TanggalLahir string `gorm:"varchar(50)" json:"tanggal_lahir"`
}
