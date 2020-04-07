// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package jwt

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an JWT/OIDC auth backend role in a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/jwt.html) for more
// information.
type AuthBackendRole struct {
	pulumi.CustomResourceState

	// The list of allowed values for redirectUri during OIDC logins.
	// Required for OIDC roles
	AllowedRedirectUris pulumi.StringArrayOutput `pulumi:"allowedRedirectUris"`
	// The unique name of the auth backend to configure.
	// Defaults to `jwt`.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// List of `aud` claims to match
	// against. Any match is sufficient.
	BoundAudiences pulumi.StringArrayOutput `pulumi:"boundAudiences"`
	// If set, a list of
	// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
	BoundCidrs pulumi.StringArrayOutput `pulumi:"boundCidrs"`
	// If set, a map of claims/values to match against.
	// The expected value may be a single string or a list of strings.
	BoundClaims pulumi.MapOutput `pulumi:"boundClaims"`
	// If set, requires that the `sub` claim matches
	// this value.
	BoundSubject pulumi.StringPtrOutput `pulumi:"boundSubject"`
	// If set, a map of claims (keys) to be copied
	// to specified metadata fields (values).
	ClaimMappings pulumi.MapOutput `pulumi:"claimMappings"`
	// The amount of leeway to add to all claims to account for clock skew, in
	// seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
	// Only applicable with "jwt" roles.
	ClockSkewLeeway pulumi.IntPtrOutput `pulumi:"clockSkewLeeway"`
	// The amount of leeway to add to expiration (`exp`) claims to account for
	// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
	// Only applicable with "jwt" roles.
	ExpirationLeeway pulumi.IntPtrOutput `pulumi:"expirationLeeway"`
	// The claim to use to uniquely identify
	// the set of groups to which the user belongs; this will be used as the names
	// for the Identity group aliases created due to a successful login. The claim
	// value must be a list of strings.
	GroupsClaim pulumi.StringPtrOutput `pulumi:"groupsClaim"`
	// (Optional; Deprecated. This field has been
	// removed since Vault 1.1. If the groups claim is not at the top level, it can
	// now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
	// A pattern of delimiters
	// used to allow the groupsClaim to live outside of the top-level JWT structure.
	// For instance, a groupsClaim of meta/user.name/groups with this field
	// set to // will expect nested structures named meta, user.name, and groups.
	// If this field was set to /./ the groups information would expect to be
	// via nested structures of meta, user, name, and groups.
	GroupsClaimDelimiterPattern pulumi.StringPtrOutput `pulumi:"groupsClaimDelimiterPattern"`
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl pulumi.IntPtrOutput `pulumi:"maxTtl"`
	// The amount of leeway to add to not before (`nbf`) claims to account for
	// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
	// Only applicable with "jwt" roles.
	NotBeforeLeeway pulumi.IntPtrOutput `pulumi:"notBeforeLeeway"`
	// If set, puts a use-count
	// limitation on the issued token.
	NumUses pulumi.IntPtrOutput `pulumi:"numUses"`
	// If set, a list of OIDC scopes to be used with an OIDC role.
	// The standard scope "openid" is automatically included and need not be specified.
	OidcScopes pulumi.StringArrayOutput `pulumi:"oidcScopes"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// The name of the role.
	RoleName pulumi.StringOutput `pulumi:"roleName"`
	// Type of role, either "oidc" (default) or "jwt".
	RoleType pulumi.StringOutput `pulumi:"roleType"`
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
	// The claim to use to uniquely identify
	// the user; this will be used as the name for the Identity entity alias created
	// due to a successful login.
	UserClaim pulumi.StringOutput `pulumi:"userClaim"`
	// Log received OIDC tokens and claims when debug-level
	// logging is active. Not recommended in production since sensitive information may be present
	// in OIDC responses.
	VerboseOidcLogging pulumi.BoolPtrOutput `pulumi:"verboseOidcLogging"`
}

// NewAuthBackendRole registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendRole(ctx *pulumi.Context,
	name string, args *AuthBackendRoleArgs, opts ...pulumi.ResourceOption) (*AuthBackendRole, error) {
	if args == nil || args.RoleName == nil {
		return nil, errors.New("missing required argument 'RoleName'")
	}
	if args == nil || args.UserClaim == nil {
		return nil, errors.New("missing required argument 'UserClaim'")
	}
	if args == nil {
		args = &AuthBackendRoleArgs{}
	}
	var resource AuthBackendRole
	err := ctx.RegisterResource("vault:jwt/authBackendRole:AuthBackendRole", name, args, &resource, opts...)
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
	err := ctx.ReadResource("vault:jwt/authBackendRole:AuthBackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendRole resources.
type authBackendRoleState struct {
	// The list of allowed values for redirectUri during OIDC logins.
	// Required for OIDC roles
	AllowedRedirectUris []string `pulumi:"allowedRedirectUris"`
	// The unique name of the auth backend to configure.
	// Defaults to `jwt`.
	Backend *string `pulumi:"backend"`
	// List of `aud` claims to match
	// against. Any match is sufficient.
	BoundAudiences []string `pulumi:"boundAudiences"`
	// If set, a list of
	// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
	BoundCidrs []string `pulumi:"boundCidrs"`
	// If set, a map of claims/values to match against.
	// The expected value may be a single string or a list of strings.
	BoundClaims map[string]interface{} `pulumi:"boundClaims"`
	// If set, requires that the `sub` claim matches
	// this value.
	BoundSubject *string `pulumi:"boundSubject"`
	// If set, a map of claims (keys) to be copied
	// to specified metadata fields (values).
	ClaimMappings map[string]interface{} `pulumi:"claimMappings"`
	// The amount of leeway to add to all claims to account for clock skew, in
	// seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
	// Only applicable with "jwt" roles.
	ClockSkewLeeway *int `pulumi:"clockSkewLeeway"`
	// The amount of leeway to add to expiration (`exp`) claims to account for
	// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
	// Only applicable with "jwt" roles.
	ExpirationLeeway *int `pulumi:"expirationLeeway"`
	// The claim to use to uniquely identify
	// the set of groups to which the user belongs; this will be used as the names
	// for the Identity group aliases created due to a successful login. The claim
	// value must be a list of strings.
	GroupsClaim *string `pulumi:"groupsClaim"`
	// (Optional; Deprecated. This field has been
	// removed since Vault 1.1. If the groups claim is not at the top level, it can
	// now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
	// A pattern of delimiters
	// used to allow the groupsClaim to live outside of the top-level JWT structure.
	// For instance, a groupsClaim of meta/user.name/groups with this field
	// set to // will expect nested structures named meta, user.name, and groups.
	// If this field was set to /./ the groups information would expect to be
	// via nested structures of meta, user, name, and groups.
	GroupsClaimDelimiterPattern *string `pulumi:"groupsClaimDelimiterPattern"`
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl *int `pulumi:"maxTtl"`
	// The amount of leeway to add to not before (`nbf`) claims to account for
	// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
	// Only applicable with "jwt" roles.
	NotBeforeLeeway *int `pulumi:"notBeforeLeeway"`
	// If set, puts a use-count
	// limitation on the issued token.
	NumUses *int `pulumi:"numUses"`
	// If set, a list of OIDC scopes to be used with an OIDC role.
	// The standard scope "openid" is automatically included and need not be specified.
	OidcScopes []string `pulumi:"oidcScopes"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period *int `pulumi:"period"`
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies []string `pulumi:"policies"`
	// The name of the role.
	RoleName *string `pulumi:"roleName"`
	// Type of role, either "oidc" (default) or "jwt".
	RoleType *string `pulumi:"roleType"`
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
	// The claim to use to uniquely identify
	// the user; this will be used as the name for the Identity entity alias created
	// due to a successful login.
	UserClaim *string `pulumi:"userClaim"`
	// Log received OIDC tokens and claims when debug-level
	// logging is active. Not recommended in production since sensitive information may be present
	// in OIDC responses.
	VerboseOidcLogging *bool `pulumi:"verboseOidcLogging"`
}

