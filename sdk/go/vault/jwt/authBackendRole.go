// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package jwt

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an JWT/OIDC auth backend role in a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/jwt.html) for more
// information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/jwt_auth_backend_role.html.markdown.
type AuthBackendRole struct {
	s *pulumi.ResourceState
}

// NewAuthBackendRole registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendRole(ctx *pulumi.Context,
	name string, args *AuthBackendRoleArgs, opts ...pulumi.ResourceOpt) (*AuthBackendRole, error) {
	if args == nil || args.BoundAudiences == nil {
		return nil, errors.New("missing required argument 'BoundAudiences'")
	}
	if args == nil || args.RoleName == nil {
		return nil, errors.New("missing required argument 'RoleName'")
	}
	if args == nil || args.UserClaim == nil {
		return nil, errors.New("missing required argument 'UserClaim'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allowedRedirectUris"] = nil
		inputs["backend"] = nil
		inputs["boundAudiences"] = nil
		inputs["boundCidrs"] = nil
		inputs["boundClaims"] = nil
		inputs["boundSubject"] = nil
		inputs["claimMappings"] = nil
		inputs["groupsClaim"] = nil
		inputs["groupsClaimDelimiterPattern"] = nil
		inputs["maxTtl"] = nil
		inputs["numUses"] = nil
		inputs["oidcScopes"] = nil
		inputs["period"] = nil
		inputs["policies"] = nil
		inputs["roleName"] = nil
		inputs["roleType"] = nil
		inputs["tokenBoundCidrs"] = nil
		inputs["tokenExplicitMaxTtl"] = nil
		inputs["tokenMaxTtl"] = nil
		inputs["tokenNoDefaultPolicy"] = nil
		inputs["tokenNumUses"] = nil
		inputs["tokenPeriod"] = nil
		inputs["tokenPolicies"] = nil
		inputs["tokenTtl"] = nil
		inputs["tokenType"] = nil
		inputs["ttl"] = nil
		inputs["userClaim"] = nil
	} else {
		inputs["allowedRedirectUris"] = args.AllowedRedirectUris
		inputs["backend"] = args.Backend
		inputs["boundAudiences"] = args.BoundAudiences
		inputs["boundCidrs"] = args.BoundCidrs
		inputs["boundClaims"] = args.BoundClaims
		inputs["boundSubject"] = args.BoundSubject
		inputs["claimMappings"] = args.ClaimMappings
		inputs["groupsClaim"] = args.GroupsClaim
		inputs["groupsClaimDelimiterPattern"] = args.GroupsClaimDelimiterPattern
		inputs["maxTtl"] = args.MaxTtl
		inputs["numUses"] = args.NumUses
		inputs["oidcScopes"] = args.OidcScopes
		inputs["period"] = args.Period
		inputs["policies"] = args.Policies
		inputs["roleName"] = args.RoleName
		inputs["roleType"] = args.RoleType
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
		inputs["userClaim"] = args.UserClaim
	}
	s, err := ctx.RegisterResource("vault:jwt/authBackendRole:AuthBackendRole", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackendRole{s: s}, nil
}

// GetAuthBackendRole gets an existing AuthBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendRole(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AuthBackendRoleState, opts ...pulumi.ResourceOpt) (*AuthBackendRole, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allowedRedirectUris"] = state.AllowedRedirectUris
		inputs["backend"] = state.Backend
		inputs["boundAudiences"] = state.BoundAudiences
		inputs["boundCidrs"] = state.BoundCidrs
		inputs["boundClaims"] = state.BoundClaims
		inputs["boundSubject"] = state.BoundSubject
		inputs["claimMappings"] = state.ClaimMappings
		inputs["groupsClaim"] = state.GroupsClaim
		inputs["groupsClaimDelimiterPattern"] = state.GroupsClaimDelimiterPattern
		inputs["maxTtl"] = state.MaxTtl
		inputs["numUses"] = state.NumUses
		inputs["oidcScopes"] = state.OidcScopes
		inputs["period"] = state.Period
		inputs["policies"] = state.Policies
		inputs["roleName"] = state.RoleName
		inputs["roleType"] = state.RoleType
		inputs["tokenBoundCidrs"] = state.TokenBoundCidrs
		inputs["tokenExplicitMaxTtl"] = state.TokenExplicitMaxTtl
		inputs["tokenMaxTtl"] = state.TokenMaxTtl
		inputs["tokenNoDefaultPolicy"] = state.TokenNoDefaultPolicy
		inputs["tokenNumUses"] = state.TokenNumUses
		inputs["tokenPeriod"] = state.TokenPeriod
		inputs["tokenPolicies"] = state.TokenPolicies
		inputs["tokenTtl"] = state.TokenTtl
		inputs["tokenType"] = state.TokenType
		inputs["ttl"] = state.Ttl
		inputs["userClaim"] = state.UserClaim
	}
	s, err := ctx.ReadResource("vault:jwt/authBackendRole:AuthBackendRole", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackendRole{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AuthBackendRole) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AuthBackendRole) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The list of allowed values for redirectUri during OIDC logins.
