// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create a role in an [Cert auth backend within Vault](https://www.vaultproject.io/docs/auth/cert.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/cert_auth_backend_role.html.markdown.
type CertAuthBackendRole struct {
	s *pulumi.ResourceState
}

// NewCertAuthBackendRole registers a new resource with the given unique name, arguments, and options.
func NewCertAuthBackendRole(ctx *pulumi.Context,
	name string, args *CertAuthBackendRoleArgs, opts ...pulumi.ResourceOpt) (*CertAuthBackendRole, error) {
	if args == nil || args.Certificate == nil {
		return nil, errors.New("missing required argument 'Certificate'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allowedCommonNames"] = nil
		inputs["allowedDnsSans"] = nil
		inputs["allowedEmailSans"] = nil
		inputs["allowedNames"] = nil
		inputs["allowedOrganizationUnits"] = nil
		inputs["allowedUriSans"] = nil
		inputs["backend"] = nil
		inputs["boundCidrs"] = nil
		inputs["certificate"] = nil
		inputs["displayName"] = nil
		inputs["maxTtl"] = nil
		inputs["name"] = nil
		inputs["period"] = nil
		inputs["policies"] = nil
		inputs["requiredExtensions"] = nil
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
	} else {
		inputs["allowedCommonNames"] = args.AllowedCommonNames
		inputs["allowedDnsSans"] = args.AllowedDnsSans
		inputs["allowedEmailSans"] = args.AllowedEmailSans
		inputs["allowedNames"] = args.AllowedNames
		inputs["allowedOrganizationUnits"] = args.AllowedOrganizationUnits
		inputs["allowedUriSans"] = args.AllowedUriSans
		inputs["backend"] = args.Backend
		inputs["boundCidrs"] = args.BoundCidrs
		inputs["certificate"] = args.Certificate
		inputs["displayName"] = args.DisplayName
		inputs["maxTtl"] = args.MaxTtl
		inputs["name"] = args.Name
		inputs["period"] = args.Period
		inputs["policies"] = args.Policies
		inputs["requiredExtensions"] = args.RequiredExtensions
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
	s, err := ctx.RegisterResource("vault:index/certAuthBackendRole:CertAuthBackendRole", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CertAuthBackendRole{s: s}, nil
}

// GetCertAuthBackendRole gets an existing CertAuthBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertAuthBackendRole(ctx *pulumi.Context,
	name string, id pulumi.ID, state *CertAuthBackendRoleState, opts ...pulumi.ResourceOpt) (*CertAuthBackendRole, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allowedCommonNames"] = state.AllowedCommonNames
		inputs["allowedDnsSans"] = state.AllowedDnsSans
		inputs["allowedEmailSans"] = state.AllowedEmailSans
		inputs["allowedNames"] = state.AllowedNames
		inputs["allowedOrganizationUnits"] = state.AllowedOrganizationUnits
		inputs["allowedUriSans"] = state.AllowedUriSans
		inputs["backend"] = state.Backend
		inputs["boundCidrs"] = state.BoundCidrs
		inputs["certificate"] = state.Certificate
		inputs["displayName"] = state.DisplayName
		inputs["maxTtl"] = state.MaxTtl
		inputs["name"] = state.Name
		inputs["period"] = state.Period
		inputs["policies"] = state.Policies
		inputs["requiredExtensions"] = state.RequiredExtensions
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
	}
	s, err := ctx.ReadResource("vault:index/certAuthBackendRole:CertAuthBackendRole", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CertAuthBackendRole{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *CertAuthBackendRole) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *CertAuthBackendRole) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Allowed the common names for authenticated client certificates
func (r *CertAuthBackendRole) AllowedCommonNames() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["allowedCommonNames"])
}

// Allowed alternative dns names for authenticated client certificates
func (r *CertAuthBackendRole) AllowedDnsSans() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["allowedDnsSans"])
}

// Allowed emails for authenticated client certificates
func (r *CertAuthBackendRole) AllowedEmailSans() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["allowedEmailSans"])
}

// Allowed subject names for authenticated client certificates
func (r *CertAuthBackendRole) AllowedNames() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["allowedNames"])
}

// Allowed organization units for authenticated client certificates
func (r *CertAuthBackendRole) AllowedOrganizationUnits() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["allowedOrganizationUnits"])
}

// Allowed URIs for authenticated client certificates
func (r *CertAuthBackendRole) AllowedUriSans() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["allowedUriSans"])
}

// Path to the mounted Cert auth backend
func (r *CertAuthBackendRole) Backend() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["backend"])
}

// Restriction usage of the
// certificates to client IPs falling within the range of the specified CIDRs
func (r *CertAuthBackendRole) BoundCidrs() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["boundCidrs"])
}

// CA certificate used to validate client certificates
func (r *CertAuthBackendRole) Certificate() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["certificate"])
}

// The name to display on tokens issued under this role.
func (r *CertAuthBackendRole) DisplayName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["displayName"])
}

// The maximum allowed lifetime of tokens
// issued using this role, provided as a number of seconds.
func (r *CertAuthBackendRole) MaxTtl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["maxTtl"])
}

