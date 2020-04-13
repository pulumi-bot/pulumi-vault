// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pkiSecret

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows setting the issuing certificate endpoints, CRL distribution points, and OCSP server endpoints that will be encoded into issued certificates.
type SecretBackendConfigUrls struct {
	pulumi.CustomResourceState

	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Specifies the URL values for the CRL Distribution Points field.
	CrlDistributionPoints pulumi.StringArrayOutput `pulumi:"crlDistributionPoints"`
	// Specifies the URL values for the Issuing Certificate field.
	IssuingCertificates pulumi.StringArrayOutput `pulumi:"issuingCertificates"`
	// Specifies the URL values for the OCSP Servers field.
	OcspServers pulumi.StringArrayOutput `pulumi:"ocspServers"`
}

// NewSecretBackendConfigUrls registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendConfigUrls(ctx *pulumi.Context,
	name string, args *SecretBackendConfigUrlsArgs, opts ...pulumi.ResourceOption) (*SecretBackendConfigUrls, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil {
		args = &SecretBackendConfigUrlsArgs{}
	}
	var resource SecretBackendConfigUrls
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendConfigUrls gets an existing SecretBackendConfigUrls resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendConfigUrls(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendConfigUrlsState, opts ...pulumi.ResourceOption) (*SecretBackendConfigUrls, error) {
	var resource SecretBackendConfigUrls
	err := ctx.ReadResource("vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendConfigUrls resources.
type secretBackendConfigUrlsState struct {
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend *string `pulumi:"backend"`
	// Specifies the URL values for the CRL Distribution Points field.
	CrlDistributionPoints []string `pulumi:"crlDistributionPoints"`
	// Specifies the URL values for the Issuing Certificate field.
	IssuingCertificates []string `pulumi:"issuingCertificates"`
	// Specifies the URL values for the OCSP Servers field.
	OcspServers []string `pulumi:"ocspServers"`
}

type SecretBackendConfigUrlsState struct {
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringPtrInput
	// Specifies the URL values for the CRL Distribution Points field.
	CrlDistributionPoints pulumi.StringArrayInput
	// Specifies the URL values for the Issuing Certificate field.
	IssuingCertificates pulumi.StringArrayInput
	// Specifies the URL values for the OCSP Servers field.
	OcspServers pulumi.StringArrayInput
}

func (SecretBackendConfigUrlsState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendConfigUrlsState)(nil)).Elem()
}

type secretBackendConfigUrlsArgs struct {
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// Specifies the URL values for the CRL Distribution Points field.
	CrlDistributionPoints []string `pulumi:"crlDistributionPoints"`
	// Specifies the URL values for the Issuing Certificate field.
	IssuingCertificates []string `pulumi:"issuingCertificates"`
	// Specifies the URL values for the OCSP Servers field.
	OcspServers []string `pulumi:"ocspServers"`
}

// The set of arguments for constructing a SecretBackendConfigUrls resource.
type SecretBackendConfigUrlsArgs struct {
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringInput
	// Specifies the URL values for the CRL Distribution Points field.
	CrlDistributionPoints pulumi.StringArrayInput
	// Specifies the URL values for the Issuing Certificate field.
	IssuingCertificates pulumi.StringArrayInput
	// Specifies the URL values for the OCSP Servers field.
	OcspServers pulumi.StringArrayInput
}

func (SecretBackendConfigUrlsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendConfigUrlsArgs)(nil)).Elem()
}
