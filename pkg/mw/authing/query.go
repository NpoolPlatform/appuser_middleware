package authing

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/app"
	entapprole "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approle"
	entappuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuser"
	entauth "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/auth"
	entauthhistory "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/authhistory"
	servicename "github.com/NpoolPlatform/appuser-middleware/pkg/servicename"
	commontracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
)

type queryHandler struct {
	*Handler
}

func (h *queryHandler) selectAuth(stm *ent.AuthQuery) {

}

func (h *queryHandler) queryAuth(cli *ent.Client) error {
	return nil
}

func (h *queryHandler) queryJoin() {

}

func (h *Handler) GetAuth(ctx context.Context) (*npool.Auth, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAuth(cli); err != nil {
			return nil
		}
		handler.queryJoin()
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})

	return nil, nil
}

func GetAuth(ctx context.Context, id string) (info *npool.Auth, err error) {
	infos := []*npool.Auth{}

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetAuth")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "auth", "db", "QueryJoin")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Auth.
			Query().
			Where(
				entauth.ID(uuid.MustParse(id)),
			)
		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetAuths", "error", err)
		return nil, err
	}

	return infos[0], nil
}

func GetAuths(ctx context.Context, appID string, offset, limit int32) (infos []*npool.Auth, total int, err error) {
	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetAuths")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "auth", "db", "QueryJoin")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Auth.
			Query().
			Where(
				entauth.AppID(uuid.MustParse(appID)),
			)

		total, err = stm.Count(ctx)
		if err != nil {
			return err
		}

		stm.
			Offset(int(offset))
		if limit > 0 {
			stm.Limit(int(limit))
		}
		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetAuths", "error", err)
		return nil, 0, err
	}

	return infos, total, nil
}

func join(stm *ent.AuthQuery) *ent.AuthSelect {
	return stm.Select(
		entauth.FieldID,
		entauth.FieldResource,
		entauth.FieldMethod,
		entauth.FieldCreatedAt,
		entauth.FieldAppID,
		entauth.FieldRoleID,
		entauth.FieldUserID,
	).Modify(func(s *sql.Selector) {
		t1 := sql.Table(entapp.Table)
		s.
			LeftJoin(t1).
			On(
				s.C(entauth.FieldAppID),
				t1.C(entapp.FieldID),
			).
			AppendSelect(
				sql.As(t1.C(entapp.FieldName), "app_name"),
				sql.As(t1.C(entapp.FieldLogo), "app_logo"),
			)

		t2 := sql.Table(entapprole.Table)
		s.
			LeftJoin(t2).
			On(
				s.C(entauth.FieldRoleID),
				t2.C(entapprole.FieldID),
			).
			AppendSelect(
				sql.As(t2.C(entapprole.FieldRole), "role_name"),
			)

		t3 := sql.Table(entappuser.Table)
		s.
			LeftJoin(t3).
			On(
				s.C(entauth.FieldUserID),
				t3.C(entappuser.FieldID),
			).
			AppendSelect(
				sql.As(t3.C(entappuser.FieldEmailAddress), "email_address"),
				sql.As(t3.C(entappuser.FieldPhoneNo), "phone_no"),
			)
	})
}

func GetHistories(ctx context.Context, appID string, offset, limit int32) (infos []*npool.History, total int, err error) {
	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetHistories")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "auth", "db", "GetHistories")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AuthHistory.
			Query().
			Where(
				entauthhistory.AppID(uuid.MustParse(appID)),
			)

		total, err = stm.Count(ctx)
		if err != nil {
			return err
		}

		stm.
			Offset(int(offset)).
			Limit(int(limit))

		return jsonH(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetHistories", "error", err)
		return nil, 0, err
	}

	return infos, total, nil
}

func jsonH(stm *ent.AuthHistoryQuery) *ent.AuthHistorySelect {
	return stm.Select(
		entauthhistory.FieldAppID,
		entauthhistory.FieldUserID,
		entauthhistory.FieldResource,
		entauthhistory.FieldMethod,
		entauthhistory.FieldAllowed,
		entauthhistory.FieldCreatedAt,
	).Modify(func(s *sql.Selector) {
		t1 := sql.Table(entapp.Table)
		s.
			LeftJoin(t1).
			On(
				s.C(entauth.FieldAppID),
				t1.C(entapp.FieldID),
			).
			AppendSelect(
				sql.As(t1.C(entapp.FieldName), "app_name"),
				sql.As(t1.C(entapp.FieldLogo), "app_logo"),
			)

		t2 := sql.Table(entappuser.Table)
		s.
			LeftJoin(t2).
			On(
				s.C(entauth.FieldUserID),
				t2.C(entappuser.FieldID),
			).
			AppendSelect(
				sql.As(t2.C(entappuser.FieldEmailAddress), "email_address"),
				sql.As(t2.C(entappuser.FieldPhoneNo), "phone_no"),
			)
	})
}
