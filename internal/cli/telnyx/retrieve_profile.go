package telnyx

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newRetrieveProfileCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "retrieve-profile",
		Short: "Retrieve Verification Profiles",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

			result, err := client.RetrieveVerifyProfile(verifyProfileId)
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

	cmd.Flags().StringVar(&verifyProfileId, "id", "", "The verification profile id")
	cmd.MarkFlagRequired("id")
	return cmd
}
