// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Function Triggers.
// For more information see [the documentation](https://www.scaleway.com/en/developers/api/serverless-functions/#path-triggers).
//
// ## Examples
//
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := scaleway.NewFunctionTrigger(ctx, "main", &scaleway.FunctionTriggerArgs{
//				FunctionId: pulumi.Any(scaleway_function.Main.Id),
//				Sqs: &scaleway.FunctionTriggerSqsArgs{
//					NamespaceId: pulumi.Any(scaleway_mnq_namespace.Main.Id),
//					Queue:       pulumi.String("MyQueue"),
//					ProjectId:   pulumi.Any(scaleway_mnq_namespace.Main.Project_id),
//					Region:      pulumi.Any(scaleway_mnq_namespace.Main.Region),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Function Triggers can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/functionTrigger:FunctionTrigger main fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type FunctionTrigger struct {
	pulumi.CustomResourceState

	// The description of the trigger.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the function to create a trigger for
	FunctionId pulumi.StringOutput `pulumi:"functionId"`
	// The unique name of the trigger. Default to a generated name.
	Name pulumi.StringOutput `pulumi:"name"`
	// `region`). The region in which the namespace should be created.
	Region pulumi.StringOutput `pulumi:"region"`
	// The configuration of the Scaleway's SQS used by the trigger
	Sqs FunctionTriggerSqsPtrOutput `pulumi:"sqs"`
}

// NewFunctionTrigger registers a new resource with the given unique name, arguments, and options.
func NewFunctionTrigger(ctx *pulumi.Context,
	name string, args *FunctionTriggerArgs, opts ...pulumi.ResourceOption) (*FunctionTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FunctionId == nil {
		return nil, errors.New("invalid value for required argument 'FunctionId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FunctionTrigger
	err := ctx.RegisterResource("scaleway:index/functionTrigger:FunctionTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionTrigger gets an existing FunctionTrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionTriggerState, opts ...pulumi.ResourceOption) (*FunctionTrigger, error) {
	var resource FunctionTrigger
	err := ctx.ReadResource("scaleway:index/functionTrigger:FunctionTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionTrigger resources.
type functionTriggerState struct {
	// The description of the trigger.
	Description *string `pulumi:"description"`
	// The ID of the function to create a trigger for
	FunctionId *string `pulumi:"functionId"`
	// The unique name of the trigger. Default to a generated name.
	Name *string `pulumi:"name"`
	// `region`). The region in which the namespace should be created.
	Region *string `pulumi:"region"`
	// The configuration of the Scaleway's SQS used by the trigger
	Sqs *FunctionTriggerSqs `pulumi:"sqs"`
}

type FunctionTriggerState struct {
	// The description of the trigger.
	Description pulumi.StringPtrInput
	// The ID of the function to create a trigger for
	FunctionId pulumi.StringPtrInput
	// The unique name of the trigger. Default to a generated name.
	Name pulumi.StringPtrInput
	// `region`). The region in which the namespace should be created.
	Region pulumi.StringPtrInput
	// The configuration of the Scaleway's SQS used by the trigger
	Sqs FunctionTriggerSqsPtrInput
}

func (FunctionTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionTriggerState)(nil)).Elem()
}

type functionTriggerArgs struct {
	// The description of the trigger.
	Description *string `pulumi:"description"`
	// The ID of the function to create a trigger for
	FunctionId string `pulumi:"functionId"`
	// The unique name of the trigger. Default to a generated name.
	Name *string `pulumi:"name"`
	// `region`). The region in which the namespace should be created.
	Region *string `pulumi:"region"`
	// The configuration of the Scaleway's SQS used by the trigger
	Sqs *FunctionTriggerSqs `pulumi:"sqs"`
}

// The set of arguments for constructing a FunctionTrigger resource.
type FunctionTriggerArgs struct {
	// The description of the trigger.
	Description pulumi.StringPtrInput
	// The ID of the function to create a trigger for
	FunctionId pulumi.StringInput
	// The unique name of the trigger. Default to a generated name.
	Name pulumi.StringPtrInput
	// `region`). The region in which the namespace should be created.
	Region pulumi.StringPtrInput
	// The configuration of the Scaleway's SQS used by the trigger
	Sqs FunctionTriggerSqsPtrInput
}

func (FunctionTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionTriggerArgs)(nil)).Elem()
}

type FunctionTriggerInput interface {
	pulumi.Input

	ToFunctionTriggerOutput() FunctionTriggerOutput
	ToFunctionTriggerOutputWithContext(ctx context.Context) FunctionTriggerOutput
}

func (*FunctionTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionTrigger)(nil)).Elem()
}

func (i *FunctionTrigger) ToFunctionTriggerOutput() FunctionTriggerOutput {
	return i.ToFunctionTriggerOutputWithContext(context.Background())
}

func (i *FunctionTrigger) ToFunctionTriggerOutputWithContext(ctx context.Context) FunctionTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionTriggerOutput)
}

