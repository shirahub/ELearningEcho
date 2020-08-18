package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Username string `json:"username" form:"username" query:"username" name:"username"`
	Password string `json:"password" form:"password" query:"password" name:"password"`
}

type UserRegist struct {
	Username string `json:"username" form:"username" query:"username" name:"username"`
	Password string `json:"password" form:"password" query:"password" name:"password"`
	Fullname string `json:"fullname" form:"fullname" query:"fullname" name:"fullname"`
}

type UserHistory struct {
	PaketId     int
	NamaPaket   string
	Nilai       int
	Waktu       time.Time
	Waktustring string
}

type UserData struct {
	UserId   int
	Username string
	Password string
	Fullname string
}

type Claims struct {
	Name   string `json:"name"`
	UserId int    `json:"userid"`
	jwt.StandardClaims
}

type PaketSoal struct {
	Id_PaketSoal int
	Tingkat      string
	Kelas        string
	Mapel        string
	Tema         string
}

// type Soal struct {
// 	Pertanyaan string `json: "pertanyaan"`
// 	Pilihan1   string `json: "pilihan1"`
// 	Pilihan2   string `json: "pilihan2"`
// 	Pilihan3   string `json: "pilihan3"`
// 	Pilihan4   string `json: "pilihan4"`
// }

type Soal struct {
	Id_Soal    int
	Pertanyaan string    `json: "pertanyaan"`
	Pilihan    [4]string `json: "pilihan"`
}

type M map[string]interface{}
