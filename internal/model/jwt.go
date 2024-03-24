package model

type StandardClaim struct {
	Jti string `json:"jti"`
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Aud string `json:"aud"`
	Exp uint64 `json:"exp"`
	Nbf uint64 `json:"nbf"`
	Iat uint64 `json:"iat"`
}

type AccessClaim struct {
	StandardClaim
	User_Id uint64       `json:"user_id" gorm:"column:user_id;not null;foreignKey:UserID;references:users(ID)"`
	Data    UserResponse `json:"data"`
}
