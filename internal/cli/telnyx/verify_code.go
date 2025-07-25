package telnyx

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/internal/telnyx"
	"github.com/spf13/cobra"
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

	result, err := client.VerifyCode(verificationId, code)
	if err != nil {
		return err
	}

	bites, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil
	}

	fmt.Println(string(bites))

	return nil
}

func init() {
	verifyCodeCmd.Flags().StringVarP(&verificationId, "verification-id", "V", "", "The verification id")
	verifyCodeCmd.Flags().StringVarP(&code, "code", "c", "", "The code")

	verifyCodeCmd.MarkFlagRequired("verification-id")
	verifyCodeCmd.MarkFlagRequired("code")
}
