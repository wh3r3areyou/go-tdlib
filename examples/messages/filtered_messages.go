package main

import (
	"fmt"
	"github.com/wh3r3areyou/go-tdlib"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// not necessary
	tdlib.SetLogVerbosityLevel(1)
	// not necessary
	tdlib.SetFilePath("./errors.txt")

	// Get APIID and APIHash - https://my.telegram.org/
	client := tdlib.NewClient(tdlib.Config{
		APIID:               "INPUT_API_ID",
		APIHash:             "INPUT_API_HASH",
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
		UseMessageDatabase:  true,
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseTestDataCenter:   false,
		DatabaseDirectory:   "./tdlib-db",
		FileDirectory:       "./tdlib-files",
		IgnoreFileNames:     false,
	})
	defer client.DestroyInstance()

	// Need wait for get state authorization
	for {
		authState, err := client.Authorize()
		if err != nil {
			panic(err)
		}
		if authState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
			fmt.Println("Success authorization!")
			break
		}
		time.Sleep(300 * time.Millisecond)
	}

	// Get self
	user, err := client.GetMe()
	if err != nil {
		panic(err)
	}

	// Filtering messages
	eventFilter := func(msg *tdlib.TdMessage) bool {
		newMsg := (*msg).(*tdlib.UpdateNewMessage)
		// For example, we will not accept messages that we have sent, but only from other users
		if newMsg.Message.SenderID.(*tdlib.MessageSenderUser).UserID != user.ID {
			return true
		}
		return false
	}

	go func() {
		// We like to get UpdateNewMessage events and with a specific FilterFunc
		receiver := client.AddEventReceiver(&tdlib.UpdateNewMessage{}, eventFilter, 5)
		for newMsg := range receiver.Chan {
			msg := (newMsg).(*tdlib.UpdateNewMessage)
			text := msg.Message.Content.(*tdlib.MessageText)
			fmt.Println("Text:", text)
		}
	}()

	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

}
