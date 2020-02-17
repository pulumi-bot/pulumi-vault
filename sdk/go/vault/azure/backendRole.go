// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package azure

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type BackendRole struct {
	pulumi.CustomResourceState

	// Application Object ID for an existing service principal that will be used instead of creating dynamic service
	// principals.
	ApplicationObjectId pulumi.StringPtrOutput `pulumi:"applicationObjectId"`
	AzureRoles BackendRoleAzureRoleArrayOutput `pulumi:"azureRoles"`
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// Human-friendly description of the mount for the backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Human-friendly description of the mount for the backend.
	MaxTtl pulumi.StringPtrOutput `pulumi:"maxTtl"`
	// Name of the role to create
	Role pulumi.StringOutput `pulumi:"role"`
	// Human-friendly description of the mount for the backend.
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
}

// NewBackendRole registers a new resource with the given unique name, arguments, and options.
func NewBackendRole(ctx *pulumi.Context,
	name string, args *BackendRoleArgs, opts ...pulumi.ResourceOption) (*BackendRole, error) {
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &BackendRoleArgs{}
	}
	var resource BackendRole
	err := ctx.RegisterResource("vault:azure/backendRole:BackendRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendRole gets an existing BackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendRoleState, opts ...pulumi.ResourceOption) (*BackendRole, error) {
	var resource BackendRole
	err := ctx.ReadResource("vault:azure/backendRole:BackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendRole resources.
type backendRoleState struct {
	// Application Object ID for an existing service principal that will be used instead of creating dynamic service
	// principals.
	ApplicationObjectId *string `pulumi:"applicationObjectId"`
	AzureRoles []BackendRoleAzureRole `pulumi:"azureRoles"`
	// Unique name of the auth backend to configure.
	Backend *string `pulumi:"backend"`
	// Human-friendly description of the mount for the backend.
	Description *string `pulumi:"description"`
	// Human-friendly description of the mount for the backend.
	MaxTtl *string `pulumi:"maxTtl"`
	// Name of the role to create
	Role *string `pulumi:"role"`
	// Human-friendly description of the mount for the backend.
	Ttl *string `pulumi:"ttl"`
}

type BackendRoleState struct {
	// Application Object ID for an existing service principal that will be used instead of creating dynamic service
	// principals.
	ApplicationObjectId pulumi.StringPtrInput
	AzureRoles BackendRoleAzureRoleArrayInput
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrInput
	// Human-friendly description of the mount for the backend.
	Description pulumi.StringPtrInput
	// Human-friendly description of the mount for the backend.
	MaxTtl pulumi.StringPtrInput
	// Name of the role to create
	Role pulumi.StringPtrInput
	// Human-friendly description of the mount for the backend.
	Ttl pulumi.StringPtrInput
}

func (BackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendRoleState)(nil)).Elem()
}

type backendRoleArgs struct {
	// Application Object ID for an existing service principal that will be used instead of creating dynamic service
	// principals.
	ApplicationObjectId *string `pulumi:"applicationObjectId"`
	AzureRoles []BackendRoleAzureRole `pulumi:"azureRoles"`
	// Unique name of the auth backend to configure.
	Backend *string `pulumi:"backend"`
	// Human-friendly description of the mount for the backend.
	Description *string `pulumi:"description"`
	// Human-friendly description of the mount for the backend.
	MaxTtl *string `pulumi:"maxTtl"`
	// Name of the role to create
	Role string `pulumi:"role"`
	// Human-friendly description of the mount for the backend.
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a BackendRole resource.
type BackendRoleArgs struct {
	// Application Object ID for an existing service principal that will be used instead of creating dynamic service
	// principals.
	ApplicationObjectId pulumi.StringPtrInput
	AzureRoles BackendRoleAzureRoleArrayInput
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrInput
	// Human-friendly description of the mount for the backend.
	Description pulumi.StringPtrInput
	// Human-friendly description of the mount for the backend.
	MaxTtl pulumi.StringPtrInput
	// Name of the role to create
	Role pulumi.StringInput
	// Human-friendly description of the mount for the backend.
	Ttl pulumi.StringPtrInput
}

func (BackendRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendRoleArgs)(nil)).Elem()
}

