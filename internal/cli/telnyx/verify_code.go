package telnyx

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newVerifyCodeCmd() *cobra.Command {
	var verificationId string
	var payload telnyx.VerifyCodePayload

	var cmd = &cobra.Command{
		Use:   "verify-code",
		Short: "Verify SMS Verification",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

			result, err := client.VerifyCode(verificationId, payload)
			if err != nil {
				return err
			}

			yamlBytes, err := yaml.Marshal(result)
			if err != nil {
				return fmt.Errorf("failed to encode as yaml: %w", err)
			}
			fmt.Println(string(yamlBytes))
			return nil
		},
	}

	cmd.Flags().StringVar(&verificationId, "id", "", "The verification id")
	cmd.Flags().StringVarP(&payload.Code, "code", "c", "", "The code")

	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("code")

	return cmd

}
