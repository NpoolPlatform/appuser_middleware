package roleuser

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	entapprole "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approle"
	entapproleuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approleuser"
	entappuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
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

func GetRoleUser(ctx context.Context, id string) (*role.RoleUser, error) {
	var err error
	var infos []*role.RoleUser

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetRoleUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id)
	span = commontracer.TraceInvoker(span, "roleuser", "db", "CRUD")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppRoleUser.
			Query().
			Where(
				entapproleuser.ID(uuid.MustParse(id)),
			)
		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetManyRoleUsers", "error", err)
		return nil, err
	}

	return infos[0], nil
}

func GetRoleUsers(ctx context.Context, appID, roleID string, offset, limit int32) ([]*role.RoleUser, uint32, error) {
	var err error
	infos := []*role.RoleUser{}
	var total int

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetManyRoleUsers")
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

	span = commontracer.TraceInvoker(span, "roleuser", "db", "CRUD")

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
		stm.
			Offset(int(offset)).
			Limit(int(limit))
		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetManyRoleUsers", "error", err)
		return nil, 0, err
	}

	return infos, uint32(total), nil
}

func GetManyRoleUsers(ctx context.Context, ids []string) ([]*role.RoleUser, uint32, error) {
	var err error
	infos := []*role.RoleUser{}
	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetManyRoleUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.StringSlice("ids", ids))
	span = commontracer.TraceInvoker(span, "roleuser", "db", "CRUD")

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
		return nil, 0, err
	}

	return infos, uint32(len(infos)), nil
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
				t1.C(entapprole.FieldGenesis),
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
				sql.As(t3.C(entappuser.FieldEmailAddress), "email_address"),
				t3.C(entappuser.FieldPhoneNo),
			)
	})
}
