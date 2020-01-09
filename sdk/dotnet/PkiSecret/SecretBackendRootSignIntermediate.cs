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
    /// Creates an PKI certificate.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_root_sign_intermediate.html.markdown.
    /// </summary>
    public partial class SecretBackendRootSignIntermediate : Pulumi.CustomResource
    {
        /// <summary>
        /// List of alternative names
        /// </summary>
        [Output("altNames")]
        public Output<ImmutableArray<string>> AltNames { get; private set; } = null!;

        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// The CA chain
        /// </summary>
        [Output("caChain")]
        public Output<string> CaChain { get; private set; } = null!;

        /// <summary>
        /// The certificate
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// CN of intermediate to create
        /// </summary>
        [Output("commonName")]
        public Output<string> CommonName { get; private set; } = null!;

        /// <summary>
        /// The country
        /// </summary>
        [Output("country")]
        public Output<string?> Country { get; private set; } = null!;

        /// <summary>
        /// The CSR
        /// </summary>
        [Output("csr")]
        public Output<string> Csr { get; private set; } = null!;

        /// <summary>
        /// Flag to exclude CN from SANs
        /// </summary>
        [Output("excludeCnFromSans")]
        public Output<bool?> ExcludeCnFromSans { get; private set; } = null!;

        /// <summary>
        /// The format of data
        /// </summary>
        [Output("format")]
        public Output<string?> Format { get; private set; } = null!;

        /// <summary>
        /// List of alternative IPs
        /// </summary>
        [Output("ipSans")]
        public Output<ImmutableArray<string>> IpSans { get; private set; } = null!;

        /// <summary>
        /// The issuing CA
        /// </summary>
        [Output("issuingCa")]
        public Output<string> IssuingCa { get; private set; } = null!;

        /// <summary>
        /// The locality
        /// </summary>
        [Output("locality")]
        public Output<string?> Locality { get; private set; } = null!;

        /// <summary>
        /// The maximum path length to encode in the generated certificate
        /// </summary>
        [Output("maxPathLength")]
        public Output<int?> MaxPathLength { get; private set; } = null!;

        /// <summary>
        /// The organization
        /// </summary>
        [Output("organization")]
        public Output<string?> Organization { get; private set; } = null!;

        /// <summary>
        /// List of other SANs
        /// </summary>
        [Output("otherSans")]
        public Output<ImmutableArray<string>> OtherSans { get; private set; } = null!;

        /// <summary>
        /// The organization unit
        /// </summary>
        [Output("ou")]
        public Output<string?> Ou { get; private set; } = null!;

        /// <summary>
        /// List of domains for which certificates are allowed to be issued
        /// </summary>
        [Output("permittedDnsDomains")]
        public Output<ImmutableArray<string>> PermittedDnsDomains { get; private set; } = null!;

        /// <summary>
        /// The postal code
        /// </summary>
        [Output("postalCode")]
        public Output<string?> PostalCode { get; private set; } = null!;

        /// <summary>
        /// The province
        /// </summary>
        [Output("province")]
        public Output<string?> Province { get; private set; } = null!;

        /// <summary>
        /// The serial
        /// </summary>
        [Output("serial")]
        public Output<string> Serial { get; private set; } = null!;

        /// <summary>
        /// The street address
        /// </summary>
        [Output("streetAddress")]
        public Output<string?> StreetAddress { get; private set; } = null!;

        /// <summary>
        /// Time to live
        /// </summary>
        [Output("ttl")]
        public Output<string?> Ttl { get; private set; } = null!;

        /// <summary>
        /// List of alternative URIs
        /// </summary>
        [Output("uriSans")]
        public Output<ImmutableArray<string>> UriSans { get; private set; } = null!;

        /// <summary>
        /// Preserve CSR values
        /// </summary>
        [Output("useCsrValues")]
        public Output<bool?> UseCsrValues { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendRootSignIntermediate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendRootSignIntermediate(string name, SecretBackendRootSignIntermediateArgs args, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendRootSignIntermediate:SecretBackendRootSignIntermediate", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendRootSignIntermediate(string name, Input<string> id, SecretBackendRootSignIntermediateState? state = null, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendRootSignIntermediate:SecretBackendRootSignIntermediate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendRootSignIntermediate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendRootSignIntermediate Get(string name, Input<string> id, SecretBackendRootSignIntermediateState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendRootSignIntermediate(name, id, state, options);
        }
    }

    public sealed class SecretBackendRootSignIntermediateArgs : Pulumi.ResourceArgs
    {
        [Input("altNames")]
        private InputList<string>? _altNames;

        /// <summary>
        /// List of alternative names
        /// </summary>
        public InputList<string> AltNames
        {
            get => _altNames ?? (_altNames = new InputList<string>());
            set => _altNames = value;
        }

        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// CN of intermediate to create
        /// </summary>
        [Input("commonName", required: true)]
        public Input<string> CommonName { get; set; } = null!;

        /// <summary>
        /// The country
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// The CSR
        /// </summary>
        [Input("csr", required: true)]
        public Input<string> Csr { get; set; } = null!;

        /// <summary>
        /// Flag to exclude CN from SANs
        /// </summary>
        [Input("excludeCnFromSans")]
        public Input<bool>? ExcludeCnFromSans { get; set; }

        /// <summary>
        /// The format of data
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        [Input("ipSans")]
        private InputList<string>? _ipSans;

        /// <summary>
        /// List of alternative IPs
        /// </summary>
        public InputList<string> IpSans
        {
            get => _ipSans ?? (_ipSans = new InputList<string>());
            set => _ipSans = value;
        }

        /// <summary>
        /// The locality
        /// </summary>
        [Input("locality")]
        public Input<string>? Locality { get; set; }

        /// <summary>
        /// The maximum path length to encode in the generated certificate
        /// </summary>
        [Input("maxPathLength")]
        public Input<int>? MaxPathLength { get; set; }

        /// <summary>
        /// The organization
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        [Input("otherSans")]
        private InputList<string>? _otherSans;

        /// <summary>
        /// List of other SANs
        /// </summary>
        public InputList<string> OtherSans
        {
            get => _otherSans ?? (_otherSans = new InputList<string>());
            set => _otherSans = value;
        }

        /// <summary>
        /// The organization unit
        /// </summary>
        [Input("ou")]
        public Input<string>? Ou { get; set; }

        [Input("permittedDnsDomains")]
        private InputList<string>? _permittedDnsDomains;

        /// <summary>
        /// List of domains for which certificates are allowed to be issued
        /// </summary>
        public InputList<string> PermittedDnsDomains
        {
            get => _permittedDnsDomains ?? (_permittedDnsDomains = new InputList<string>());
            set => _permittedDnsDomains = value;
        }

        /// <summary>
        /// The postal code
        /// </summary>
        [Input("postalCode")]
        public Input<string>? PostalCode { get; set; }

        /// <summary>
        /// The province
        /// </summary>
        [Input("province")]
        public Input<string>? Province { get; set; }

        /// <summary>
        /// The street address
        /// </summary>
        [Input("streetAddress")]
        public Input<string>? StreetAddress { get; set; }

        /// <summary>
        /// Time to live
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        [Input("uriSans")]
        private InputList<string>? _uriSans;

        /// <summary>
        /// List of alternative URIs
        /// </summary>
        public InputList<string> UriSans
        {
            get => _uriSans ?? (_uriSans = new InputList<string>());
            set => _uriSans = value;
        }

        /// <summary>
        /// Preserve CSR values
        /// </summary>
        [Input("useCsrValues")]
        public Input<bool>? UseCsrValues { get; set; }

        public SecretBackendRootSignIntermediateArgs()
        {
        }
    }

    public sealed class SecretBackendRootSignIntermediateState : Pulumi.ResourceArgs
    {
        [Input("altNames")]
        private InputList<string>? _altNames;

        /// <summary>
        /// List of alternative names
        /// </summary>
        public InputList<string> AltNames
        {
            get => _altNames ?? (_altNames = new InputList<string>());
            set => _altNames = value;
        }

        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The CA chain
        /// </summary>
        [Input("caChain")]
        public Input<string>? CaChain { get; set; }

        /// <summary>
        /// The certificate
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// CN of intermediate to create
        /// </summary>
        [Input("commonName")]
        public Input<string>? CommonName { get; set; }

        /// <summary>
        /// The country
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// The CSR
        /// </summary>
        [Input("csr")]
        public Input<string>? Csr { get; set; }

        /// <summary>
        /// Flag to exclude CN from SANs
        /// </summary>
        [Input("excludeCnFromSans")]
        public Input<bool>? ExcludeCnFromSans { get; set; }

        /// <summary>
        /// The format of data
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        [Input("ipSans")]
        private InputList<string>? _ipSans;

        /// <summary>
        /// List of alternative IPs
        /// </summary>
        public InputList<string> IpSans
        {
            get => _ipSans ?? (_ipSans = new InputList<string>());
            set => _ipSans = value;
        }

        /// <summary>
        /// The issuing CA
        /// </summary>
        [Input("issuingCa")]
        public Input<string>? IssuingCa { get; set; }

        /// <summary>
        /// The locality
        /// </summary>
        [Input("locality")]
        public Input<string>? Locality { get; set; }

        /// <summary>
        /// The maximum path length to encode in the generated certificate
        /// </summary>
        [Input("maxPathLength")]
        public Input<int>? MaxPathLength { get; set; }

        /// <summary>
        /// The organization
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        [Input("otherSans")]
        private InputList<string>? _otherSans;

        /// <summary>
        /// List of other SANs
        /// </summary>
        public InputList<string> OtherSans
        {
            get => _otherSans ?? (_otherSans = new InputList<string>());
            set => _otherSans = value;
        }

        /// <summary>
        /// The organization unit
        /// </summary>
        [Input("ou")]
        public Input<string>? Ou { get; set; }

        [Input("permittedDnsDomains")]
        private InputList<string>? _permittedDnsDomains;

        /// <summary>
        /// List of domains for which certificates are allowed to be issued
        /// </summary>
        public InputList<string> PermittedDnsDomains
        {
            get => _permittedDnsDomains ?? (_permittedDnsDomains = new InputList<string>());
            set => _permittedDnsDomains = value;
        }

        /// <summary>
        /// The postal code
        /// </summary>
        [Input("postalCode")]
        public Input<string>? PostalCode { get; set; }

        /// <summary>
        /// The province
        /// </summary>
        [Input("province")]
        public Input<string>? Province { get; set; }

        /// <summary>
        /// The serial
        /// </summary>
        [Input("serial")]
        public Input<string>? Serial { get; set; }

        /// <summary>
        /// The street address
        /// </summary>
        [Input("streetAddress")]
        public Input<string>? StreetAddress { get; set; }

        /// <summary>
        /// Time to live
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        [Input("uriSans")]
        private InputList<string>? _uriSans;

        /// <summary>
        /// List of alternative URIs
        /// </summary>
        public InputList<string> UriSans
        {
            get => _uriSans ?? (_uriSans = new InputList<string>());
            set => _uriSans = value;
        }

        /// <summary>
        /// Preserve CSR values
        /// </summary>
        [Input("useCsrValues")]
        public Input<bool>? UseCsrValues { get; set; }

        public SecretBackendRootSignIntermediateState()
        {
        }
    }
}
