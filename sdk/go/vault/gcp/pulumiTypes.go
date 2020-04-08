// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gcp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type SecretRolesetBinding struct {
	// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
	Resource string `pulumi:"resource"`
	// List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
	Roles []string `pulumi:"roles"`
}

type SecretRolesetBindingInput interface {
	pulumi.Input

	ToSecretRolesetBindingOutput() SecretRolesetBindingOutput
	ToSecretRolesetBindingOutputWithContext(context.Context) SecretRolesetBindingOutput
}

type SecretRolesetBindingArgs struct {
	// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
	Resource pulumi.StringInput `pulumi:"resource"`
	// List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
	Roles pulumi.StringArrayInput `pulumi:"roles"`
}

func (SecretRolesetBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretRolesetBinding)(nil)).Elem()
}

func (i SecretRolesetBindingArgs) ToSecretRolesetBindingOutput() SecretRolesetBindingOutput {
	return i.ToSecretRolesetBindingOutputWithContext(context.Background())
}

func (i SecretRolesetBindingArgs) ToSecretRolesetBindingOutputWithContext(ctx context.Context) SecretRolesetBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRolesetBindingOutput)
}

type SecretRolesetBindingArrayInput interface {
	pulumi.Input

	ToSecretRolesetBindingArrayOutput() SecretRolesetBindingArrayOutput
	ToSecretRolesetBindingArrayOutputWithContext(context.Context) SecretRolesetBindingArrayOutput
}

type SecretRolesetBindingArray []SecretRolesetBindingInput

func (SecretRolesetBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretRolesetBinding)(nil)).Elem()
}

func (i SecretRolesetBindingArray) ToSecretRolesetBindingArrayOutput() SecretRolesetBindingArrayOutput {
	return i.ToSecretRolesetBindingArrayOutputWithContext(context.Background())
}

func (i SecretRolesetBindingArray) ToSecretRolesetBindingArrayOutputWithContext(ctx context.Context) SecretRolesetBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRolesetBindingArrayOutput)
}

type SecretRolesetBindingOutput struct{ *pulumi.OutputState }

func (SecretRolesetBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretRolesetBinding)(nil)).Elem()
}

func (o SecretRolesetBindingOutput) ToSecretRolesetBindingOutput() SecretRolesetBindingOutput {
	return o
}

func (o SecretRolesetBindingOutput) ToSecretRolesetBindingOutputWithContext(ctx context.Context) SecretRolesetBindingOutput {
	return o
}

// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
func (o SecretRolesetBindingOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v SecretRolesetBinding) string { return v.Resource }).(pulumi.StringOutput)
}

// List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
func (o SecretRolesetBindingOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecretRolesetBinding) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

type SecretRolesetBindingArrayOutput struct{ *pulumi.OutputState }

func (SecretRolesetBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretRolesetBinding)(nil)).Elem()
}

func (o SecretRolesetBindingArrayOutput) ToSecretRolesetBindingArrayOutput() SecretRolesetBindingArrayOutput {
	return o
}

func (o SecretRolesetBindingArrayOutput) ToSecretRolesetBindingArrayOutputWithContext(ctx context.Context) SecretRolesetBindingArrayOutput {
	return o
}

func (o SecretRolesetBindingArrayOutput) Index(i pulumi.IntInput) SecretRolesetBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretRolesetBinding {
		return vs[0].([]SecretRolesetBinding)[vs[1].(int)]
	}).(SecretRolesetBindingOutput)
}

type SecretRolesetBindingArgs struct {
	// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
	Resource string `pulumi:"resource"`
	// List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
	Roles []string `pulumi:"roles"`
}

type SecretRolesetBindingArgsInput interface {
	pulumi.Input

	ToSecretRolesetBindingArgsOutput() SecretRolesetBindingArgsOutput
	ToSecretRolesetBindingArgsOutputWithContext(context.Context) SecretRolesetBindingArgsOutput
}

type SecretRolesetBindingArgsArgs struct {
	// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
	Resource pulumi.StringInput `pulumi:"resource"`
	// List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
	Roles pulumi.StringArrayInput `pulumi:"roles"`
}

func (SecretRolesetBindingArgsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretRolesetBindingArgs)(nil)).Elem()
}

func (i SecretRolesetBindingArgsArgs) ToSecretRolesetBindingArgsOutput() SecretRolesetBindingArgsOutput {
	return i.ToSecretRolesetBindingArgsOutputWithContext(context.Background())
}

func (i SecretRolesetBindingArgsArgs) ToSecretRolesetBindingArgsOutputWithContext(ctx context.Context) SecretRolesetBindingArgsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRolesetBindingArgsOutput)
}

type SecretRolesetBindingArgsArrayInput interface {
	pulumi.Input

	ToSecretRolesetBindingArgsArrayOutput() SecretRolesetBindingArgsArrayOutput
	ToSecretRolesetBindingArgsArrayOutputWithContext(context.Context) SecretRolesetBindingArgsArrayOutput
}

type SecretRolesetBindingArgsArray []SecretRolesetBindingArgsInput

func (SecretRolesetBindingArgsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretRolesetBindingArgs)(nil)).Elem()
}

func (i SecretRolesetBindingArgsArray) ToSecretRolesetBindingArgsArrayOutput() SecretRolesetBindingArgsArrayOutput {
	return i.ToSecretRolesetBindingArgsArrayOutputWithContext(context.Background())
}

func (i SecretRolesetBindingArgsArray) ToSecretRolesetBindingArgsArrayOutputWithContext(ctx context.Context) SecretRolesetBindingArgsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRolesetBindingArgsArrayOutput)
}

type SecretRolesetBindingArgsOutput struct{ *pulumi.OutputState }

func (SecretRolesetBindingArgsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretRolesetBindingArgs)(nil)).Elem()
}

func (o SecretRolesetBindingArgsOutput) ToSecretRolesetBindingArgsOutput() SecretRolesetBindingArgsOutput {
	return o
}

func (o SecretRolesetBindingArgsOutput) ToSecretRolesetBindingArgsOutputWithContext(ctx context.Context) SecretRolesetBindingArgsOutput {
	return o
}

// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
func (o SecretRolesetBindingArgsOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v SecretRolesetBindingArgs) string { return v.Resource }).(pulumi.StringOutput)
}

// List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
func (o SecretRolesetBindingArgsOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecretRolesetBindingArgs) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

type SecretRolesetBindingArgsArrayOutput struct{ *pulumi.OutputState }

func (SecretRolesetBindingArgsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretRolesetBindingArgs)(nil)).Elem()
}

func (o SecretRolesetBindingArgsArrayOutput) ToSecretRolesetBindingArgsArrayOutput() SecretRolesetBindingArgsArrayOutput {
	return o
}

func (o SecretRolesetBindingArgsArrayOutput) ToSecretRolesetBindingArgsArrayOutputWithContext(ctx context.Context) SecretRolesetBindingArgsArrayOutput {
	return o
}

func (o SecretRolesetBindingArgsArrayOutput) Index(i pulumi.IntInput) SecretRolesetBindingArgsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretRolesetBindingArgs {
		return vs[0].([]SecretRolesetBindingArgs)[vs[1].(int)]
	}).(SecretRolesetBindingArgsOutput)
}

type SecretRolesetBindingState struct {
	// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
	Resource string `pulumi:"resource"`
	// List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
	Roles []string `pulumi:"roles"`
}

type SecretRolesetBindingStateInput interface {
	pulumi.Input

	ToSecretRolesetBindingStateOutput() SecretRolesetBindingStateOutput
	ToSecretRolesetBindingStateOutputWithContext(context.Context) SecretRolesetBindingStateOutput
}

type SecretRolesetBindingStateArgs struct {
	// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
	Resource pulumi.StringInput `pulumi:"resource"`
	// List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
	Roles pulumi.StringArrayInput `pulumi:"roles"`
}

func (SecretRolesetBindingStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretRolesetBindingState)(nil)).Elem()
}

func (i SecretRolesetBindingStateArgs) ToSecretRolesetBindingStateOutput() SecretRolesetBindingStateOutput {
	return i.ToSecretRolesetBindingStateOutputWithContext(context.Background())
}

func (i SecretRolesetBindingStateArgs) ToSecretRolesetBindingStateOutputWithContext(ctx context.Context) SecretRolesetBindingStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRolesetBindingStateOutput)
}

type SecretRolesetBindingStateArrayInput interface {
	pulumi.Input

	ToSecretRolesetBindingStateArrayOutput() SecretRolesetBindingStateArrayOutput
	ToSecretRolesetBindingStateArrayOutputWithContext(context.Context) SecretRolesetBindingStateArrayOutput
}

type SecretRolesetBindingStateArray []SecretRolesetBindingStateInput

func (SecretRolesetBindingStateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretRolesetBindingState)(nil)).Elem()
}

func (i SecretRolesetBindingStateArray) ToSecretRolesetBindingStateArrayOutput() SecretRolesetBindingStateArrayOutput {
	return i.ToSecretRolesetBindingStateArrayOutputWithContext(context.Background())
}

func (i SecretRolesetBindingStateArray) ToSecretRolesetBindingStateArrayOutputWithContext(ctx context.Context) SecretRolesetBindingStateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRolesetBindingStateArrayOutput)
}

type SecretRolesetBindingStateOutput struct{ *pulumi.OutputState }

func (SecretRolesetBindingStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretRolesetBindingState)(nil)).Elem()
}

func (o SecretRolesetBindingStateOutput) ToSecretRolesetBindingStateOutput() SecretRolesetBindingStateOutput {
	return o
}

func (o SecretRolesetBindingStateOutput) ToSecretRolesetBindingStateOutputWithContext(ctx context.Context) SecretRolesetBindingStateOutput {
	return o
}

// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
func (o SecretRolesetBindingStateOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v SecretRolesetBindingState) string { return v.Resource }).(pulumi.StringOutput)
}

// List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
func (o SecretRolesetBindingStateOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecretRolesetBindingState) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

type SecretRolesetBindingStateArrayOutput struct{ *pulumi.OutputState }

func (SecretRolesetBindingStateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretRolesetBindingState)(nil)).Elem()
}

func (o SecretRolesetBindingStateArrayOutput) ToSecretRolesetBindingStateArrayOutput() SecretRolesetBindingStateArrayOutput {
	return o
}

func (o SecretRolesetBindingStateArrayOutput) ToSecretRolesetBindingStateArrayOutputWithContext(ctx context.Context) SecretRolesetBindingStateArrayOutput {
	return o
}

func (o SecretRolesetBindingStateArrayOutput) Index(i pulumi.IntInput) SecretRolesetBindingStateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretRolesetBindingState {
		return vs[0].([]SecretRolesetBindingState)[vs[1].(int)]
	}).(SecretRolesetBindingStateOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretRolesetBindingOutput{})
	pulumi.RegisterOutputType(SecretRolesetBindingArrayOutput{})
	pulumi.RegisterOutputType(SecretRolesetBindingArgsOutput{})
	pulumi.RegisterOutputType(SecretRolesetBindingArgsArrayOutput{})
	pulumi.RegisterOutputType(SecretRolesetBindingStateOutput{})
	pulumi.RegisterOutputType(SecretRolesetBindingStateArrayOutput{})
}
