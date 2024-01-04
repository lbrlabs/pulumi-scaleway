// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Function Token.
// For more information see [the documentation](https://developers.scaleway.com/en/products/functions/api/#tokens-26b085).
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
//			mainFunctionNamespace, err := scaleway.NewFunctionNamespace(ctx, "mainFunctionNamespace", nil)
//			if err != nil {
//				return err
//			}
//			mainFunction, err := scaleway.NewFunction(ctx, "mainFunction", &scaleway.FunctionArgs{
//				NamespaceId: mainFunctionNamespace.ID(),
//				Runtime:     pulumi.String("go118"),
//				Handler:     pulumi.String("Handle"),
//				Privacy:     pulumi.String("private"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewFunctionToken(ctx, "namespace", &scaleway.FunctionTokenArgs{
//				NamespaceId: mainFunctionNamespace.ID(),
//				ExpiresAt:   pulumi.String("2022-10-18T11:35:15+02:00"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewFunctionToken(ctx, "function", &scaleway.FunctionTokenArgs{
//				FunctionId: mainFunction.ID(),
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
// Tokens can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/functionToken:FunctionToken main fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type FunctionToken struct {
	pulumi.CustomResourceState

	// The description of the token.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The expiration date of the token.
	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// The ID of the function.
	//
	// > Only one of `namespaceId` or `functionId` must be set.
	FunctionId pulumi.StringPtrOutput `pulumi:"functionId"`
	// The ID of the function namespace.
	NamespaceId pulumi.StringPtrOutput `pulumi:"namespaceId"`
	// `region`). The region in which the namespace should be created.
	//
	// > **Important** Updates to any fields will recreate the token.
	Region pulumi.StringOutput `pulumi:"region"`
	// The token.
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewFunctionToken registers a new resource with the given unique name, arguments, and options.
func NewFunctionToken(ctx *pulumi.Context,
	name string, args *FunctionTokenArgs, opts ...pulumi.ResourceOption) (*FunctionToken, error) {
	if args == nil {
		args = &FunctionTokenArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FunctionToken
	err := ctx.RegisterResource("scaleway:index/functionToken:FunctionToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionToken gets an existing FunctionToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionTokenState, opts ...pulumi.ResourceOption) (*FunctionToken, error) {
	var resource FunctionToken
	err := ctx.ReadResource("scaleway:index/functionToken:FunctionToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionToken resources.
type functionTokenState struct {
	// The description of the token.
	Description *string `pulumi:"description"`
	// The expiration date of the token.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The ID of the function.
	//
	// > Only one of `namespaceId` or `functionId` must be set.
	FunctionId *string `pulumi:"functionId"`
	// The ID of the function namespace.
	NamespaceId *string `pulumi:"namespaceId"`
	// `region`). The region in which the namespace should be created.
	//
	// > **Important** Updates to any fields will recreate the token.
	Region *string `pulumi:"region"`
	// The token.
	Token *string `pulumi:"token"`
}

type FunctionTokenState struct {
	// The description of the token.
	Description pulumi.StringPtrInput
	// The expiration date of the token.
	ExpiresAt pulumi.StringPtrInput
	// The ID of the function.
	//
	// > Only one of `namespaceId` or `functionId` must be set.
	FunctionId pulumi.StringPtrInput
	// The ID of the function namespace.
	NamespaceId pulumi.StringPtrInput
	// `region`). The region in which the namespace should be created.
	//
	// > **Important** Updates to any fields will recreate the token.
	Region pulumi.StringPtrInput
	// The token.
	Token pulumi.StringPtrInput
}

func (FunctionTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionTokenState)(nil)).Elem()
}

type functionTokenArgs struct {
	// The description of the token.
	Description *string `pulumi:"description"`
	// The expiration date of the token.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The ID of the function.
	//
	// > Only one of `namespaceId` or `functionId` must be set.
	FunctionId *string `pulumi:"functionId"`
	// The ID of the function namespace.
	NamespaceId *string `pulumi:"namespaceId"`
	// `region`). The region in which the namespace should be created.
	//
	// > **Important** Updates to any fields will recreate the token.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a FunctionToken resource.
type FunctionTokenArgs struct {
	// The description of the token.
	Description pulumi.StringPtrInput
	// The expiration date of the token.
	ExpiresAt pulumi.StringPtrInput
	// The ID of the function.
	//
	// > Only one of `namespaceId` or `functionId` must be set.
	FunctionId pulumi.StringPtrInput
	// The ID of the function namespace.
	NamespaceId pulumi.StringPtrInput
	// `region`). The region in which the namespace should be created.
	//
	// > **Important** Updates to any fields will recreate the token.
	Region pulumi.StringPtrInput
}

func (FunctionTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionTokenArgs)(nil)).Elem()
}

type FunctionTokenInput interface {
	pulumi.Input

	ToFunctionTokenOutput() FunctionTokenOutput
	ToFunctionTokenOutputWithContext(ctx context.Context) FunctionTokenOutput
}

func (*FunctionToken) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionToken)(nil)).Elem()
}

func (i *FunctionToken) ToFunctionTokenOutput() FunctionTokenOutput {
	return i.ToFunctionTokenOutputWithContext(context.Background())
}

