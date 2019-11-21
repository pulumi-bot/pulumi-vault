// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Consul
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/consul_secret_backend.html.markdown.
    /// </summary>
    public partial class SecretBackend : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// The default TTL for credentials issued by this backend.
        /// </summary>
        [Output("defaultLeaseTtlSeconds")]
        public Output<int?> DefaultLeaseTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// A human-friendly description for this backend.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The maximum TTL that can be requested
        /// for credentials issued by this backend.
        /// </summary>
        [Output("maxLeaseTtlSeconds")]
        public Output<int?> MaxLeaseTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// Specifies the URL scheme to use. Defaults to `http`.
        /// </summary>
        [Output("scheme")]
        public Output<string?> Scheme { get; private set; } = null!;

        /// <summary>
        /// The Consul management token this backend should use to issue new tokens.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackend(string name, SecretBackendArgs args, CustomResourceOptions? options = null)
            : base("vault:consul/secretBackend:SecretBackend", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SecretBackend(string name, Input<string> id, SecretBackendState? state = null, CustomResourceOptions? options = null)
            : base("vault:consul/secretBackend:SecretBackend", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackend resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackend Get(string name, Input<string> id, SecretBackendState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackend(name, id, state, options);
        }
    }

    public sealed class SecretBackendArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// The default TTL for credentials issued by this backend.
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        /// <summary>
        /// A human-friendly description for this backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The maximum TTL that can be requested
        /// for credentials issued by this backend.
        /// </summary>
        [Input("maxLeaseTtlSeconds")]
        public Input<int>? MaxLeaseTtlSeconds { get; set; }

        /// <summary>
        /// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Specifies the URL scheme to use. Defaults to `http`.
        /// </summary>
        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        /// <summary>
        /// The Consul management token this backend should use to issue new tokens.
        /// </summary>
        [Input("token", required: true)]
        public Input<string> Token { get; set; } = null!;

        public SecretBackendArgs()
        {
        }
    }

    public sealed class SecretBackendState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The default TTL for credentials issued by this backend.
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        /// <summary>
        /// A human-friendly description for this backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The maximum TTL that can be requested
        /// for credentials issued by this backend.
        /// </summary>
        [Input("maxLeaseTtlSeconds")]
        public Input<int>? MaxLeaseTtlSeconds { get; set; }

        /// <summary>
        /// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Specifies the URL scheme to use. Defaults to `http`.
        /// </summary>
        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        /// <summary>
        /// The Consul management token this backend should use to issue new tokens.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public SecretBackendState()
        {
        }
    }
}
