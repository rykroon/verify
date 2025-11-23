package sinch

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/sinch"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newGetVerificationCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get-verification",
		Short: "Get Verification",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := sinch.NewClient(nil, os.Getenv("SINCH_APP_KEY"), os.Getenv("SINCH_APP_SECRET"))

			result, err := client.GetVerificationById(id)
			if err != nil {
				return err
			}

			yamlBytes, err := yaml.Marshal(result)
			if err != nil {
				return fmt.Errorf("failed to encode to yaml: %w", err)
			}
			fmt.Println(string(yamlBytes))
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "ID of the verification request")
	return cmd
}
