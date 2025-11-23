package telnyx

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newListProfilesCmd() *cobra.Command {
	var params telnyx.ListVerifyProfilesParams

	var cmd = &cobra.Command{
		Use:   "list-profiles",
		Short: "List verification profiles.",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

			result, err := client.ListVerifyProfiles(params)
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

	cmd.Flags().IntVar(&params.PageNumber, "page-number", 0, "The page number")
	cmd.Flags().IntVar(&params.PageSize, "page-size", 0, "The page size")
	return cmd
}
