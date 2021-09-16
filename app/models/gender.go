package models

type DBGender struct {
	ID     uint
	Gender string
}

type JGender struct {
	ID     int64  `json:"id"`
	Gender string `json:"gender"`
}

type RGender struct {
	Gender string `json:"gender"`
}
