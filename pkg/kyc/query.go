package kyc

import (
	"context"

	entappuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"

	"entgo.io/ent/dialect/sql"
	crudkyc "github.com/NpoolPlatform/appuser-manager/pkg/crud/kyc"
	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	entapp "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	entkyc "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/kyc"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
)

func GetKyc(ctx context.Context, id string) (*kyc.Kyc, error) {
	var err error
	var infos []*kyc.Kyc

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetKycs")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id)
	span = commontracer.TraceInvoker(span, "kyc", "db", "CRUD")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Kyc.
			Query().
			Where(
				entkyc.ID(uuid.MustParse(id)),
			)
		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetKyc", "error", err)
		return nil, err
	}

	return infos[0], nil
}

func GetKycs(ctx context.Context, conds *kyc.Conds, offset, limit int32) ([]*kyc.Kyc, uint32, error) {
	var err error
	infos := []*kyc.Kyc{}
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetKycs")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceOffsetLimit(span, int(offset), int(limit))
	span = commontracer.TraceInvoker(span, "kyc", "db", "CRUD")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm, err := crudkyc.SetQueryConds(conds.GetConds(), cli)
		if err != nil {
			return err
		}

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
		logger.Sugar().Errorw("GetKycs", "error", err)
		return nil, 0, err
	}

	return infos, uint32(total), nil
}

func join(stm *ent.KycQuery) *ent.KycSelect {
	return stm.Select(
		entkyc.FieldID,
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
		entkyc.FieldAppID,
		entkyc.FieldUserID,
	).Modify(func(s *sql.Selector) {
		t1 := sql.Table(entapp.Table)
		s.
			LeftJoin(t1).
			On(
				s.C(entkyc.FieldAppID),
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
				s.C(entkyc.FieldUserID),
				t2.C(entappuser.FieldID),
			).
			AppendSelect(
				sql.As(t2.C(entappuser.FieldEmailAddress), "email_address"),
				sql.As(t2.C(entappuser.FieldPhoneNo), "phone_no"),
			)
	})
}