type AuthBackendRoleState struct {
	// The list of allowed values for redirectUri during OIDC logins.
	// Required for OIDC roles
	AllowedRedirectUris pulumi.StringArrayInput
	// The unique name of the auth backend to configure.
	// Defaults to `jwt`.
	Backend pulumi.StringPtrInput
	// List of `aud` claims to match
	// against. Any match is sufficient.
	BoundAudiences pulumi.StringArrayInput
	// If set, a list of
	// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
	BoundCidrs pulumi.StringArrayInput
	// If set, a map of claims/values to match against.
	// The expected value may be a single string or a list of strings.
	BoundClaims pulumi.MapInput
	// If set, requires that the `sub` claim matches
	// this value.
	BoundSubject pulumi.StringPtrInput
	// If set, a map of claims (keys) to be copied
	// to specified metadata fields (values).
	ClaimMappings pulumi.MapInput
	// The amount of leeway to add to all claims to account for clock skew, in
	// seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
	// Only applicable with "jwt" roles.
	ClockSkewLeeway pulumi.IntPtrInput
	// The amount of leeway to add to expiration (`exp`) claims to account for
	// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
	// Only applicable with "jwt" roles.
	ExpirationLeeway pulumi.IntPtrInput
	// The claim to use to uniquely identify
	// the set of groups to which the user belongs; this will be used as the names
	// for the Identity group aliases created due to a successful login. The claim
	// value must be a list of strings.
	GroupsClaim pulumi.StringPtrInput
	// (Optional; Deprecated. This field has been
	// removed since Vault 1.1. If the groups claim is not at the top level, it can
	// now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
	// A pattern of delimiters
	// used to allow the groupsClaim to live outside of the top-level JWT structure.
	// For instance, a groupsClaim of meta/user.name/groups with this field
	// set to // will expect nested structures named meta, user.name, and groups.
	// If this field was set to /./ the groups information would expect to be
	// via nested structures of meta, user, name, and groups.
	GroupsClaimDelimiterPattern pulumi.StringPtrInput
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl pulumi.IntPtrInput
	// The amount of leeway to add to not before (`nbf`) claims to account for
	// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
	// Only applicable with "jwt" roles.
	NotBeforeLeeway pulumi.IntPtrInput
	// If set, puts a use-count
	// limitation on the issued token.
	NumUses pulumi.IntPtrInput
	// If set, a list of OIDC scopes to be used with an OIDC role.
	// The standard scope "openid" is automatically included and need not be specified.
	OidcScopes pulumi.StringArrayInput
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period pulumi.IntPtrInput
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies pulumi.StringArrayInput
	// The name of the role.
	RoleName pulumi.StringPtrInput
	// Type of role, either "oidc" (default) or "jwt".
	RoleType pulumi.StringPtrInput
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
	// The claim to use to uniquely identify
	// the user; this will be used as the name for the Identity entity alias created
	// due to a successful login.
	UserClaim pulumi.StringPtrInput
	// Log received OIDC tokens and claims when debug-level
	// logging is active. Not recommended in production since sensitive information may be present
	// in OIDC responses.
	VerboseOidcLogging pulumi.BoolPtrInput
}

