# YubiKey Setup Guide
It is recommended to use a YubiKey 5 with touch enabled so every time your secrets need to be
decrypted, you must physically tap on a device other than your laptop. This provides great protection against
software-based attacks.

Follow this guide to create a GPG key directly on your YubiKey and enable touch-to-sign.

## Change card pin
First off, you're going to want to change the PIN of your YubiKey. Note that the PIN is not limited to numbers, it can be a password.

1. Plug in YubiKey
2. Execute `gpg --change-pin`
3. Change the card PIN first. The default is `123456`.
4. Change the admin PIN. The default is `12345678`.

You will need to enter the card PIN any time you unlock your computer and you first access the OpenPGP module (i.e. your 
laptop going to sleep will cause the session to be lost).

**WARNING:** If you enter in the incorrect PIN three times, the card will be locked. You can then unlock it with the 
admin PIN. If the admin PIN is entered incorrectly three times, all GPG data will be entirely wiped.

## Generate GPG key
This will generate a GPG key directly on your YubiKey, making it impossible to steal the private key data. However, it
also means there is no backup. If you lose your YubiKey, your secrets within gpgenv will be lost. We recommend making a
backup of your secrets elsewhere such as in your password manager, or better yet, an offline backup like a USB key.

1. Plug in your YubiKey
2. Execute `gpg --card-edit`
3. Type `admin` and press enter
4. Type `generate` and press enter
5. When asked `Make off-card backup of encryption key? (Y/n)`, type `n`
6. If asked for your pin, enter your card PIN
7. Enter `0` for validity period
8. Enter `y` to confirm validity
9. Enter your Git username (`git config user.name`)
10. Enter the email you use to commit (`git config user.email`)
11. Don’t enter a comment
12. Enter `O` to confirm
13. If asked for your admin pin, enter the admin pin.
14. Your YubiKey will begin blinking for roughly one minute

## Enable touch-to-sign
You'll need the `ykman` tool from YubiKey to enable touch-to-sign. You can install it with brew or PIP. https://developers.yubico.com/yubikey-manager/

Once you've got ykman installed:
1. Plug in your YubiKey
2. Run `ykman openpgp keys set-touch aut fixed`
3. Run `ykman openpgp keys set-touch sig fixed`
4. Run `ykman openpgp keys set-touch enc fixed`

Now you can use your YubiKey in combination with gpgenv to keep your secrets secure!