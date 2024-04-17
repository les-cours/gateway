package service

import "github.com/dgrijalva/jwt-go"

type Permissions struct {
	Bots     bool `json:"bots"`
	Triggers bool `json:"triggers"`
	Tickets  bool `json:"tickets"`
	Profiles bool `json:"profiles"`
	Kbas     bool `json:"kbas"`
	Settings bool `json:"settings"`
}

type UserToken struct {
	jwt.StandardClaims
	Username      string      `json:"username"`
	AccountID     string      `json:"accountID"`
	Email         string      `json:"email"`
	ID            string      `json:"id"`
	CoBrowsing    bool        `json:"coBrowsing"`
	ScreenShare   bool        `json:"screenShare"`
	AudioDownload bool        `json:"audioDownload"`
	VideoDownload bool        `json:"videoDownload"`
	Create        Permissions `json:"create"`
	Update        Permissions `json:"update"`
	Read          Permissions `json:"read"`
	Delete        Permissions `json:"delete"`
}
