# gpgenv
***Securely store and pass secrets as environment variables to other applications!***

## Purpose
It's extremely common for CLI applications to consume secrets via environment variables.
Many people will add these environment variables to their bash profile for convenience.

However, this means two things:

1. Your secrets are passed to every application you run from your terminal
2. Your secrets are accessible on the disk to any software you run

You could store it as a separate script to source when needed, but that still leaves your secrets
on your hard drive.

**gpgenv** aims to solve this by acting as a wrapper around applications that injects your secrets
as environment variables from a GPG-encrypted store.

When combined with an OpenPGP smartcard like a YubiKey 5, your secrets are secured by a second
factor of authentication, protecting you against any software-based attacks.

## Installation

TODO

## Usage

### Configuration
Configuration of gpgenv is performed with `gpgenv edit`. This will decrypt the store and open it in an editor.

Sets of environment variables are called contexts.

Example:
```yaml
contexts:
  terraform:
    CLOUDFLARE_API_KEY: supersecret
    CLOUDFLARE_API_TOKEN: loremipsum
    SUMOLOGIC_ACCESSID: alsoasecret
    SUMOLOGIC_ACCESSKEY: beepboop
  anothercontext:
    ANOTHER_VAR: sosecret
    SECRET: VALUE
  athirdcontext:
    HELLO: WORLD
    LOREM: IPSUM
  
key-email: 86763948+notfromstatefarm@users.noreply.github.com
```

`key-email` should be the email of the GPG key you wish to encrypt and decrypt with.

If you'd like to change the editor from vim, pass the `EDITOR` environment variable with the editor you wish to use i.e. `EDITOR=nano gpgenv edit`.

### Running commands

Simply prepend the command you wish to run with `gpgenv context-name`. For example, to run `terraform plan` with the `tf` context, run:
```shell
gpgenv tf terraform plan
```
### Aliasing
If you'd like to avoid executing the gpgenv command directly, you can set up an alias function in your bash profile.

For example, if you'd like to alias `terraform` to always use the `tf` context, add the following to your bash profile:

```shell
terraform() { gpgenv tf terraform "$@" }
```