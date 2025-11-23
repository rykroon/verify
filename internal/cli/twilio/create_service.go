package twilio

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newCreateServiceCmd() *cobra.Command {
	var form twilio.CreateServiceForm

	var cmd = &cobra.Command{
		Use:   "create-service",
		Short: "Create Service",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := twilio.NewClient(
				os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"), nil,
			)

			result, err := client.CreateService(form)
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

	cmd.Flags().StringVarP(&form.FriendlyName, "friendly-name", "n", "", "A descriptive name for the service")
	cmd.Flags().IntVar(&form.CodeLength, "code-length", 0, "The length of the code")
	cmd.MarkFlagRequired("friendly-name")

	return cmd
}
