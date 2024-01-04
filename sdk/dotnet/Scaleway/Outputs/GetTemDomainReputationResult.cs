// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Scaleway.Outputs
{

    [OutputType]
    public sealed class GetTemDomainReputationResult
    {
        public readonly int PreviousScore;
        public readonly string PreviousScoredAt;
        public readonly int Score;
        public readonly string ScoredAt;
        public readonly string Status;

        [OutputConstructor]
        private GetTemDomainReputationResult(
            int previousScore,

            string previousScoredAt,

            int score,

            string scoredAt,

            string status)
        {
            PreviousScore = previousScore;
            PreviousScoredAt = previousScoredAt;
            Score = score;
            ScoredAt = scoredAt;
            Status = status;
        }
    }
}
