// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type AuthBackendConfig struct {
	pulumi.CustomResourceState

	// The path the Azure auth backend being configured was
	// mounted at.  Defaults to `azure`.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// The client id for credentials to query the Azure APIs.
	// Currently read permissions to query compute resources are required.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The client secret for credentials to query the
	// Azure APIs.
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// The Azure cloud environment. Valid values:
	// AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
	// AzureGermanCloud.  Defaults to `AzurePublicCloud`.
	Environment pulumi.StringPtrOutput `pulumi:"environment"`
	// The configured URL for the application registered in
	// Azure Active Directory.
	Resource pulumi.StringOutput `pulumi:"resource"`
	// The tenant id for the Azure Active Directory
	// organization.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewAuthBackendConfig registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendConfig(ctx *pulumi.Context,
	name string, args *AuthBackendConfigArgs, opts ...pulumi.ResourceOption) (*AuthBackendConfig, error) {
	if args == nil || args.Resource == nil {
		return nil, errors.New("missing required argument 'Resource'")
	}
	if args == nil || args.TenantId == nil {
		return nil, errors.New("missing required argument 'TenantId'")
	}
	if args == nil {
		args = &AuthBackendConfigArgs{}
	}
	var resource AuthBackendConfig
	err := ctx.RegisterResource("vault:azure/authBackendConfig:AuthBackendConfig", name, args, &resource, opts...)
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
	err := ctx.ReadResource("vault:azure/authBackendConfig:AuthBackendConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendConfig resources.
type authBackendConfigState struct {
	// The path the Azure auth backend being configured was
	// mounted at.  Defaults to `azure`.
	Backend *string `pulumi:"backend"`
	// The client id for credentials to query the Azure APIs.
	// Currently read permissions to query compute resources are required.
	ClientId *string `pulumi:"clientId"`
	// The client secret for credentials to query the
	// Azure APIs.
	ClientSecret *string `pulumi:"clientSecret"`
	// The Azure cloud environment. Valid values:
	// AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
	// AzureGermanCloud.  Defaults to `AzurePublicCloud`.
	Environment *string `pulumi:"environment"`
	// The configured URL for the application registered in
	// Azure Active Directory.
	Resource *string `pulumi:"resource"`
	// The tenant id for the Azure Active Directory
	// organization.
	TenantId *string `pulumi:"tenantId"`
}

type AuthBackendConfigState struct {
	// The path the Azure auth backend being configured was
	// mounted at.  Defaults to `azure`.
	Backend pulumi.StringPtrInput
	// The client id for credentials to query the Azure APIs.
	// Currently read permissions to query compute resources are required.
	ClientId pulumi.StringPtrInput
	// The client secret for credentials to query the
	// Azure APIs.
	ClientSecret pulumi.StringPtrInput
	// The Azure cloud environment. Valid values:
	// AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
	// AzureGermanCloud.  Defaults to `AzurePublicCloud`.
	Environment pulumi.StringPtrInput
	// The configured URL for the application registered in
	// Azure Active Directory.
	Resource pulumi.StringPtrInput
	// The tenant id for the Azure Active Directory
	// organization.
	TenantId pulumi.StringPtrInput
}

func (AuthBackendConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendConfigState)(nil)).Elem()
}

type authBackendConfigArgs struct {
	// The path the Azure auth backend being configured was
	// mounted at.  Defaults to `azure`.
	Backend *string `pulumi:"backend"`
	// The client id for credentials to query the Azure APIs.
	// Currently read permissions to query compute resources are required.
	ClientId *string `pulumi:"clientId"`
	// The client secret for credentials to query the
	// Azure APIs.
	ClientSecret *string `pulumi:"clientSecret"`
	// The Azure cloud environment. Valid values:
	// AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
	// AzureGermanCloud.  Defaults to `AzurePublicCloud`.
	Environment *string `pulumi:"environment"`
	// The configured URL for the application registered in
	// Azure Active Directory.
	Resource string `pulumi:"resource"`
	// The tenant id for the Azure Active Directory
	// organization.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a AuthBackendConfig resource.
type AuthBackendConfigArgs struct {
	// The path the Azure auth backend being configured was
	// mounted at.  Defaults to `azure`.
	Backend pulumi.StringPtrInput
	// The client id for credentials to query the Azure APIs.
	// Currently read permissions to query compute resources are required.
	ClientId pulumi.StringPtrInput
	// The client secret for credentials to query the
	// Azure APIs.
	ClientSecret pulumi.StringPtrInput
	// The Azure cloud environment. Valid values:
	// AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
	// AzureGermanCloud.  Defaults to `AzurePublicCloud`.
	Environment pulumi.StringPtrInput
	// The configured URL for the application registered in
	// Azure Active Directory.
	Resource pulumi.StringInput
	// The tenant id for the Azure Active Directory
	// organization.
	TenantId pulumi.StringInput
}

func (AuthBackendConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendConfigArgs)(nil)).Elem()
}
