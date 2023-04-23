package entities

type User struct {
	DefaultModel
}

type ResponseGoogleUserInfo struct {
	ID            string `json:"sub" gorm:"primarykey"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Locale        string `json:"locale"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	DefaultModelNonId
}

type FlickrUserResponse struct {
	ID       string
	Username string
	Fullname string
}
