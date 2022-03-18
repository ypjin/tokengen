package main

import (
	"fmt"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/ypjin/tokengen/authentication"
	"github.com/ypjin/tokengen/authentication/token"

	"k8s.io/apiserver/pkg/authentication/user"
)

func main() {

	opts := authentication.NewOptions()

	opts.JwtSecret = "h4ZQqIlE8s2td9CMnCgtqLnm7qqFT1YJ"

	opts.OAuthOptions.Issuer = "https://tryit"
	// opts.OAuthOptions.AccessTokenMaxAge = 0

	issuer, err := token.NewIssuer(opts)

	if err != nil {
		panic(err)
	}

	issueRequest := token.IssueRequest{
		User: &user.DefaultInfo{
			Name:   "yamluser",
			UID:    "yamluserID",
			Groups: []string{},
			Extra: map[string][]string{
				"purpose": []string{"fortest"},
			},
		},
		ExpiresIn: time.Minute * 30,
		Claims: token.Claims{
			// TokenType:         token.AccessToken,
			TokenType:         token.StaticToken,
			Email:             "ymaluser@163.com",
			Name:              "yamluser_claimname",
			PreferredUsername: "yamluser_claimpreferredname",
			Username:          "yamluser_claimusername",
			Scopes:            []string{"token"},
			StandardClaims: jwt.StandardClaims{
				Audience: []string{"arscli"},
				// ExpiresAt: ,
			},
		},
	}

	// {"exp":1647336046,"iat":1647332446,"iss":"kubesphere","sub":"admin","token_type":"access_token","username":"admin"}
	// {"exp":1647337846,"iat":1647332446,"iss":"kubesphere","sub":"admin","token_type":"refresh_token","username":"admin"}

	token, err := issuer.IssueTo(&issueRequest)
	if err != nil {
		panic(err)
	}

	fmt.Printf("token: %s", token)

}
