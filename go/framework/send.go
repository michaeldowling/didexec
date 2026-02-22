package framework

import (
	"context"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/crypto"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

func (fw *Framework) Invite(guestUser id.UserID) (id.RoomID, error) {

	resp, err := fw.Client.CreateRoom(context.Background(), &mautrix.ReqCreateRoom{

		Name:       uuid.NewString(),
		Visibility: "private",
		IsDirect:   true,
		Preset:     "trusted_private_chat",
		Invite:     []id.UserID{guestUser},
		Topic:      "Collaboration Room",
	})
	if err != nil {
		log.Fatalf("Failed to create room: %v", err)
		return "", err
	}

	//_, err = fw.Client.InviteUser(context.Background(), resp.RoomID, &mautrix.ReqInviteUser{UserID: guestUser})
	//if err != nil {
	//	log.Fatalf("Failed to invite user: %v", err)
	//	return "", err
	//}

	return resp.RoomID, nil

}

func (fw *Framework) Send(roomId id.RoomID, userId id.UserID, msg string) error {

	messageContent := event.MessageEventContent{
		MsgType: event.MsgText,
		Body:    msg,
		To:      userId,
	}

	_, err := fw.Client.SendMessageEvent(context.Background(), roomId, event.EventMessage, messageContent)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
		return err
	}

	return nil

}

func (fw *Framework) VerifyWithRecoveryKey(machine *crypto.OlmMachine, recoveryKey string) (err error) {
	ctx := context.Background()

	keyId, keyData, err := machine.SSSS.GetDefaultKeyData(ctx)
	if err != nil {
		return
	}
	key, err := keyData.VerifyRecoveryKey(keyId, recoveryKey)
	if err != nil {
		return
	}
	err = machine.FetchCrossSigningKeysFromSSSS(ctx, key)
	if err != nil {
		return
	}
	err = machine.SignOwnDevice(ctx, machine.OwnIdentity())
	if err != nil {
		return
	}
	err = machine.SignOwnMasterKey(ctx)

	return
}
