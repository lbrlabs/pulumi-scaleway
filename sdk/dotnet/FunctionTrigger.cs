// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway
{
    /// <summary>
    /// Creates and manages Scaleway Function Triggers.
    /// For more information see [the documentation](https://www.scaleway.com/en/developers/api/serverless-functions/#path-triggers).
    /// 
    /// ## Examples
    /// 
    /// ### SQS
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.FunctionTrigger("main", new()
    ///     {
    ///         FunctionId = scaleway_function.Main.Id,
    ///         Sqs = new Scaleway.Inputs.FunctionTriggerSqsArgs
    ///         {
    ///             ProjectId = scaleway_mnq_sqs.Main.Project_id,
    ///             Queue = "MyQueue",
    ///             Region = scaleway_mnq_sqs.Main.Region,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Function Triggers can be imported using the `{region}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/functionTrigger:FunctionTrigger main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/functionTrigger:FunctionTrigger")]
    public partial class FunctionTrigger : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the trigger.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the function to create a trigger for
        /// </summary>
        [Output("functionId")]
        public Output<string> FunctionId { get; private set; } = null!;

        /// <summary>
        /// The unique name of the trigger. Default to a generated name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The configuration for the Scaleway's Nats used by the trigger
        /// </summary>
        [Output("nats")]
        public Output<Outputs.FunctionTriggerNats?> Nats { get; private set; } = null!;

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The configuration of the Scaleway's SQS used by the trigger
        /// </summary>
        [Output("sqs")]
        public Output<Outputs.FunctionTriggerSqs?> Sqs { get; private set; } = null!;


        /// <summary>
        /// Create a FunctionTrigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FunctionTrigger(string name, FunctionTriggerArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/functionTrigger:FunctionTrigger", name, args ?? new FunctionTriggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FunctionTrigger(string name, Input<string> id, FunctionTriggerState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/functionTrigger:FunctionTrigger", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FunctionTrigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FunctionTrigger Get(string name, Input<string> id, FunctionTriggerState? state = null, CustomResourceOptions? options = null)
        {
            return new FunctionTrigger(name, id, state, options);
        }
    }

    public sealed class FunctionTriggerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the trigger.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the function to create a trigger for
        /// </summary>
        [Input("functionId", required: true)]
        public Input<string> FunctionId { get; set; } = null!;

        /// <summary>
        /// The unique name of the trigger. Default to a generated name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The configuration for the Scaleway's Nats used by the trigger
        /// </summary>
        [Input("nats")]
        public Input<Inputs.FunctionTriggerNatsArgs>? Nats { get; set; }

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The configuration of the Scaleway's SQS used by the trigger
        /// </summary>
        [Input("sqs")]
        public Input<Inputs.FunctionTriggerSqsArgs>? Sqs { get; set; }

        public FunctionTriggerArgs()
        {
        }
        public static new FunctionTriggerArgs Empty => new FunctionTriggerArgs();
    }

    public sealed class FunctionTriggerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the trigger.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the function to create a trigger for
        /// </summary>
        [Input("functionId")]
        public Input<string>? FunctionId { get; set; }

        /// <summary>
        /// The unique name of the trigger. Default to a generated name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The configuration for the Scaleway's Nats used by the trigger
        /// </summary>
        [Input("nats")]
        public Input<Inputs.FunctionTriggerNatsGetArgs>? Nats { get; set; }

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The configuration of the Scaleway's SQS used by the trigger
        /// </summary>
        [Input("sqs")]
        public Input<Inputs.FunctionTriggerSqsGetArgs>? Sqs { get; set; }

        public FunctionTriggerState()
        {
        }
        public static new FunctionTriggerState Empty => new FunctionTriggerState();
    }
}