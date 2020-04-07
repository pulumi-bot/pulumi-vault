// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package pkiSecret

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type SecretBackendIntermediateCertRequest struct {
	pulumi.CustomResourceState

	// List of alternative names
	AltNames pulumi.StringArrayOutput `pulumi:"altNames"`
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringOutput `pulumi:"backend"`
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
	// The number of bits to use
	KeyBits pulumi.IntPtrOutput `pulumi:"keyBits"`
	// The desired key type
	KeyType pulumi.StringPtrOutput `pulumi:"keyType"`
	// The locality
	Locality pulumi.StringPtrOutput `pulumi:"locality"`
	// The organization
	Organization pulumi.StringPtrOutput `pulumi:"organization"`
	// List of other SANs
	OtherSans pulumi.StringArrayOutput `pulumi:"otherSans"`
	// The organization unit
	Ou pulumi.StringPtrOutput `pulumi:"ou"`
	// The postal code
	PostalCode pulumi.StringPtrOutput `pulumi:"postalCode"`
	// The private key
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The private key format
	PrivateKeyFormat pulumi.StringPtrOutput `pulumi:"privateKeyFormat"`
	// The private key type
	PrivateKeyType pulumi.StringOutput `pulumi:"privateKeyType"`
	// The province
	Province pulumi.StringPtrOutput `pulumi:"province"`
	// The street address
	StreetAddress pulumi.StringPtrOutput `pulumi:"streetAddress"`
	// Type of intermediate to create. Must be either \"exported\" or \"internal\"
	Type pulumi.StringOutput `pulumi:"type"`
	// List of alternative URIs
	UriSans pulumi.StringArrayOutput `pulumi:"uriSans"`
}

// NewSecretBackendIntermediateCertRequest registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendIntermediateCertRequest(ctx *pulumi.Context,
	name string, args *SecretBackendIntermediateCertRequestArgs, opts ...pulumi.ResourceOption) (*SecretBackendIntermediateCertRequest, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil || args.CommonName == nil {
		return nil, errors.New("missing required argument 'CommonName'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &SecretBackendIntermediateCertRequestArgs{}
	}
	var resource SecretBackendIntermediateCertRequest
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendIntermediateCertRequest:SecretBackendIntermediateCertRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendIntermediateCertRequest gets an existing SecretBackendIntermediateCertRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendIntermediateCertRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendIntermediateCertRequestState, opts ...pulumi.ResourceOption) (*SecretBackendIntermediateCertRequest, error) {
	var resource SecretBackendIntermediateCertRequest
	err := ctx.ReadResource("vault:pkiSecret/secretBackendIntermediateCertRequest:SecretBackendIntermediateCertRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendIntermediateCertRequest resources.
type secretBackendIntermediateCertRequestState struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// The PKI secret backend the resource belongs to.
	Backend *string `pulumi:"backend"`
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
	// The number of bits to use
	KeyBits *int `pulumi:"keyBits"`
	// The desired key type
	KeyType *string `pulumi:"keyType"`
	// The locality
	Locality *string `pulumi:"locality"`
	// The organization
	Organization *string `pulumi:"organization"`
	// List of other SANs
	OtherSans []string `pulumi:"otherSans"`
	// The organization unit
	Ou *string `pulumi:"ou"`
	// The postal code
	PostalCode *string `pulumi:"postalCode"`
	// The private key
	PrivateKey *string `pulumi:"privateKey"`
	// The private key format
	PrivateKeyFormat *string `pulumi:"privateKeyFormat"`
	// The private key type
	PrivateKeyType *string `pulumi:"privateKeyType"`
	// The province
	Province *string `pulumi:"province"`
	// The street address
	StreetAddress *string `pulumi:"streetAddress"`
	// Type of intermediate to create. Must be either \"exported\" or \"internal\"
	Type *string `pulumi:"type"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
}

type SecretBackendIntermediateCertRequestState struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringPtrInput
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
	// The number of bits to use
	KeyBits pulumi.IntPtrInput
	// The desired key type
	KeyType pulumi.StringPtrInput
	// The locality
	Locality pulumi.StringPtrInput
	// The organization
	Organization pulumi.StringPtrInput
	// List of other SANs
	OtherSans pulumi.StringArrayInput
	// The organization unit
	Ou pulumi.StringPtrInput
	// The postal code
	PostalCode pulumi.StringPtrInput
	// The private key
	PrivateKey pulumi.StringPtrInput
	// The private key format
	PrivateKeyFormat pulumi.StringPtrInput
	// The private key type
	PrivateKeyType pulumi.StringPtrInput
	// The province
	Province pulumi.StringPtrInput
	// The street address
	StreetAddress pulumi.StringPtrInput
	// Type of intermediate to create. Must be either \"exported\" or \"internal\"
	Type pulumi.StringPtrInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
}

func (SecretBackendIntermediateCertRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendIntermediateCertRequestState)(nil)).Elem()
}

type secretBackendIntermediateCertRequestArgs struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// The PKI secret backend the resource belongs to.
	Backend string `pulumi:"backend"`
	// CN of intermediate to create
	CommonName string `pulumi:"commonName"`
	// The country
	Country *string `pulumi:"country"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans *bool `pulumi:"excludeCnFromSans"`
	// The format of data
	Format *string `pulumi:"format"`
	// List of alternative IPs
	IpSans []string `pulumi:"ipSans"`
	// The number of bits to use
	KeyBits *int `pulumi:"keyBits"`
	// The desired key type
	KeyType *string `pulumi:"keyType"`
	// The locality
	Locality *string `pulumi:"locality"`
	// The organization
	Organization *string `pulumi:"organization"`
	// List of other SANs
	OtherSans []string `pulumi:"otherSans"`
	// The organization unit
	Ou *string `pulumi:"ou"`
	// The postal code
	PostalCode *string `pulumi:"postalCode"`
	// The private key format
	PrivateKeyFormat *string `pulumi:"privateKeyFormat"`
	// The province
	Province *string `pulumi:"province"`
	// The street address
	StreetAddress *string `pulumi:"streetAddress"`
	// Type of intermediate to create. Must be either \"exported\" or \"internal\"
	Type string `pulumi:"type"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
}

// The set of arguments for constructing a SecretBackendIntermediateCertRequest resource.
type SecretBackendIntermediateCertRequestArgs struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringInput
	// CN of intermediate to create
	CommonName pulumi.StringInput
	// The country
	Country pulumi.StringPtrInput
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrInput
	// The format of data
	Format pulumi.StringPtrInput
	// List of alternative IPs
	IpSans pulumi.StringArrayInput
	// The number of bits to use
	KeyBits pulumi.IntPtrInput
	// The desired key type
	KeyType pulumi.StringPtrInput
	// The locality
	Locality pulumi.StringPtrInput
	// The organization
	Organization pulumi.StringPtrInput
	// List of other SANs
	OtherSans pulumi.StringArrayInput
	// The organization unit
	Ou pulumi.StringPtrInput
	// The postal code
	PostalCode pulumi.StringPtrInput
	// The private key format
	PrivateKeyFormat pulumi.StringPtrInput
	// The province
	Province pulumi.StringPtrInput
	// The street address
	StreetAddress pulumi.StringPtrInput
	// Type of intermediate to create. Must be either \"exported\" or \"internal\"
	Type pulumi.StringInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
}

func (SecretBackendIntermediateCertRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendIntermediateCertRequestArgs)(nil)).Elem()
}
