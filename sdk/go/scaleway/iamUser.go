// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway IAM Users.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/iam/#path-users-list-users-of-an-organization).
//
// ## Examples
//
// ### Basic
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
//			_, err := scaleway.NewIamUser(ctx, "basic", &scaleway.IamUserArgs{
//				Email: pulumi.String("test@test.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// IAM users can be imported using the `{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/iamUser:IamUser basic 11111111-1111-1111-1111-111111111111
// ```
type IamUser struct {
	pulumi.CustomResourceState

	// The ID of the account root user associated with the user.
	AccountRootUserId pulumi.StringOutput `pulumi:"accountRootUserId"`
	// The date and time of the creation of the iam user.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Whether the iam user is deletable.
	Deletable pulumi.BoolOutput `pulumi:"deletable"`
	// The email of the IAM user.
	Email pulumi.StringOutput `pulumi:"email"`
	// The date of the last login.
	LastLoginAt pulumi.StringOutput `pulumi:"lastLoginAt"`
	// Whether the MFA is enabled.
	Mfa pulumi.BoolOutput `pulumi:"mfa"`
	// `organizationId`) The ID of the organization the user is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The status of user invitation. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
	Status pulumi.StringOutput `pulumi:"status"`
	// The type of user. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
	Type pulumi.StringOutput `pulumi:"type"`
	// The date and time of the last update of the iam user.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewIamUser registers a new resource with the given unique name, arguments, and options.
func NewIamUser(ctx *pulumi.Context,
	name string, args *IamUserArgs, opts ...pulumi.ResourceOption) (*IamUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IamUser
	err := ctx.RegisterResource("scaleway:index/iamUser:IamUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamUser gets an existing IamUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamUserState, opts ...pulumi.ResourceOption) (*IamUser, error) {
	var resource IamUser
	err := ctx.ReadResource("scaleway:index/iamUser:IamUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamUser resources.
type iamUserState struct {
	// The ID of the account root user associated with the user.
	AccountRootUserId *string `pulumi:"accountRootUserId"`
	// The date and time of the creation of the iam user.
	CreatedAt *string `pulumi:"createdAt"`
	// Whether the iam user is deletable.
	Deletable *bool `pulumi:"deletable"`
	// The email of the IAM user.
	Email *string `pulumi:"email"`
	// The date of the last login.
	LastLoginAt *string `pulumi:"lastLoginAt"`
	// Whether the MFA is enabled.
	Mfa *bool `pulumi:"mfa"`
	// `organizationId`) The ID of the organization the user is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The status of user invitation. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
	Status *string `pulumi:"status"`
	// The type of user. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
	Type *string `pulumi:"type"`
	// The date and time of the last update of the iam user.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type IamUserState struct {
	// The ID of the account root user associated with the user.
	AccountRootUserId pulumi.StringPtrInput
	// The date and time of the creation of the iam user.
	CreatedAt pulumi.StringPtrInput
	// Whether the iam user is deletable.
	Deletable pulumi.BoolPtrInput
	// The email of the IAM user.
	Email pulumi.StringPtrInput
	// The date of the last login.
	LastLoginAt pulumi.StringPtrInput
	// Whether the MFA is enabled.
	Mfa pulumi.BoolPtrInput
	// `organizationId`) The ID of the organization the user is associated with.
	OrganizationId pulumi.StringPtrInput
	// The status of user invitation. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
	Status pulumi.StringPtrInput
	// The type of user. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
	Type pulumi.StringPtrInput
	// The date and time of the last update of the iam user.
	UpdatedAt pulumi.StringPtrInput
}

func (IamUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamUserState)(nil)).Elem()
}

type iamUserArgs struct {
	// The email of the IAM user.
	Email string `pulumi:"email"`
	// `organizationId`) The ID of the organization the user is associated with.
	OrganizationId *string `pulumi:"organizationId"`
}

// The set of arguments for constructing a IamUser resource.
type IamUserArgs struct {
	// The email of the IAM user.
	Email pulumi.StringInput
	// `organizationId`) The ID of the organization the user is associated with.
	OrganizationId pulumi.StringPtrInput
}

func (IamUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamUserArgs)(nil)).Elem()
}

type IamUserInput interface {
	pulumi.Input

	ToIamUserOutput() IamUserOutput
	ToIamUserOutputWithContext(ctx context.Context) IamUserOutput
}

func (*IamUser) ElementType() reflect.Type {
	return reflect.TypeOf((**IamUser)(nil)).Elem()
}

func (i *IamUser) ToIamUserOutput() IamUserOutput {
	return i.ToIamUserOutputWithContext(context.Background())
}

