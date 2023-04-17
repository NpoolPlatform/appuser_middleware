package app

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"entgo.io/ent/dialect/sql"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	"go.opentelemetry.io/otel/attribute"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	entapp "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	ctrl "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appcontrol"
	banapp "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banapp"
	servicename "github.com/NpoolPlatform/appuser-middleware/pkg/servicename"
	scodes "go.opentelemetry.io/otel/codes"

	"go.opentelemetry.io/otel"

	"github.com/google/uuid"

	ctrlpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
)

func (h *Handler) GetApp(ctx context.Context) (*npool.App, error) {
	return GetApp(ctx, h.ID.String())
}

func GetApp(ctx context.Context, id string) (*npool.App, error) {
	var err error
	infos := []*npool.App{}

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetApp")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id)

	span = commontracer.TraceInvoker(span, "app", "db", "CRUD")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			App.
			Query().
			Where(
				entapp.ID(uuid.MustParse(id)),
			).
			Limit(1)
		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetApp", "error", err)
		return nil, err
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("no record")
	}
	if len(infos) > 1 {
		logger.Sugar().Errorw("GetApp", "too many records")
		return nil, fmt.Errorf("too many records")
	}

	infos = expand(infos)

	return infos[0], nil
}

func GetApps(ctx context.Context, offset, limit int32) ([]*npool.App, error) {
	var err error
	infos := []*npool.App{}

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetApps")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceOffsetLimit(span, int(offset), int(limit))

	span = commontracer.TraceInvoker(span, "app", "db", "CRUD")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			App.
			Query().
			Offset(int(offset)).
			Limit(int(limit))

		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetApps", "error", err)
		return nil, err
	}

	infos = expand(infos)

	return infos, nil
}

func GetUserApps(ctx context.Context, userID string, offset, limit int32) ([]*npool.App, int, error) {
	var err error
	infos := []*npool.App{}
	var total int

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetUserApps")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.String("UserID", userID))

	span = commontracer.TraceOffsetLimit(span, int(offset), int(limit))

	span = commontracer.TraceInvoker(span, "app", "db", "CRUD")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			App.
			Query().
			Where(
				entapp.CreatedBy(uuid.MustParse(userID)),
			)

		total, err = stm.Count(ctx)
		if err != nil {
			logger.Sugar().Errorw("GetUserApps", "error", err)
			return err
		}

		stm.
			Offset(int(offset)).
			Limit(int(limit))

		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetUserApps", "error", err)
		return nil, 0, err
	}

	infos = expand(infos)

	return infos, total, nil
}

func GetManyApps(ctx context.Context, ids []string) ([]*npool.App, int, error) {
	var err error
	infos := []*npool.App{}
	var total int

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetManyApps")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.StringSlice("ids", ids))

	idsU := []uuid.UUID{}
	for _, val := range ids {
		idsU = append(idsU, uuid.MustParse(val))
	}

	span = commontracer.TraceInvoker(span, "app", "db", "CRUD")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			App.
			Query().
			Where(
				entapp.IDIn(idsU...),
			)
		total, err = stm.Count(ctx)
		if err != nil {
			logger.Sugar().Errorw("GetManyApps", "error", err)
			return err
		}
		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetManyApps", "error", err)
		return nil, 0, err
	}

	infos = expand(infos)

	return infos, total, nil
}

func join(stm *ent.AppQuery) *ent.AppSelect {
	return stm.Select(
		entapp.FieldID,
		entapp.FieldLogo,
		entapp.FieldName,
		entapp.FieldCreatedBy,
		entapp.FieldCreatedAt,
		entapp.FieldDescription,
	).Modify(func(s *sql.Selector) {
		t1 := sql.Table(banapp.Table)
		s.
			LeftJoin(t1).
			On(
				s.C(entapp.FieldID),
				t1.C(banapp.FieldAppID),
			).
			AppendSelect(
				sql.As(t1.C(banapp.FieldID), "ban_app_id"),
				sql.As(t1.C(banapp.FieldMessage), "ban_message"),
			)

		t2 := sql.Table(ctrl.Table)
		s.
			LeftJoin(t2).
			On(
				s.C(entapp.FieldID),
				t2.C(ctrl.FieldAppID),
			).
			AppendSelect(
				t2.C(ctrl.FieldSignupMethods),
				t2.C(ctrl.FieldExternSigninMethods),
				t2.C(ctrl.FieldRecaptchaMethod),
				t2.C(ctrl.FieldKycEnable),
				t2.C(ctrl.FieldSigninVerifyEnable),
				t2.C(ctrl.FieldInvitationCodeMust),
				t2.C(ctrl.FieldCreateInvitationCodeWhen),
				t2.C(ctrl.FieldMaxTypedCouponsPerOrder),
				t2.C(ctrl.FieldMaintaining),
				t2.C(ctrl.FieldCommitButtonTargets),
			)
	})
}

func expand(infos []*npool.App) []*npool.App {
	for key, info := range infos {
		info.CreateInvitationCodeWhen =
			ctrlpb.CreateInvitationCodeWhen(
				ctrlpb.CreateInvitationCodeWhen_value[info.CreateInvitationCodeWhenStr],
			)
		_ = json.Unmarshal([]byte(info.CommitButtonTargetsStr), &infos[key].CommitButtonTargets)
	}
	return infos
}
