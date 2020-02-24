# shared-2fa

This tool helps you to share a virtual TOTP MFA-device with a team by saving the intial seed at AWS SecretsManager.

```
Available Commands:
  add         Create a new entry
  delete      Delete an entry
  generate    Generate a token for the given entry
  help        Help about any command
  list        A brief description of your command
```

## add
```
Create a new entry

Usage:
  shared-2fa add <name> <seed>
```

## delete
```
Delete an entry

Usage:
  shared-2fa delete <name>
```

## generate
```
Generate a token for the given entry

Usage:
  shared-2fa generate <name>
```

## list
```
List all available entries

Usage:
  shared-2fa list
```
