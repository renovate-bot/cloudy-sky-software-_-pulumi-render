// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Inputs
{

    public sealed class StaticSiteServiceDetailsParentServerPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public StaticSiteServiceDetailsParentServerPropertiesArgs()
        {
        }
        public static new StaticSiteServiceDetailsParentServerPropertiesArgs Empty => new StaticSiteServiceDetailsParentServerPropertiesArgs();
    }
}
