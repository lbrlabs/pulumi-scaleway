// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Outputs
{

    [OutputType]
    public sealed class GetKubernetesClusterAutoscalerConfigResult
    {
        /// <summary>
        /// True if detecting similar node groups and balance the number of nodes between them is enabled.
        /// </summary>
        public readonly bool BalanceSimilarNodeGroups;
        /// <summary>
        /// True if the scale down feature of the autoscaler is disabled.
        /// </summary>
        public readonly bool DisableScaleDown;
        /// <summary>
        /// The type of resource estimator used in scale up.
        /// </summary>
        public readonly string Estimator;
        /// <summary>
        /// The type of node group expander be used in scale up.
        /// </summary>
        public readonly string Expander;
        /// <summary>
        /// Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they don't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
        /// </summary>
        public readonly int ExpendablePodsPriorityCutoff;
        /// <summary>
        /// True if ignoring DaemonSet pods when calculating resource utilization for scaling down is enabled.
        /// </summary>
        public readonly bool IgnoreDaemonsetsUtilization;
        public readonly int MaxGracefulTerminationSec;
        /// <summary>
        /// The duration after scale up that scale down evaluation resumes.
        /// </summary>
        public readonly string ScaleDownDelayAfterAdd;
        /// <summary>
        /// The duration a node should be unneeded before it is eligible for scale down.
        /// </summary>
        public readonly string ScaleDownUnneededTime;
        public readonly double ScaleDownUtilizationThreshold;

        [OutputConstructor]
        private GetKubernetesClusterAutoscalerConfigResult(
            bool balanceSimilarNodeGroups,

            bool disableScaleDown,

            string estimator,

            string expander,

            int expendablePodsPriorityCutoff,

            bool ignoreDaemonsetsUtilization,

            int maxGracefulTerminationSec,

            string scaleDownDelayAfterAdd,

            string scaleDownUnneededTime,

            double scaleDownUtilizationThreshold)
        {
            BalanceSimilarNodeGroups = balanceSimilarNodeGroups;
            DisableScaleDown = disableScaleDown;
            Estimator = estimator;
            Expander = expander;
            ExpendablePodsPriorityCutoff = expendablePodsPriorityCutoff;
            IgnoreDaemonsetsUtilization = ignoreDaemonsetsUtilization;
            MaxGracefulTerminationSec = maxGracefulTerminationSec;
            ScaleDownDelayAfterAdd = scaleDownDelayAfterAdd;
            ScaleDownUnneededTime = scaleDownUnneededTime;
            ScaleDownUtilizationThreshold = scaleDownUtilizationThreshold;
        }
    }
}
