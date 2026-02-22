package main

import (
	"context"
	"os"
	"sync"

	"github.com/michaeldowling/didexec/go/framework"
	log "github.com/sirupsen/logrus"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
)

func main() {

	log.Debugf("Removing Existing CryptoDB")
	os.Remove("crypto.db")
	os.Remove("crypto.db-shm")
	os.Remove("crypto.db-wal")

	log.Debugf("Starting connection")
	fw, err := framework.Connect("matrix.openreserve.io", "alpha", "alphapassword", "")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Debugf("Connected: %v", fw.Client.DeviceID)

	log.Debugf("Inviting External User")
	roomId, err := fw.Invite("@mzero:openreserve.io")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Debugf("Invited External User to Room: %v", roomId)

	err = fw.Send(roomId, "@mzero@openreserve.io", "Hello World!")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Lets wait for a response and do things
	syncer := mautrix.NewDefaultSyncer()
	fw.Client.Syncer = syncer

	syncErrChan := make(chan error)
	go func() {
		if cerr := fw.Client.Sync(); cerr != nil {
			syncErrChan <- cerr
			return
		}
	}()

	readyChan := make(chan bool)
	var wg sync.WaitGroup
	syncer.OnSync(func(ctx context.Context, resp *mautrix.RespSync, since string) bool {
		wg.Go(func() {
			log.Infof("Synced to %s", since)
			readyChan <- true
		})

		return true
	})

	log.Println("Waiting for sync to receive first event from the encrypted room...")
	<-readyChan
	log.Println("Sync received")

	//recoveryKey := os.Getenv("RECOVERY_KEY")
	//err = fw.VerifyWithRecoveryKey(fw.Crypto.Machine(), recoveryKey)
	//if err != nil {
	//	log.Fatal("Error verifying with recovery key: %v", err)
	//	os.Exit(1)
	//}

	syncer.OnEventType(event.EventMessage, func(ctx context.Context, event *event.Event) {
		log.Infof("Received event: %s", event.Type)
		log.Infof("Content: %s", event.Content)

	})

	// wait for error
	syncErr := <-syncErrChan
	if syncErr != nil {
		log.Fatal(syncErr)
		os.Exit(1)
	}

}
