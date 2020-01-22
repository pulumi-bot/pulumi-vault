// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package pkiSecret

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an PKI certificate.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_root_sign_intermediate.html.markdown.
type SecretBackendRootSignIntermediate struct {
	pulumi.CustomResourceState

	// List of alternative names
	AltNames pulumi.StringArrayOutput `pulumi:"altNames"`
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// The CA chain
	CaChain pulumi.StringOutput `pulumi:"caChain"`
	// The certificate
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// CN of intermediate to create
	CommonName pulumi.StringOutput `pulumi:"commonName"`
	// The country
	Country pulumi.StringPtrOutput `pulumi:"country"`
	// The CSR
	Csr pulumi.StringOutput `pulumi:"csr"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrOutput `pulumi:"excludeCnFromSans"`
	// The format of data
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// List of alternative IPs
	IpSans pulumi.StringArrayOutput `pulumi:"ipSans"`
	// The issuing CA
	IssuingCa pulumi.StringOutput `pulumi:"issuingCa"`
	// The locality
	Locality pulumi.StringPtrOutput `pulumi:"locality"`
	// The maximum path length to encode in the generated certificate
	MaxPathLength pulumi.IntPtrOutput `pulumi:"maxPathLength"`
	// The organization
	Organization pulumi.StringPtrOutput `pulumi:"organization"`
	// List of other SANs
	OtherSans pulumi.StringArrayOutput `pulumi:"otherSans"`
	// The organization unit
	Ou pulumi.StringPtrOutput `pulumi:"ou"`
	// List of domains for which certificates are allowed to be issued
	PermittedDnsDomains pulumi.StringArrayOutput `pulumi:"permittedDnsDomains"`
	// The postal code
	PostalCode pulumi.StringPtrOutput `pulumi:"postalCode"`
	// The province
	Province pulumi.StringPtrOutput `pulumi:"province"`
	// The serial
	Serial pulumi.StringOutput `pulumi:"serial"`
	// The street address
	StreetAddress pulumi.StringPtrOutput `pulumi:"streetAddress"`
	// Time to live
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// List of alternative URIs
	UriSans pulumi.StringArrayOutput `pulumi:"uriSans"`
	// Preserve CSR values
	UseCsrValues pulumi.BoolPtrOutput `pulumi:"useCsrValues"`
}

// NewSecretBackendRootSignIntermediate registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendRootSignIntermediate(ctx *pulumi.Context,
	name string, args *SecretBackendRootSignIntermediateArgs, opts ...pulumi.ResourceOption) (*SecretBackendRootSignIntermediate, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil || args.CommonName == nil {
		return nil, errors.New("missing required argument 'CommonName'")
	}
	if args == nil || args.Csr == nil {
		return nil, errors.New("missing required argument 'Csr'")
	}
	if args == nil {
		args = &SecretBackendRootSignIntermediateArgs{}
	}
	var resource SecretBackendRootSignIntermediate
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendRootSignIntermediate:SecretBackendRootSignIntermediate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendRootSignIntermediate gets an existing SecretBackendRootSignIntermediate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendRootSignIntermediate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendRootSignIntermediateState, opts ...pulumi.ResourceOption) (*SecretBackendRootSignIntermediate, error) {
	var resource SecretBackendRootSignIntermediate
	err := ctx.ReadResource("vault:pkiSecret/secretBackendRootSignIntermediate:SecretBackendRootSignIntermediate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendRootSignIntermediate resources.
type secretBackendRootSignIntermediateState struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// The PKI secret backend the resource belongs to.
	Backend *string `pulumi:"backend"`
	// The CA chain
	CaChain *string `pulumi:"caChain"`
	// The certificate
	Certificate *string `pulumi:"certificate"`
	// CN of intermediate to create
	CommonName *string `pulumi:"commonName"`
	// The country
	Country *string `pulumi:"country"`
	// The CSR
	Csr *string `pulumi:"csr"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans *bool `pulumi:"excludeCnFromSans"`
	// The format of data
	Format *string `pulumi:"format"`
	// List of alternative IPs
	IpSans []string `pulumi:"ipSans"`
	// The issuing CA
	IssuingCa *string `pulumi:"issuingCa"`
	// The locality
	Locality *string `pulumi:"locality"`
	// The maximum path length to encode in the generated certificate
	MaxPathLength *int `pulumi:"maxPathLength"`
	// The organization
	Organization *string `pulumi:"organization"`
	// List of other SANs
	OtherSans []string `pulumi:"otherSans"`
	// The organization unit
	Ou *string `pulumi:"ou"`
	// List of domains for which certificates are allowed to be issued
	PermittedDnsDomains []string `pulumi:"permittedDnsDomains"`
	// The postal code
	PostalCode *string `pulumi:"postalCode"`
	// The province
	Province *string `pulumi:"province"`
	// The serial
	Serial *string `pulumi:"serial"`
	// The street address
	StreetAddress *string `pulumi:"streetAddress"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
	// Preserve CSR values
	UseCsrValues *bool `pulumi:"useCsrValues"`
}

type SecretBackendRootSignIntermediateState struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringPtrInput
	// The CA chain
	CaChain pulumi.StringPtrInput
	// The certificate
	Certificate pulumi.StringPtrInput
	// CN of intermediate to create
	CommonName pulumi.StringPtrInput
	// The country
	Country pulumi.StringPtrInput
	// The CSR
	Csr pulumi.StringPtrInput
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrInput
	// The format of data
	Format pulumi.StringPtrInput
	// List of alternative IPs
	IpSans pulumi.StringArrayInput
	// The issuing CA
	IssuingCa pulumi.StringPtrInput
	// The locality
	Locality pulumi.StringPtrInput
	// The maximum path length to encode in the generated certificate
	MaxPathLength pulumi.IntPtrInput
	// The organization
	Organization pulumi.StringPtrInput
	// List of other SANs
	OtherSans pulumi.StringArrayInput
	// The organization unit
	Ou pulumi.StringPtrInput
	// List of domains for which certificates are allowed to be issued
	PermittedDnsDomains pulumi.StringArrayInput
	// The postal code
	PostalCode pulumi.StringPtrInput
	// The province
	Province pulumi.StringPtrInput
	// The serial
	Serial pulumi.StringPtrInput
	// The street address
	StreetAddress pulumi.StringPtrInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
	// Preserve CSR values
	UseCsrValues pulumi.BoolPtrInput
}

func (SecretBackendRootSignIntermediateState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRootSignIntermediateState)(nil)).Elem()
}

type secretBackendRootSignIntermediateArgs struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// The PKI secret backend the resource belongs to.
	Backend string `pulumi:"backend"`
	// CN of intermediate to create
	CommonName string `pulumi:"commonName"`
	// The country
	Country *string `pulumi:"country"`
	// The CSR
	Csr string `pulumi:"csr"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans *bool `pulumi:"excludeCnFromSans"`
	// The format of data
	Format *string `pulumi:"format"`
	// List of alternative IPs
	IpSans []string `pulumi:"ipSans"`
	// The locality
	Locality *string `pulumi:"locality"`
	// The maximum path length to encode in the generated certificate
	MaxPathLength *int `pulumi:"maxPathLength"`
	// The organization
	Organization *string `pulumi:"organization"`
	// List of other SANs
	OtherSans []string `pulumi:"otherSans"`
	// The organization unit
	Ou *string `pulumi:"ou"`
	// List of domains for which certificates are allowed to be issued
	PermittedDnsDomains []string `pulumi:"permittedDnsDomains"`
	// The postal code
	PostalCode *string `pulumi:"postalCode"`
	// The province
	Province *string `pulumi:"province"`
	// The street address
	StreetAddress *string `pulumi:"streetAddress"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
	// Preserve CSR values
	UseCsrValues *bool `pulumi:"useCsrValues"`
}

// The set of arguments for constructing a SecretBackendRootSignIntermediate resource.
type SecretBackendRootSignIntermediateArgs struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringInput
	// CN of intermediate to create
	CommonName pulumi.StringInput
	// The country
	Country pulumi.StringPtrInput
	// The CSR
	Csr pulumi.StringInput
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrInput
	// The format of data
	Format pulumi.StringPtrInput
	// List of alternative IPs
	IpSans pulumi.StringArrayInput
	// The locality
	Locality pulumi.StringPtrInput
	// The maximum path length to encode in the generated certificate
	MaxPathLength pulumi.IntPtrInput
	// The organization
	Organization pulumi.StringPtrInput
	// List of other SANs
	OtherSans pulumi.StringArrayInput
	// The organization unit
	Ou pulumi.StringPtrInput
	// List of domains for which certificates are allowed to be issued
	PermittedDnsDomains pulumi.StringArrayInput
	// The postal code
	PostalCode pulumi.StringPtrInput
	// The province
	Province pulumi.StringPtrInput
	// The street address
	StreetAddress pulumi.StringPtrInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
	// Preserve CSR values
	UseCsrValues pulumi.BoolPtrInput
}

func (SecretBackendRootSignIntermediateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRootSignIntermediateArgs)(nil)).Elem()
}

