// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Aws
{
    public static partial class Invokes
    {
        /// <summary>
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/aws_access_credentials.html.markdown.
        /// </summary>
        public static Task<GetAccessCredentialsResult> GetAccessCredentials(GetAccessCredentialsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccessCredentialsResult>("vault:aws/getAccessCredentials:getAccessCredentials", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetAccessCredentialsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path to the AWS secret backend to
        /// read credentials from, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// The name of the AWS secret backend role to read
        /// credentials from, with no leading or trailing `/`s.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The type of credentials to read. Defaults
        /// to `"creds"`, which just returns an AWS Access Key ID and Secret
        /// Key. Can also be set to `"sts"`, which will return a security token
        /// in addition to the keys.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetAccessCredentialsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetAccessCredentialsResult
    {
        /// <summary>
        /// The AWS Access Key ID returned by Vault.
        /// </summary>
        public readonly string AccessKey;
        public readonly string Backend;
        /// <summary>
        /// The duration of the secret lease, in seconds relative
        /// to the time the data was requested. Once this time has passed any plan
        /// generated with this data may fail to apply.
        /// </summary>
        public readonly int LeaseDuration;
        /// <summary>
        /// The lease identifier assigned by Vault.
        /// </summary>
        public readonly string LeaseId;
        public readonly bool LeaseRenewable;
        public readonly string LeaseStartTime;
        public readonly string Role;
        /// <summary>
        /// The AWS Secret Key returned by Vault.
        /// </summary>
        public readonly string SecretKey;
        /// <summary>
        /// The STS token returned by Vault, if any.
        /// </summary>
        public readonly string SecurityToken;
        public readonly string? Type;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAccessCredentialsResult(
            string accessKey,
            string backend,
            int leaseDuration,
            string leaseId,
            bool leaseRenewable,
            string leaseStartTime,
            string role,
            string secretKey,
            string securityToken,
            string? type,
            string id)
        {
            AccessKey = accessKey;
            Backend = backend;
            LeaseDuration = leaseDuration;
            LeaseId = leaseId;
            LeaseRenewable = leaseRenewable;
            LeaseStartTime = leaseStartTime;
            Role = role;
            SecretKey = secretKey;
            SecurityToken = securityToken;
            Type = type;
            Id = id;
        }
    }
}
