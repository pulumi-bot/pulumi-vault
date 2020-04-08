// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource for managing an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).
type AuthBackend struct {
	pulumi.CustomResourceState

	// The accessor for this auth mount.
	Accessor pulumi.StringOutput `pulumi:"accessor"`
	// DN of object to bind when performing user search
	Binddn pulumi.StringOutput `pulumi:"binddn"`
	// Password to use with `binddn` when performing user search
	Bindpass pulumi.StringOutput `pulumi:"bindpass"`
	// Trusted CA to validate TLS certificate
	Certificate  pulumi.StringOutput `pulumi:"certificate"`
	DenyNullBind pulumi.BoolOutput   `pulumi:"denyNullBind"`
	// Description for the LDAP auth backend mount
	Description pulumi.StringOutput `pulumi:"description"`
	Discoverdn  pulumi.BoolOutput   `pulumi:"discoverdn"`
	// LDAP attribute to follow on objects returned by groupfilter
	Groupattr pulumi.StringOutput `pulumi:"groupattr"`
	// Base DN under which to perform group search
	Groupdn pulumi.StringOutput `pulumi:"groupdn"`
	// Go template used to construct group membership query
	Groupfilter pulumi.StringOutput `pulumi:"groupfilter"`
	// Control whether or TLS certificates must be validated
	InsecureTls pulumi.BoolOutput `pulumi:"insecureTls"`
	// Path to mount the LDAP auth backend under
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// Control use of TLS when conecting to LDAP
	Starttls pulumi.BoolOutput `pulumi:"starttls"`
	// Maximum acceptable version of TLS
	TlsMaxVersion pulumi.StringOutput `pulumi:"tlsMaxVersion"`
	// Minimum acceptable version of TLS
	TlsMinVersion pulumi.StringOutput `pulumi:"tlsMinVersion"`
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
	// The userPrincipalDomain used to construct UPN string
	Upndomain pulumi.StringOutput `pulumi:"upndomain"`
	// The URL of the LDAP server
	Url pulumi.StringOutput `pulumi:"url"`
	// Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
	UseTokenGroups pulumi.BoolOutput `pulumi:"useTokenGroups"`
	// Attribute on user object matching username passed in
	Userattr pulumi.StringOutput `pulumi:"userattr"`
	// Base DN under which to perform user search
	Userdn pulumi.StringOutput `pulumi:"userdn"`
}

