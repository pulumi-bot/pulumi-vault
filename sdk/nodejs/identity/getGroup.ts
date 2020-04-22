// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getGroup(args?: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("vault:identity/getGroup:getGroup", {
        "aliasId": args.aliasId,
        "aliasMountAccessor": args.aliasMountAccessor,
        "aliasName": args.aliasName,
        "groupId": args.groupId,
        "groupName": args.groupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    /**
     * ID of the alias.
     */
    readonly aliasId?: string;
    /**
     * Accessor of the mount to which the alias belongs to.
     * This should be supplied in conjunction with `aliasName`.
     */
    readonly aliasMountAccessor?: string;
    /**
     * Name of the alias. This should be supplied in conjunction with
     * `aliasMountAccessor`.
     */
    readonly aliasName?: string;
    /**
     * ID of the group.
     */
    readonly groupId?: string;
    /**
     * Name of the group.
     */
    readonly groupName?: string;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    /**
     * Canonical ID of the Alias
     */
    readonly aliasCanonicalId: string;
    /**
     * Creation time of the Alias
     */
    readonly aliasCreationTime: string;
    readonly aliasId: string;
    /**
     * Last update time of the alias
     */
    readonly aliasLastUpdateTime: string;
    /**
     * List of canonical IDs merged with this alias
     */
    readonly aliasMergedFromCanonicalIds: string[];
    /**
     * Arbitrary metadata
     */
    readonly aliasMetadata: {[key: string]: any};
    readonly aliasMountAccessor: string;
    /**
     * Authentication mount path which this alias belongs to
     */
    readonly aliasMountPath: string;
    /**
     * Authentication mount type which this alias belongs to
     */
    readonly aliasMountType: string;
    readonly aliasName: string;
    /**
     * Creation timestamp of the group
     */
    readonly creationTime: string;
    /**
     * A string containing the full data payload retrieved from
     * Vault, serialized in JSON format.
     */
    readonly dataJson: string;
    readonly groupId: string;
    readonly groupName: string;
    /**
     * Last updated time of the group
     */
    readonly lastUpdateTime: string;
    /**
     * List of Entity IDs which are members of this group
     */
    readonly memberEntityIds: string[];
    /**
     * List of Group IDs which are members of this group
     */
    readonly memberGroupIds: string[];
    /**
     * Arbitrary metadata
     */
    readonly metadata: {[key: string]: any};
    /**
     * Modify index of the group
     */
    readonly modifyIndex: number;
    /**
     * Namespace of which the group is part of
     */
    readonly namespaceId: string;
    /**
     * List of Group IDs which are parents of this group.
     */
    readonly parentGroupIds: string[];
    /**
     * List of policies attached to the group
     */
    readonly policies: string[];
    /**
     * Type of group
     */
    readonly type: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
