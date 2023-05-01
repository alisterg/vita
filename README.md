vīta

_latin - life_

Record everything about your life. Currently using DynamoDB for data storage / retrieval.

# SETUP

## Prerequisites

- Go >= 1.20.3 `brew install go`
- Sam CLI `brew install aws-sam-cli`
- AWS CLI credentials setup

## entry_types.json

The file `cli/entry_types.json` contains your personal data types that the CLI application
will recognise.

Example:

```json
{ "location": ["City name"] }
```

This will allow you to enter a `location` entry into the application; it will prompt you for
the 'City name'.

## Infrastructure

In the `infra` directory:

1. `sam package`
2. `sam deploy --stack-name MyStackName --capabilities CAPABILITY_IAM`

This will create the CloudFormation stack containing a DynamoDB table and IAM role to access it.

# USAGE

todo

# APP DETAILS

Using Cobra for managing CLI commands and flags:

https://github.com/spf13/cobra/blob/main/user_guide.md
