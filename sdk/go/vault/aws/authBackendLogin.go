// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package aws

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Logs into a Vault server using an AWS auth backend. Login can be
// accomplished using a signed identity request from IAM or using ec2
// instance metadata. For more information, see the [Vault
// documentation](https://www.vaultproject.io/docs/auth/aws.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/aws_auth_backend_login.html.md.
type AuthBackendLogin struct {
	pulumi.CustomResourceState

	// The token's accessor.
	Accessor pulumi.StringOutput `pulumi:"accessor"`
	// The authentication type used to generate this token.
	AuthType pulumi.StringOutput `pulumi:"authType"`
	// The unique name of the AWS auth backend. Defaults to
	// 'aws'.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// The token returned by Vault.
	ClientToken pulumi.StringOutput `pulumi:"clientToken"`
	// The HTTP method used in the signed IAM
	// request.
	IamHttpRequestMethod pulumi.StringPtrOutput `pulumi:"iamHttpRequestMethod"`
	// The base64-encoded body of the signed
	// request.
	IamRequestBody pulumi.StringPtrOutput `pulumi:"iamRequestBody"`
	// The base64-encoded, JSON serialized
	// representation of the GetCallerIdentity HTTP request headers.
	IamRequestHeaders pulumi.StringPtrOutput `pulumi:"iamRequestHeaders"`
	// The base64-encoded HTTP URL used in the signed
	// request.
	IamRequestUrl pulumi.StringPtrOutput `pulumi:"iamRequestUrl"`
	// The base64-encoded EC2 instance identity document to
	// authenticate with. Can be retrieved from the EC2 metadata server.
	Identity pulumi.StringPtrOutput `pulumi:"identity"`
	// The duration in seconds the token will be valid, relative
	// to the time in `leaseStartTime`.
	LeaseDuration pulumi.IntOutput `pulumi:"leaseDuration"`
	// Time at which the lease was read, using the clock of the system where Terraform was running
	LeaseStartTime pulumi.StringOutput `pulumi:"leaseStartTime"`
	// A map of information returned by the Vault server about the
	// authentication used to generate this token.
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// The unique nonce to be used for login requests. Can be
	// set to a user-specified value, or will contain the server-generated value
	// once a token is issued. EC2 instances can only acquire a single token until
	// the whitelist is tidied again unless they keep track of this nonce.
	Nonce pulumi.StringOutput `pulumi:"nonce"`
	// The PKCS#7 signature of the identity document to
	// authenticate with, with all newline characters removed. Can be retrieved from
	// the EC2 metadata server.
	Pkcs7 pulumi.StringPtrOutput `pulumi:"pkcs7"`
	// The Vault policies assigned to this token.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// Set to true if the token can be extended through renewal.
	Renewable pulumi.BoolOutput `pulumi:"renewable"`
	// The name of the AWS auth backend role to create tokens
	// against.
	Role pulumi.StringOutput `pulumi:"role"`
	// The base64-encoded SHA256 RSA signature of the
	// instance identity document to authenticate with, with all newline characters
	// removed. Can be retrieved from the EC2 metadata server.
	Signature pulumi.StringPtrOutput `pulumi:"signature"`
}

