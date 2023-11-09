package user

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	usercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entappuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuser"
	entappuserthirdparty "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuserthirdparty"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"github.com/google/uuid"
)

type existHandler struct {
	*Handler
	stm            *ent.AppUserSelect
	infos          []*npool.User
	total          uint32
	joinThirdParty bool
}

func (h *existHandler) selectAppUser(stm *ent.AppUserQuery) {
	h.stm = stm.Select(entappuser.FieldID)
}

func (h *existHandler) queryAppUser(cli *ent.Client) {
	stm := cli.AppUser.
		Query().
		Where(entappuser.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappuser.ID(*h.ID))
	}
	if h.AppID != nil {
		stm.Where(entappuser.AppID(*h.AppID))
	}
	if h.EntID != nil {
		stm.Where(entappuser.EntID(*h.EntID))
	}
	h.selectAppUser(stm)
}

func (h *existHandler) queryAppUserByConds(cli *ent.Client) error {
	stm, err := usercrud.SetQueryConds(cli.AppUser.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.selectAppUser(stm)
	return nil
}

func (h *existHandler) queryJoinAppUserThirdParty(s *sql.Selector) error {
	if !h.joinThirdParty {
		return nil
	}

	t := sql.Table(entappuserthirdparty.Table)
	s.LeftJoin(t).
		On(
			s.C(entappuser.FieldEntID),
			t.C(entappuserthirdparty.FieldUserID),
		).
		On(
			s.C(entappuser.FieldAppID),
			t.C(entappuserthirdparty.FieldAppID),
		).
		On(
			s.C(entappuser.FieldDeletedAt),
			t.C(entappuserthirdparty.FieldDeletedAt),
		).
		AppendSelect(
			sql.As(t.C(entappuserthirdparty.FieldThirdPartyID), "third_party_id"),
			sql.As(t.C(entappuserthirdparty.FieldThirdPartyUserID), "third_party_user_id"),
			sql.As(t.C(entappuserthirdparty.FieldThirdPartyUsername), "third_party_username"),
			sql.As(t.C(entappuserthirdparty.FieldThirdPartyAvatar), "third_party_avatar"),
		)

	if h.Conds != nil && h.Conds.ThirdPartyID != nil {
		thirdPartyID, ok := h.Conds.ThirdPartyID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid oauth thirdpartyid")
		}
		s.Where(
			sql.EQ(t.C(entappuserthirdparty.FieldThirdPartyID), thirdPartyID),
		)
	}

	if h.Conds != nil && h.Conds.ThirdPartyUserID != nil {
		thirdPartyUserID, ok := h.Conds.ThirdPartyUserID.Val.(string)
		if !ok {
			return fmt.Errorf("invalid oauth thirdpartyuserid")
		}
		s.Where(
			sql.EQ(t.C(entappuserthirdparty.FieldThirdPartyUserID), thirdPartyUserID),
		)
	}

	return nil
}

func (h *existHandler) queryJoin() error {
	var err error
	h.stm.Modify(func(s *sql.Selector) {
		err = h.queryJoinAppUserThirdParty(s)
	})
	return err
}

func (h *Handler) ExistUser(ctx context.Context) (exist bool, err error) {
	handler := &existHandler{
		Handler:        h,
		joinThirdParty: false,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryAppUser(cli)
		if err := handler.queryJoin(); err != nil {
			return err
		}
		exist, err = handler.stm.Exist(_ctx)
		return err
	})
	return exist, err
}

func (h *Handler) ExistUserConds(ctx context.Context) (exist bool, err error) {
	handler := &existHandler{
		Handler:        h,
		joinThirdParty: false,
	}
	if h.Conds != nil && (h.Conds.ThirdPartyID != nil || h.Conds.ThirdPartyUserID != nil) {
		handler.joinThirdParty = true
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppUserByConds(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		exist, err = handler.stm.Exist(_ctx)
		return err
	})
	return exist, err
}
