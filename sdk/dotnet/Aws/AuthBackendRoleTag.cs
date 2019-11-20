// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Aws
{
    /// <summary>
    /// Reads role tag information from an AWS auth backend in Vault. 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/aws_auth_backend_role_tag.html.markdown.
    /// </summary>
    public partial class AuthBackendRoleTag : Pulumi.CustomResource
    {
        /// <summary>
        /// If set, allows migration of the underlying instances where the client resides. Use with caution.
        /// </summary>
        [Output("allowInstanceMigration")]
        public Output<bool?> AllowInstanceMigration { get; private set; } = null!;

        /// <summary>
        /// The path to the AWS auth backend to
        /// read role tags from, with no leading or trailing `/`s. Defaults to "aws".
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// If set, only allows a single token to be granted per instance ID.
        /// </summary>
        [Output("disallowReauthentication")]
        public Output<bool?> DisallowReauthentication { get; private set; } = null!;

        /// <summary>
        /// Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string?> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The maximum TTL of the tokens issued using this role.
        /// </summary>
        [Output("maxTtl")]
        public Output<string?> MaxTtl { get; private set; } = null!;

        /// <summary>
        /// The policies to be associated with the tag. Must be a subset of the policies associated with the role.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;

        /// <summary>
        /// The name of the AWS auth backend role to read
        /// role tags from, with no leading or trailing `/`s.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The key of the role tag.
        /// </summary>
        [Output("tagKey")]
        public Output<string> TagKey { get; private set; } = null!;

        /// <summary>
        /// The value to set the role key.
        /// </summary>
        [Output("tagValue")]
        public Output<string> TagValue { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendRoleTag resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendRoleTag(string name, AuthBackendRoleTagArgs args, CustomResourceOptions? options = null)
            : base("vault:aws/authBackendRoleTag:AuthBackendRoleTag", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendRoleTag(string name, Input<string> id, AuthBackendRoleTagState? state = null, CustomResourceOptions? options = null)
            : base("vault:aws/authBackendRoleTag:AuthBackendRoleTag", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthBackendRoleTag resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendRoleTag Get(string name, Input<string> id, AuthBackendRoleTagState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendRoleTag(name, id, state, options);
        }
    }

    public sealed class AuthBackendRoleTagArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, allows migration of the underlying instances where the client resides. Use with caution.
        /// </summary>
        [Input("allowInstanceMigration")]
        public Input<bool>? AllowInstanceMigration { get; set; }

        /// <summary>
        /// The path to the AWS auth backend to
        /// read role tags from, with no leading or trailing `/`s. Defaults to "aws".
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// If set, only allows a single token to be granted per instance ID.
        /// </summary>
        [Input("disallowReauthentication")]
        public Input<bool>? DisallowReauthentication { get; set; }

        /// <summary>
        /// Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The maximum TTL of the tokens issued using this role.
        /// </summary>
        [Input("maxTtl")]
        public Input<string>? MaxTtl { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// The policies to be associated with the tag. Must be a subset of the policies associated with the role.
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// The name of the AWS auth backend role to read
        /// role tags from, with no leading or trailing `/`s.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public AuthBackendRoleTagArgs()
        {
        }
    }

    public sealed class AuthBackendRoleTagState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, allows migration of the underlying instances where the client resides. Use with caution.
        /// </summary>
        [Input("allowInstanceMigration")]
        public Input<bool>? AllowInstanceMigration { get; set; }

        /// <summary>
        /// The path to the AWS auth backend to
        /// read role tags from, with no leading or trailing `/`s. Defaults to "aws".
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// If set, only allows a single token to be granted per instance ID.
        /// </summary>
        [Input("disallowReauthentication")]
        public Input<bool>? DisallowReauthentication { get; set; }

        /// <summary>
        /// Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The maximum TTL of the tokens issued using this role.
        /// </summary>
        [Input("maxTtl")]
        public Input<string>? MaxTtl { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// The policies to be associated with the tag. Must be a subset of the policies associated with the role.
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// The name of the AWS auth backend role to read
        /// role tags from, with no leading or trailing `/`s.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The key of the role tag.
        /// </summary>
        [Input("tagKey")]
        public Input<string>? TagKey { get; set; }

        /// <summary>
        /// The value to set the role key.
        /// </summary>
        [Input("tagValue")]
        public Input<string>? TagValue { get; set; }

        public AuthBackendRoleTagState()
        {
        }
    }
}
