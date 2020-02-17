// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Identity
{
    public static partial class Invokes
    {
        public static Task<GetEntityResult> GetEntity(GetEntityArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEntityResult>("vault:identity/getEntity:getEntity", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetEntityArgs : Pulumi.InvokeArgs
    {
        [Input("aliasId")]
        public string? AliasId { get; set; }

        [Input("aliasMountAccessor")]
        public string? AliasMountAccessor { get; set; }

        [Input("aliasName")]
        public string? AliasName { get; set; }

        [Input("entityId")]
        public string? EntityId { get; set; }

        [Input("entityName")]
        public string? EntityName { get; set; }

        public GetEntityArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetEntityResult
    {
        public readonly string AliasId;
        public readonly string AliasMountAccessor;
        public readonly string AliasName;
        public readonly ImmutableArray<Outputs.GetEntityAliasesResult> Aliases;
        public readonly string CreationTime;
        public readonly string DataJson;
        public readonly ImmutableArray<string> DirectGroupIds;
        public readonly bool Disabled;
        public readonly string EntityId;
        public readonly string EntityName;
        public readonly ImmutableArray<string> GroupIds;
        public readonly ImmutableArray<string> InheritedGroupIds;
        public readonly string LastUpdateTime;
        public readonly ImmutableArray<string> MergedEntityIds;
        public readonly ImmutableDictionary<string, object> Metadata;
        public readonly string NamespaceId;
        public readonly ImmutableArray<string> Policies;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetEntityResult(
            string aliasId,
            string aliasMountAccessor,
            string aliasName,
            ImmutableArray<Outputs.GetEntityAliasesResult> aliases,
            string creationTime,
            string dataJson,
            ImmutableArray<string> directGroupIds,
            bool disabled,
            string entityId,
            string entityName,
            ImmutableArray<string> groupIds,
            ImmutableArray<string> inheritedGroupIds,
            string lastUpdateTime,
            ImmutableArray<string> mergedEntityIds,
            ImmutableDictionary<string, object> metadata,
            string namespaceId,
            ImmutableArray<string> policies,
            string id)
        {
            AliasId = aliasId;
            AliasMountAccessor = aliasMountAccessor;
            AliasName = aliasName;
            Aliases = aliases;
            CreationTime = creationTime;
            DataJson = dataJson;
            DirectGroupIds = directGroupIds;
            Disabled = disabled;
            EntityId = entityId;
            EntityName = entityName;
            GroupIds = groupIds;
            InheritedGroupIds = inheritedGroupIds;
            LastUpdateTime = lastUpdateTime;
            MergedEntityIds = mergedEntityIds;
            Metadata = metadata;
            NamespaceId = namespaceId;
            Policies = policies;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetEntityAliasesResult
    {
        public readonly string CanonicalId;
        public readonly string CreationTime;
        public readonly string Id;
        public readonly string LastUpdateTime;
        public readonly ImmutableArray<string> MergedFromCanonicalIds;
        public readonly ImmutableDictionary<string, object> Metadata;
        public readonly string MountAccessor;
        public readonly string MountPath;
        public readonly string MountType;
        public readonly string Name;

        [OutputConstructor]
        private GetEntityAliasesResult(
            string canonicalId,
            string creationTime,
            string id,
            string lastUpdateTime,
            ImmutableArray<string> mergedFromCanonicalIds,
            ImmutableDictionary<string, object> metadata,
            string mountAccessor,
            string mountPath,
            string mountType,
            string name)
        {
            CanonicalId = canonicalId;
            CreationTime = creationTime;
            Id = id;
            LastUpdateTime = lastUpdateTime;
            MergedFromCanonicalIds = mergedFromCanonicalIds;
            Metadata = metadata;
            MountAccessor = mountAccessor;
            MountPath = mountPath;
            MountType = mountType;
            Name = name;
        }
    }
    }
}