// Required for OIDC roles
func (r *AuthBackendRole) AllowedRedirectUris() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["allowedRedirectUris"])
}

// The unique name of the auth backend to configure.
// Defaults to `jwt`.
func (r *AuthBackendRole) Backend() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["backend"])
}

// List of `aud` claims to match
// against. Any match is sufficient.
func (r *AuthBackendRole) BoundAudiences() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["boundAudiences"])
}

// If set, a list of
// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
func (r *AuthBackendRole) BoundCidrs() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["boundCidrs"])
}

// If set, a map of claims/values to match against.
// The expected value may be a single string or a list of strings.
func (r *AuthBackendRole) BoundClaims() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["boundClaims"])
}

// If set, requires that the `sub` claim matches
// this value.
func (r *AuthBackendRole) BoundSubject() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["boundSubject"])
}

// If set, a map of claims (keys) to be copied
// to specified metadata fields (values).
func (r *AuthBackendRole) ClaimMappings() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["claimMappings"])
}

// The claim to use to uniquely identify
// the set of groups to which the user belongs; this will be used as the names
// for the Identity group aliases created due to a successful login. The claim
// value must be a list of strings.
func (r *AuthBackendRole) GroupsClaim() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["groupsClaim"])
}

// (Optional; Deprecated. This field has been
// removed since Vault 1.1. If the groups claim is not at the top level, it can
// now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
// A pattern of delimiters
// used to allow the groupsClaim to live outside of the top-level JWT structure.
// For instance, a groupsClaim of meta/user.name/groups with this field
// set to // will expect nested structures named meta, user.name, and groups.
// If this field was set to /./ the groups information would expect to be
// via nested structures of meta, user, name, and groups.
func (r *AuthBackendRole) GroupsClaimDelimiterPattern() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["groupsClaimDelimiterPattern"])
}

// The maximum allowed lifetime of tokens
// issued using this role, provided as a number of seconds.
func (r *AuthBackendRole) MaxTtl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["maxTtl"])
}

// If set, puts a use-count
// limitation on the issued token.
func (r *AuthBackendRole) NumUses() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["numUses"])
}

// If set, a list of OIDC scopes to be used with an OIDC role.
// The standard scope "openid" is automatically included and need not be specified.
func (r *AuthBackendRole) OidcScopes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["oidcScopes"])
}

// If set, indicates that the
// token generated using this role should never expire. The token should be renewed within the
// duration specified by this value. At each renewal, the token's TTL will be set to the
// value of this field. The maximum allowed lifetime of token issued using this
// role. Specified as a number of seconds.
func (r *AuthBackendRole) Period() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["period"])
}

// An array of strings
// specifying the policies to be set on tokens issued using this role.
func (r *AuthBackendRole) Policies() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["policies"])
}

// The name of the role.
func (r *AuthBackendRole) RoleName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["roleName"])
}

// Type of role, either "oidc" (default) or "jwt".
func (r *AuthBackendRole) RoleType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["roleType"])
}

// List of CIDR blocks; if set, specifies blocks of IP
// addresses which can authenticate successfully, and ties the resulting token to these blocks
// as well.
func (r *AuthBackendRole) TokenBoundCidrs() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["tokenBoundCidrs"])
}

// If set, will encode an
// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
// `tokenMaxTtl` would otherwise allow a renewal.
func (r *AuthBackendRole) TokenExplicitMaxTtl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["tokenExplicitMaxTtl"])
}

// The maximum lifetime for generated tokens in number of seconds.
// Its current value will be referenced at renewal time.
func (r *AuthBackendRole) TokenMaxTtl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["tokenMaxTtl"])
}

// If set, the default policy will not be set on
// generated tokens; otherwise it will be added to the policies set in token_policies.
func (r *AuthBackendRole) TokenNoDefaultPolicy() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["tokenNoDefaultPolicy"])
}

// The
// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
// if any, in number of seconds to set on the token.
func (r *AuthBackendRole) TokenNumUses() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["tokenNumUses"])
}

// Generated Token's Period
func (r *AuthBackendRole) TokenPeriod() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["tokenPeriod"])
}

// List of policies to encode onto generated tokens. Depending
// on the auth method, this list may be supplemented by user/group/other values.
func (r *AuthBackendRole) TokenPolicies() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["tokenPolicies"])
}

// The incremental lifetime for generated tokens in number of seconds.
// Its current value will be referenced at renewal time.
func (r *AuthBackendRole) TokenTtl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["tokenTtl"])
}

// The type of token that should be generated. Can be `service`,
// `batch`, or `default` to use the mount's tuned default (which unless changed will be
// `service` tokens). For token store roles, there are two additional possibilities:
// `default-service` and `default-batch` which specify the type to return unless the client
// requests a different type at generation time.
func (r *AuthBackendRole) TokenType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tokenType"])
}

// The TTL period of tokens issued
// using this role, provided as a number of seconds.
func (r *AuthBackendRole) Ttl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ttl"])
}

