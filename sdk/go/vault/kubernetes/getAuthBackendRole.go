// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Reads the Role of an Kubernetes from a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/api/auth/kubernetes/index.html#read-role) for more
// information.
func LookupAuthBackendRole(ctx *pulumi.Context, args *LookupAuthBackendRoleArgs, opts ...pulumi.InvokeOption) (*LookupAuthBackendRoleResult, error) {
	var rv LookupAuthBackendRoleResult
	err := ctx.Invoke("vault:kubernetes/getAuthBackendRole:getAuthBackendRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthBackendRole.
type LookupAuthBackendRoleArgs struct {
	// (Optional) Audience claim to verify in the JWT.
	Audience *string `pulumi:"audience"`
	// The unique name for the Kubernetes backend the role to
	// retrieve Role attributes for resides in. Defaults to "kubernetes".
	Backend    *string  `pulumi:"backend"`
	BoundCidrs []string `pulumi:"boundCidrs"`
	MaxTtl     *int     `pulumi:"maxTtl"`
	NumUses    *int     `pulumi:"numUses"`
	Period     *int     `pulumi:"period"`
	Policies   []string `pulumi:"policies"`
	// The name of the role to retrieve the Role attributes for.
	RoleName string `pulumi:"roleName"`
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs []string `pulumi:"tokenBoundCidrs"`
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl *int `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy *bool `pulumi:"tokenNoDefaultPolicy"`
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses *int `pulumi:"tokenNumUses"`
	// (Optional) If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod *int `pulumi:"tokenPeriod"`
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies []string `pulumi:"tokenPolicies"`
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl *int `pulumi:"tokenTtl"`
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType *string `pulumi:"tokenType"`
	Ttl       *int    `pulumi:"ttl"`
}

// A collection of values returned by getAuthBackendRole.
type LookupAuthBackendRoleResult struct {
	// (Optional) Audience claim to verify in the JWT.
	Audience   *string  `pulumi:"audience"`
	Backend    *string  `pulumi:"backend"`
	BoundCidrs []string `pulumi:"boundCidrs"`
	// List of service account names able to access this role. If set to "*" all names are allowed, both this and boundServiceAccountNamespaces can not be "*".
	BoundServiceAccountNames []string `pulumi:"boundServiceAccountNames"`
	// List of namespaces allowed to access this role. If set to "*" all namespaces are allowed, both this and boundServiceAccountNames can not be set to "*".
	BoundServiceAccountNamespaces []string `pulumi:"boundServiceAccountNamespaces"`
	// id is the provider-assigned unique ID for this managed resource.
	Id       string   `pulumi:"id"`
	MaxTtl   *int     `pulumi:"maxTtl"`
	NumUses  *int     `pulumi:"numUses"`
	Period   *int     `pulumi:"period"`
	Policies []string `pulumi:"policies"`
	RoleName string   `pulumi:"roleName"`
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs []string `pulumi:"tokenBoundCidrs"`
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl *int `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy *bool `pulumi:"tokenNoDefaultPolicy"`
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses *int `pulumi:"tokenNumUses"`
	// (Optional) If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod *int `pulumi:"tokenPeriod"`
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies []string `pulumi:"tokenPolicies"`
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl *int `pulumi:"tokenTtl"`
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType *string `pulumi:"tokenType"`
	Ttl       *int    `pulumi:"ttl"`
}
