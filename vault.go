package vault

import (
	"errors"
	"fmt"
	"strings"

	hashivault "github.com/hashicorp/vault/api"
)

type LDAP struct {
	Client *hashivault.Client
}

type LDAPLoginOptions struct {
	Username  string
	Password  string
	MountPath string
}

func (l *LDAP) LdapLogin(options LDAPLoginOptions) (*hashivault.Secret, error) {
	ldapCreds := map[string]interface{}{
		"password": options.Password,
	}
	pathFormatString := "auth/ldap/login/%s"
	if options.MountPath != "" {
		pathFormatString = "auth/" + strings.Trim(options.MountPath, "/") + "/login/%s"
	}
	normalizedPath := fmt.Sprintf(pathFormatString, options.Username)

	authSecret, err := l.Client.Logical().Write(normalizedPath, ldapCreds)

	if err != nil {
		return nil, err
	}

	if authSecret == nil {
		return nil, errors.New("empty response from vault ldap")
	}

	l.Client.SetToken(authSecret.Auth.ClientToken)

	return authSecret, nil
}

func NewLdapClient(address string) (*LDAP, error) {

	config := hashivault.DefaultConfig()
	config.Address = address
	client, err := hashivault.NewClient(config)

	if err != nil {
		return nil, err
	}

	return &LDAP{
		Client: client,
	}, nil
}

func (l *LDAP) ReadSecretKey(path string, key string) (*string, error) {
	s, err := l.Client.Logical().Read(path)
	if err != nil {
		return nil, err
	}
	value := fmt.Sprintf("%v", s.Data[key])
	return &value, nil
}
