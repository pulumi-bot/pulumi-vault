// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getEntity(args?: GetEntityArgs, opts?: pulumi.InvokeOptions): Promise<GetEntityResult> & GetEntityResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetEntityResult> = pulumi.runtime.invoke("vault:identity/getEntity:getEntity", {
        "aliasId": args.aliasId,
        "aliasMountAccessor": args.aliasMountAccessor,
        "aliasName": args.aliasName,
        "entityId": args.entityId,
        "entityName": args.entityName,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getEntity.
 */
export interface GetEntityArgs {
    readonly aliasId?: string;
    readonly aliasMountAccessor?: string;
    readonly aliasName?: string;
    readonly entityId?: string;
    readonly entityName?: string;
}

/**
 * A collection of values returned by getEntity.
 */
export interface GetEntityResult {
    readonly aliasId: string;
    readonly aliasMountAccessor: string;
    readonly aliasName: string;
    readonly aliases: outputs.identity.GetEntityAlias[];
    readonly creationTime: string;
    readonly dataJson: string;
    readonly directGroupIds: string[];
    readonly disabled: boolean;
    readonly entityId: string;
    readonly entityName: string;
    readonly groupIds: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly inheritedGroupIds: string[];
    readonly lastUpdateTime: string;
    readonly mergedEntityIds: string[];
    readonly metadata: {[key: string]: any};
    readonly namespaceId: string;
    readonly policies: string[];
}
