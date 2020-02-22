package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		svc := secretsmanager.New(session.New())
		input := &secretsmanager.CreateSecretInput{
			ClientRequestToken: aws.String("EXAMPLE1-90ab-cdef-fedc-ba987SECRET1"),
			Description:        aws.String("My test database secret created with the CLI"),
			Name:               aws.String("MyTestDatabaseSecret"),
			SecretString:       aws.String("{\"username\":\"david\",\"password\":\"BnQw!XDWgaEeT9XGTT29\"}"),
		}

		result, err := svc.CreateSecret(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case secretsmanager.ErrCodeInvalidParameterException:
					fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
				case secretsmanager.ErrCodeInvalidRequestException:
					fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
				case secretsmanager.ErrCodeLimitExceededException:
					fmt.Println(secretsmanager.ErrCodeLimitExceededException, aerr.Error())
				case secretsmanager.ErrCodeEncryptionFailure:
					fmt.Println(secretsmanager.ErrCodeEncryptionFailure, aerr.Error())
				case secretsmanager.ErrCodeResourceExistsException:
					fmt.Println(secretsmanager.ErrCodeResourceExistsException, aerr.Error())
				case secretsmanager.ErrCodeResourceNotFoundException:
					fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
				case secretsmanager.ErrCodeMalformedPolicyDocumentException:
					fmt.Println(secretsmanager.ErrCodeMalformedPolicyDocumentException, aerr.Error())
				case secretsmanager.ErrCodeInternalServiceError:
					fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
				case secretsmanager.ErrCodePreconditionNotMetException:
					fmt.Println(secretsmanager.ErrCodePreconditionNotMetException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
