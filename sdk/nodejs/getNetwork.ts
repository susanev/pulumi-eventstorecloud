// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieves data for an existing `Network` resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as eventstorecloud from "@pulumi/eventstorecloud";
 *
 * const example = eventstorecloud.getNetwork({
 *     name: "Example Network",
 *     projectId: _var.project_id,
 * });
 * export const networkCidr = example.then(example => example.cidrBlock);
 * ```
 */
export function getNetwork(args: GetNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("eventstorecloud:index/getNetwork:getNetwork", {
        "name": args.name,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetwork.
 */
export interface GetNetworkArgs {
    name: string;
    projectId: string;
}

/**
 * A collection of values returned by getNetwork.
 */
export interface GetNetworkResult {
    /**
     * Address space of the network in CIDR block notation
     */
    readonly cidrBlock: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly projectId: string;
    /**
     * Provider region in which to provision the network
     */
    readonly region: string;
    /**
     * Cloud Provider in which to provision the network.
     */
    readonly resourceProvider: string;
}

export function getNetworkOutput(args: GetNetworkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkResult> {
    return pulumi.output(args).apply(a => getNetwork(a, opts))
}

/**
 * A collection of arguments for invoking getNetwork.
 */
export interface GetNetworkOutputArgs {
    name: pulumi.Input<string>;
    projectId: pulumi.Input<string>;
}
