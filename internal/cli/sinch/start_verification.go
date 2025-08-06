package sinch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/sinch"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var startVerificationCmd = &cobra.Command{
	Use:   "start-verification",
	Short: "Start Verification",
	Long:  `See API documentation: https://developers.sinch.com/docs/verification/api-reference/verification/tag/Verifications-start/#tag/Verifications-start/operation/StartVerification`,
	RunE:  runStartVerification,
}

type startVerificationParams struct {
	identityType     string
	identityEndpoint string
	method           string
}

var svp startVerificationParams

func runStartVerification(cmd *cobra.Command, args []string) error {
	client := sinch.NewClient(os.Getenv("SINCH_APP_KEY"), os.Getenv("SINCH_APP_SECRET"))

	params := sinch.NewStartVerificationParams()
	params.SetIdentityType(svp.identityType)
	params.SetIdentityEndpoint(svp.identityEndpoint)
	params.SetMethod(svp.method)

	req, err := client.NewStartVerificationRequest(params)
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

	yamlBytes, err := yaml.Marshal(m)
	if err != nil {
		return nil
	}

	fmt.Println(string(yamlBytes))
	return nil
}

func init() {
	startVerificationCmd.Flags().StringVarP(&svp.identityType, "identity-type", "t", "number", "")
	startVerificationCmd.Flags().StringVarP(&svp.identityEndpoint, "identity-endpoint", "e", "", "E.164 formatted phone number")
	startVerificationCmd.Flags().StringVarP(&svp.method, "method", "m", "sms", "The type of the verification request")
	startVerificationCmd.MarkFlagRequired("identity-endpoint")
}