// NewAuthBackend registers a new resource with the given unique name, arguments, and options.
func NewAuthBackend(ctx *pulumi.Context,
	name string, args *AuthBackendArgs, opts ...pulumi.ResourceOption) (*AuthBackend, error) {
	if args == nil || args.Url == nil {
		return nil, errors.New("missing required argument 'Url'")
	}
	if args == nil {
		args = &AuthBackendArgs{}
	}
	var resource AuthBackend
	err := ctx.RegisterResource("vault:ldap/authBackend:AuthBackend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackend gets an existing AuthBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendState, opts ...pulumi.ResourceOption) (*AuthBackend, error) {
	var resource AuthBackend
	err := ctx.ReadResource("vault:ldap/authBackend:AuthBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackend resources.
type authBackendState struct {
	// The accessor for this auth mount.
	Accessor *string `pulumi:"accessor"`
	// DN of object to bind when performing user search
	Binddn *string `pulumi:"binddn"`
	// Password to use with `binddn` when performing user search
	Bindpass *string `pulumi:"bindpass"`
	// Trusted CA to validate TLS certificate
	Certificate  *string `pulumi:"certificate"`
	DenyNullBind *bool   `pulumi:"denyNullBind"`
	// Description for the LDAP auth backend mount
	Description *string `pulumi:"description"`
	Discoverdn  *bool   `pulumi:"discoverdn"`
	// LDAP attribute to follow on objects returned by groupfilter
	Groupattr *string `pulumi:"groupattr"`
	// Base DN under which to perform group search
	Groupdn *string `pulumi:"groupdn"`
	// Go template used to construct group membership query
	Groupfilter *string `pulumi:"groupfilter"`
	// Control whether or TLS certificates must be validated
	InsecureTls *bool `pulumi:"insecureTls"`
	// Path to mount the LDAP auth backend under
	Path *string `pulumi:"path"`
	// Control use of TLS when conecting to LDAP
	Starttls *bool `pulumi:"starttls"`
	// Maximum acceptable version of TLS
	TlsMaxVersion *string `pulumi:"tlsMaxVersion"`
	// Minimum acceptable version of TLS
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
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
	// The userPrincipalDomain used to construct UPN string
	Upndomain *string `pulumi:"upndomain"`
	// The URL of the LDAP server
	Url *string `pulumi:"url"`
	// Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
	UseTokenGroups *bool `pulumi:"useTokenGroups"`
	// Attribute on user object matching username passed in
	Userattr *string `pulumi:"userattr"`
	// Base DN under which to perform user search
	Userdn *string `pulumi:"userdn"`
}

type AuthBackendState struct {
	// The accessor for this auth mount.
	Accessor pulumi.StringPtrInput
	// DN of object to bind when performing user search
	Binddn pulumi.StringPtrInput
	// Password to use with `binddn` when performing user search
	Bindpass pulumi.StringPtrInput
	// Trusted CA to validate TLS certificate
	Certificate  pulumi.StringPtrInput
	DenyNullBind pulumi.BoolPtrInput
	// Description for the LDAP auth backend mount
	Description pulumi.StringPtrInput
	Discoverdn  pulumi.BoolPtrInput
	// LDAP attribute to follow on objects returned by groupfilter
	Groupattr pulumi.StringPtrInput
	// Base DN under which to perform group search
	Groupdn pulumi.StringPtrInput
	// Go template used to construct group membership query
	Groupfilter pulumi.StringPtrInput
	// Control whether or TLS certificates must be validated
	InsecureTls pulumi.BoolPtrInput
	// Path to mount the LDAP auth backend under
	Path pulumi.StringPtrInput
	// Control use of TLS when conecting to LDAP
	Starttls pulumi.BoolPtrInput
	// Maximum acceptable version of TLS
	TlsMaxVersion pulumi.StringPtrInput
	// Minimum acceptable version of TLS
	TlsMinVersion pulumi.StringPtrInput
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
	// The userPrincipalDomain used to construct UPN string
	Upndomain pulumi.StringPtrInput
	// The URL of the LDAP server
	Url pulumi.StringPtrInput
	// Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
	UseTokenGroups pulumi.BoolPtrInput
	// Attribute on user object matching username passed in
	Userattr pulumi.StringPtrInput
	// Base DN under which to perform user search
	Userdn pulumi.StringPtrInput
}

func (AuthBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendState)(nil)).Elem()
}

type authBackendArgs struct {
	// DN of object to bind when performing user search
	Binddn *string `pulumi:"binddn"`
	// Password to use with `binddn` when performing user search
	Bindpass *string `pulumi:"bindpass"`
	// Trusted CA to validate TLS certificate
	Certificate  *string `pulumi:"certificate"`
	DenyNullBind *bool   `pulumi:"denyNullBind"`
	// Description for the LDAP auth backend mount
	Description *string `pulumi:"description"`
	Discoverdn  *bool   `pulumi:"discoverdn"`
	// LDAP attribute to follow on objects returned by groupfilter
	Groupattr *string `pulumi:"groupattr"`
	// Base DN under which to perform group search
	Groupdn *string `pulumi:"groupdn"`
	// Go template used to construct group membership query
	Groupfilter *string `pulumi:"groupfilter"`
	// Control whether or TLS certificates must be validated
	InsecureTls *bool `pulumi:"insecureTls"`
	// Path to mount the LDAP auth backend under
	Path *string `pulumi:"path"`
	// Control use of TLS when conecting to LDAP
	Starttls *bool `pulumi:"starttls"`
	// Maximum acceptable version of TLS
	TlsMaxVersion *string `pulumi:"tlsMaxVersion"`
	// Minimum acceptable version of TLS
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
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
	// The userPrincipalDomain used to construct UPN string
	Upndomain *string `pulumi:"upndomain"`
	// The URL of the LDAP server
	Url string `pulumi:"url"`
	// Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
	UseTokenGroups *bool `pulumi:"useTokenGroups"`
	// Attribute on user object matching username passed in
	Userattr *string `pulumi:"userattr"`
	// Base DN under which to perform user search
	Userdn *string `pulumi:"userdn"`
}

// The set of arguments for constructing a AuthBackend resource.
type AuthBackendArgs struct {
	// DN of object to bind when performing user search
	Binddn pulumi.StringPtrInput
	// Password to use with `binddn` when performing user search
	Bindpass pulumi.StringPtrInput
	// Trusted CA to validate TLS certificate
	Certificate  pulumi.StringPtrInput
	DenyNullBind pulumi.BoolPtrInput
	// Description for the LDAP auth backend mount
	Description pulumi.StringPtrInput
	Discoverdn  pulumi.BoolPtrInput
	// LDAP attribute to follow on objects returned by groupfilter
	Groupattr pulumi.StringPtrInput
	// Base DN under which to perform group search
	Groupdn pulumi.StringPtrInput
	// Go template used to construct group membership query
	Groupfilter pulumi.StringPtrInput
	// Control whether or TLS certificates must be validated
	InsecureTls pulumi.BoolPtrInput
	// Path to mount the LDAP auth backend under
	Path pulumi.StringPtrInput
	// Control use of TLS when conecting to LDAP
	Starttls pulumi.BoolPtrInput
	// Maximum acceptable version of TLS
	TlsMaxVersion pulumi.StringPtrInput
	// Minimum acceptable version of TLS
	TlsMinVersion pulumi.StringPtrInput
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
	// The userPrincipalDomain used to construct UPN string
	Upndomain pulumi.StringPtrInput
	// The URL of the LDAP server
	Url pulumi.StringInput
	// Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
	UseTokenGroups pulumi.BoolPtrInput
	// Attribute on user object matching username passed in
	Userattr pulumi.StringPtrInput
	// Base DN under which to perform user search
	Userdn pulumi.StringPtrInput
}

func (AuthBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendArgs)(nil)).Elem()
}
