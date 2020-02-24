# shared-2fa

This tool helps you to share a virtual TOTP MFA-device with a team by saving the intial seed at AWS SecretsManager.

## Pricing
SecretsManager is charged both per secret and per 10.000 API calls. Every secret costs 0.40\$/Month which will be the
main operational costs of this tool. 10k API calls will cost you 0.05$, which you probably never reach in a month even
with a larger people using this tool multiple times a day.

Because every value in AWS SecretsManager is a JSON object we could save all
seeds in a single key/value pair and cap the monthly costs 0.40\$/month by
this, but would loose the feature of fine-grained access control with an IAM
rule.

## IAM permissions
With this policy one could use all features of this tool. If you want people to just have read access, just remove the 
`Delete` and `Create` actions.
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": [
                "secretsmanager:GetSecretValue",
                "secretsmanager:DeleteSecret"
            ],
            "Resource": "arn:aws:secretsmanager:eu-central-1:<your-account-id>:secret:2fa-*"
        },
        {
            "Sid": "VisualEditor1",
            "Effect": "Allow",
            "Action": [
                "secretsmanager:CreateSecret",
                "secretsmanager:ListSecrets"
            ],
            "Resource": "*"
        }
    ]
}
```

## Usage
```
Available Commands:
  add         Create a new entry
  delete      Delete an entry
  generate    Generate a token for the given entry
  help        Help about any command
  list        A brief description of your command
```

### add
```
Create a new entry

Usage:
  shared-2fa add <name> <seed>
```

### delete
```
Delete an entry

Usage:
  shared-2fa delete <name>
```

### generate
```
Generate a token for the given entry

Usage:
  shared-2fa generate <name>
```

### list
```
List all available entries

Usage:
  shared-2fa list
```
