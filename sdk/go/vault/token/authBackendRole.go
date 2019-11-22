// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package token

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages Token auth backend role in a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/token.html) for more
// information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/token_auth_backend_role.html.markdown.
type AuthBackendRole struct {
	s *pulumi.ResourceState
}

// NewAuthBackendRole registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendRole(ctx *pulumi.Context,
	name string, args *AuthBackendRoleArgs, opts ...pulumi.ResourceOpt) (*AuthBackendRole, error) {
	if args == nil || args.RoleName == nil {
		return nil, errors.New("missing required argument 'RoleName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allowedPolicies"] = nil
		inputs["boundCidrs"] = nil
		inputs["disallowedPolicies"] = nil
		inputs["explicitMaxTtl"] = nil
		inputs["orphan"] = nil
		inputs["pathSuffix"] = nil
		inputs["period"] = nil
		inputs["renewable"] = nil
		inputs["roleName"] = nil
		inputs["tokenBoundCidrs"] = nil
		inputs["tokenExplicitMaxTtl"] = nil
		inputs["tokenMaxTtl"] = nil
		inputs["tokenNoDefaultPolicy"] = nil
		inputs["tokenNumUses"] = nil
		inputs["tokenPeriod"] = nil
		inputs["tokenPolicies"] = nil
		inputs["tokenTtl"] = nil
		inputs["tokenType"] = nil
	} else {
		inputs["allowedPolicies"] = args.AllowedPolicies
		inputs["boundCidrs"] = args.BoundCidrs
		inputs["disallowedPolicies"] = args.DisallowedPolicies
		inputs["explicitMaxTtl"] = args.ExplicitMaxTtl
		inputs["orphan"] = args.Orphan
		inputs["pathSuffix"] = args.PathSuffix
		inputs["period"] = args.Period
		inputs["renewable"] = args.Renewable
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
	}
	s, err := ctx.RegisterResource("vault:token/authBackendRole:AuthBackendRole", name, true, inputs, opts...)
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
		inputs["allowedPolicies"] = state.AllowedPolicies
		inputs["boundCidrs"] = state.BoundCidrs
		inputs["disallowedPolicies"] = state.DisallowedPolicies
		inputs["explicitMaxTtl"] = state.ExplicitMaxTtl
		inputs["orphan"] = state.Orphan
		inputs["pathSuffix"] = state.PathSuffix
		inputs["period"] = state.Period
		inputs["renewable"] = state.Renewable
		inputs["roleName"] = state.RoleName
		inputs["tokenBoundCidrs"] = state.TokenBoundCidrs
		inputs["tokenExplicitMaxTtl"] = state.TokenExplicitMaxTtl
		inputs["tokenMaxTtl"] = state.TokenMaxTtl
		inputs["tokenNoDefaultPolicy"] = state.TokenNoDefaultPolicy
		inputs["tokenNumUses"] = state.TokenNumUses
		inputs["tokenPeriod"] = state.TokenPeriod
		inputs["tokenPolicies"] = state.TokenPolicies
		inputs["tokenTtl"] = state.TokenTtl
		inputs["tokenType"] = state.TokenType
	}
	s, err := ctx.ReadResource("vault:token/authBackendRole:AuthBackendRole", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackendRole{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AuthBackendRole) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AuthBackendRole) ID() pulumi.IDOutput {
	return r.s.ID()
}

// List of allowed policies for given role.
func (r *AuthBackendRole) AllowedPolicies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["allowedPolicies"])
}

// If set, a list of
// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
func (r *AuthBackendRole) BoundCidrs() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["boundCidrs"])
}

// List of disallowed policies for given role.
func (r *AuthBackendRole) DisallowedPolicies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["disallowedPolicies"])
}

// Number of seconds after which issued tokens can no longer be renewed.
func (r *AuthBackendRole) ExplicitMaxTtl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["explicitMaxTtl"])
}

// If true, tokens created against this policy will be orphan tokens.
func (r *AuthBackendRole) Orphan() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["orphan"])
}

// Tokens created against this role will have the given suffix as part of their path in addition to the role name.
func (r *AuthBackendRole) PathSuffix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["pathSuffix"])
}