func (i *IamUser) ToIamUserOutputWithContext(ctx context.Context) IamUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamUserOutput)
}

// IamUserArrayInput is an input type that accepts IamUserArray and IamUserArrayOutput values.
// You can construct a concrete instance of `IamUserArrayInput` via:
//
//	IamUserArray{ IamUserArgs{...} }
type IamUserArrayInput interface {
	pulumi.Input

	ToIamUserArrayOutput() IamUserArrayOutput
	ToIamUserArrayOutputWithContext(context.Context) IamUserArrayOutput
}

type IamUserArray []IamUserInput

func (IamUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamUser)(nil)).Elem()
}

func (i IamUserArray) ToIamUserArrayOutput() IamUserArrayOutput {
	return i.ToIamUserArrayOutputWithContext(context.Background())
}

func (i IamUserArray) ToIamUserArrayOutputWithContext(ctx context.Context) IamUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamUserArrayOutput)
}

// IamUserMapInput is an input type that accepts IamUserMap and IamUserMapOutput values.
// You can construct a concrete instance of `IamUserMapInput` via:
//
//	IamUserMap{ "key": IamUserArgs{...} }
type IamUserMapInput interface {
	pulumi.Input

	ToIamUserMapOutput() IamUserMapOutput
	ToIamUserMapOutputWithContext(context.Context) IamUserMapOutput
}

type IamUserMap map[string]IamUserInput

func (IamUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamUser)(nil)).Elem()
}

func (i IamUserMap) ToIamUserMapOutput() IamUserMapOutput {
	return i.ToIamUserMapOutputWithContext(context.Background())
}

func (i IamUserMap) ToIamUserMapOutputWithContext(ctx context.Context) IamUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamUserMapOutput)
}

type IamUserOutput struct{ *pulumi.OutputState }

func (IamUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamUser)(nil)).Elem()
}

func (o IamUserOutput) ToIamUserOutput() IamUserOutput {
	return o
}

func (o IamUserOutput) ToIamUserOutputWithContext(ctx context.Context) IamUserOutput {
	return o
}

// The ID of the account root user associated with the user.
func (o IamUserOutput) AccountRootUserId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.AccountRootUserId }).(pulumi.StringOutput)
}

// The date and time of the creation of the iam user.
func (o IamUserOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Whether the iam user is deletable.
func (o IamUserOutput) Deletable() pulumi.BoolOutput {
	return o.ApplyT(func(v *IamUser) pulumi.BoolOutput { return v.Deletable }).(pulumi.BoolOutput)
}

// The email of the IAM user.
func (o IamUserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The date of the last login.
func (o IamUserOutput) LastLoginAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.LastLoginAt }).(pulumi.StringOutput)
}

// Whether the MFA is enabled.
func (o IamUserOutput) Mfa() pulumi.BoolOutput {
	return o.ApplyT(func(v *IamUser) pulumi.BoolOutput { return v.Mfa }).(pulumi.BoolOutput)
}

// `organizationId`) The ID of the organization the user is associated with.
func (o IamUserOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The status of user invitation. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
func (o IamUserOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The type of user. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
func (o IamUserOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The date and time of the last update of the iam user.
func (o IamUserOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type IamUserArrayOutput struct{ *pulumi.OutputState }

func (IamUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamUser)(nil)).Elem()
}

func (o IamUserArrayOutput) ToIamUserArrayOutput() IamUserArrayOutput {
	return o
}

func (o IamUserArrayOutput) ToIamUserArrayOutputWithContext(ctx context.Context) IamUserArrayOutput {
	return o
}

func (o IamUserArrayOutput) Index(i pulumi.IntInput) IamUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamUser {
		return vs[0].([]*IamUser)[vs[1].(int)]
	}).(IamUserOutput)
}

type IamUserMapOutput struct{ *pulumi.OutputState }

func (IamUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamUser)(nil)).Elem()
}

func (o IamUserMapOutput) ToIamUserMapOutput() IamUserMapOutput {
	return o
}

func (o IamUserMapOutput) ToIamUserMapOutputWithContext(ctx context.Context) IamUserMapOutput {
	return o
}

func (o IamUserMapOutput) MapIndex(k pulumi.StringInput) IamUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamUser {
		return vs[0].(map[string]*IamUser)[vs[1].(string)]
	}).(IamUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamUserInput)(nil)).Elem(), &IamUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamUserArrayInput)(nil)).Elem(), IamUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamUserMapInput)(nil)).Elem(), IamUserMap{})
	pulumi.RegisterOutputType(IamUserOutput{})
	pulumi.RegisterOutputType(IamUserArrayOutput{})
	pulumi.RegisterOutputType(IamUserMapOutput{})
}
