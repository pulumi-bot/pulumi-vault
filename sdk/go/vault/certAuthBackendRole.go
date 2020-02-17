// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vault

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create a role in an [Cert auth backend within Vault](https://www.vaultproject.io/docs/auth/cert.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/cert_auth_backend_role.html.markdown.
type CertAuthBackendRole struct {
	pulumi.CustomResourceState

	// Allowed the common names for authenticated client certificates
	AllowedCommonNames pulumi.StringArrayOutput `pulumi:"allowedCommonNames"`
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans pulumi.StringArrayOutput `pulumi:"allowedDnsSans"`
	// Allowed emails for authenticated client certificates
	AllowedEmailSans pulumi.StringArrayOutput `pulumi:"allowedEmailSans"`
	// Allowed subject names for authenticated client certificates
	AllowedNames pulumi.StringArrayOutput `pulumi:"allowedNames"`
	// Allowed organization units for authenticated client certificates
	AllowedOrganizationUnits pulumi.StringArrayOutput `pulumi:"allowedOrganizationUnits"`
	// Allowed URIs for authenticated client certificates
	AllowedUriSans pulumi.StringArrayOutput `pulumi:"allowedUriSans"`
	// Path to the mounted Cert auth backend
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// Restriction usage of the
	// certificates to client IPs falling within the range of the specified CIDRs
	BoundCidrs pulumi.StringArrayOutput `pulumi:"boundCidrs"`
	// CA certificate used to validate client certificates
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// The name to display on tokens issued under this role.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl pulumi.StringOutput `pulumi:"maxTtl"`
	// Name of the role
	Name pulumi.StringOutput `pulumi:"name"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period pulumi.StringOutput `pulumi:"period"`
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// TLS extensions required on client certificates
	RequiredExtensions pulumi.StringArrayOutput `pulumi:"requiredExtensions"`
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
	Ttl pulumi.StringOutput `pulumi:"ttl"`
}

// NewCertAuthBackendRole registers a new resource with the given unique name, arguments, and options.
func NewCertAuthBackendRole(ctx *pulumi.Context,
	name string, args *CertAuthBackendRoleArgs, opts ...pulumi.ResourceOption) (*CertAuthBackendRole, error) {
	if args == nil || args.Certificate == nil {
		return nil, errors.New("missing required argument 'Certificate'")
	}
	if args == nil {
		args = &CertAuthBackendRoleArgs{}
	}
	var resource CertAuthBackendRole
	err := ctx.RegisterResource("vault:index/certAuthBackendRole:CertAuthBackendRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertAuthBackendRole gets an existing CertAuthBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertAuthBackendRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertAuthBackendRoleState, opts ...pulumi.ResourceOption) (*CertAuthBackendRole, error) {
	var resource CertAuthBackendRole
	err := ctx.ReadResource("vault:index/certAuthBackendRole:CertAuthBackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertAuthBackendRole resources.
type certAuthBackendRoleState struct {
	// Allowed the common names for authenticated client certificates
	AllowedCommonNames []string `pulumi:"allowedCommonNames"`
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans []string `pulumi:"allowedDnsSans"`
	// Allowed emails for authenticated client certificates
	AllowedEmailSans []string `pulumi:"allowedEmailSans"`
	// Allowed subject names for authenticated client certificates
	AllowedNames []string `pulumi:"allowedNames"`
	// Allowed organization units for authenticated client certificates
	AllowedOrganizationUnits []string `pulumi:"allowedOrganizationUnits"`
	// Allowed URIs for authenticated client certificates
	AllowedUriSans []string `pulumi:"allowedUriSans"`
	// Path to the mounted Cert auth backend
	Backend *string `pulumi:"backend"`
	// Restriction usage of the
	// certificates to client IPs falling within the range of the specified CIDRs
	BoundCidrs []string `pulumi:"boundCidrs"`
	// CA certificate used to validate client certificates
	Certificate *string `pulumi:"certificate"`
	// The name to display on tokens issued under this role.
	DisplayName *string `pulumi:"displayName"`
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl *string `pulumi:"maxTtl"`
	// Name of the role
	Name *string `pulumi:"name"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period *string `pulumi:"period"`
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies []string `pulumi:"policies"`
	// TLS extensions required on client certificates
	RequiredExtensions []string `pulumi:"requiredExtensions"`
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
	Ttl *string `pulumi:"ttl"`
}

type CertAuthBackendRoleState struct {
	// Allowed the common names for authenticated client certificates
	AllowedCommonNames pulumi.StringArrayInput
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans pulumi.StringArrayInput
	// Allowed emails for authenticated client certificates
	AllowedEmailSans pulumi.StringArrayInput
	// Allowed subject names for authenticated client certificates
	AllowedNames pulumi.StringArrayInput
	// Allowed organization units for authenticated client certificates
	AllowedOrganizationUnits pulumi.StringArrayInput
	// Allowed URIs for authenticated client certificates
	AllowedUriSans pulumi.StringArrayInput
	// Path to the mounted Cert auth backend
	Backend pulumi.StringPtrInput
	// Restriction usage of the
	// certificates to client IPs falling within the range of the specified CIDRs
	BoundCidrs pulumi.StringArrayInput
	// CA certificate used to validate client certificates
	Certificate pulumi.StringPtrInput
	// The name to display on tokens issued under this role.
	DisplayName pulumi.StringPtrInput
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl pulumi.StringPtrInput
	// Name of the role
	Name pulumi.StringPtrInput
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period pulumi.StringPtrInput
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies pulumi.StringArrayInput
	// TLS extensions required on client certificates
	RequiredExtensions pulumi.StringArrayInput
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
	Ttl pulumi.StringPtrInput
}

func (CertAuthBackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*certAuthBackendRoleState)(nil)).Elem()
}

