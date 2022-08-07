package user

import (
	"context"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approleuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusercontrol"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserthirdparty"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
)

func DeleteUser(ctx context.Context, userID uuid.UUID) error {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	commontracer.TraceID(span, userID.String())
	span = commontracer.TraceInvoker(span, "user", "db", "DeleteTx")

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		err = tx.AppUser.DeleteOneID(userID).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.AppUserExtra.Delete().Where(appuserextra.UserID(userID)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.AppUserControl.Delete().Where(appusercontrol.UserID(userID)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.AppUserSecret.Delete().Where(appusersecret.UserID(userID)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.AppUserThirdParty.Delete().Where(appuserthirdparty.UserID(userID)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.AppRoleUser.Delete().Where(approleuser.UserID(userID)).Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return err
}
