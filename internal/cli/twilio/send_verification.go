package twilio

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var sendVerificationCmd = &cobra.Command{
	Use:   "send-verification",
	Short: "Send Verification",
	Long:  ``,
	RunE:  runSendVerificationCmd,
}

type sendVerificationParams struct {
	serviceSid string
	*twilio.SendVerificationParams
}

var svp *sendVerificationParams

func runSendVerificationCmd(cmd *cobra.Command, args []string) error {
	client := twilio.NewClient(nil, os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))

	resp, err := client.SendVerification(svp.serviceSid, svp.SendVerificationParams)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)

	var m map[string]any
	if err := json.Unmarshal(resp.Body, &m); err != nil {
		return err
	}

	rawYaml, err := yaml.Marshal(m)
	if err != nil {
		return nil
	}

	fmt.Println(string(rawYaml))

	return nil
}

func init() {
	svp = &sendVerificationParams{"", &twilio.SendVerificationParams{}}
	sendVerificationCmd.Flags().StringVar(&svp.serviceSid, "service-sid", "", "The SID of the verification Service.")
	sendVerificationCmd.Flags().StringVar(&svp.To, "to", "", "The phone number or email to verify.")
	sendVerificationCmd.Flags().StringVar(&svp.Channel, "channel", "", "The verification method to use.")
	sendVerificationCmd.MarkFlagRequired("service-sid")
	sendVerificationCmd.MarkFlagRequired("to")
	sendVerificationCmd.MarkFlagRequired("channel")
}
