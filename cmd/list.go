package cmd

import (
	"fmt"
	"strings"

	"github.com/codemonauts/shared-2fa/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available entries",
	Run: func(cmd *cobra.Command, args []string) {
		svc := secretsmanager.New(session.New(&aws.Config{
			Region: aws.String(config.AWS_REGION),
		}))
		input := &secretsmanager.ListSecretsInput{}
		data, err := svc.ListSecrets(input)
		if err != nil {
			fmt.Errorf("error reading from secretsmanager: %s\n", err.Error())
			return
		}
		fmt.Println("Found the following 2FA entries:")
		for _, entry := range data.SecretList {
			if strings.HasPrefix(*entry.Name, config.SECRETS_PREFIX) {
				cleanName := strings.Replace(*entry.Name, config.SECRETS_PREFIX, "", 1)
				fmt.Printf("- %s\n", cleanName)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
