package sinch

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/sinch"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var getVerificationCmd = &cobra.Command{
	Use:   "get-verification",
	Short: "Get Verification",
	Long:  ``,
	RunE:  runGetVerification,
}

type getVerificationParams struct {
	id string
}

var gvp reportVerificationParams

func runGetVerification(cmd *cobra.Command, args []string) error {
	client := sinch.NewClient(os.Getenv("SINCH_APP_KEY"), os.Getenv("SINCH_APP_SECRET"))

	jsonBytes, err := client.GetVerificationById(rvp.id)
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
	getVerificationCmd.Flags().StringVar(&rvp.id, "id", "", "ID of the verification request")
}
