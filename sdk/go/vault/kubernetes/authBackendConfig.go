// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package kubernetes

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Kubernetes auth backend config in a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/kubernetes.html) for more
// information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/kubernetes_auth_backend_config.md.
type AuthBackendConfig struct {
	pulumi.CustomResourceState

	// Unique name of the kubernetes backend to configure.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer. 
	Issuer pulumi.StringPtrOutput `pulumi:"issuer"`
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert pulumi.StringPtrOutput `pulumi:"kubernetesCaCert"`
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost pulumi.StringOutput `pulumi:"kubernetesHost"`
	// List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys pulumi.StringArrayOutput `pulumi:"pemKeys"`
	// A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	TokenReviewerJwt pulumi.StringPtrOutput `pulumi:"tokenReviewerJwt"`
}

// NewAuthBackendConfig registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendConfig(ctx *pulumi.Context,
	name string, args *AuthBackendConfigArgs, opts ...pulumi.ResourceOption) (*AuthBackendConfig, error) {
	if args == nil || args.KubernetesHost == nil {
		return nil, errors.New("missing required argument 'KubernetesHost'")
	}
	if args == nil {
		args = &AuthBackendConfigArgs{}
	}
	var resource AuthBackendConfig
	err := ctx.RegisterResource("vault:kubernetes/authBackendConfig:AuthBackendConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendConfig gets an existing AuthBackendConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendConfigState, opts ...pulumi.ResourceOption) (*AuthBackendConfig, error) {
	var resource AuthBackendConfig
	err := ctx.ReadResource("vault:kubernetes/authBackendConfig:AuthBackendConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendConfig resources.
type authBackendConfigState struct {
	// Unique name of the kubernetes backend to configure.
	Backend *string `pulumi:"backend"`
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer. 
	Issuer *string `pulumi:"issuer"`
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert *string `pulumi:"kubernetesCaCert"`
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost *string `pulumi:"kubernetesHost"`
	// List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys []string `pulumi:"pemKeys"`
	// A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	TokenReviewerJwt *string `pulumi:"tokenReviewerJwt"`
}

type AuthBackendConfigState struct {
	// Unique name of the kubernetes backend to configure.
	Backend pulumi.StringPtrInput
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer. 
	Issuer pulumi.StringPtrInput
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert pulumi.StringPtrInput
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost pulumi.StringPtrInput
	// List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys pulumi.StringArrayInput
	// A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	TokenReviewerJwt pulumi.StringPtrInput
}

func (AuthBackendConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendConfigState)(nil)).Elem()
}

type authBackendConfigArgs struct {
	// Unique name of the kubernetes backend to configure.
	Backend *string `pulumi:"backend"`
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer. 
	Issuer *string `pulumi:"issuer"`
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert *string `pulumi:"kubernetesCaCert"`
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost string `pulumi:"kubernetesHost"`
	// List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys []string `pulumi:"pemKeys"`
	// A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	TokenReviewerJwt *string `pulumi:"tokenReviewerJwt"`
}

// The set of arguments for constructing a AuthBackendConfig resource.
type AuthBackendConfigArgs struct {
	// Unique name of the kubernetes backend to configure.
	Backend pulumi.StringPtrInput
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer. 
	Issuer pulumi.StringPtrInput
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert pulumi.StringPtrInput
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost pulumi.StringInput
	// List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys pulumi.StringArrayInput
	// A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	TokenReviewerJwt pulumi.StringPtrInput
}

func (AuthBackendConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendConfigArgs)(nil)).Elem()
}