func (AuthBackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoleState)(nil)).Elem()
}

type authBackendRoleArgs struct {
	// The list of allowed values for redirectUri during OIDC logins.
	// Required for OIDC roles
	AllowedRedirectUris []string `pulumi:"allowedRedirectUris"`
	// The unique name of the auth backend to configure.
	// Defaults to `jwt`.
	Backend *string `pulumi:"backend"`
	// List of `aud` claims to match
	// against. Any match is sufficient.
	BoundAudiences []string `pulumi:"boundAudiences"`
	// If set, a list of
	// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
	BoundCidrs []string `pulumi:"boundCidrs"`
	// If set, a map of claims/values to match against.
	// The expected value may be a single string or a list of strings.
	BoundClaims map[string]interface{} `pulumi:"boundClaims"`
	// If set, requires that the `sub` claim matches
	// this value.
	BoundSubject *string `pulumi:"boundSubject"`
	// If set, a map of claims (keys) to be copied
	// to specified metadata fields (values).
	ClaimMappings map[string]interface{} `pulumi:"claimMappings"`
	// The amount of leeway to add to all claims to account for clock skew, in
	// seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
	// Only applicable with "jwt" roles.
	ClockSkewLeeway *int `pulumi:"clockSkewLeeway"`
	// The amount of leeway to add to expiration (`exp`) claims to account for
	// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
	// Only applicable with "jwt" roles.
	ExpirationLeeway *int `pulumi:"expirationLeeway"`
	// The claim to use to uniquely identify
	// the set of groups to which the user belongs; this will be used as the names
	// for the Identity group aliases created due to a successful login. The claim
	// value must be a list of strings.
	GroupsClaim *string `pulumi:"groupsClaim"`
	// (Optional; Deprecated. This field has been
	// removed since Vault 1.1. If the groups claim is not at the top level, it can
	// now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
	// A pattern of delimiters
	// used to allow the groupsClaim to live outside of the top-level JWT structure.
	// For instance, a groupsClaim of meta/user.name/groups with this field
	// set to // will expect nested structures named meta, user.name, and groups.
	// If this field was set to /./ the groups information would expect to be
	// via nested structures of meta, user, name, and groups.
	GroupsClaimDelimiterPattern *string `pulumi:"groupsClaimDelimiterPattern"`
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl *int `pulumi:"maxTtl"`
	// The amount of leeway to add to not before (`nbf`) claims to account for
	// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
	// Only applicable with "jwt" roles.
	NotBeforeLeeway *int `pulumi:"notBeforeLeeway"`
	// If set, puts a use-count
	// limitation on the issued token.
	NumUses *int `pulumi:"numUses"`
	// If set, a list of OIDC scopes to be used with an OIDC role.
	// The standard scope "openid" is automatically included and need not be specified.
	OidcScopes []string `pulumi:"oidcScopes"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period *int `pulumi:"period"`
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies []string `pulumi:"policies"`
	// The name of the role.
	RoleName string `pulumi:"roleName"`
	// Type of role, either "oidc" (default) or "jwt".
	RoleType *string `pulumi:"roleType"`
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
	// The claim to use to uniquely identify
	// the user; this will be used as the name for the Identity entity alias created
	// due to a successful login.
	UserClaim string `pulumi:"userClaim"`
	// Log received OIDC tokens and claims when debug-level
	// logging is active. Not recommended in production since sensitive information may be present
	// in OIDC responses.
	VerboseOidcLogging *bool `pulumi:"verboseOidcLogging"`
}

// The set of arguments for constructing a AuthBackendRole resource.
type AuthBackendRoleArgs struct {
	// The list of allowed values for redirectUri during OIDC logins.
	// Required for OIDC roles
	AllowedRedirectUris pulumi.StringArrayInput
	// The unique name of the auth backend to configure.
	// Defaults to `jwt`.
	Backend pulumi.StringPtrInput
	// List of `aud` claims to match
	// against. Any match is sufficient.
	BoundAudiences pulumi.StringArrayInput
	// If set, a list of
	// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
	BoundCidrs pulumi.StringArrayInput
	// If set, a map of claims/values to match against.
	// The expected value may be a single string or a list of strings.
	BoundClaims pulumi.MapInput
	// If set, requires that the `sub` claim matches
	// this value.
	BoundSubject pulumi.StringPtrInput
	// If set, a map of claims (keys) to be copied
	// to specified metadata fields (values).
	ClaimMappings pulumi.MapInput
	// The amount of leeway to add to all claims to account for clock skew, in
	// seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
	// Only applicable with "jwt" roles.
	ClockSkewLeeway pulumi.IntPtrInput
	// The amount of leeway to add to expiration (`exp`) claims to account for
	// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
	// Only applicable with "jwt" roles.
	ExpirationLeeway pulumi.IntPtrInput
	// The claim to use to uniquely identify
	// the set of groups to which the user belongs; this will be used as the names
	// for the Identity group aliases created due to a successful login. The claim
	// value must be a list of strings.
	GroupsClaim pulumi.StringPtrInput
	// (Optional; Deprecated. This field has been
	// removed since Vault 1.1. If the groups claim is not at the top level, it can
	// now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
	// A pattern of delimiters
	// used to allow the groupsClaim to live outside of the top-level JWT structure.
	// For instance, a groupsClaim of meta/user.name/groups with this field
	// set to // will expect nested structures named meta, user.name, and groups.
	// If this field was set to /./ the groups information would expect to be
	// via nested structures of meta, user, name, and groups.
	GroupsClaimDelimiterPattern pulumi.StringPtrInput
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl pulumi.IntPtrInput
	// The amount of leeway to add to not before (`nbf`) claims to account for
	// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
	// Only applicable with "jwt" roles.
	NotBeforeLeeway pulumi.IntPtrInput
	// If set, puts a use-count
	// limitation on the issued token.
	NumUses pulumi.IntPtrInput
	// If set, a list of OIDC scopes to be used with an OIDC role.
	// The standard scope "openid" is automatically included and need not be specified.
	OidcScopes pulumi.StringArrayInput
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period pulumi.IntPtrInput
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies pulumi.StringArrayInput
	// The name of the role.
	RoleName pulumi.StringInput
	// Type of role, either "oidc" (default) or "jwt".
	RoleType pulumi.StringPtrInput
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
	// The claim to use to uniquely identify
	// the user; this will be used as the name for the Identity entity alias created
	// due to a successful login.
	UserClaim pulumi.StringInput
	// Log received OIDC tokens and claims when debug-level
	// logging is active. Not recommended in production since sensitive information may be present
	// in OIDC responses.
	VerboseOidcLogging pulumi.BoolPtrInput
}

func (AuthBackendRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoleArgs)(nil)).Elem()
}
