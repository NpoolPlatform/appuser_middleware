package role

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	entapprole "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approle"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	entapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/app"
	servicename "github.com/NpoolPlatform/appuser-middleware/pkg/servicename"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
)

func GetRole(ctx context.Context, id string) (*role.Role, error) {
	var err error
	var infos []*role.Role

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetRoles")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id)
	span = commontracer.TraceInvoker(span, "role", "db", "CRUD")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppRole.
			Query().
			Where(
				entapprole.ID(uuid.MustParse(id)),
			)
		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetRoles", "error", err)
		return nil, err
	}

	return infos[0], nil
}

func GetRoles(ctx context.Context, appID string, offset, limit int32) ([]*role.Role, uint32, error) {
	var err error
	infos := []*role.Role{}
	var total int

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetRoles")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.String("AppID", appID))
	span = commontracer.TraceOffsetLimit(span, int(offset), int(limit))
	span = commontracer.TraceInvoker(span, "role", "db", "CRUD")

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
		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetRoles", "error", err)
		return nil, 0, err
	}

	return infos, uint32(total), nil
}

func GetManyRoles(ctx context.Context, ids []string) ([]*role.Role, uint32, error) {
	var err error
	infos := []*role.Role{}

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetManyRoles")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.StringSlice("ids", ids))
	span = commontracer.TraceInvoker(span, "app", "db", "CRUD")

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
		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetManyRoles", "error", err)
		return nil, 0, err
	}

	return infos, uint32(len(infos)), nil
}

func join(stm *ent.AppRoleQuery) *ent.AppRoleSelect {
	return stm.Select(
		entapprole.FieldID,
		entapprole.FieldCreatedBy,
		entapprole.FieldRole,
		entapprole.FieldDescription,
		entapprole.FieldDefault,
		entapprole.FieldGenesis,
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
