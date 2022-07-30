package user

type User struct {
	ID                                 string `json:"id"`
	AppID                              string `json:"app_id"`
	EmailAddress                       string `json:"email_address"`
	PhoneNO                            string `json:"phone_no"`
	ImportFromApp                      string `json:"import_from_app"`
	CreatedAt                          uint32 `json:"created_at"`
	Username                           string `json:"username"`
	AddressFields                      string `json:"address_fields"`
	Gender                             string `json:"gender"`
	PostalCode                         string `json:"postal_code"`
	Age                                uint32 `json:"age"`
	Birthday                           uint32 `json:"birthday"`
	Avatar                             string `json:"avatar"`
	Organization                       string `json:"organization"`
	FirstName                          string `json:"first_name"`
	LastName                           string `json:"last_name"`
	IDNumber                           string `json:"id_number"`
	SigninVerifyByGoogleAuthentication uint32 `json:"signin_verify_by_google_authentication"`
	GoogleAuthenticationVerified       uint32 `json:"google_authentication_verified"`
	BanAppUserID                       string `json:"ban_app_user_id"`
	BanAppUserMessage                  string `json:"ban_app_user_message"`
	HasGoogleSecret                    string `json:"has_google_secret"`
	Roles                              string `json:"roles"`
}
