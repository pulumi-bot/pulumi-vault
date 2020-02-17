// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package database

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type SecretBackendRole struct {
	pulumi.CustomResourceState

	// The path of the Database Secret Backend the role belongs to.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Database statements to execute to create and configure a user.
	CreationStatements pulumi.StringArrayOutput `pulumi:"creationStatements"`
	// Database connection to use for this role.
	DbName pulumi.StringOutput `pulumi:"dbName"`
	// Default TTL for leases associated with this role, in seconds.
	DefaultTtl pulumi.IntPtrOutput `pulumi:"defaultTtl"`
	// Maximum TTL for leases associated with this role, in seconds.
	MaxTtl pulumi.IntPtrOutput `pulumi:"maxTtl"`
	// Unique name for the role.
	Name pulumi.StringOutput `pulumi:"name"`
	// Database statements to execute to renew a user.
	RenewStatements pulumi.StringArrayOutput `pulumi:"renewStatements"`
	// Database statements to execute to revoke a user.
	RevocationStatements pulumi.StringArrayOutput `pulumi:"revocationStatements"`
	// Database statements to execute to rollback a create operation in the event of an error.
	RollbackStatements pulumi.StringArrayOutput `pulumi:"rollbackStatements"`
}

// NewSecretBackendRole registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendRole(ctx *pulumi.Context,
	name string, args *SecretBackendRoleArgs, opts ...pulumi.ResourceOption) (*SecretBackendRole, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil || args.CreationStatements == nil {
		return nil, errors.New("missing required argument 'CreationStatements'")
	}
	if args == nil || args.DbName == nil {
		return nil, errors.New("missing required argument 'DbName'")
	}
	if args == nil {
		args = &SecretBackendRoleArgs{}
	}
	var resource SecretBackendRole
	err := ctx.RegisterResource("vault:database/secretBackendRole:SecretBackendRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendRole gets an existing SecretBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendRoleState, opts ...pulumi.ResourceOption) (*SecretBackendRole, error) {
	var resource SecretBackendRole
	err := ctx.ReadResource("vault:database/secretBackendRole:SecretBackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendRole resources.
type secretBackendRoleState struct {
	// The path of the Database Secret Backend the role belongs to.
	Backend *string `pulumi:"backend"`
	// Database statements to execute to create and configure a user.
	CreationStatements []string `pulumi:"creationStatements"`
	// Database connection to use for this role.
	DbName *string `pulumi:"dbName"`
	// Default TTL for leases associated with this role, in seconds.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// Maximum TTL for leases associated with this role, in seconds.
	MaxTtl *int `pulumi:"maxTtl"`
	// Unique name for the role.
	Name *string `pulumi:"name"`
	// Database statements to execute to renew a user.
	RenewStatements []string `pulumi:"renewStatements"`
	// Database statements to execute to revoke a user.
	RevocationStatements []string `pulumi:"revocationStatements"`
	// Database statements to execute to rollback a create operation in the event of an error.
	RollbackStatements []string `pulumi:"rollbackStatements"`
}

type SecretBackendRoleState struct {
	// The path of the Database Secret Backend the role belongs to.
	Backend pulumi.StringPtrInput
	// Database statements to execute to create and configure a user.
	CreationStatements pulumi.StringArrayInput
	// Database connection to use for this role.
	DbName pulumi.StringPtrInput
	// Default TTL for leases associated with this role, in seconds.
	DefaultTtl pulumi.IntPtrInput
	// Maximum TTL for leases associated with this role, in seconds.
	MaxTtl pulumi.IntPtrInput
	// Unique name for the role.
	Name pulumi.StringPtrInput
	// Database statements to execute to renew a user.
	RenewStatements pulumi.StringArrayInput
	// Database statements to execute to revoke a user.
	RevocationStatements pulumi.StringArrayInput
	// Database statements to execute to rollback a create operation in the event of an error.
	RollbackStatements pulumi.StringArrayInput
}

func (SecretBackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRoleState)(nil)).Elem()
}

type secretBackendRoleArgs struct {
	// The path of the Database Secret Backend the role belongs to.
	Backend string `pulumi:"backend"`
	// Database statements to execute to create and configure a user.
	CreationStatements []string `pulumi:"creationStatements"`
	// Database connection to use for this role.
	DbName string `pulumi:"dbName"`
	// Default TTL for leases associated with this role, in seconds.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// Maximum TTL for leases associated with this role, in seconds.
	MaxTtl *int `pulumi:"maxTtl"`
	// Unique name for the role.
	Name *string `pulumi:"name"`
	// Database statements to execute to renew a user.
	RenewStatements []string `pulumi:"renewStatements"`
	// Database statements to execute to revoke a user.
	RevocationStatements []string `pulumi:"revocationStatements"`
	// Database statements to execute to rollback a create operation in the event of an error.
	RollbackStatements []string `pulumi:"rollbackStatements"`
}

// The set of arguments for constructing a SecretBackendRole resource.
type SecretBackendRoleArgs struct {
	// The path of the Database Secret Backend the role belongs to.
	Backend pulumi.StringInput
	// Database statements to execute to create and configure a user.
	CreationStatements pulumi.StringArrayInput
	// Database connection to use for this role.
	DbName pulumi.StringInput
	// Default TTL for leases associated with this role, in seconds.
	DefaultTtl pulumi.IntPtrInput
	// Maximum TTL for leases associated with this role, in seconds.
	MaxTtl pulumi.IntPtrInput
	// Unique name for the role.
	Name pulumi.StringPtrInput
	// Database statements to execute to renew a user.
	RenewStatements pulumi.StringArrayInput
	// Database statements to execute to revoke a user.
	RevocationStatements pulumi.StringArrayInput
	// Database statements to execute to rollback a create operation in the event of an error.
	RollbackStatements pulumi.StringArrayInput
}

func (SecretBackendRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRoleArgs)(nil)).Elem()
}

