// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class CustomDomain extends pulumi.CustomResource {
    /**
     * Get an existing CustomDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomDomain {
        return new CustomDomain(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'render:services:CustomDomain';

    /**
     * Returns true if the given object is an instance of CustomDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomDomain.__pulumiType;
    }

    public /*out*/ readonly createdAt!: pulumi.Output<string | undefined>;
    public /*out*/ readonly domainType!: pulumi.Output<enums.services.DomainType>;
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly publicSuffix!: pulumi.Output<string | undefined>;
    public /*out*/ readonly redirectForName!: pulumi.Output<string>;
    public /*out*/ readonly server!: pulumi.Output<outputs.services.ServerProperties>;
    public /*out*/ readonly verificationStatus!: pulumi.Output<enums.services.VerificationStatus>;

    /**
     * Create a CustomDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomDomainArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["domainType"] = undefined /*out*/;
            resourceInputs["publicSuffix"] = undefined /*out*/;
            resourceInputs["redirectForName"] = undefined /*out*/;
            resourceInputs["server"] = undefined /*out*/;
            resourceInputs["verificationStatus"] = undefined /*out*/;
        } else {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["domainType"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["publicSuffix"] = undefined /*out*/;
            resourceInputs["redirectForName"] = undefined /*out*/;
            resourceInputs["server"] = undefined /*out*/;
            resourceInputs["verificationStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomDomain resource.
 */
export interface CustomDomainArgs {
    name: pulumi.Input<string>;
    /**
     * (Required) The ID of the service
     */
    serviceId?: pulumi.Input<string>;
}
