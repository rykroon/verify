package telnyx

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newCreateProfileCmd() *cobra.Command {
	var params telnyx.CreateVerifyProfileParams

	var cmd = &cobra.Command{
		Use:   "create-profile",
		Short: "Create Verification Profile",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))
			result, err := client.CreateVerifyProfile(params)
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

	cmd.Flags().StringVarP(&params.Name, "name", "n", "", "The name of the profile")
	// createProfileCmd.Flags().StringVarP(&cvpp.Sms.AppName, "app-name", "a", "", "The Nname of the application")
	cmd.MarkFlagRequired("name")

	return cmd
}
