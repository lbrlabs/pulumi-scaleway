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
    public sealed class GetBillingInvoicesInvoiceResult
    {
        /// <summary>
        /// The payment time limit, set according to the Organization's payment conditions (RFC 3339 format).
        /// </summary>
        public readonly string DueDate;
        /// <summary>
        /// The associated invoice ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Invoices with the given type are listed. Valid values are `periodic` and `purchase`.
        /// </summary>
        public readonly string InvoiceType;
        /// <summary>
        /// The date when the invoice was sent to the customer (RFC 3339 format).
        /// </summary>
        public readonly string IssuedDate;
        /// <summary>
        /// The invoice number.
        /// </summary>
        public readonly int Number;
        /// <summary>
        /// The start date of the billing period (RFC 3339 format).
        /// </summary>
        public readonly string StartDate;
        /// <summary>
        /// The total amount, taxed.
        /// </summary>
        public readonly string TotalTaxed;
        /// <summary>
        /// The total amount, untaxed.
        /// </summary>
        public readonly string TotalUntaxed;

        [OutputConstructor]
        private GetBillingInvoicesInvoiceResult(
            string dueDate,

            string id,

            string invoiceType,

            string issuedDate,

            int number,

            string startDate,

            string totalTaxed,

            string totalUntaxed)
        {
            DueDate = dueDate;
            Id = id;
            InvoiceType = invoiceType;
            IssuedDate = issuedDate;
            Number = number;
            StartDate = startDate;
            TotalTaxed = totalTaxed;
            TotalUntaxed = totalUntaxed;
        }
    }
}
