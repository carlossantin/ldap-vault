# LDAP-VAULT

The goal of this project is to allow reading secrets from Hashicorp Vault using LDAP credentials

## How to import this project:
```
go get github.com/carlossantin/ldap-vault
```

## Example of usage:

```
fmt.Print("Enter your LDAP username: ")
var username string
fmt.Scanln(&username)

fmt.Print("Enter your LDAP password: ")
bytePass, err := term.ReadPassword(int(syscall.Stdin))
fmt.Println()
if err != nil {
    log.Fatalf("unable to read ldap password: %v", err)
}

ldapClient, err := vault.NewLdapClient("https://vault.domain.com:443")
if err != nil {
    log.Fatalf("unable to initialize Vault client: %v", err)
}

ldapOptions := vault.LDAPLoginOptions{
    Username: username,
    Password: string(bytePass),
}
_, err = ldapClient.LdapLogin(ldapOptions)
if err != nil {
    log.Fatalf("Error login in vault: %v", err)
}

// Read database password and connect in database
dbPass, err := ldapClient.ReadSecretKey("secret/db/mongodb.app", "password")
if err != nil {
    log.Fatalf("Error reading database password from vault: %v", err)
}
```