type certAuthBackendRoleArgs struct {
	// Allowed the common names for authenticated client certificates
	AllowedCommonNames []string `pulumi:"allowedCommonNames"`
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans []string `pulumi:"allowedDnsSans"`
	// Allowed emails for authenticated client certificates
	AllowedEmailSans []string `pulumi:"allowedEmailSans"`
	// Allowed subject names for authenticated client certificates
	AllowedNames []string `pulumi:"allowedNames"`
	// Allowed organization units for authenticated client certificates
	AllowedOrganizationUnits []string `pulumi:"allowedOrganizationUnits"`
	// Allowed URIs for authenticated client certificates
	AllowedUriSans []string `pulumi:"allowedUriSans"`
	// Path to the mounted Cert auth backend
	Backend *string `pulumi:"backend"`
	// Restriction usage of the
	// certificates to client IPs falling within the range of the specified CIDRs
	BoundCidrs []string `pulumi:"boundCidrs"`
	// CA certificate used to validate client certificates
	Certificate string `pulumi:"certificate"`
	// The name to display on tokens issued under this role.
	DisplayName *string `pulumi:"displayName"`
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl *string `pulumi:"maxTtl"`
	// Name of the role
	Name *string `pulumi:"name"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period *string `pulumi:"period"`
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies []string `pulumi:"policies"`
	// TLS extensions required on client certificates
	RequiredExtensions []string `pulumi:"requiredExtensions"`
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
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a CertAuthBackendRole resource.
type CertAuthBackendRoleArgs struct {
	// Allowed the common names for authenticated client certificates
	AllowedCommonNames pulumi.StringArrayInput
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans pulumi.StringArrayInput
	// Allowed emails for authenticated client certificates
	AllowedEmailSans pulumi.StringArrayInput
	// Allowed subject names for authenticated client certificates
	AllowedNames pulumi.StringArrayInput
	// Allowed organization units for authenticated client certificates
	AllowedOrganizationUnits pulumi.StringArrayInput
	// Allowed URIs for authenticated client certificates
	AllowedUriSans pulumi.StringArrayInput
	// Path to the mounted Cert auth backend
	Backend pulumi.StringPtrInput
	// Restriction usage of the
	// certificates to client IPs falling within the range of the specified CIDRs
	BoundCidrs pulumi.StringArrayInput
	// CA certificate used to validate client certificates
	Certificate pulumi.StringInput
	// The name to display on tokens issued under this role.
	DisplayName pulumi.StringPtrInput
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl pulumi.StringPtrInput
	// Name of the role
	Name pulumi.StringPtrInput
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period pulumi.StringPtrInput
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies pulumi.StringArrayInput
	// TLS extensions required on client certificates
	RequiredExtensions pulumi.StringArrayInput
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
	Ttl pulumi.StringPtrInput
}

func (CertAuthBackendRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certAuthBackendRoleArgs)(nil)).Elem()
}

