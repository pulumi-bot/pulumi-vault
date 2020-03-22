// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.PkiSecret
{
    /// <summary>
    /// Allows setting the issuing certificate endpoints, CRL distribution points, and OCSP server endpoints that will be encoded into issued certificates.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_config_urls.html.md.
    /// </summary>
    public partial class SecretBackendConfigUrls : Pulumi.CustomResource
    {
        /// <summary>
        /// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Specifies the URL values for the CRL Distribution Points field.
        /// </summary>
        [Output("crlDistributionPoints")]
        public Output<ImmutableArray<string>> CrlDistributionPoints { get; private set; } = null!;

        /// <summary>
        /// Specifies the URL values for the Issuing Certificate field.
        /// </summary>
        [Output("issuingCertificates")]
        public Output<ImmutableArray<string>> IssuingCertificates { get; private set; } = null!;

        /// <summary>
        /// Specifies the URL values for the OCSP Servers field.
        /// </summary>
        [Output("ocspServers")]
        public Output<ImmutableArray<string>> OcspServers { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendConfigUrls resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendConfigUrls(string name, SecretBackendConfigUrlsArgs args, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendConfigUrls(string name, Input<string> id, SecretBackendConfigUrlsState? state = null, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendConfigUrls resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendConfigUrls Get(string name, Input<string> id, SecretBackendConfigUrlsState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendConfigUrls(name, id, state, options);
        }
    }

    public sealed class SecretBackendConfigUrlsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        [Input("crlDistributionPoints")]
        private InputList<string>? _crlDistributionPoints;

        /// <summary>
        /// Specifies the URL values for the CRL Distribution Points field.
        /// </summary>
        public InputList<string> CrlDistributionPoints
        {
            get => _crlDistributionPoints ?? (_crlDistributionPoints = new InputList<string>());
            set => _crlDistributionPoints = value;
        }

        [Input("issuingCertificates")]
        private InputList<string>? _issuingCertificates;

        /// <summary>
        /// Specifies the URL values for the Issuing Certificate field.
        /// </summary>
        public InputList<string> IssuingCertificates
        {
            get => _issuingCertificates ?? (_issuingCertificates = new InputList<string>());
            set => _issuingCertificates = value;
        }

        [Input("ocspServers")]
        private InputList<string>? _ocspServers;

        /// <summary>
        /// Specifies the URL values for the OCSP Servers field.
        /// </summary>
        public InputList<string> OcspServers
        {
            get => _ocspServers ?? (_ocspServers = new InputList<string>());
            set => _ocspServers = value;
        }

        public SecretBackendConfigUrlsArgs()
        {
        }
    }

    public sealed class SecretBackendConfigUrlsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("crlDistributionPoints")]
        private InputList<string>? _crlDistributionPoints;

        /// <summary>
        /// Specifies the URL values for the CRL Distribution Points field.
        /// </summary>
        public InputList<string> CrlDistributionPoints
        {
            get => _crlDistributionPoints ?? (_crlDistributionPoints = new InputList<string>());
            set => _crlDistributionPoints = value;
        }

        [Input("issuingCertificates")]
        private InputList<string>? _issuingCertificates;

        /// <summary>
        /// Specifies the URL values for the Issuing Certificate field.
        /// </summary>
        public InputList<string> IssuingCertificates
        {
            get => _issuingCertificates ?? (_issuingCertificates = new InputList<string>());
            set => _issuingCertificates = value;
        }

        [Input("ocspServers")]
        private InputList<string>? _ocspServers;

        /// <summary>
        /// Specifies the URL values for the OCSP Servers field.
        /// </summary>
        public InputList<string> OcspServers
        {
            get => _ocspServers ?? (_ocspServers = new InputList<string>());
            set => _ocspServers = value;
        }

        public SecretBackendConfigUrlsState()
        {
        }
    }
}
