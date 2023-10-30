// Code generated by ent, DO NOT EDIT.

package appuserthirdparty

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// ThirdPartyUserID applies equality check predicate on the "third_party_user_id" field. It's identical to ThirdPartyUserIDEQ.
func ThirdPartyUserID(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyID applies equality check predicate on the "third_party_id" field. It's identical to ThirdPartyIDEQ.
func ThirdPartyID(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyUsername applies equality check predicate on the "third_party_username" field. It's identical to ThirdPartyUsernameEQ.
func ThirdPartyUsername(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyAvatar applies equality check predicate on the "third_party_avatar" field. It's identical to ThirdPartyAvatarEQ.
func ThirdPartyAvatar(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyAvatar), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// ThirdPartyUserIDEQ applies the EQ predicate on the "third_party_user_id" field.
func ThirdPartyUserIDEQ(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDNEQ applies the NEQ predicate on the "third_party_user_id" field.
func ThirdPartyUserIDNEQ(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDIn applies the In predicate on the "third_party_user_id" field.
func ThirdPartyUserIDIn(vs ...string) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldThirdPartyUserID), v...))
	})
}

// ThirdPartyUserIDNotIn applies the NotIn predicate on the "third_party_user_id" field.
func ThirdPartyUserIDNotIn(vs ...string) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldThirdPartyUserID), v...))
	})
}

// ThirdPartyUserIDGT applies the GT predicate on the "third_party_user_id" field.
func ThirdPartyUserIDGT(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDGTE applies the GTE predicate on the "third_party_user_id" field.
func ThirdPartyUserIDGTE(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDLT applies the LT predicate on the "third_party_user_id" field.
func ThirdPartyUserIDLT(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDLTE applies the LTE predicate on the "third_party_user_id" field.
func ThirdPartyUserIDLTE(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDContains applies the Contains predicate on the "third_party_user_id" field.
func ThirdPartyUserIDContains(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDHasPrefix applies the HasPrefix predicate on the "third_party_user_id" field.
func ThirdPartyUserIDHasPrefix(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDHasSuffix applies the HasSuffix predicate on the "third_party_user_id" field.
func ThirdPartyUserIDHasSuffix(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDIsNil applies the IsNil predicate on the "third_party_user_id" field.
func ThirdPartyUserIDIsNil() predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldThirdPartyUserID)))
	})
}

// ThirdPartyUserIDNotNil applies the NotNil predicate on the "third_party_user_id" field.
func ThirdPartyUserIDNotNil() predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldThirdPartyUserID)))
	})
}

// ThirdPartyUserIDEqualFold applies the EqualFold predicate on the "third_party_user_id" field.
func ThirdPartyUserIDEqualFold(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDContainsFold applies the ContainsFold predicate on the "third_party_user_id" field.
func ThirdPartyUserIDContainsFold(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyIDEQ applies the EQ predicate on the "third_party_id" field.
func ThirdPartyIDEQ(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDNEQ applies the NEQ predicate on the "third_party_id" field.
func ThirdPartyIDNEQ(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDIn applies the In predicate on the "third_party_id" field.
func ThirdPartyIDIn(vs ...uuid.UUID) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldThirdPartyID), v...))
	})
}

// ThirdPartyIDNotIn applies the NotIn predicate on the "third_party_id" field.
func ThirdPartyIDNotIn(vs ...uuid.UUID) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldThirdPartyID), v...))
	})
}

// ThirdPartyIDGT applies the GT predicate on the "third_party_id" field.
func ThirdPartyIDGT(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDGTE applies the GTE predicate on the "third_party_id" field.
func ThirdPartyIDGTE(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDLT applies the LT predicate on the "third_party_id" field.
func ThirdPartyIDLT(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDLTE applies the LTE predicate on the "third_party_id" field.
func ThirdPartyIDLTE(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDIsNil applies the IsNil predicate on the "third_party_id" field.
func ThirdPartyIDIsNil() predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldThirdPartyID)))
	})
}

// ThirdPartyIDNotNil applies the NotNil predicate on the "third_party_id" field.
func ThirdPartyIDNotNil() predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldThirdPartyID)))
	})
}

// ThirdPartyUsernameEQ applies the EQ predicate on the "third_party_username" field.
func ThirdPartyUsernameEQ(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameNEQ applies the NEQ predicate on the "third_party_username" field.
func ThirdPartyUsernameNEQ(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameIn applies the In predicate on the "third_party_username" field.
func ThirdPartyUsernameIn(vs ...string) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldThirdPartyUsername), v...))
	})
}