// Name of the role
func (r *CertAuthBackendRole) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// If set, indicates that the
// token generated using this role should never expire. The token should be renewed within the
// duration specified by this value. At each renewal, the token's TTL will be set to the
// value of this field. The maximum allowed lifetime of token issued using this
// role. Specified as a number of seconds.
func (r *CertAuthBackendRole) Period() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["period"])
}

// An array of strings
// specifying the policies to be set on tokens issued using this role.
func (r *CertAuthBackendRole) Policies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["policies"])
}

// TLS extensions required on client certificates
func (r *CertAuthBackendRole) RequiredExtensions() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["requiredExtensions"])
}

// List of CIDR blocks; if set, specifies blocks of IP
// addresses which can authenticate successfully, and ties the resulting token to these blocks
// as well.
func (r *CertAuthBackendRole) TokenBoundCidrs() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["tokenBoundCidrs"])
}

// If set, will encode an
// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
// `tokenMaxTtl` would otherwise allow a renewal.
func (r *CertAuthBackendRole) TokenExplicitMaxTtl() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["tokenExplicitMaxTtl"])
}

// The maximum lifetime for generated tokens in number of seconds.
// Its current value will be referenced at renewal time.
func (r *CertAuthBackendRole) TokenMaxTtl() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["tokenMaxTtl"])
}

// If set, the default policy will not be set on
// generated tokens; otherwise it will be added to the policies set in token_policies.
func (r *CertAuthBackendRole) TokenNoDefaultPolicy() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["tokenNoDefaultPolicy"])
}

// The
// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
// if any, in number of seconds to set on the token.
func (r *CertAuthBackendRole) TokenNumUses() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["tokenNumUses"])
}

// Generated Token's Period
func (r *CertAuthBackendRole) TokenPeriod() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["tokenPeriod"])
}

// List of policies to encode onto generated tokens. Depending
// on the auth method, this list may be supplemented by user/group/other values.
func (r *CertAuthBackendRole) TokenPolicies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["tokenPolicies"])
}

// The incremental lifetime for generated tokens in number of seconds.
// Its current value will be referenced at renewal time.
func (r *CertAuthBackendRole) TokenTtl() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["tokenTtl"])
}

// The type of token that should be generated. Can be `service`,
// `batch`, or `default` to use the mount's tuned default (which unless changed will be
// `service` tokens). For token store roles, there are two additional possibilities:
// `default-service` and `default-batch` which specify the type to return unless the client
// requests a different type at generation time.
func (r *CertAuthBackendRole) TokenType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["tokenType"])
}

// The TTL period of tokens issued
// using this role, provided as a number of seconds.
func (r *CertAuthBackendRole) Ttl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ttl"])
}

// Input properties used for looking up and filtering CertAuthBackendRole resources.
type CertAuthBackendRoleState struct {
	// Allowed the common names for authenticated client certificates
	AllowedCommonNames interface{}
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans interface{}
	// Allowed emails for authenticated client certificates
	AllowedEmailSans interface{}
	// Allowed subject names for authenticated client certificates
	AllowedNames interface{}
	// Allowed organization units for authenticated client certificates
	AllowedOrganizationUnits interface{}
	// Allowed URIs for authenticated client certificates
	AllowedUriSans interface{}
	// Path to the mounted Cert auth backend
	Backend interface{}
	// Restriction usage of the
	// certificates to client IPs falling within the range of the specified CIDRs
	BoundCidrs interface{}
	// CA certificate used to validate client certificates
	Certificate interface{}
	// The name to display on tokens issued under this role.
	DisplayName interface{}
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl interface{}
	// Name of the role
	Name interface{}
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. The maximum allowed lifetime of token issued using this
	// role. Specified as a number of seconds.
	Period interface{}
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies interface{}
	// TLS extensions required on client certificates
	RequiredExtensions interface{}
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
}

// The set of arguments for constructing a CertAuthBackendRole resource.
type CertAuthBackendRoleArgs struct {
	// Allowed the common names for authenticated client certificates
	AllowedCommonNames interface{}
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans interface{}
	// Allowed emails for authenticated client certificates
	AllowedEmailSans interface{}
	// Allowed subject names for authenticated client certificates
	AllowedNames interface{}
	// Allowed organization units for authenticated client certificates
	AllowedOrganizationUnits interface{}
	// Allowed URIs for authenticated client certificates
	AllowedUriSans interface{}
	// Path to the mounted Cert auth backend
	Backend interface{}
	// Restriction usage of the
	// certificates to client IPs falling within the range of the specified CIDRs
	BoundCidrs interface{}
	// CA certificate used to validate client certificates
	Certificate interface{}
	// The name to display on tokens issued under this role.
	DisplayName interface{}
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl interface{}
	// Name of the role
	Name interface{}
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. The maximum allowed lifetime of token issued using this
	// role. Specified as a number of seconds.
	Period interface{}
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies interface{}
	// TLS extensions required on client certificates
	RequiredExtensions interface{}
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
}
