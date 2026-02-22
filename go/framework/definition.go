package framework

import (
	"context"

	log "github.com/sirupsen/logrus"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/crypto/cryptohelper"
	"maunium.net/go/mautrix/id"
)

type Framework struct {
	Client *mautrix.Client
	Crypto *cryptohelper.CryptoHelper
}

func Connect(homeserver string, userId string, pass string, deviceId string) (*Framework, error) {

	var userIdObj = id.UserID(userId)
	client, err := mautrix.NewClient(homeserver, userIdObj, "")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}

	resp, err := client.Login(context.Background(), &mautrix.ReqLogin{
		Type: mautrix.AuthTypePassword,
		Identifier: mautrix.UserIdentifier{
			User: userId,
			Type: mautrix.IdentifierTypeUser,
		},
		Password:           pass,
		StoreCredentials:   true,
		StoreHomeserverURL: true,
		DeviceID:           id.DeviceID(deviceId),
	})
	if err != nil {
		log.Fatalf("Failed to login: %v", err)
		return nil, err
	}

	log.Printf("Device ID: %s", resp.DeviceID)
	log.Printf("Access Token: %s", resp.AccessToken)
	log.Printf("UserID: %s", resp.UserID)

	fw := Framework{Client: client}
	fw.Client.DeviceID = resp.DeviceID

	//cryptoHelper, err := SetupCrypto(client)
	//if err != nil {
	//	log.Fatalf("Failed to setup crypto: %v", err)
	//	return nil, err
	//}
	//
	//fw.Client.Crypto = cryptoHelper
	//fw.Crypto = cryptoHelper

	return &fw, nil

}

func SetupCrypto(client *mautrix.Client) (*cryptohelper.CryptoHelper, error) {

	// note that the key doesn't have to be a string, you can directly generate random bytes and store them somewhere in a binary form
	const pickleKeyString = "NnSHJguDSW7vtSshQJh2Yny4zQHc6Wyf"

	// remember to use a secure key for the pickle key in production
	pickleKey := []byte(pickleKeyString)

	// this is a path to the SQLite database you will use to store various data about your bot
	dbPath := "crypto.db"

	helper, err := cryptohelper.NewCryptoHelper(client, pickleKey, dbPath)
	if err != nil {
		return nil, err
	}

	// initialize the database and other stuff
	err = helper.Init(context.Background())
	if err != nil {
		return nil, err
	}

	return helper, nil

}
