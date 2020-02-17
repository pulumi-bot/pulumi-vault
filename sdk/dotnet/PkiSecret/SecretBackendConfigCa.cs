// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.PkiSecret
{
    public partial class SecretBackendConfigCa : Pulumi.CustomResource
    {
        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// The key and certificate PEM bundle.
        /// </summary>
        [Output("pemBundle")]
        public Output<string> PemBundle { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendConfigCa resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendConfigCa(string name, SecretBackendConfigCaArgs args, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendConfigCa:SecretBackendConfigCa", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendConfigCa(string name, Input<string> id, SecretBackendConfigCaState? state = null, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendConfigCa:SecretBackendConfigCa", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendConfigCa resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendConfigCa Get(string name, Input<string> id, SecretBackendConfigCaState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendConfigCa(name, id, state, options);
        }
    }

    public sealed class SecretBackendConfigCaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// The key and certificate PEM bundle.
        /// </summary>
        [Input("pemBundle", required: true)]
        public Input<string> PemBundle { get; set; } = null!;

        public SecretBackendConfigCaArgs()
        {
        }
    }

    public sealed class SecretBackendConfigCaState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The key and certificate PEM bundle.
        /// </summary>
        [Input("pemBundle")]
        public Input<string>? PemBundle { get; set; }

        public SecretBackendConfigCaState()
        {
        }
    }
}
