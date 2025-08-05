package sinch

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/sinch"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var reportVerificationCmd = &cobra.Command{
	Use:   "report-verification",
	Short: "Report Verification",
	Long:  `See API documentation: https://developers.sinch.com/docs/verification/api-reference/verification/tag/Verifications-report/`,
	RunE:  runReportVerification,
}

type reportVerificationParams struct {
	id     string
	method string
	code   string
}

var rvp reportVerificationParams

func runReportVerification(cmd *cobra.Command, args []string) error {
	client := sinch.NewClient(os.Getenv("SINCH_APP_KEY"), os.Getenv("SINCH_APP_SECRET"))
	params := sinch.NewReportVerificationParams()

	params.SetCode(rvp.code)
	params.SetMethod(rvp.method)

	jsonBytes, err := client.ReportVerificationById(rvp.id, params)
	if err != nil {
		return err
	}

	var m map[string]any
	if err := json.Unmarshal(jsonBytes, &m); err != nil {
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
	reportVerificationCmd.Flags().StringVarP(&rvp.id, "id", "i", "", "Verification ID of the verification request")
	reportVerificationCmd.Flags().StringVarP(&rvp.method, "method", "m", "", "The type of the verification")
	reportVerificationCmd.Flags().StringVarP(&rvp.code, "code", "c", "", "The code which was received by the user submitting the SMS verification")
	reportVerificationCmd.MarkFlagRequired("id")
	reportVerificationCmd.MarkFlagRequired("method")
	reportVerificationCmd.MarkFlagRequired("code")
}
