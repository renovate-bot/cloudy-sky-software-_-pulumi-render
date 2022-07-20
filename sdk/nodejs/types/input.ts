// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";

import * as utilities from "../utilities";

export namespace services {
    export interface ServerPropertiesArgs {
        id?: pulumi.Input<string>;
        name?: pulumi.Input<string>;
    }

    export interface ServiceDiskArgs {
        mountPath: pulumi.Input<string>;
        name: pulumi.Input<string>;
        sizeGB?: pulumi.Input<number>;
    }
    /**
     * serviceDiskArgsProvideDefaults sets the appropriate defaults for ServiceDiskArgs
     */
    export function serviceDiskArgsProvideDefaults(val: ServiceDiskArgs): ServiceDiskArgs {
        return {
            ...val,
            sizeGB: (val.sizeGB) ?? 1,
        };
    }

    export interface ServiceDockerDetailsArgs {
        dockerCommand?: pulumi.Input<string>;
        dockerContext?: pulumi.Input<string>;
        dockerfilePath?: pulumi.Input<string>;
    }
    /**
     * serviceDockerDetailsArgsProvideDefaults sets the appropriate defaults for ServiceDockerDetailsArgs
     */
    export function serviceDockerDetailsArgsProvideDefaults(val: ServiceDockerDetailsArgs): ServiceDockerDetailsArgs {
        return {
            ...val,
            dockerfilePath: (val.dockerfilePath) ?? "./Dockerfile",
        };
    }

    export interface ServiceNativeEnvironmentDetailsArgs {
        buildCommand: pulumi.Input<string>;
        startCommand: pulumi.Input<string>;
    }

    /**
     * A service header object
     */
    export interface ServiceServiceHeaderArgs {
        name: pulumi.Input<string>;
        path: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    /**
     * A static website service
     */
    export interface ServiceStaticSiteArgs {
        buildCommand?: pulumi.Input<string>;
        headers?: pulumi.Input<pulumi.Input<inputs.services.ServiceServiceHeaderArgs>[]>;
        parentServer?: pulumi.Input<inputs.services.ServiceStaticSiteParentServerPropertiesArgs>;
        publishPath?: pulumi.Input<string>;
        pullRequestPreviewsEnabled?: pulumi.Input<enums.services.ServiceStaticSitePullRequestPreviewsEnabled>;
        routes?: pulumi.Input<pulumi.Input<inputs.services.ServiceStaticSiteRouteArgs>[]>;
        /**
         * The HTTPS service URL. A subdomain of onrender.com, by default.
         */
        url?: pulumi.Input<string>;
    }
    /**
     * serviceStaticSiteArgsProvideDefaults sets the appropriate defaults for ServiceStaticSiteArgs
     */
    export function serviceStaticSiteArgsProvideDefaults(val: ServiceStaticSiteArgs): ServiceStaticSiteArgs {
        return {
            ...val,
            publishPath: (val.publishPath) ?? "public",
            pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
        };
    }

    export interface ServiceStaticSiteParentServerPropertiesArgs {
        id?: pulumi.Input<string>;
        name?: pulumi.Input<string>;
    }

    /**
     * A route object for a static site
     */
    export interface ServiceStaticSiteRouteArgs {
        destination: pulumi.Input<string>;
        source: pulumi.Input<string>;
        type: pulumi.Input<enums.services.ServiceStaticSiteRouteType>;
    }

    export interface ServiceWebServiceArgs {
        disk?: pulumi.Input<inputs.services.ServiceDiskArgs>;
        env: pulumi.Input<enums.services.ServiceWebServiceEnv>;
        envSpecificDetails?: pulumi.Input<inputs.services.ServiceDockerDetailsArgs | inputs.services.ServiceNativeEnvironmentDetailsArgs>;
        healthCheckPath?: pulumi.Input<string>;
        numInstances?: pulumi.Input<number>;
        plan?: pulumi.Input<enums.services.ServiceWebServicePlan>;
        pullRequestPreviewsEnabled?: pulumi.Input<enums.services.ServiceWebServicePullRequestPreviewsEnabled>;
        region?: pulumi.Input<enums.services.ServiceWebServiceRegion>;
    }
    /**
     * serviceWebServiceArgsProvideDefaults sets the appropriate defaults for ServiceWebServiceArgs
     */
    export function serviceWebServiceArgsProvideDefaults(val: ServiceWebServiceArgs): ServiceWebServiceArgs {
        return {
            ...val,
            disk: (val.disk ? pulumi.output(val.disk).apply(inputs.services.serviceDiskArgsProvideDefaults) : undefined),
            numInstances: (val.numInstances) ?? 1,
            plan: (val.plan) ?? "starter",
            pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
            region: (val.region) ?? "oregon",
        };
    }
}
