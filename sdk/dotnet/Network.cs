// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Eventstorecloud
{
    /// <summary>
    /// Manages VPC (network) resources in Event Store Cloud
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Eventstorecloud = Pulumi.Eventstorecloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleProject = new Eventstorecloud.Project("exampleProject", new Eventstorecloud.ProjectArgs
    ///         {
    ///         });
    ///         var exampleNetwork = new Eventstorecloud.Network("exampleNetwork", new Eventstorecloud.NetworkArgs
    ///         {
    ///             ProjectId = exampleProject.Id,
    ///             ResourceProvider = "aws",
    ///             Region = "us-west-2",
    ///             CidrBlock = "172.21.0.0/16",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [EventstorecloudResourceType("eventstorecloud:index/network:Network")]
    public partial class Network : Pulumi.CustomResource
    {
        /// <summary>
        /// Address space of the network in CIDR block notation
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// Human-friendly name for the network
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Project ID
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Provider region in which to provision the network
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Cloud Provider in which to provision the network.
        /// </summary>
        [Output("resourceProvider")]
        public Output<string> ResourceProvider { get; private set; } = null!;


        /// <summary>
        /// Create a Network resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Network(string name, NetworkArgs args, CustomResourceOptions? options = null)
            : base("eventstorecloud:index/network:Network", name, args ?? new NetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Network(string name, Input<string> id, NetworkState? state = null, CustomResourceOptions? options = null)
            : base("eventstorecloud:index/network:Network", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Network resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Network Get(string name, Input<string> id, NetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new Network(name, id, state, options);
        }
    }

    public sealed class NetworkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address space of the network in CIDR block notation
        /// </summary>
        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        /// <summary>
        /// Human-friendly name for the network
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Project ID
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Provider region in which to provision the network
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// Cloud Provider in which to provision the network.
        /// </summary>
        [Input("resourceProvider", required: true)]
        public Input<string> ResourceProvider { get; set; } = null!;

        public NetworkArgs()
        {
        }
    }

    public sealed class NetworkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address space of the network in CIDR block notation
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Human-friendly name for the network
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Project ID
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Provider region in which to provision the network
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Cloud Provider in which to provision the network.
        /// </summary>
        [Input("resourceProvider")]
        public Input<string>? ResourceProvider { get; set; }

        public NetworkState()
        {
        }
    }
}
