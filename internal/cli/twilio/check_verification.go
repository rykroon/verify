package twilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var checkVerificationCmd = &cobra.Command{
	Use:   "check-verification",
	Short: "Check Verification",
	Long:  ``,
	RunE:  runCheckVerificationCmd,
}

type checkVerificationParams struct {
	serviceSid      string
	to              string
	verificationSid string
	code            string
}

var cvp checkVerificationParams

func runCheckVerificationCmd(cmd *cobra.Command, args []string) error {
	client := twilio.NewClient(os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))
	params := twilio.NewCheckVerificationParams()

	if cvp.to != "" {
		params.SetTo(cvp.to)
	} else if cvp.verificationSid != "" {
		params.SetVerificationSid(cvp.verificationSid)
	}
	params.SetCode(cvp.code)

	req, err := client.NewCheckVerificationRequest(cvp.serviceSid, params)
	if err != nil {
		return err
	}

	resp, err := utils.DoAndReadAll(http.DefaultClient, req)
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
	checkVerificationCmd.Flags().StringVarP(&cvp.serviceSid, "service-sid", "s", "", "The SID of the verification Service.")
	checkVerificationCmd.Flags().StringVarP(&cvp.to, "to", "t", "", "The phone number or email to verify.")
	checkVerificationCmd.Flags().StringVarP(&cvp.verificationSid, "verification-sid", "V", "", "A SID that uniquely identifies the Verification Check.")
	checkVerificationCmd.Flags().StringVarP(&cvp.code, "code", "c", "", "The 4-10 character string being verified.")
	checkVerificationCmd.MarkFlagRequired("service-sid")
	checkVerificationCmd.MarkFlagsOneRequired("to", "verification-sid")
	checkVerificationCmd.MarkFlagRequired("code")
}
