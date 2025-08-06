package telnyx

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
)

var retrieveProfileCmd = &cobra.Command{
	Use:   "retrieve-profile",
	Short: "Retrieve Verification Profiles",
	Long:  ``,
	RunE:  runRetrieveProfile,
}

func runRetrieveProfile(cmd *cobra.Command, args []string) error {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

	resp, err := client.RetrieveVerifyProfile(verifyProfileId)
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
	retrieveProfileCmd.Flags().StringVar(&verifyProfileId, "id", "", "The verification profile id")
	retrieveProfileCmd.MarkFlagRequired("id")
}
