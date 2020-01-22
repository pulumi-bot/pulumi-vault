// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package kubernetes

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Kubernetes auth backend role in a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/kubernetes.html) for more
// information.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/kubernetes_auth_backend_role.html.markdown.
type AuthBackendRole struct {
	pulumi.CustomResourceState

	// Audience claim to verify in the JWT.
	Audience pulumi.StringPtrOutput `pulumi:"audience"`
	// Unique name of the kubernetes backend to configure.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// If set, a list of
	// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
	BoundCidrs pulumi.StringArrayOutput `pulumi:"boundCidrs"`
	// List of service account names able to access this role. If set to `["*"]` all names are allowed, both this and boundServiceAccountNamespaces can not be "*".
	BoundServiceAccountNames pulumi.StringArrayOutput `pulumi:"boundServiceAccountNames"`
	// List of namespaces allowed to access this role. If set to `["*"]` all namespaces are allowed, both this and boundServiceAccountNames can not be set to "*".
	BoundServiceAccountNamespaces pulumi.StringArrayOutput `pulumi:"boundServiceAccountNamespaces"`
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl pulumi.IntPtrOutput `pulumi:"maxTtl"`
	// If set, puts a use-count
	// limitation on the issued token.
	NumUses pulumi.IntPtrOutput `pulumi:"numUses"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// Name of the role.
	RoleName pulumi.StringOutput `pulumi:"roleName"`
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs pulumi.StringArrayOutput `pulumi:"tokenBoundCidrs"`
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl pulumi.IntPtrOutput `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl pulumi.IntPtrOutput `pulumi:"tokenMaxTtl"`
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy pulumi.BoolPtrOutput `pulumi:"tokenNoDefaultPolicy"`
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses pulumi.IntPtrOutput `pulumi:"tokenNumUses"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod pulumi.IntPtrOutput `pulumi:"tokenPeriod"`
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies pulumi.StringArrayOutput `pulumi:"tokenPolicies"`
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl pulumi.IntPtrOutput `pulumi:"tokenTtl"`
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType pulumi.StringPtrOutput `pulumi:"tokenType"`
	// The TTL period of tokens issued
	// using this role, provided as a number of seconds.
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
}

// NewAuthBackendRole registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendRole(ctx *pulumi.Context,
	name string, args *AuthBackendRoleArgs, opts ...pulumi.ResourceOption) (*AuthBackendRole, error) {
	if args == nil || args.BoundServiceAccountNames == nil {
		return nil, errors.New("missing required argument 'BoundServiceAccountNames'")
	}
	if args == nil || args.BoundServiceAccountNamespaces == nil {
		return nil, errors.New("missing required argument 'BoundServiceAccountNamespaces'")
	}
	if args == nil || args.RoleName == nil {
		return nil, errors.New("missing required argument 'RoleName'")
	}
	if args == nil {
		args = &AuthBackendRoleArgs{}
	}
	var resource AuthBackendRole
	err := ctx.RegisterResource("vault:kubernetes/authBackendRole:AuthBackendRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendRole gets an existing AuthBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendRoleState, opts ...pulumi.ResourceOption) (*AuthBackendRole, error) {
	var resource AuthBackendRole
	err := ctx.ReadResource("vault:kubernetes/authBackendRole:AuthBackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendRole resources.
type authBackendRoleState struct {
	// Audience claim to verify in the JWT.
	Audience *string `pulumi:"audience"`
	// Unique name of the kubernetes backend to configure.
	Backend *string `pulumi:"backend"`
	// If set, a list of
	// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
	BoundCidrs []string `pulumi:"boundCidrs"`
	// List of service account names able to access this role. If set to `["*"]` all names are allowed, both this and boundServiceAccountNamespaces can not be "*".
	BoundServiceAccountNames []string `pulumi:"boundServiceAccountNames"`
	// List of namespaces allowed to access this role. If set to `["*"]` all namespaces are allowed, both this and boundServiceAccountNames can not be set to "*".
	BoundServiceAccountNamespaces []string `pulumi:"boundServiceAccountNamespaces"`
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl *int `pulumi:"maxTtl"`
	// If set, puts a use-count
	// limitation on the issued token.
	NumUses *int `pulumi:"numUses"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period *int `pulumi:"period"`
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies []string `pulumi:"policies"`
	// Name of the role.
	RoleName *string `pulumi:"roleName"`
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
	// If set, indicates that the
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
	// The TTL period of tokens issued
	// using this role, provided as a number of seconds.
	Ttl *int `pulumi:"ttl"`
}

