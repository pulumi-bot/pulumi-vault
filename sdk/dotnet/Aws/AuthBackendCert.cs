// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Aws
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/aws_auth_backend_cert.html.markdown.
    /// </summary>
    public partial class AuthBackendCert : Pulumi.CustomResource
    {
        /// <summary>
        /// The  Base64 encoded AWS Public key required to
        /// verify PKCS7 signature of the EC2 instance metadata. You can find this key in
        /// the [AWS
        /// documentation](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html).
        /// </summary>
        [Output("awsPublicCert")]
        public Output<string> AwsPublicCert { get; private set; } = null!;

        /// <summary>
        /// The path the AWS auth backend being configured was
        /// mounted at.  Defaults to `aws`.
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// The name of the certificate.
        /// </summary>
        [Output("certName")]
        public Output<string> CertName { get; private set; } = null!;

        /// <summary>
        /// Either "pkcs7" or "identity", indicating the type of
        /// document which can be verified using the given certificate. Defaults to
        /// "pkcs7".
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendCert resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendCert(string name, AuthBackendCertArgs args, CustomResourceOptions? options = null)
            : base("vault:aws/authBackendCert:AuthBackendCert", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendCert(string name, Input<string> id, AuthBackendCertState? state = null, CustomResourceOptions? options = null)
            : base("vault:aws/authBackendCert:AuthBackendCert", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AuthBackendCert resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendCert Get(string name, Input<string> id, AuthBackendCertState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendCert(name, id, state, options);
        }
    }

    public sealed class AuthBackendCertArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The  Base64 encoded AWS Public key required to
        /// verify PKCS7 signature of the EC2 instance metadata. You can find this key in
        /// the [AWS
        /// documentation](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html).
        /// </summary>
        [Input("awsPublicCert", required: true)]
        public Input<string> AwsPublicCert { get; set; } = null!;

        /// <summary>
        /// The path the AWS auth backend being configured was
        /// mounted at.  Defaults to `aws`.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The name of the certificate.
        /// </summary>
        [Input("certName", required: true)]
        public Input<string> CertName { get; set; } = null!;

        /// <summary>
        /// Either "pkcs7" or "identity", indicating the type of
        /// document which can be verified using the given certificate. Defaults to
        /// "pkcs7".
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AuthBackendCertArgs()
        {
        }
    }

    public sealed class AuthBackendCertState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The  Base64 encoded AWS Public key required to
        /// verify PKCS7 signature of the EC2 instance metadata. You can find this key in
        /// the [AWS
        /// documentation](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html).
        /// </summary>
        [Input("awsPublicCert")]
        public Input<string>? AwsPublicCert { get; set; }

        /// <summary>
        /// The path the AWS auth backend being configured was
        /// mounted at.  Defaults to `aws`.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The name of the certificate.
        /// </summary>
        [Input("certName")]
        public Input<string>? CertName { get; set; }

        /// <summary>
        /// Either "pkcs7" or "identity", indicating the type of
        /// document which can be verified using the given certificate. Defaults to
        /// "pkcs7".
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AuthBackendCertState()
        {
        }
    }
}