// FunctionTriggerArrayInput is an input type that accepts FunctionTriggerArray and FunctionTriggerArrayOutput values.
// You can construct a concrete instance of `FunctionTriggerArrayInput` via:
//
//	FunctionTriggerArray{ FunctionTriggerArgs{...} }
type FunctionTriggerArrayInput interface {
	pulumi.Input

	ToFunctionTriggerArrayOutput() FunctionTriggerArrayOutput
	ToFunctionTriggerArrayOutputWithContext(context.Context) FunctionTriggerArrayOutput
}

type FunctionTriggerArray []FunctionTriggerInput

func (FunctionTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionTrigger)(nil)).Elem()
}

func (i FunctionTriggerArray) ToFunctionTriggerArrayOutput() FunctionTriggerArrayOutput {
	return i.ToFunctionTriggerArrayOutputWithContext(context.Background())
}

func (i FunctionTriggerArray) ToFunctionTriggerArrayOutputWithContext(ctx context.Context) FunctionTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionTriggerArrayOutput)
}

// FunctionTriggerMapInput is an input type that accepts FunctionTriggerMap and FunctionTriggerMapOutput values.
// You can construct a concrete instance of `FunctionTriggerMapInput` via:
//
//	FunctionTriggerMap{ "key": FunctionTriggerArgs{...} }
type FunctionTriggerMapInput interface {
	pulumi.Input

	ToFunctionTriggerMapOutput() FunctionTriggerMapOutput
	ToFunctionTriggerMapOutputWithContext(context.Context) FunctionTriggerMapOutput
}

type FunctionTriggerMap map[string]FunctionTriggerInput

func (FunctionTriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionTrigger)(nil)).Elem()
}

func (i FunctionTriggerMap) ToFunctionTriggerMapOutput() FunctionTriggerMapOutput {
	return i.ToFunctionTriggerMapOutputWithContext(context.Background())
}

func (i FunctionTriggerMap) ToFunctionTriggerMapOutputWithContext(ctx context.Context) FunctionTriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionTriggerMapOutput)
}

type FunctionTriggerOutput struct{ *pulumi.OutputState }

func (FunctionTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionTrigger)(nil)).Elem()
}

func (o FunctionTriggerOutput) ToFunctionTriggerOutput() FunctionTriggerOutput {
	return o
}

func (o FunctionTriggerOutput) ToFunctionTriggerOutputWithContext(ctx context.Context) FunctionTriggerOutput {
	return o
}

// The description of the trigger.
func (o FunctionTriggerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionTrigger) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the function to create a trigger for
func (o FunctionTriggerOutput) FunctionId() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionTrigger) pulumi.StringOutput { return v.FunctionId }).(pulumi.StringOutput)
}

// The unique name of the trigger. Default to a generated name.
func (o FunctionTriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionTrigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// `region`). The region in which the namespace should be created.
func (o FunctionTriggerOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionTrigger) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The configuration of the Scaleway's SQS used by the trigger
func (o FunctionTriggerOutput) Sqs() FunctionTriggerSqsPtrOutput {
	return o.ApplyT(func(v *FunctionTrigger) FunctionTriggerSqsPtrOutput { return v.Sqs }).(FunctionTriggerSqsPtrOutput)
}

type FunctionTriggerArrayOutput struct{ *pulumi.OutputState }

func (FunctionTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionTrigger)(nil)).Elem()
}

func (o FunctionTriggerArrayOutput) ToFunctionTriggerArrayOutput() FunctionTriggerArrayOutput {
	return o
}

func (o FunctionTriggerArrayOutput) ToFunctionTriggerArrayOutputWithContext(ctx context.Context) FunctionTriggerArrayOutput {
	return o
}

func (o FunctionTriggerArrayOutput) Index(i pulumi.IntInput) FunctionTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FunctionTrigger {
		return vs[0].([]*FunctionTrigger)[vs[1].(int)]
	}).(FunctionTriggerOutput)
}

type FunctionTriggerMapOutput struct{ *pulumi.OutputState }

func (FunctionTriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionTrigger)(nil)).Elem()
}

func (o FunctionTriggerMapOutput) ToFunctionTriggerMapOutput() FunctionTriggerMapOutput {
	return o
}

func (o FunctionTriggerMapOutput) ToFunctionTriggerMapOutputWithContext(ctx context.Context) FunctionTriggerMapOutput {
	return o
}

func (o FunctionTriggerMapOutput) MapIndex(k pulumi.StringInput) FunctionTriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FunctionTrigger {
		return vs[0].(map[string]*FunctionTrigger)[vs[1].(string)]
	}).(FunctionTriggerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionTriggerInput)(nil)).Elem(), &FunctionTrigger{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionTriggerArrayInput)(nil)).Elem(), FunctionTriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionTriggerMapInput)(nil)).Elem(), FunctionTriggerMap{})
	pulumi.RegisterOutputType(FunctionTriggerOutput{})
	pulumi.RegisterOutputType(FunctionTriggerArrayOutput{})
	pulumi.RegisterOutputType(FunctionTriggerMapOutput{})
}
