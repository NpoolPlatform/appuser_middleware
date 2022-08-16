package roleuser

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	entapp "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	entapprole "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approle"
	entapproleuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approleuser"
	entappuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
)

func GetRoleUsers(ctx context.Context, appID, roleID string, offset, limit int32) ([]*role.RoleUser, int, error) {
	var err error
	infos := []*role.RoleUser{}
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetManyRoleUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.String("AppID", appID))
	span.SetAttributes(attribute.String("RoleID", roleID))
	commontracer.TraceOffsetLimit(span, int(offset), int(limit))

	span = commontracer.TraceInvoker(span, "app", "db", "query join")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppRoleUser.
			Query().
			Where(
				entapproleuser.AppID(uuid.MustParse(appID)),
				entapproleuser.RoleID(uuid.MustParse(roleID)),
			)
		total, err = stm.Count(ctx)
		if err != nil {
			return err
		}
		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetManyRoleUsers", "error", err)
		return nil, 0, err
	}

	return infos, total, nil
}

func GetManyRoleUsers(ctx context.Context, ids []string) ([]*role.RoleUser, error) {
	var err error
	infos := []*role.RoleUser{}
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetManyRoleUsers")
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
			AppRoleUser.
			Query().
			Where(
				entapproleuser.IDIn(idsU...),
			)
		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetManyRoleUsers", "error", err)
		return nil, err
	}

	return infos, nil
}

func join(stm *ent.AppRoleUserQuery) *ent.AppRoleUserSelect {
	return stm.Select(
		entapproleuser.FieldID,
	).Modify(func(s *sql.Selector) {
		t1 := sql.Table(entapprole.Table)
		s.
			LeftJoin(t1).
			On(
				s.C(entapproleuser.FieldRoleID),
				t1.C(entapprole.FieldID),
			).
			AppendSelect(
				t1.C(entapprole.FieldCreatedBy),
				t1.C(entapprole.FieldRole),
				t1.C(entapprole.FieldDescription),
				t1.C(entapprole.FieldDefault),
			)

		t2 := sql.Table(entapp.Table)
		s.
			LeftJoin(t2).
			On(
				s.C(entapproleuser.FieldAppID),
				t2.C(entapp.FieldID),
			).
			AppendSelect(
				sql.As(t2.C(entapp.FieldID), "app_id"),
				sql.As(t2.C(entapp.FieldName), "app_name"),
				sql.As(t2.C(entapp.FieldLogo), "app_logo"),
			)
		t3 := sql.Table(entappuser.Table)
		s.
			LeftJoin(t3).
			On(
				s.C(entapproleuser.FieldUserID),
				t3.C(entappuser.FieldID),
			).
			AppendSelect(
				sql.As(t3.C(entappuser.FieldID), "user_id"),
				sql.As(t3.C(entappuser.FieldEmailAddress), "app_name"),
				t3.C(entappuser.FieldPhoneNo),
			)
	})
}
