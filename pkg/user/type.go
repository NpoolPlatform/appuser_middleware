package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID                                 uuid.UUID `sql:"id"`
	AppID                              uuid.UUID `sql:"app_id"`
	EmailAddress                       string    `sql:"email_address"`
	PhoneNO                            string    `sql:"phone_no"`
	ImportedFromAppID                  uuid.UUID `sql:"import_from_app"`
	ImportedFromAppName                string    `sql:"imported_from_app_name"`
	ImportedFromAppLogo                string    `sql:"imported_from_app_logo"`
	ImportedFromAppHome                string    `sql:"imported_from_app_home"`
	CreatedAt                          uint32    `sql:"created_at"`
	Username                           string    `sql:"username"`
	AddressFields                      string    `sql:"address_fields"`
	Gender                             string    `sql:"gender"`
	PostalCode                         string    `sql:"postal_code"`
	Age                                uint32    `sql:"age"`
	Birthday                           uint32    `sql:"birthday"`
	Avatar                             string    `sql:"avatar"`
	Organization                       string    `sql:"organization"`
	FirstName                          string    `sql:"first_name"`
	LastName                           string    `sql:"last_name"`
	IDNumber                           string    `sql:"id_number"`
	SigninVerifyByGoogleAuthentication int       `sql:"signin_verify_by_google_authentication"`
	GoogleAuthenticationVerified       int       `sql:"google_authentication_verified"`
	Banned                             bool      `sql:"banned"`
	BanAppUserID                       uuid.UUID `sql:"ban_app_user_id"`
	BanMessage                         string    `sql:"ban_message"`
	HasGoogleSecret                    bool      `sql:"has_google_secret"`
	Roles                              []string  `sql:"roles"`
}
