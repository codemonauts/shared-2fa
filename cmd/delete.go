package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codemonauts/shared-2fa/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "Delete an entry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Please type 'YES' if you really want to delete the entry for %q\n", name)
		text, _ := reader.ReadString('\n')
		if strings.TrimSuffix(text, "\n") == "YES" {
			svc := secretsmanager.New(session.New(&aws.Config{
				Region: aws.String(config.AWS_REGION),
			}))
			input := &secretsmanager.DeleteSecretInput{
				SecretId: aws.String(fmt.Sprintf("%s%s", config.SECRETS_PREFIX, name)),
			}
			_, err := svc.DeleteSecret(input)
			if err != nil {
				fmt.Printf("Couldn't delete entry: %s\n", err.Error())
				return
			}

		} else {
			fmt.Println("Aborting.")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
