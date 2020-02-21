// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Generic
{
    public partial class Endpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// String containing a JSON-encoded object that will be
        /// written to the given path as the secret data.
        /// </summary>
        [Output("dataJson")]
        public Output<string> DataJson { get; private set; } = null!;

        /// <summary>
        /// Don't attempt to delete the path from Vault if true
        /// </summary>
        [Output("disableDelete")]
        public Output<bool?> DisableDelete { get; private set; } = null!;

        /// <summary>
        /// True/false. Set this to true if your vault
        /// authentication is not able to read the data or if the endpoint does
        /// not support the `GET` method. Setting this to `true` will break drift
        /// detection. You should set this to `true` for endpoints that are
        /// write-only. Defaults to false.
        /// </summary>
        [Output("disableRead")]
        public Output<bool?> DisableRead { get; private set; } = null!;

        /// <summary>
        /// When reading, disregard fields not present in data_json
        /// </summary>
        [Output("ignoreAbsentFields")]
        public Output<bool?> IgnoreAbsentFields { get; private set; } = null!;

        /// <summary>
        /// The full logical path at which to write the given
        /// data. Consult each backend's documentation to see which endpoints
        /// support the `PUT` methods and to determine whether they also support
        /// `DELETE` and `GET`.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// Map of strings returned by write operation
        /// </summary>
        [Output("writeData")]
        public Output<ImmutableDictionary<string, string>> WriteData { get; private set; } = null!;

        /// <summary>
        /// JSON data returned by write operation
        /// </summary>
        [Output("writeDataJson")]
        public Output<string> WriteDataJson { get; private set; } = null!;

        /// <summary>
        /// Top-level fields returned by write to persist in state
        /// </summary>
        [Output("writeFields")]
        public Output<ImmutableArray<string>> WriteFields { get; private set; } = null!;


        /// <summary>
        /// Create a Endpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Endpoint(string name, EndpointArgs args, CustomResourceOptions? options = null)
            : base("vault:generic/endpoint:Endpoint", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Endpoint(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
            : base("vault:generic/endpoint:Endpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Endpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Endpoint Get(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new Endpoint(name, id, state, options);
        }
    }

    public sealed class EndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// String containing a JSON-encoded object that will be
        /// written to the given path as the secret data.
        /// </summary>
        [Input("dataJson", required: true)]
        public Input<string> DataJson { get; set; } = null!;

        /// <summary>
        /// Don't attempt to delete the path from Vault if true
        /// </summary>
        [Input("disableDelete")]
        public Input<bool>? DisableDelete { get; set; }

        /// <summary>
        /// True/false. Set this to true if your vault
        /// authentication is not able to read the data or if the endpoint does
        /// not support the `GET` method. Setting this to `true` will break drift
        /// detection. You should set this to `true` for endpoints that are
        /// write-only. Defaults to false.
        /// </summary>
        [Input("disableRead")]
        public Input<bool>? DisableRead { get; set; }

        /// <summary>
        /// When reading, disregard fields not present in data_json
        /// </summary>
        [Input("ignoreAbsentFields")]
        public Input<bool>? IgnoreAbsentFields { get; set; }

        /// <summary>
        /// The full logical path at which to write the given
        /// data. Consult each backend's documentation to see which endpoints
        /// support the `PUT` methods and to determine whether they also support
        /// `DELETE` and `GET`.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("writeFields")]
        private InputList<string>? _writeFields;

        /// <summary>
        /// Top-level fields returned by write to persist in state
        /// </summary>
        public InputList<string> WriteFields
        {
            get => _writeFields ?? (_writeFields = new InputList<string>());
            set => _writeFields = value;
        }

        public EndpointArgs()
        {
        }
    }

    public sealed class EndpointState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// String containing a JSON-encoded object that will be
        /// written to the given path as the secret data.
        /// </summary>
        [Input("dataJson")]
        public Input<string>? DataJson { get; set; }

        /// <summary>
        /// Don't attempt to delete the path from Vault if true
        /// </summary>
        [Input("disableDelete")]
        public Input<bool>? DisableDelete { get; set; }

        /// <summary>
        /// True/false. Set this to true if your vault
        /// authentication is not able to read the data or if the endpoint does
        /// not support the `GET` method. Setting this to `true` will break drift
        /// detection. You should set this to `true` for endpoints that are
        /// write-only. Defaults to false.
        /// </summary>
        [Input("disableRead")]
        public Input<bool>? DisableRead { get; set; }

        /// <summary>
        /// When reading, disregard fields not present in data_json
        /// </summary>
        [Input("ignoreAbsentFields")]
        public Input<bool>? IgnoreAbsentFields { get; set; }

        /// <summary>
        /// The full logical path at which to write the given
        /// data. Consult each backend's documentation to see which endpoints
        /// support the `PUT` methods and to determine whether they also support
        /// `DELETE` and `GET`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("writeData")]
        private InputMap<string>? _writeData;

        /// <summary>
        /// Map of strings returned by write operation
        /// </summary>
        public InputMap<string> WriteData
        {
            get => _writeData ?? (_writeData = new InputMap<string>());
            set => _writeData = value;
        }

        /// <summary>
        /// JSON data returned by write operation
        /// </summary>
        [Input("writeDataJson")]
        public Input<string>? WriteDataJson { get; set; }

        [Input("writeFields")]
        private InputList<string>? _writeFields;

        /// <summary>
        /// Top-level fields returned by write to persist in state
        /// </summary>
        public InputList<string> WriteFields
        {
            get => _writeFields ?? (_writeFields = new InputList<string>());
            set => _writeFields = value;
        }

        public EndpointState()
        {
        }
    }
}