// ThirdPartyUsernameNotIn applies the NotIn predicate on the "third_party_username" field.
func ThirdPartyUsernameNotIn(vs ...string) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldThirdPartyUsername), v...))
	})
}

// ThirdPartyUsernameGT applies the GT predicate on the "third_party_username" field.
func ThirdPartyUsernameGT(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameGTE applies the GTE predicate on the "third_party_username" field.
func ThirdPartyUsernameGTE(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameLT applies the LT predicate on the "third_party_username" field.
func ThirdPartyUsernameLT(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameLTE applies the LTE predicate on the "third_party_username" field.
func ThirdPartyUsernameLTE(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameContains applies the Contains predicate on the "third_party_username" field.
func ThirdPartyUsernameContains(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameHasPrefix applies the HasPrefix predicate on the "third_party_username" field.
func ThirdPartyUsernameHasPrefix(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameHasSuffix applies the HasSuffix predicate on the "third_party_username" field.
func ThirdPartyUsernameHasSuffix(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameIsNil applies the IsNil predicate on the "third_party_username" field.
func ThirdPartyUsernameIsNil() predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldThirdPartyUsername)))
	})
}

// ThirdPartyUsernameNotNil applies the NotNil predicate on the "third_party_username" field.
func ThirdPartyUsernameNotNil() predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldThirdPartyUsername)))
	})
}

// ThirdPartyUsernameEqualFold applies the EqualFold predicate on the "third_party_username" field.
func ThirdPartyUsernameEqualFold(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameContainsFold applies the ContainsFold predicate on the "third_party_username" field.
func ThirdPartyUsernameContainsFold(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyAvatarEQ applies the EQ predicate on the "third_party_avatar" field.
func ThirdPartyAvatarEQ(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyAvatar), v))
	})
}

// ThirdPartyAvatarNEQ applies the NEQ predicate on the "third_party_avatar" field.
func ThirdPartyAvatarNEQ(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThirdPartyAvatar), v))
	})
}

// ThirdPartyAvatarIn applies the In predicate on the "third_party_avatar" field.
func ThirdPartyAvatarIn(vs ...string) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldThirdPartyAvatar), v...))
	})
}

// ThirdPartyAvatarNotIn applies the NotIn predicate on the "third_party_avatar" field.
func ThirdPartyAvatarNotIn(vs ...string) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldThirdPartyAvatar), v...))
	})
}

// ThirdPartyAvatarGT applies the GT predicate on the "third_party_avatar" field.
func ThirdPartyAvatarGT(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThirdPartyAvatar), v))
	})
}

// ThirdPartyAvatarGTE applies the GTE predicate on the "third_party_avatar" field.
func ThirdPartyAvatarGTE(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThirdPartyAvatar), v))
	})
}

// ThirdPartyAvatarLT applies the LT predicate on the "third_party_avatar" field.
func ThirdPartyAvatarLT(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThirdPartyAvatar), v))
	})
}

// ThirdPartyAvatarLTE applies the LTE predicate on the "third_party_avatar" field.
func ThirdPartyAvatarLTE(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThirdPartyAvatar), v))
	})
}

// ThirdPartyAvatarContains applies the Contains predicate on the "third_party_avatar" field.
func ThirdPartyAvatarContains(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldThirdPartyAvatar), v))
	})
}

// ThirdPartyAvatarHasPrefix applies the HasPrefix predicate on the "third_party_avatar" field.
func ThirdPartyAvatarHasPrefix(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldThirdPartyAvatar), v))
	})
}

// ThirdPartyAvatarHasSuffix applies the HasSuffix predicate on the "third_party_avatar" field.
func ThirdPartyAvatarHasSuffix(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldThirdPartyAvatar), v))
	})
}

// ThirdPartyAvatarIsNil applies the IsNil predicate on the "third_party_avatar" field.
func ThirdPartyAvatarIsNil() predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldThirdPartyAvatar)))
	})
}

// ThirdPartyAvatarNotNil applies the NotNil predicate on the "third_party_avatar" field.
func ThirdPartyAvatarNotNil() predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldThirdPartyAvatar)))
	})
}

// ThirdPartyAvatarEqualFold applies the EqualFold predicate on the "third_party_avatar" field.
func ThirdPartyAvatarEqualFold(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldThirdPartyAvatar), v))
	})
}

// ThirdPartyAvatarContainsFold applies the ContainsFold predicate on the "third_party_avatar" field.
func ThirdPartyAvatarContainsFold(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldThirdPartyAvatar), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppUserThirdParty) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppUserThirdParty) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppUserThirdParty) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		p(s.Not())
	})
}
