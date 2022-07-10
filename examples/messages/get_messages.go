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

	go func() {
		// We like to get UpdateNewMessage events without FilterFunc
		receiver := client.AddEventReceiver(&tdlib.UpdateNewMessage{}, func(msg *tdlib.TdMessage) bool {
			return true
		}, 5)
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