type AuthBackendRoleState struct {
	// Audience claim to verify in the JWT.
	Audience pulumi.StringPtrInput
	// Unique name of the kubernetes backend to configure.
	Backend pulumi.StringPtrInput
	// If set, a list of
	// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
	BoundCidrs pulumi.StringArrayInput
	// List of service account names able to access this role. If set to `["*"]` all names are allowed, both this and boundServiceAccountNamespaces can not be "*".
	BoundServiceAccountNames pulumi.StringArrayInput
	// List of namespaces allowed to access this role. If set to `["*"]` all namespaces are allowed, both this and boundServiceAccountNames can not be set to "*".
	BoundServiceAccountNamespaces pulumi.StringArrayInput
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl pulumi.IntPtrInput
	// If set, puts a use-count
	// limitation on the issued token.
	NumUses pulumi.IntPtrInput
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period pulumi.IntPtrInput
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies pulumi.StringArrayInput
	// Name of the role.
	RoleName pulumi.StringPtrInput
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs pulumi.StringArrayInput
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl pulumi.IntPtrInput
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl pulumi.IntPtrInput
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy pulumi.BoolPtrInput
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses pulumi.IntPtrInput
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod pulumi.IntPtrInput
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies pulumi.StringArrayInput
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl pulumi.IntPtrInput
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType pulumi.StringPtrInput
	// The TTL period of tokens issued
	// using this role, provided as a number of seconds.
	Ttl pulumi.IntPtrInput
}

func (AuthBackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoleState)(nil)).Elem()
}

type authBackendRoleArgs struct {
	// Audience claim to verify in the JWT.
	Audience *string `pulumi:"audience"`
	// Unique name of the kubernetes backend to configure.
	Backend *string `pulumi:"backend"`
	// If set, a list of
	// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
	BoundCidrs []string `pulumi:"boundCidrs"`
	// List of service account names able to access this role. If set to `["*"]` all names are allowed, both this and boundServiceAccountNamespaces can not be "*".
	BoundServiceAccountNames []string `pulumi:"boundServiceAccountNames"`
	// List of namespaces allowed to access this role. If set to `["*"]` all namespaces are allowed, both this and boundServiceAccountNames can not be set to "*".
	BoundServiceAccountNamespaces []string `pulumi:"boundServiceAccountNamespaces"`
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl *int `pulumi:"maxTtl"`
	// If set, puts a use-count
	// limitation on the issued token.
	NumUses *int `pulumi:"numUses"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period *int `pulumi:"period"`
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies []string `pulumi:"policies"`
	// Name of the role.
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
	// If set, indicates that the
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
	// The TTL period of tokens issued
	// using this role, provided as a number of seconds.
	Ttl *int `pulumi:"ttl"`
}

// The set of arguments for constructing a AuthBackendRole resource.
type AuthBackendRoleArgs struct {
	// Audience claim to verify in the JWT.
	Audience pulumi.StringPtrInput
	// Unique name of the kubernetes backend to configure.
	Backend pulumi.StringPtrInput
	// If set, a list of
	// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
	BoundCidrs pulumi.StringArrayInput
	// List of service account names able to access this role. If set to `["*"]` all names are allowed, both this and boundServiceAccountNamespaces can not be "*".
	BoundServiceAccountNames pulumi.StringArrayInput
	// List of namespaces allowed to access this role. If set to `["*"]` all namespaces are allowed, both this and boundServiceAccountNames can not be set to "*".
	BoundServiceAccountNamespaces pulumi.StringArrayInput
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl pulumi.IntPtrInput
	// If set, puts a use-count
	// limitation on the issued token.
	NumUses pulumi.IntPtrInput
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period pulumi.IntPtrInput
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies pulumi.StringArrayInput
	// Name of the role.
	RoleName pulumi.StringInput
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs pulumi.StringArrayInput
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl pulumi.IntPtrInput
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl pulumi.IntPtrInput
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy pulumi.BoolPtrInput
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses pulumi.IntPtrInput
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod pulumi.IntPtrInput
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies pulumi.StringArrayInput
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl pulumi.IntPtrInput
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType pulumi.StringPtrInput
	// The TTL period of tokens issued
	// using this role, provided as a number of seconds.
	Ttl pulumi.IntPtrInput
}

func (AuthBackendRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoleArgs)(nil)).Elem()
}