// NewAuthBackendLogin registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendLogin(ctx *pulumi.Context,
	name string, args *AuthBackendLoginArgs, opts ...pulumi.ResourceOption) (*AuthBackendLogin, error) {
	if args == nil {
		args = &AuthBackendLoginArgs{}
	}
	var resource AuthBackendLogin
	err := ctx.RegisterResource("vault:aws/authBackendLogin:AuthBackendLogin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendLogin gets an existing AuthBackendLogin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendLogin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendLoginState, opts ...pulumi.ResourceOption) (*AuthBackendLogin, error) {
	var resource AuthBackendLogin
	err := ctx.ReadResource("vault:aws/authBackendLogin:AuthBackendLogin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendLogin resources.
type authBackendLoginState struct {
	// The token's accessor.
	Accessor *string `pulumi:"accessor"`
	// The authentication type used to generate this token.
	AuthType *string `pulumi:"authType"`
	// The unique name of the AWS auth backend. Defaults to
	// 'aws'.
	Backend *string `pulumi:"backend"`
	// The token returned by Vault.
	ClientToken *string `pulumi:"clientToken"`
	// The HTTP method used in the signed IAM
	// request.
	IamHttpRequestMethod *string `pulumi:"iamHttpRequestMethod"`
	// The base64-encoded body of the signed
	// request.
	IamRequestBody *string `pulumi:"iamRequestBody"`
	// The base64-encoded, JSON serialized
	// representation of the GetCallerIdentity HTTP request headers.
	IamRequestHeaders *string `pulumi:"iamRequestHeaders"`
	// The base64-encoded HTTP URL used in the signed
	// request.
	IamRequestUrl *string `pulumi:"iamRequestUrl"`
	// The base64-encoded EC2 instance identity document to
	// authenticate with. Can be retrieved from the EC2 metadata server.
	Identity *string `pulumi:"identity"`
	// The duration in seconds the token will be valid, relative
	// to the time in `leaseStartTime`.
	LeaseDuration *int `pulumi:"leaseDuration"`
	// Time at which the lease was read, using the clock of the system where Terraform was running
	LeaseStartTime *string `pulumi:"leaseStartTime"`
	// A map of information returned by the Vault server about the
	// authentication used to generate this token.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The unique nonce to be used for login requests. Can be
	// set to a user-specified value, or will contain the server-generated value
	// once a token is issued. EC2 instances can only acquire a single token until
	// the whitelist is tidied again unless they keep track of this nonce.
	Nonce *string `pulumi:"nonce"`
	// The PKCS#7 signature of the identity document to
	// authenticate with, with all newline characters removed. Can be retrieved from
	// the EC2 metadata server.
	Pkcs7 *string `pulumi:"pkcs7"`
	// The Vault policies assigned to this token.
	Policies []string `pulumi:"policies"`
	// Set to true if the token can be extended through renewal.
	Renewable *bool `pulumi:"renewable"`
	// The name of the AWS auth backend role to create tokens
	// against.
	Role *string `pulumi:"role"`
	// The base64-encoded SHA256 RSA signature of the
	// instance identity document to authenticate with, with all newline characters
	// removed. Can be retrieved from the EC2 metadata server.
	Signature *string `pulumi:"signature"`
}

type AuthBackendLoginState struct {
	// The token's accessor.
	Accessor pulumi.StringPtrInput
	// The authentication type used to generate this token.
	AuthType pulumi.StringPtrInput
	// The unique name of the AWS auth backend. Defaults to
	// 'aws'.
	Backend pulumi.StringPtrInput
	// The token returned by Vault.
	ClientToken pulumi.StringPtrInput
	// The HTTP method used in the signed IAM
	// request.
	IamHttpRequestMethod pulumi.StringPtrInput
	// The base64-encoded body of the signed
	// request.
	IamRequestBody pulumi.StringPtrInput
	// The base64-encoded, JSON serialized
	// representation of the GetCallerIdentity HTTP request headers.
	IamRequestHeaders pulumi.StringPtrInput
	// The base64-encoded HTTP URL used in the signed
	// request.
	IamRequestUrl pulumi.StringPtrInput
	// The base64-encoded EC2 instance identity document to
	// authenticate with. Can be retrieved from the EC2 metadata server.
	Identity pulumi.StringPtrInput
	// The duration in seconds the token will be valid, relative
	// to the time in `leaseStartTime`.
	LeaseDuration pulumi.IntPtrInput
	// Time at which the lease was read, using the clock of the system where Terraform was running
	LeaseStartTime pulumi.StringPtrInput
	// A map of information returned by the Vault server about the
	// authentication used to generate this token.
	Metadata pulumi.MapInput
	// The unique nonce to be used for login requests. Can be
	// set to a user-specified value, or will contain the server-generated value
	// once a token is issued. EC2 instances can only acquire a single token until
	// the whitelist is tidied again unless they keep track of this nonce.
	Nonce pulumi.StringPtrInput
	// The PKCS#7 signature of the identity document to
	// authenticate with, with all newline characters removed. Can be retrieved from
	// the EC2 metadata server.
	Pkcs7 pulumi.StringPtrInput
	// The Vault policies assigned to this token.
	Policies pulumi.StringArrayInput
	// Set to true if the token can be extended through renewal.
	Renewable pulumi.BoolPtrInput
	// The name of the AWS auth backend role to create tokens
	// against.
	Role pulumi.StringPtrInput
	// The base64-encoded SHA256 RSA signature of the
	// instance identity document to authenticate with, with all newline characters
	// removed. Can be retrieved from the EC2 metadata server.
	Signature pulumi.StringPtrInput
}

func (AuthBackendLoginState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendLoginState)(nil)).Elem()
}

type authBackendLoginArgs struct {
	// The unique name of the AWS auth backend. Defaults to
	// 'aws'.
	Backend *string `pulumi:"backend"`
	// The HTTP method used in the signed IAM
	// request.
	IamHttpRequestMethod *string `pulumi:"iamHttpRequestMethod"`
	// The base64-encoded body of the signed
	// request.
	IamRequestBody *string `pulumi:"iamRequestBody"`
	// The base64-encoded, JSON serialized
	// representation of the GetCallerIdentity HTTP request headers.
	IamRequestHeaders *string `pulumi:"iamRequestHeaders"`
	// The base64-encoded HTTP URL used in the signed
	// request.
	IamRequestUrl *string `pulumi:"iamRequestUrl"`
	// The base64-encoded EC2 instance identity document to
	// authenticate with. Can be retrieved from the EC2 metadata server.
	Identity *string `pulumi:"identity"`
	// The unique nonce to be used for login requests. Can be
	// set to a user-specified value, or will contain the server-generated value
	// once a token is issued. EC2 instances can only acquire a single token until
	// the whitelist is tidied again unless they keep track of this nonce.
	Nonce *string `pulumi:"nonce"`
	// The PKCS#7 signature of the identity document to
	// authenticate with, with all newline characters removed. Can be retrieved from
	// the EC2 metadata server.
	Pkcs7 *string `pulumi:"pkcs7"`
	// The name of the AWS auth backend role to create tokens
	// against.
	Role *string `pulumi:"role"`
	// The base64-encoded SHA256 RSA signature of the
	// instance identity document to authenticate with, with all newline characters
	// removed. Can be retrieved from the EC2 metadata server.
	Signature *string `pulumi:"signature"`
}

// The set of arguments for constructing a AuthBackendLogin resource.
type AuthBackendLoginArgs struct {
	// The unique name of the AWS auth backend. Defaults to
	// 'aws'.
	Backend pulumi.StringPtrInput
	// The HTTP method used in the signed IAM
	// request.
	IamHttpRequestMethod pulumi.StringPtrInput
	// The base64-encoded body of the signed
	// request.
	IamRequestBody pulumi.StringPtrInput
	// The base64-encoded, JSON serialized
	// representation of the GetCallerIdentity HTTP request headers.
	IamRequestHeaders pulumi.StringPtrInput
	// The base64-encoded HTTP URL used in the signed
	// request.
	IamRequestUrl pulumi.StringPtrInput
	// The base64-encoded EC2 instance identity document to
	// authenticate with. Can be retrieved from the EC2 metadata server.
	Identity pulumi.StringPtrInput
	// The unique nonce to be used for login requests. Can be
	// set to a user-specified value, or will contain the server-generated value
	// once a token is issued. EC2 instances can only acquire a single token until
	// the whitelist is tidied again unless they keep track of this nonce.
	Nonce pulumi.StringPtrInput
	// The PKCS#7 signature of the identity document to
	// authenticate with, with all newline characters removed. Can be retrieved from
	// the EC2 metadata server.
	Pkcs7 pulumi.StringPtrInput
	// The name of the AWS auth backend role to create tokens
	// against.
	Role pulumi.StringPtrInput
	// The base64-encoded SHA256 RSA signature of the
	// instance identity document to authenticate with, with all newline characters
	// removed. Can be retrieved from the EC2 metadata server.
	Signature pulumi.StringPtrInput
}

func (AuthBackendLoginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendLoginArgs)(nil)).Elem()
}