func (i *FunctionToken) ToFunctionTokenOutputWithContext(ctx context.Context) FunctionTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionTokenOutput)
}

// FunctionTokenArrayInput is an input type that accepts FunctionTokenArray and FunctionTokenArrayOutput values.
// You can construct a concrete instance of `FunctionTokenArrayInput` via:
//
//	FunctionTokenArray{ FunctionTokenArgs{...} }
type FunctionTokenArrayInput interface {
	pulumi.Input

	ToFunctionTokenArrayOutput() FunctionTokenArrayOutput
	ToFunctionTokenArrayOutputWithContext(context.Context) FunctionTokenArrayOutput
}

type FunctionTokenArray []FunctionTokenInput

func (FunctionTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionToken)(nil)).Elem()
}

func (i FunctionTokenArray) ToFunctionTokenArrayOutput() FunctionTokenArrayOutput {
	return i.ToFunctionTokenArrayOutputWithContext(context.Background())
}

func (i FunctionTokenArray) ToFunctionTokenArrayOutputWithContext(ctx context.Context) FunctionTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionTokenArrayOutput)
}

// FunctionTokenMapInput is an input type that accepts FunctionTokenMap and FunctionTokenMapOutput values.
// You can construct a concrete instance of `FunctionTokenMapInput` via:
//
//	FunctionTokenMap{ "key": FunctionTokenArgs{...} }
type FunctionTokenMapInput interface {
	pulumi.Input

	ToFunctionTokenMapOutput() FunctionTokenMapOutput
	ToFunctionTokenMapOutputWithContext(context.Context) FunctionTokenMapOutput
}

type FunctionTokenMap map[string]FunctionTokenInput

func (FunctionTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionToken)(nil)).Elem()
}

func (i FunctionTokenMap) ToFunctionTokenMapOutput() FunctionTokenMapOutput {
	return i.ToFunctionTokenMapOutputWithContext(context.Background())
}

func (i FunctionTokenMap) ToFunctionTokenMapOutputWithContext(ctx context.Context) FunctionTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionTokenMapOutput)
}

type FunctionTokenOutput struct{ *pulumi.OutputState }

func (FunctionTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionToken)(nil)).Elem()
}

func (o FunctionTokenOutput) ToFunctionTokenOutput() FunctionTokenOutput {
	return o
}

func (o FunctionTokenOutput) ToFunctionTokenOutputWithContext(ctx context.Context) FunctionTokenOutput {
	return o
}

// The description of the token.
func (o FunctionTokenOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionToken) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The expiration date of the token.
func (o FunctionTokenOutput) ExpiresAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionToken) pulumi.StringPtrOutput { return v.ExpiresAt }).(pulumi.StringPtrOutput)
}

// The ID of the function.
//
// > Only one of `namespaceId` or `functionId` must be set.
func (o FunctionTokenOutput) FunctionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionToken) pulumi.StringPtrOutput { return v.FunctionId }).(pulumi.StringPtrOutput)
}

// The ID of the function namespace.
func (o FunctionTokenOutput) NamespaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionToken) pulumi.StringPtrOutput { return v.NamespaceId }).(pulumi.StringPtrOutput)
}

// `region`). The region in which the namespace should be created.
//
// > **Important** Updates to any fields will recreate the token.
func (o FunctionTokenOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionToken) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The token.
func (o FunctionTokenOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionToken) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

type FunctionTokenArrayOutput struct{ *pulumi.OutputState }

func (FunctionTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionToken)(nil)).Elem()
}

func (o FunctionTokenArrayOutput) ToFunctionTokenArrayOutput() FunctionTokenArrayOutput {
	return o
}

func (o FunctionTokenArrayOutput) ToFunctionTokenArrayOutputWithContext(ctx context.Context) FunctionTokenArrayOutput {
	return o
}

func (o FunctionTokenArrayOutput) Index(i pulumi.IntInput) FunctionTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FunctionToken {
		return vs[0].([]*FunctionToken)[vs[1].(int)]
	}).(FunctionTokenOutput)
}

type FunctionTokenMapOutput struct{ *pulumi.OutputState }

func (FunctionTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionToken)(nil)).Elem()
}

func (o FunctionTokenMapOutput) ToFunctionTokenMapOutput() FunctionTokenMapOutput {
	return o
}

func (o FunctionTokenMapOutput) ToFunctionTokenMapOutputWithContext(ctx context.Context) FunctionTokenMapOutput {
	return o
}

func (o FunctionTokenMapOutput) MapIndex(k pulumi.StringInput) FunctionTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FunctionToken {
		return vs[0].(map[string]*FunctionToken)[vs[1].(string)]
	}).(FunctionTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionTokenInput)(nil)).Elem(), &FunctionToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionTokenArrayInput)(nil)).Elem(), FunctionTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionTokenMapInput)(nil)).Elem(), FunctionTokenMap{})
	pulumi.RegisterOutputType(FunctionTokenOutput{})
	pulumi.RegisterOutputType(FunctionTokenArrayOutput{})
	pulumi.RegisterOutputType(FunctionTokenMapOutput{})
}
