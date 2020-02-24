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

var addCmd = &cobra.Command{
	Use:   "add <name> <seed>",
	Short: "Create a new entry",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		secretID := fmt.Sprintf("%s%s", config.SECRETS_PREFIX, name)
		token := strings.ReplaceAll(strings.ToUpper(args[1]), " ", "")

		svc := secretsmanager.New(session.New(&aws.Config{
			Region: aws.String(config.AWS_REGION),
		}))
		input := &secretsmanager.CreateSecretInput{
			Name:         aws.String(secretID),
			SecretString: aws.String(fmt.Sprintf("{\"seed\":\"%s\"}", token)),
		}

		_, err := svc.CreateSecret(input)
		if err != nil {
			fmt.Printf("error while writing to secretsmanager: %s\n", err.Error())
			return
		}
		fmt.Println("Saved successfully.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
