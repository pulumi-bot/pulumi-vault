// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Identity
{
    public partial class EntityAlias : Pulumi.CustomResource
    {
        /// <summary>
        /// Entity ID to which this alias belongs to.
        /// </summary>
        [Output("canonicalId")]
        public Output<string> CanonicalId { get; private set; } = null!;

        /// <summary>
        /// Accessor of the mount to which the alias should belong to.
        /// </summary>
        [Output("mountAccessor")]
        public Output<string> MountAccessor { get; private set; } = null!;

        /// <summary>
        /// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a EntityAlias resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EntityAlias(string name, EntityAliasArgs args, CustomResourceOptions? options = null)
            : base("vault:identity/entityAlias:EntityAlias", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private EntityAlias(string name, Input<string> id, EntityAliasState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/entityAlias:EntityAlias", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EntityAlias resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EntityAlias Get(string name, Input<string> id, EntityAliasState? state = null, CustomResourceOptions? options = null)
        {
            return new EntityAlias(name, id, state, options);
        }
    }

    public sealed class EntityAliasArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Entity ID to which this alias belongs to.
        /// </summary>
        [Input("canonicalId", required: true)]
        public Input<string> CanonicalId { get; set; } = null!;

        /// <summary>
        /// Accessor of the mount to which the alias should belong to.
        /// </summary>
        [Input("mountAccessor", required: true)]
        public Input<string> MountAccessor { get; set; } = null!;

        /// <summary>
        /// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public EntityAliasArgs()
        {
        }
    }

    public sealed class EntityAliasState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Entity ID to which this alias belongs to.
        /// </summary>
        [Input("canonicalId")]
        public Input<string>? CanonicalId { get; set; }

        /// <summary>
        /// Accessor of the mount to which the alias should belong to.
        /// </summary>
        [Input("mountAccessor")]
        public Input<string>? MountAccessor { get; set; }

        /// <summary>
        /// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public EntityAliasState()
        {
        }
    }
}
