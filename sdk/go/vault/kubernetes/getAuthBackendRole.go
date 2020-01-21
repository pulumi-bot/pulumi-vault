// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Reads the Role of an Kubernetes from a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/api/auth/kubernetes/index.html#read-role) for more
// information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/kubernetes_auth_backend_role.html.markdown.
func LookupAuthBackendRole(ctx *pulumi.Context, args *GetAuthBackendRoleArgs) (*GetAuthBackendRoleResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["audience"] = args.Audience
		inputs["backend"] = args.Backend
		inputs["boundCidrs"] = args.BoundCidrs
		inputs["maxTtl"] = args.MaxTtl
		inputs["numUses"] = args.NumUses
		inputs["period"] = args.Period
		inputs["policies"] = args.Policies
		inputs["roleName"] = args.RoleName
		inputs["tokenBoundCidrs"] = args.TokenBoundCidrs
		inputs["tokenExplicitMaxTtl"] = args.TokenExplicitMaxTtl
		inputs["tokenMaxTtl"] = args.TokenMaxTtl
		inputs["tokenNoDefaultPolicy"] = args.TokenNoDefaultPolicy
		inputs["tokenNumUses"] = args.TokenNumUses
		inputs["tokenPeriod"] = args.TokenPeriod
		inputs["tokenPolicies"] = args.TokenPolicies
		inputs["tokenTtl"] = args.TokenTtl
		inputs["tokenType"] = args.TokenType
		inputs["ttl"] = args.Ttl
	}
	outputs, err := ctx.Invoke("vault:kubernetes/getAuthBackendRole:getAuthBackendRole", inputs)
	if err != nil {
		return nil, err
	}
	return &GetAuthBackendRoleResult{
		Audience: outputs["audience"],
		Backend: outputs["backend"],
		BoundCidrs: outputs["boundCidrs"],
		BoundServiceAccountNames: outputs["boundServiceAccountNames"],
		BoundServiceAccountNamespaces: outputs["boundServiceAccountNamespaces"],
		MaxTtl: outputs["maxTtl"],
		NumUses: outputs["numUses"],
		Period: outputs["period"],
		Policies: outputs["policies"],
		RoleName: outputs["roleName"],
		TokenBoundCidrs: outputs["tokenBoundCidrs"],
		TokenExplicitMaxTtl: outputs["tokenExplicitMaxTtl"],
		TokenMaxTtl: outputs["tokenMaxTtl"],
		TokenNoDefaultPolicy: outputs["tokenNoDefaultPolicy"],
		TokenNumUses: outputs["tokenNumUses"],
		TokenPeriod: outputs["tokenPeriod"],
		TokenPolicies: outputs["tokenPolicies"],
		TokenTtl: outputs["tokenTtl"],
		TokenType: outputs["tokenType"],
		Ttl: outputs["ttl"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getAuthBackendRole.
type GetAuthBackendRoleArgs struct {
	Audience interface{}
	// The unique name for the Kubernetes backend the role to
	// retrieve Role attributes for resides in. Defaults to "kubernetes".
	Backend interface{}
	// DeprecationMessage: use `token_bound_cidrs` instead if you are running Vault >= 1.2
	BoundCidrs interface{}
	// DeprecationMessage: use `token_max_ttl` instead if you are running Vault >= 1.2
	MaxTtl interface{}
	// DeprecationMessage: use `token_num_uses` instead if you are running Vault >= 1.2
	NumUses interface{}
	// DeprecationMessage: use `token_period` instead if you are running Vault >= 1.2
	Period interface{}
	// DeprecationMessage: use `token_policies` instead if you are running Vault >= 1.2
	Policies interface{}
	// The name of the role to retrieve the Role attributes for.
	RoleName interface{}
	TokenBoundCidrs interface{}
	TokenExplicitMaxTtl interface{}
	TokenMaxTtl interface{}
	TokenNoDefaultPolicy interface{}
	TokenNumUses interface{}
	TokenPeriod interface{}
	TokenPolicies interface{}
	TokenTtl interface{}
	TokenType interface{}
	// DeprecationMessage: use `token_ttl` instead if you are running Vault >= 1.2
	Ttl interface{}
}

// A collection of values returned by getAuthBackendRole.
type GetAuthBackendRoleResult struct {
	// (Optional) Audience claim to verify in the JWT.
	Audience interface{}
	Backend interface{}
	// DeprecationMessage: use `token_bound_cidrs` instead if you are running Vault >= 1.2
	BoundCidrs interface{}
	// List of service account names able to access this role. If set to "*" all names are allowed, both this and boundServiceAccountNamespaces can not be "*".
	BoundServiceAccountNames interface{}
	// List of namespaces allowed to access this role. If set to "*" all namespaces are allowed, both this and boundServiceAccountNames can not be set to "*".
	BoundServiceAccountNamespaces interface{}
	// DeprecationMessage: use `token_max_ttl` instead if you are running Vault >= 1.2
	MaxTtl interface{}
	// DeprecationMessage: use `token_num_uses` instead if you are running Vault >= 1.2
	NumUses interface{}
	// DeprecationMessage: use `token_period` instead if you are running Vault >= 1.2
	Period interface{}
	// DeprecationMessage: use `token_policies` instead if you are running Vault >= 1.2
	Policies interface{}
	RoleName interface{}
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs interface{}
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl interface{}
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl interface{}
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy interface{}
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses interface{}
	// (Optional) If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod interface{}
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies interface{}
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl interface{}
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType interface{}
	// DeprecationMessage: use `token_ttl` instead if you are running Vault >= 1.2
	Ttl interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
