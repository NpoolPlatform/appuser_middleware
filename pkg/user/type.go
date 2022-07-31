package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID                                 uuid.UUID `json:"id"`
	AppID                              uuid.UUID `json:"app_id"`
	EmailAddress                       string    `json:"email_address"`
	PhoneNO                            string    `json:"phone_no"`
	ImportedFromAppID                  uuid.UUID `json:"imported_from_app_id"`
	ImportedFromAppName                string    `sql:"imported_from_app_name"`
	ImportedFromAppLogo                string    `sql:"imported_from_app_logo"`
	ImportedFromAppHome                string    `sql:"imported_from_app_home"`
	CreatedAt                          uint32    `json:"created_at"`
	Username                           string    `json:"username"`
	AddressFields                      string    `json:"address_fields"`
	Gender                             string    `json:"gender"`
	PostalCode                         string    `json:"postal_code"`
	Age                                uint32    `json:"age"`
	Birthday                           uint32    `json:"birthday"`
	Avatar                             string    `json:"avatar"`
	Organization                       string    `json:"organization"`
	FirstName                          string    `json:"first_name"`
	LastName                           string    `json:"last_name"`
	IDNumber                           string    `json:"id_number"`
	SigninVerifyByGoogleAuthentication int       `json:"signin_verify_by_google_authentication"`
	GoogleAuthenticationVerified       int       `json:"google_authentication_verified"`
	Banned                             bool      `json:"banned"`
	BanAppUserID                       uuid.UUID `json:"ban_app_user_id"`
	BanAppUserMessage                  string    `json:"ban_app_user_message"`
	HasGoogleSecret                    bool      `json:"has_google_secret"`
	Roles                              []string  `json:"roles"`
}
