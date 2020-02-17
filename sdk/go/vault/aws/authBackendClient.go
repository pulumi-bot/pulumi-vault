// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package aws

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type AuthBackendClient struct {
	pulumi.CustomResourceState

	// AWS Access key with permissions to query AWS APIs.
	AccessKey pulumi.StringPtrOutput `pulumi:"accessKey"`
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// URL to override the default generated endpoint for making AWS EC2 API calls.
	Ec2Endpoint pulumi.StringPtrOutput `pulumi:"ec2Endpoint"`
	// URL to override the default generated endpoint for making AWS IAM API calls.
	IamEndpoint pulumi.StringPtrOutput `pulumi:"iamEndpoint"`
	// The value to require in the X-Vault-AWS-IAM-Server-ID header as part of GetCallerIdentity requests that are used in the
	// iam auth method.
	IamServerIdHeaderValue pulumi.StringPtrOutput `pulumi:"iamServerIdHeaderValue"`
	// AWS Secret key with permissions to query AWS APIs.
	SecretKey pulumi.StringPtrOutput `pulumi:"secretKey"`
	// URL to override the default generated endpoint for making AWS STS API calls.
	StsEndpoint pulumi.StringPtrOutput `pulumi:"stsEndpoint"`
}

// NewAuthBackendClient registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendClient(ctx *pulumi.Context,
	name string, args *AuthBackendClientArgs, opts ...pulumi.ResourceOption) (*AuthBackendClient, error) {
	if args == nil {
		args = &AuthBackendClientArgs{}
	}
	var resource AuthBackendClient
	err := ctx.RegisterResource("vault:aws/authBackendClient:AuthBackendClient", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendClient gets an existing AuthBackendClient resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendClientState, opts ...pulumi.ResourceOption) (*AuthBackendClient, error) {
	var resource AuthBackendClient
	err := ctx.ReadResource("vault:aws/authBackendClient:AuthBackendClient", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendClient resources.
type authBackendClientState struct {
	// AWS Access key with permissions to query AWS APIs.
	AccessKey *string `pulumi:"accessKey"`
	// Unique name of the auth backend to configure.
	Backend *string `pulumi:"backend"`
	// URL to override the default generated endpoint for making AWS EC2 API calls.
	Ec2Endpoint *string `pulumi:"ec2Endpoint"`
	// URL to override the default generated endpoint for making AWS IAM API calls.
	IamEndpoint *string `pulumi:"iamEndpoint"`
	// The value to require in the X-Vault-AWS-IAM-Server-ID header as part of GetCallerIdentity requests that are used in the
	// iam auth method.
	IamServerIdHeaderValue *string `pulumi:"iamServerIdHeaderValue"`
	// AWS Secret key with permissions to query AWS APIs.
	SecretKey *string `pulumi:"secretKey"`
	// URL to override the default generated endpoint for making AWS STS API calls.
	StsEndpoint *string `pulumi:"stsEndpoint"`
}

type AuthBackendClientState struct {
	// AWS Access key with permissions to query AWS APIs.
	AccessKey pulumi.StringPtrInput
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrInput
	// URL to override the default generated endpoint for making AWS EC2 API calls.
	Ec2Endpoint pulumi.StringPtrInput
	// URL to override the default generated endpoint for making AWS IAM API calls.
	IamEndpoint pulumi.StringPtrInput
	// The value to require in the X-Vault-AWS-IAM-Server-ID header as part of GetCallerIdentity requests that are used in the
	// iam auth method.
	IamServerIdHeaderValue pulumi.StringPtrInput
	// AWS Secret key with permissions to query AWS APIs.
	SecretKey pulumi.StringPtrInput
	// URL to override the default generated endpoint for making AWS STS API calls.
	StsEndpoint pulumi.StringPtrInput
}

func (AuthBackendClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendClientState)(nil)).Elem()
}

type authBackendClientArgs struct {
	// AWS Access key with permissions to query AWS APIs.
	AccessKey *string `pulumi:"accessKey"`
	// Unique name of the auth backend to configure.
	Backend *string `pulumi:"backend"`
	// URL to override the default generated endpoint for making AWS EC2 API calls.
	Ec2Endpoint *string `pulumi:"ec2Endpoint"`
	// URL to override the default generated endpoint for making AWS IAM API calls.
	IamEndpoint *string `pulumi:"iamEndpoint"`
	// The value to require in the X-Vault-AWS-IAM-Server-ID header as part of GetCallerIdentity requests that are used in the
	// iam auth method.
	IamServerIdHeaderValue *string `pulumi:"iamServerIdHeaderValue"`
	// AWS Secret key with permissions to query AWS APIs.
	SecretKey *string `pulumi:"secretKey"`
	// URL to override the default generated endpoint for making AWS STS API calls.
	StsEndpoint *string `pulumi:"stsEndpoint"`
}

// The set of arguments for constructing a AuthBackendClient resource.
type AuthBackendClientArgs struct {
	// AWS Access key with permissions to query AWS APIs.
	AccessKey pulumi.StringPtrInput
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrInput
	// URL to override the default generated endpoint for making AWS EC2 API calls.
	Ec2Endpoint pulumi.StringPtrInput
	// URL to override the default generated endpoint for making AWS IAM API calls.
	IamEndpoint pulumi.StringPtrInput
	// The value to require in the X-Vault-AWS-IAM-Server-ID header as part of GetCallerIdentity requests that are used in the
	// iam auth method.
	IamServerIdHeaderValue pulumi.StringPtrInput
	// AWS Secret key with permissions to query AWS APIs.
	SecretKey pulumi.StringPtrInput
	// URL to override the default generated endpoint for making AWS STS API calls.
	StsEndpoint pulumi.StringPtrInput
}

func (AuthBackendClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendClientArgs)(nil)).Elem()
}

