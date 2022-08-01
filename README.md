# aws-vault-prompt

## Example zsh Config

```sh
export AWS_VAULT_PROMPT="$fg_bold[blue]aws:($fg_bold[cyan]%[1]s$fg_bold[blue])$reset_color "
export AWS_VAULT_PROMPT_EXPIRING_SOON="$fg_bold[blue]aws:($fg_bold[yellow]%[1]s ⏱  %[3]dm$fg_bold[blue])$reset_color "
export AWS_VAULT_PROMPT_EXPIRED="$fg_bold[blue]aws:($fg_bold[red]%[1]s ⚠️  $fg_bold[blue])$reset_color "
PROMPT=$PROMPT'$(~/projects/aws-vault-prompt/aws-vault-prompt)'
```