// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package aws

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/aws_access_credentials.html.md.
func GetAccessCredentials(ctx *pulumi.Context, args *GetAccessCredentialsArgs, opts ...pulumi.InvokeOption) (*GetAccessCredentialsResult, error) {
	var rv GetAccessCredentialsResult
	err := ctx.Invoke("vault:aws/getAccessCredentials:getAccessCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessCredentials.
type GetAccessCredentialsArgs struct {
	// The path to the AWS secret backend to
	// read credentials from, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// The name of the AWS secret backend role to read
	// credentials from, with no leading or trailing `/`s.
	Role string `pulumi:"role"`
	// The specific AWS ARN to use
	// from the configured role. If the role does not have multiple ARNs, this does
	// not need to be specified.
	RoleArn *string `pulumi:"roleArn"`
	// The type of credentials to read. Defaults
	// to `"creds"`, which just returns an AWS Access Key ID and Secret
	// Key. Can also be set to `"sts"`, which will return a security token
	// in addition to the keys.
	Type *string `pulumi:"type"`
}


// A collection of values returned by getAccessCredentials.
type GetAccessCredentialsResult struct {
	// The AWS Access Key ID returned by Vault.
	AccessKey string `pulumi:"accessKey"`
	Backend string `pulumi:"backend"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The duration of the secret lease, in seconds relative
	// to the time the data was requested. Once this time has passed any plan
	// generated with this data may fail to apply.
	LeaseDuration int `pulumi:"leaseDuration"`
	// The lease identifier assigned by Vault.
	LeaseId string `pulumi:"leaseId"`
	LeaseRenewable bool `pulumi:"leaseRenewable"`
	LeaseStartTime string `pulumi:"leaseStartTime"`
	Role string `pulumi:"role"`
	RoleArn *string `pulumi:"roleArn"`
	// The AWS Secret Key returned by Vault.
	SecretKey string `pulumi:"secretKey"`
	// The STS token returned by Vault, if any.
	SecurityToken string `pulumi:"securityToken"`
	Type *string `pulumi:"type"`
}

