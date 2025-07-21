package main

import (
	"fmt"

	"github.com/rykroon/verify/internal/telnyx"
)

func main() {
	fmt.Println("Hello World!")
	client := telnyx.NewClient("")
	params := &telnyx.TriggerSmsVerificationParams{
		PhoneNumber:     "+17325132147",
		VerifyProfileId: "4900018e-c5c2-5b29-0fc3-ece4b31251e5",
	}
	result, err := client.TriggerSmsVerification(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
