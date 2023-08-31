package main

import (
	"context"
	"fmt"
	client "github.com/ory/kratos-client-go"
	"os"
)

func main() {

	id := "devang.chhajed@synthix.com" // string | ID must be set to the ID of identity you want to get
	includeCredential := []string{id}  // []string | Include Credentials in Response  Currently, only `oidc` is supported. This will return the initial OAuth 2.0 Access, Refresh and (optionally) OpenID Connect ID Token. (optional)

	headers := make(map[string]string)

	token := "{get your token from gcloud auth print-identity-token from console}"

	headers["Authorization"] = "Bearer " + token

	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: "https://frontier-idp-backend-service-p57kvjg3ta-ey.a.run.app", // Kratos Admin API
		},
	}
	configuration.DefaultHeader = headers
	apiClient := client.NewAPIClient(configuration)

	resp, r, err := apiClient.IdentityApi.GetIdentity(context.Background(), "").IncludeCredential(includeCredential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentity`: Identity
	fmt.Fprintf(os.Stdout, "Response from `IdentityApi.GetIdentity`: %v\n", resp)
}
