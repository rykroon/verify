package telnyx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var verifyCodeCmd = &cobra.Command{
	Use:   "verify-code",
	Short: "Verify SMS Verification",
	Long:  ``,
	RunE:  runVerifyCode,
}

var verificationId string
var code string

func runVerifyCode(cmd *cobra.Command, args []string) error {
	client := telnyx.NewClient(os.Getenv("TELNYX_API_KEY"))

	req, err := client.NewVerifyCodeRequest(verificationId, code)
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
		return err
	}

	fmt.Println(string(yamlBytes))
	return nil
}

func init() {
	verifyCodeCmd.Flags().StringVarP(&verificationId, "verification-id", "V", "", "The verification id")
	verifyCodeCmd.Flags().StringVarP(&code, "code", "c", "", "The code")

	verifyCodeCmd.MarkFlagRequired("verification-id")
	verifyCodeCmd.MarkFlagRequired("code")
}