// If set, indicates that the
// token generated using this role should never expire. The token should be renewed within the
// duration specified by this value. At each renewal, the token's TTL will be set to the
// value of this field. The maximum allowed lifetime of token issued using this
// role. Specified as a number of seconds.
func (r *AuthBackendRole) Period() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["period"])
}

// Wether to disable the ability of the token to be renewed past its initial TTL.
func (r *AuthBackendRole) Renewable() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["renewable"])
}

// The name of the role.
func (r *AuthBackendRole) RoleName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["roleName"])
}

// List of CIDR blocks; if set, specifies blocks of IP
// addresses which can authenticate successfully, and ties the resulting token to these blocks
// as well.
func (r *AuthBackendRole) TokenBoundCidrs() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["tokenBoundCidrs"])
}

// If set, will encode an
// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
// `tokenMaxTtl` would otherwise allow a renewal.
func (r *AuthBackendRole) TokenExplicitMaxTtl() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["tokenExplicitMaxTtl"])
}

// The maximum lifetime for generated tokens in number of seconds.
// Its current value will be referenced at renewal time.
func (r *AuthBackendRole) TokenMaxTtl() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["tokenMaxTtl"])
}

// If set, the default policy will not be set on
// generated tokens; otherwise it will be added to the policies set in token_policies.
func (r *AuthBackendRole) TokenNoDefaultPolicy() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["tokenNoDefaultPolicy"])
}

// The
// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
// if any, in number of seconds to set on the token.
func (r *AuthBackendRole) TokenNumUses() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["tokenNumUses"])
}

// Generated Token's Period
func (r *AuthBackendRole) TokenPeriod() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["tokenPeriod"])
}

// Generated Token's Policies
func (r *AuthBackendRole) TokenPolicies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["tokenPolicies"])
}

// The incremental lifetime for generated tokens in number of seconds.
// Its current value will be referenced at renewal time.
func (r *AuthBackendRole) TokenTtl() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["tokenTtl"])
}

// The type of token that should be generated. Can be `service`,
// `batch`, or `default` to use the mount's tuned default (which unless changed will be
// `service` tokens). For token store roles, there are two additional possibilities:
// `default-service` and `default-batch` which specify the type to return unless the client
// requests a different type at generation time.
func (r *AuthBackendRole) TokenType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["tokenType"])
}

// Input properties used for looking up and filtering AuthBackendRole resources.
type AuthBackendRoleState struct {
	// List of allowed policies for given role.
	AllowedPolicies interface{}
	// If set, a list of
	// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
	BoundCidrs interface{}
	// List of disallowed policies for given role.
	DisallowedPolicies interface{}
	// Number of seconds after which issued tokens can no longer be renewed.
	ExplicitMaxTtl interface{}
	// If true, tokens created against this policy will be orphan tokens.
	Orphan interface{}
	// Tokens created against this role will have the given suffix as part of their path in addition to the role name.
	PathSuffix interface{}
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. The maximum allowed lifetime of token issued using this
	// role. Specified as a number of seconds.
	Period interface{}
	// Wether to disable the ability of the token to be renewed past its initial TTL.
	Renewable interface{}
	// The name of the role.
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
	// Generated Token's Period
	TokenPeriod interface{}
	// Generated Token's Policies
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
}

// The set of arguments for constructing a AuthBackendRole resource.
type AuthBackendRoleArgs struct {
	// List of allowed policies for given role.
	AllowedPolicies interface{}
	// If set, a list of
	// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
	BoundCidrs interface{}
	// List of disallowed policies for given role.
	DisallowedPolicies interface{}
	// Number of seconds after which issued tokens can no longer be renewed.
	ExplicitMaxTtl interface{}
	// If true, tokens created against this policy will be orphan tokens.
	Orphan interface{}
	// Tokens created against this role will have the given suffix as part of their path in addition to the role name.
	PathSuffix interface{}
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. The maximum allowed lifetime of token issued using this
	// role. Specified as a number of seconds.
	Period interface{}
	// Wether to disable the ability of the token to be renewed past its initial TTL.
	Renewable interface{}
	// The name of the role.
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
	// Generated Token's Period
	TokenPeriod interface{}
	// Generated Token's Policies
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
}