// The claim to use to uniquely identify
// the user; this will be used as the name for the Identity entity alias created
// due to a successful login.
func (r *AuthBackendRole) UserClaim() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userClaim"])
}

// Input properties used for looking up and filtering AuthBackendRole resources.
type AuthBackendRoleState struct {
	// The list of allowed values for redirectUri during OIDC logins.
	// Required for OIDC roles
	AllowedRedirectUris interface{}
	// The unique name of the auth backend to configure.
	// Defaults to `jwt`.
	Backend interface{}
	// List of `aud` claims to match
	// against. Any match is sufficient.
	BoundAudiences interface{}
	// If set, a list of
	// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
	BoundCidrs interface{}
	// If set, a map of claims/values to match against.
	// The expected value may be a single string or a list of strings.
	BoundClaims interface{}
	// If set, requires that the `sub` claim matches
	// this value.
	BoundSubject interface{}
	// If set, a map of claims (keys) to be copied
	// to specified metadata fields (values).
	ClaimMappings interface{}
	// The claim to use to uniquely identify
	// the set of groups to which the user belongs; this will be used as the names
	// for the Identity group aliases created due to a successful login. The claim
	// value must be a list of strings.
	GroupsClaim interface{}
	// (Optional; Deprecated. This field has been
	// removed since Vault 1.1. If the groups claim is not at the top level, it can
	// now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
	// A pattern of delimiters
	// used to allow the groupsClaim to live outside of the top-level JWT structure.
	// For instance, a groupsClaim of meta/user.name/groups with this field
	// set to // will expect nested structures named meta, user.name, and groups.
	// If this field was set to /./ the groups information would expect to be
	// via nested structures of meta, user, name, and groups.
	GroupsClaimDelimiterPattern interface{}
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl interface{}
	// If set, puts a use-count
	// limitation on the issued token.
	NumUses interface{}
	// If set, a list of OIDC scopes to be used with an OIDC role.
	// The standard scope "openid" is automatically included and need not be specified.
	OidcScopes interface{}
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. The maximum allowed lifetime of token issued using this
	// role. Specified as a number of seconds.
	Period interface{}
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies interface{}
	// The name of the role.
	RoleName interface{}
	// Type of role, either "oidc" (default) or "jwt".
	RoleType interface{}
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
	// Generated Token's Period
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
	// The TTL period of tokens issued
	// using this role, provided as a number of seconds.
	Ttl interface{}
	// The claim to use to uniquely identify
	// the user; this will be used as the name for the Identity entity alias created
	// due to a successful login.
	UserClaim interface{}
}

// The set of arguments for constructing a AuthBackendRole resource.
type AuthBackendRoleArgs struct {
	// The list of allowed values for redirectUri during OIDC logins.
	// Required for OIDC roles
	AllowedRedirectUris interface{}
	// The unique name of the auth backend to configure.
	// Defaults to `jwt`.
	Backend interface{}
	// List of `aud` claims to match
	// against. Any match is sufficient.
	BoundAudiences interface{}
	// If set, a list of
	// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
	BoundCidrs interface{}
	// If set, a map of claims/values to match against.
	// The expected value may be a single string or a list of strings.
	BoundClaims interface{}
	// If set, requires that the `sub` claim matches
	// this value.
	BoundSubject interface{}
	// If set, a map of claims (keys) to be copied
	// to specified metadata fields (values).
	ClaimMappings interface{}
	// The claim to use to uniquely identify
	// the set of groups to which the user belongs; this will be used as the names
	// for the Identity group aliases created due to a successful login. The claim
	// value must be a list of strings.
	GroupsClaim interface{}
	// (Optional; Deprecated. This field has been
	// removed since Vault 1.1. If the groups claim is not at the top level, it can
	// now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
	// A pattern of delimiters
	// used to allow the groupsClaim to live outside of the top-level JWT structure.
	// For instance, a groupsClaim of meta/user.name/groups with this field
	// set to // will expect nested structures named meta, user.name, and groups.
	// If this field was set to /./ the groups information would expect to be
	// via nested structures of meta, user, name, and groups.
	GroupsClaimDelimiterPattern interface{}
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl interface{}
	// If set, puts a use-count
	// limitation on the issued token.
	NumUses interface{}
	// If set, a list of OIDC scopes to be used with an OIDC role.
	// The standard scope "openid" is automatically included and need not be specified.
	OidcScopes interface{}
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. The maximum allowed lifetime of token issued using this
	// role. Specified as a number of seconds.
	Period interface{}
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies interface{}
	// The name of the role.
	RoleName interface{}
	// Type of role, either "oidc" (default) or "jwt".
	RoleType interface{}
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
	// Generated Token's Period
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
	// The TTL period of tokens issued
	// using this role, provided as a number of seconds.
	Ttl interface{}
	// The claim to use to uniquely identify
	// the user; this will be used as the name for the Identity entity alias created
	// due to a successful login.
	UserClaim interface{}
}
