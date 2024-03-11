// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Get information about Scaleway Load-Balancer Routes.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-route).
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ip01, err := scaleway.NewLoadbalancerIp(ctx, "ip01", nil)
//			if err != nil {
//				return err
//			}
//			lb01, err := scaleway.NewLoadbalancer(ctx, "lb01", &scaleway.LoadbalancerArgs{
//				IpId: ip01.ID(),
//				Type: pulumi.String("lb-s"),
//			})
//			if err != nil {
//				return err
//			}
//			bkd01, err := scaleway.NewLoadbalancerBackend(ctx, "bkd01", &scaleway.LoadbalancerBackendArgs{
//				LbId:            lb01.ID(),
//				ForwardProtocol: pulumi.String("tcp"),
//				ForwardPort:     pulumi.Int(80),
//				ProxyProtocol:   pulumi.String("none"),
//			})
//			if err != nil {
//				return err
//			}
//			frt01, err := scaleway.NewLoadbalancerFrontend(ctx, "frt01", &scaleway.LoadbalancerFrontendArgs{
//				LbId:        lb01.ID(),
//				BackendId:   bkd01.ID(),
//				InboundPort: pulumi.Int(80),
//			})
//			if err != nil {
//				return err
//			}
//			rt01, err := scaleway.NewLoadbalancerRoute(ctx, "rt01", &scaleway.LoadbalancerRouteArgs{
//				FrontendId: frt01.ID(),
//				BackendId:  bkd01.ID(),
//				MatchSni:   pulumi.String("sni.scaleway.com"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = scaleway.GetLbRouteOutput(ctx, scaleway.GetLbRouteOutputArgs{
//				RouteId: rt01.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetLbRoute(ctx *pulumi.Context, args *GetLbRouteArgs, opts ...pulumi.InvokeOption) (*GetLbRouteResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLbRouteResult
	err := ctx.Invoke("scaleway:index/getLbRoute:getLbRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLbRoute.
type GetLbRouteArgs struct {
	// The route id.
	RouteId string `pulumi:"routeId"`
}

// A collection of values returned by getLbRoute.
type GetLbRouteResult struct {
	BackendId  string `pulumi:"backendId"`
	CreatedAt  string `pulumi:"createdAt"`
	FrontendId string `pulumi:"frontendId"`
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	MatchHostHeader string `pulumi:"matchHostHeader"`
	MatchSni        string `pulumi:"matchSni"`
	RouteId         string `pulumi:"routeId"`
	UpdatedAt       string `pulumi:"updatedAt"`
}

func GetLbRouteOutput(ctx *pulumi.Context, args GetLbRouteOutputArgs, opts ...pulumi.InvokeOption) GetLbRouteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLbRouteResult, error) {
			args := v.(GetLbRouteArgs)
			r, err := GetLbRoute(ctx, &args, opts...)
			var s GetLbRouteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLbRouteResultOutput)
}

// A collection of arguments for invoking getLbRoute.
type GetLbRouteOutputArgs struct {
	// The route id.
	RouteId pulumi.StringInput `pulumi:"routeId"`
}

func (GetLbRouteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLbRouteArgs)(nil)).Elem()
}

// A collection of values returned by getLbRoute.
type GetLbRouteResultOutput struct{ *pulumi.OutputState }

func (GetLbRouteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLbRouteResult)(nil)).Elem()
}

func (o GetLbRouteResultOutput) ToGetLbRouteResultOutput() GetLbRouteResultOutput {
	return o
}

func (o GetLbRouteResultOutput) ToGetLbRouteResultOutputWithContext(ctx context.Context) GetLbRouteResultOutput {
	return o
}

func (o GetLbRouteResultOutput) BackendId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbRouteResult) string { return v.BackendId }).(pulumi.StringOutput)
}

func (o GetLbRouteResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbRouteResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o GetLbRouteResultOutput) FrontendId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbRouteResult) string { return v.FrontendId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLbRouteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbRouteResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLbRouteResultOutput) MatchHostHeader() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbRouteResult) string { return v.MatchHostHeader }).(pulumi.StringOutput)
}

func (o GetLbRouteResultOutput) MatchSni() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbRouteResult) string { return v.MatchSni }).(pulumi.StringOutput)
}

func (o GetLbRouteResultOutput) RouteId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbRouteResult) string { return v.RouteId }).(pulumi.StringOutput)
}

func (o GetLbRouteResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbRouteResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLbRouteResultOutput{})
}
