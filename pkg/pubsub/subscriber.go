package pubsub

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entpubsubmsg "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/pubsubmessage"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	authhistory "github.com/NpoolPlatform/appuser-middleware/pkg/pubsub/authing/history"
	loginhistory "github.com/NpoolPlatform/appuser-middleware/pkg/pubsub/user/login/history"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	"github.com/google/uuid"
)

var subscriber *pubsub.Subscriber
var publisher *pubsub.Publisher

// TODO: here we should call from DB transaction context
func finish(ctx context.Context, msg *pubsub.Msg, err error) error {
	state := basetypes.MsgState_StateSuccess
	if err != nil {
		state = basetypes.MsgState_StateFail
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.
			PubsubMessage.
			Create().
			SetEntID(msg.UID).
			SetMessageID(msg.MID).
			SetArguments(msg.Body).
			SetState(state.String())
		if msg.RID != nil {
			c.SetRespToID(*msg.RID)
		}
		if msg.UnID != nil {
			c.SetUndoID(*msg.UnID)
		}
		_, err = c.Save(ctx)
		return err
	})
}

func prepare(mid, body string) (req interface{}, err error) {
	switch mid {
	case basetypes.MsgID_CreateLoginHistoryReq.String():
		req, err = loginhistory.Prepare(body)
	case basetypes.MsgID_CreateAuthHistoryReq.String():
		req, err = authhistory.Prepare(body)
	default:
		return nil, nil
	}

	if err != nil {
		logger.Sugar().Errorw(
			"handler",
			"MID", mid,
			"Body", body,
		)
		return nil, err
	}

	return req, nil
}

// Query a request message
//  Return
//   bool   appliable == true, caller should go ahead to apply this message
//   error  error message
func statReq(ctx context.Context, mid string, uid uuid.UUID) (bool, error) {
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_, err = cli.
			PubsubMessage.
			Query().
			Where(
				entpubsubmsg.EntID(uid),
			).
			Only(_ctx)
		return err
	})

	switch err {
	case nil:
	default:
		if ent.IsNotFound(err) {
			return true, nil
		}
		logger.Sugar().Warnw(
			"stat",
			"MID", mid,
			"UID", uid,
			"Error", err,
		)
		return false, err
	}

	return false, nil
}

// Query a message state in database
//  Return
//   bool    appliable == true, caller should go ahead to apply this message
//   error   error message
func statMsg(ctx context.Context, mid string, uid uuid.UUID, rid *uuid.UUID) (bool, error) { //nolint
	switch mid {
	case basetypes.MsgID_CreateLoginHistoryReq.String():
		fallthrough //nolint
	case basetypes.MsgID_CreateAuthHistoryReq.String():
		return statReq(ctx, mid, uid)
	default:
		return false, fmt.Errorf("invalid message")
	}
}

// Stat if message in right status, and is appliable
//  Return
//   bool    appliable == true, the message needs to be applied
//   error   error happens
func stat(ctx context.Context, mid string, uid uuid.UUID, rid *uuid.UUID) (bool, error) {
	return statMsg(ctx, mid, uid, rid)
}

// Process will consume the message and return consuming state
//  Return
//   error   reason of error, if nil, means the message should be acked
func process(ctx context.Context, mid string, req interface{}) (err error) {
	switch mid {
	case basetypes.MsgID_CreateLoginHistoryReq.String():
		err = loginhistory.Apply(ctx, req)
	case basetypes.MsgID_CreateAuthHistoryReq.String():
		err = authhistory.Apply(ctx, req)
	default:
		return nil
	}
	return err
}

// No matter what handler return, the message will be acked, unless handler halt
// If handler halt, the service will be restart, all message will be requeue
func handler(ctx context.Context, msg *pubsub.Msg) (err error) {
	var req interface{}
	var appliable bool

	defer func() {
		msg.Ack()
		if req != nil && appliable {
			_ = finish(ctx, msg, err)
		}
	}()

	req, err = prepare(msg.MID, msg.Body)
	if err != nil {
		return err
	}
	if req == nil {
		return nil
	}

	appliable, err = stat(ctx, msg.MID, msg.UID, msg.RID)
	if err != nil {
		return err
	}
	if !appliable {
		return nil
	}

	err = process(ctx, msg.MID, req)
	return err
}

func Subscribe(ctx context.Context) (err error) {
	subscriber, err = pubsub.NewSubscriber()
	if err != nil {
		return err
	}

	publisher, err = pubsub.NewPublisher()
	if err != nil {
		return err
	}

	return subscriber.Subscribe(ctx, handler)
}

func Shutdown(ctx context.Context) error {
	if subscriber != nil {
		subscriber.Close()
	}
	if publisher != nil {
		publisher.Close()
	}

	return nil
}
