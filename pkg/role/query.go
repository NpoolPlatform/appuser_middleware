package role

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	entapp "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	entapprole "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approle"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
)

func GetRoles(ctx context.Context, appID string, offset, limit int32) ([]*role.Role, int, error) {
	var err error
	infos := []*role.Role{}
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetRoles")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.String("AppID", appID))
	span = commontracer.TraceOffsetLimit(span, int(offset), int(limit))
	span = commontracer.TraceInvoker(span, "app", "db", "query join")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppRole.
			Query().
			Where(
				entapprole.AppID(uuid.MustParse(appID)),
			)

		total, err = stm.Count(ctx)
		if err != nil {
			return err
		}

		stm.
			Offset(int(offset)).
			Limit(int(limit))
		return joinRole(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetRoles", "error", err)
		return nil, 0, err
	}

	return infos, total, nil
}

func GetManyRoles(ctx context.Context, ids []string) ([]*role.Role, error) {
	var err error
	infos := []*role.Role{}

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetManyRoles")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.StringSlice("ids", ids))
	span = commontracer.TraceInvoker(span, "app", "db", "query join")

	idsU := []uuid.UUID{}
	for _, val := range ids {
		idsU = append(idsU, uuid.MustParse(val))
	}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppRole.
			Query().
			Where(
				entapprole.IDIn(idsU...),
			)
		return joinRole(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetManyRoles", "error", err)
		return nil, err
	}

	return infos, nil
}

func joinRole(stm *ent.AppRoleQuery) *ent.AppRoleSelect {
	return stm.Select(
		entapprole.FieldID,
		entapprole.FieldCreatedBy,
		entapprole.FieldRole,
		entapprole.FieldDescription,
		entapprole.FieldDefault,
	).Modify(func(s *sql.Selector) {
		t1 := sql.Table(entapp.Table)
		s.
			LeftJoin(t1).
			On(
				s.C(entapprole.FieldAppID),
				t1.C(entapp.FieldID),
			).
			AppendSelect(
				sql.As(t1.C(entapp.FieldID), "app_id"),
				sql.As(t1.C(entapp.FieldName), "app_name"),
				sql.As(t1.C(entapp.FieldLogo), "app_logo"),
				t1.C(entapp.FieldCreatedAt),
			)
	})
}