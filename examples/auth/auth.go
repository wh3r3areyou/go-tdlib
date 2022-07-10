package main

import (
	"fmt"
	"github.com/wh3r3areyou/go-tdlib"
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

	// Logic auth
	for {
		currentState, _ := client.Authorize()
		if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPhoneNumberType {
			fmt.Print("Enter phone: ")
			var number string
			fmt.Scanln(&number)
			_, err := client.SendPhoneNumber(number)
			if err != nil {
				fmt.Printf("Error sending phone number: %v", err)
			}
		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitCodeType {
			fmt.Print("Enter code: ")
			var code string
			fmt.Scanln(&code)
			_, err := client.SendAuthCode(code)
			if err != nil {
				fmt.Printf("Error sending auth code : %v", err)
			}
		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPasswordType {
			fmt.Print("Enter Password: ")
			var password string
			fmt.Scanln(&password)
			_, err := client.SendAuthPassword(password)
			if err != nil {
				fmt.Printf("Error sending auth password: %v", err)
			}
		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
			fmt.Println("Authorization Ready!")
			break
		}
	}

	// rawUpdates gets all updates (not only messages)
	rawUpdates := client.GetRawUpdatesChannel(100)
	for update := range rawUpdates {
		fmt.Println(update.Data)
		fmt.Print("\n\n")
	}
}
