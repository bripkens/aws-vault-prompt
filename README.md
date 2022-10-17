# aws-vault-prompt

A zsh prompt for [99designs/aws-vault](https://github.com/99designs/aws-vault) to understand what the active AWS session is and when it expires.

## Installation

Download the latest release from the [releases page](https://github.com/bripkens/aws-vault-prompt/releases) and put it on your `$PATH`.

## Configuration

The prompt is configured with the help of the `AWS_VAULT_TERM_PROMPT`, `AWS_VAULT_TERM_PROMPT_EXPIRING_SOON` and `AWS_VAULT_TERM_PROMPT_EXPIRED` environment variables.
All environments have to be golang format strings. You can use the following placeholders:

 - `%[1]s`: Name of the AWS vault
 - `%[2]s`: Name of the default AWS region
 - `%[3]d`: Remaining session duration in minutes.

## Example zsh Config

Make sure to place these lines after: `source $ZSH/oh-my-zsh.sh`

```sh
export AWS_VAULT_TERM_PROMPT="$fg_bold[blue]aws:($fg_bold[cyan]%[1]s$fg_bold[blue])$reset_color "
export AWS_VAULT_TERM_PROMPT_EXPIRING_SOON="$fg_bold[blue]aws:($fg_bold[yellow]%[1]s ⏱  %[3]dm$fg_bold[blue])$reset_color "
export AWS_VAULT_TERM_PROMPT_EXPIRED="$fg_bold[blue]aws:($fg_bold[red]%[1]s ⚠️  $fg_bold[blue])$reset_color "
PROMPT=$PROMPT'$(aws-vault-prompt)'
```

This configuration will result in the following prompt:

### Valid Session

<img width="161" alt="Screenshot 2022-10-17 at 08 36 16" src="https://user-images.githubusercontent.com/596443/196105703-13492b71-3c7c-4e2b-a1e8-1b5759762583.png">

### Session that exprires soon
<img width="205" alt="Screenshot 2022-10-17 at 08 36 26" src="https://user-images.githubusercontent.com/596443/196105708-c53d4d6e-c18b-4822-ac6d-8efee08a8f0f.png">

### Expired Session
<img width="192" alt="Screenshot 2022-10-17 at 08 36 33" src="https://user-images.githubusercontent.com/596443/196105711-1ba149a4-9819-4033-9a47-e0c9fa28afd7.png">
