package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/codemonauts/shared-2fa/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/xlzd/gotp"
)

type Entry struct {
	Seed string `json:"seed"`
}

var generateCmd = &cobra.Command{
	Use:   "generate <name>",
	Short: "Generate a token for the given entry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		secretID := fmt.Sprintf("%s%s", config.SECRETS_PREFIX, name)
		svc := secretsmanager.New(session.New(&aws.Config{
			Region: aws.String(config.AWS_REGION),
		}))
		input := &secretsmanager.GetSecretValueInput{
			SecretId: aws.String(secretID),
		}

		result, err := svc.GetSecretValue(input)
		if err != nil {
			fmt.Printf("Couldn't get entry from secrets manager: %s\n", err.Error())
			return
		}

		var e Entry
		json.Unmarshal([]byte(*result.SecretString), &e)
		totp := gotp.NewDefaultTOTP(e.Seed)
		key, expiration := totp.NowWithExpiration()
		exp := expiration - time.Now().Unix()
		fmt.Printf("%s (%s)\n", key, colorExpiration(exp))
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func colorExpiration(duration int64) string {
	s := fmt.Sprintf("Expires in %ds", duration)
	if duration < 10 {
		return color.RedString(s)
	} else {
		return s
	}
}
