package app

import (
	"github.com/google/uuid"
)

type App struct {
	ID                 uuid.UUID `sql:"id"`
	CreatedBy          uuid.UUID `sql:"created_by"`
	Name               string    `sql:"name"`
	Logo               string    `sql:"logo"`
	Description        string    `sql:"description"`
	Banned             bool      `sql:"banned"`
	BanAppID           uuid.UUID `sql:"ban_app_id"`
	BanMessage         string    `sql:"ban_message"`
	SignupMethods      string    `sql:"signup_methods"`
	ExtSigninMethods   string    `sql:"extern_signin_methods"`
	RecaptchaMethod    string    `sql:"recaptcha_method"`
	KycEnable          int       `sql:"kyc_enable"`
	SigninVerifyEnable int       `sql:"signin_verify_enable"`
	InvitationCodeMust int       `sql:"invitation_code_must"`
	CreatedAt          uint32    `sql:"created_at"`
}
