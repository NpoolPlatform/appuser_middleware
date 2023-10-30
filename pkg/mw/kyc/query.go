package kyc

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	kyccrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/kyc"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	entapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/app"
	entappuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuser"
	entkyc "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/kyc"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

type queryHandler struct {
	*Handler
	stm   *ent.KycSelect
	infos []*npool.Kyc
	total uint32
}

func (h *queryHandler) selectKyc(stm *ent.KycQuery) {
	h.stm = stm.Select(
		entkyc.FieldID,
		entkyc.FieldAppID,
		entkyc.FieldUserID,
		entkyc.FieldDocumentType,
		entkyc.FieldIDNumber,
		entkyc.FieldFrontImg,
		entkyc.FieldBackImg,
		entkyc.FieldSelfieImg,
		entkyc.FieldEntityType,
		entkyc.FieldReviewID,
		entkyc.FieldState,
		entkyc.FieldCreatedAt,
		entkyc.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryKyc(cli *ent.Client) error {
	h.selectKyc(
		cli.Kyc.
			Query().
			Where(
				entkyc.EntID(*h.EntID),
				entkyc.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryKycs(ctx context.Context, cli *ent.Client) error {
	stm, err := kyccrud.SetQueryConds(cli.Kyc.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectKyc(stm)
	return nil
}

func (h *queryHandler) queryJoinApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entkyc.FieldAppID),
			t.C(entapp.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entapp.FieldName), "app_name"),
			sql.As(t.C(entapp.FieldLogo), "app_logo"),
		)
}

func (h *queryHandler) queryJoinAppUser(s *sql.Selector) {
	t := sql.Table(entappuser.Table)
	s.LeftJoin(t).
		On(
			s.C(entkyc.FieldAppID),
			t.C(entappuser.FieldAppID),
		).
		On(
			s.C(entkyc.FieldUserID),
			t.C(entappuser.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entappuser.FieldEmailAddress), "email_address"),
			sql.As(t.C(entappuser.FieldPhoneNo), "phone_no"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinApp(s)
		h.queryJoinAppUser(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.DocumentType = basetypes.KycDocumentType(basetypes.KycDocumentType_value[info.DocumentTypeStr])
		info.EntityType = basetypes.KycEntityType(basetypes.KycEntityType_value[info.EntityTypeStr])
		info.State = basetypes.KycState(basetypes.KycState_value[info.StateStr])
	}
}

func (h *Handler) GetKyc(ctx context.Context) (*npool.Kyc, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryKyc(cli); err != nil {
			return err
		}
		handler.queryJoin()
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetKycs(ctx context.Context) ([]*npool.Kyc, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryKycs(ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		if err := handler.scan(ctx); err != nil {
			return nil
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
