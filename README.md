# aws-vault-prompt

A zsh prompt for 99designs/aws-vault to understand what the active AWS session is and when it expires.

## Example zsh Config

```sh
export AWS_VAULT_PROMPT="$fg_bold[blue]aws:($fg_bold[cyan]%[1]s$fg_bold[blue])$reset_color "
export AWS_VAULT_PROMPT_EXPIRING_SOON="$fg_bold[blue]aws:($fg_bold[yellow]%[1]s ⏱  %[3]dm$fg_bold[blue])$reset_color "
export AWS_VAULT_PROMPT_EXPIRED="$fg_bold[blue]aws:($fg_bold[red]%[1]s ⚠️  $fg_bold[blue])$reset_color "
PROMPT=$PROMPT'$(~/projects/aws-vault-prompt/aws-vault-prompt)'
```