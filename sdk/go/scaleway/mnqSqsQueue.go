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

// Creates and manages Scaleway Messaging and queuing SQS Queues.
// For further information please check
// our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/how-to/create-manage-queues/)
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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mainMnqSqs, err := scaleway.NewMnqSqs(ctx, "mainMnqSqs", nil)
//			if err != nil {
//				return err
//			}
//			mainMnqSqsCredentials, err := scaleway.NewMnqSqsCredentials(ctx, "mainMnqSqsCredentials", &scaleway.MnqSqsCredentialsArgs{
//				ProjectId: mainMnqSqs.ProjectId,
//				Permissions: &scaleway.MnqSqsCredentialsPermissionsArgs{
//					CanManage:  pulumi.Bool(false),
//					CanReceive: pulumi.Bool(true),
//					CanPublish: pulumi.Bool(false),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewMnqSqsQueue(ctx, "mainMnqSqsQueue", &scaleway.MnqSqsQueueArgs{
//				ProjectId: mainMnqSqs.ProjectId,
//				Endpoint:  mainMnqSqs.Endpoint,
//				AccessKey: mainMnqSqsCredentials.AccessKey,
//				SecretKey: mainMnqSqsCredentials.SecretKey,
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
// SQS queues can be imported using the `{region}/{project-id}/{queue-name}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/mnqSqsQueue:MnqSqsQueue main fr-par/11111111111111111111111111111111/my-queue
//
// ```
type MnqSqsQueue struct {
	pulumi.CustomResourceState

	// The access key of the SQS queue.
	AccessKey pulumi.StringOutput `pulumi:"accessKey"`
	// Specifies whether to enable content-based deduplication. Defaults to `false`.
	ContentBasedDeduplication pulumi.BoolOutput `pulumi:"contentBasedDeduplication"`
	// The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `http://sqs-sns.mnq.{region}.scw.cloud`.
	Endpoint pulumi.StringPtrOutput `pulumi:"endpoint"`
	// Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
	FifoQueue pulumi.BoolOutput `pulumi:"fifoQueue"`
	// The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
	MessageMaxAge pulumi.IntPtrOutput `pulumi:"messageMaxAge"`
	// The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
	MessageMaxSize pulumi.IntPtrOutput `pulumi:"messageMaxSize"`
	// The unique name of the sqs queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// `projectId`) The ID of the project the sqs is enabled for.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
	ReceiveWaitTimeSeconds pulumi.IntPtrOutput `pulumi:"receiveWaitTimeSeconds"`
	// `region`). The region in which sqs is enabled.
	Region pulumi.StringOutput `pulumi:"region"`
	// The secret key of the SQS queue.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// The URL of the queue.
	Url pulumi.StringOutput `pulumi:"url"`
	// The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
	VisibilityTimeoutSeconds pulumi.IntPtrOutput `pulumi:"visibilityTimeoutSeconds"`
}

// NewMnqSqsQueue registers a new resource with the given unique name, arguments, and options.
func NewMnqSqsQueue(ctx *pulumi.Context,
	name string, args *MnqSqsQueueArgs, opts ...pulumi.ResourceOption) (*MnqSqsQueue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessKey == nil {
		return nil, errors.New("invalid value for required argument 'AccessKey'")
	}
	if args.SecretKey == nil {
		return nil, errors.New("invalid value for required argument 'SecretKey'")
	}
	if args.AccessKey != nil {
		args.AccessKey = pulumi.ToSecret(args.AccessKey).(pulumi.StringInput)
	}
	if args.SecretKey != nil {
		args.SecretKey = pulumi.ToSecret(args.SecretKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessKey",
		"secretKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MnqSqsQueue
	err := ctx.RegisterResource("scaleway:index/mnqSqsQueue:MnqSqsQueue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMnqSqsQueue gets an existing MnqSqsQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMnqSqsQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MnqSqsQueueState, opts ...pulumi.ResourceOption) (*MnqSqsQueue, error) {
	var resource MnqSqsQueue
	err := ctx.ReadResource("scaleway:index/mnqSqsQueue:MnqSqsQueue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MnqSqsQueue resources.
type mnqSqsQueueState struct {
	// The access key of the SQS queue.
	AccessKey *string `pulumi:"accessKey"`
	// Specifies whether to enable content-based deduplication. Defaults to `false`.
	ContentBasedDeduplication *bool `pulumi:"contentBasedDeduplication"`
	// The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `http://sqs-sns.mnq.{region}.scw.cloud`.
	Endpoint *string `pulumi:"endpoint"`
	// Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
	FifoQueue *bool `pulumi:"fifoQueue"`
	// The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
	MessageMaxAge *int `pulumi:"messageMaxAge"`
	// The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
	MessageMaxSize *int `pulumi:"messageMaxSize"`
	// The unique name of the sqs queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// `projectId`) The ID of the project the sqs is enabled for.
	ProjectId *string `pulumi:"projectId"`
	// The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
	ReceiveWaitTimeSeconds *int `pulumi:"receiveWaitTimeSeconds"`
	// `region`). The region in which sqs is enabled.
	Region *string `pulumi:"region"`
	// The secret key of the SQS queue.
	SecretKey *string `pulumi:"secretKey"`
	// The URL of the queue.
	Url *string `pulumi:"url"`
	// The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
	VisibilityTimeoutSeconds *int `pulumi:"visibilityTimeoutSeconds"`
}

type MnqSqsQueueState struct {
	// The access key of the SQS queue.
	AccessKey pulumi.StringPtrInput
	// Specifies whether to enable content-based deduplication. Defaults to `false`.
	ContentBasedDeduplication pulumi.BoolPtrInput
	// The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `http://sqs-sns.mnq.{region}.scw.cloud`.
	Endpoint pulumi.StringPtrInput
	// Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
	FifoQueue pulumi.BoolPtrInput
	// The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
	MessageMaxAge pulumi.IntPtrInput
	// The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
	MessageMaxSize pulumi.IntPtrInput
	// The unique name of the sqs queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// `projectId`) The ID of the project the sqs is enabled for.
	ProjectId pulumi.StringPtrInput
	// The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
	ReceiveWaitTimeSeconds pulumi.IntPtrInput
	// `region`). The region in which sqs is enabled.
	Region pulumi.StringPtrInput
	// The secret key of the SQS queue.
	SecretKey pulumi.StringPtrInput
	// The URL of the queue.
	Url pulumi.StringPtrInput
	// The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
	VisibilityTimeoutSeconds pulumi.IntPtrInput
}

func (MnqSqsQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqSqsQueueState)(nil)).Elem()
}

type mnqSqsQueueArgs struct {
	// The access key of the SQS queue.
	AccessKey string `pulumi:"accessKey"`
	// Specifies whether to enable content-based deduplication. Defaults to `false`.
	ContentBasedDeduplication *bool `pulumi:"contentBasedDeduplication"`
	// The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `http://sqs-sns.mnq.{region}.scw.cloud`.
	Endpoint *string `pulumi:"endpoint"`
	// Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
	FifoQueue *bool `pulumi:"fifoQueue"`
	// The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
	MessageMaxAge *int `pulumi:"messageMaxAge"`
	// The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
	MessageMaxSize *int `pulumi:"messageMaxSize"`
	// The unique name of the sqs queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// `projectId`) The ID of the project the sqs is enabled for.
	ProjectId *string `pulumi:"projectId"`
	// The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
	ReceiveWaitTimeSeconds *int `pulumi:"receiveWaitTimeSeconds"`
	// `region`). The region in which sqs is enabled.
	Region *string `pulumi:"region"`
	// The secret key of the SQS queue.
	SecretKey string `pulumi:"secretKey"`
	// The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
	VisibilityTimeoutSeconds *int `pulumi:"visibilityTimeoutSeconds"`
}

// The set of arguments for constructing a MnqSqsQueue resource.
type MnqSqsQueueArgs struct {
	// The access key of the SQS queue.
	AccessKey pulumi.StringInput
	// Specifies whether to enable content-based deduplication. Defaults to `false`.
	ContentBasedDeduplication pulumi.BoolPtrInput
	// The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `http://sqs-sns.mnq.{region}.scw.cloud`.
	Endpoint pulumi.StringPtrInput
	// Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
	FifoQueue pulumi.BoolPtrInput
	// The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
	MessageMaxAge pulumi.IntPtrInput
	// The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
	MessageMaxSize pulumi.IntPtrInput
	// The unique name of the sqs queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// `projectId`) The ID of the project the sqs is enabled for.
	ProjectId pulumi.StringPtrInput
	// The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
	ReceiveWaitTimeSeconds pulumi.IntPtrInput
	// `region`). The region in which sqs is enabled.
	Region pulumi.StringPtrInput
	// The secret key of the SQS queue.
	SecretKey pulumi.StringInput
	// The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
	VisibilityTimeoutSeconds pulumi.IntPtrInput
}

func (MnqSqsQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqSqsQueueArgs)(nil)).Elem()
}

type MnqSqsQueueInput interface {
	pulumi.Input

	ToMnqSqsQueueOutput() MnqSqsQueueOutput
	ToMnqSqsQueueOutputWithContext(ctx context.Context) MnqSqsQueueOutput
}

func (*MnqSqsQueue) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqSqsQueue)(nil)).Elem()
}

func (i *MnqSqsQueue) ToMnqSqsQueueOutput() MnqSqsQueueOutput {
	return i.ToMnqSqsQueueOutputWithContext(context.Background())
}

func (i *MnqSqsQueue) ToMnqSqsQueueOutputWithContext(ctx context.Context) MnqSqsQueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqSqsQueueOutput)
}

// MnqSqsQueueArrayInput is an input type that accepts MnqSqsQueueArray and MnqSqsQueueArrayOutput values.
// You can construct a concrete instance of `MnqSqsQueueArrayInput` via:
//
//	MnqSqsQueueArray{ MnqSqsQueueArgs{...} }
type MnqSqsQueueArrayInput interface {
	pulumi.Input

	ToMnqSqsQueueArrayOutput() MnqSqsQueueArrayOutput
	ToMnqSqsQueueArrayOutputWithContext(context.Context) MnqSqsQueueArrayOutput
}

type MnqSqsQueueArray []MnqSqsQueueInput

func (MnqSqsQueueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqSqsQueue)(nil)).Elem()
}

func (i MnqSqsQueueArray) ToMnqSqsQueueArrayOutput() MnqSqsQueueArrayOutput {
	return i.ToMnqSqsQueueArrayOutputWithContext(context.Background())
}

func (i MnqSqsQueueArray) ToMnqSqsQueueArrayOutputWithContext(ctx context.Context) MnqSqsQueueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqSqsQueueArrayOutput)
}

// MnqSqsQueueMapInput is an input type that accepts MnqSqsQueueMap and MnqSqsQueueMapOutput values.
// You can construct a concrete instance of `MnqSqsQueueMapInput` via:
//
//	MnqSqsQueueMap{ "key": MnqSqsQueueArgs{...} }
type MnqSqsQueueMapInput interface {
	pulumi.Input

	ToMnqSqsQueueMapOutput() MnqSqsQueueMapOutput
	ToMnqSqsQueueMapOutputWithContext(context.Context) MnqSqsQueueMapOutput
}

type MnqSqsQueueMap map[string]MnqSqsQueueInput

func (MnqSqsQueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqSqsQueue)(nil)).Elem()
}

func (i MnqSqsQueueMap) ToMnqSqsQueueMapOutput() MnqSqsQueueMapOutput {
	return i.ToMnqSqsQueueMapOutputWithContext(context.Background())
}

func (i MnqSqsQueueMap) ToMnqSqsQueueMapOutputWithContext(ctx context.Context) MnqSqsQueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqSqsQueueMapOutput)
}

type MnqSqsQueueOutput struct{ *pulumi.OutputState }

func (MnqSqsQueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqSqsQueue)(nil)).Elem()
}

func (o MnqSqsQueueOutput) ToMnqSqsQueueOutput() MnqSqsQueueOutput {
	return o
}

func (o MnqSqsQueueOutput) ToMnqSqsQueueOutputWithContext(ctx context.Context) MnqSqsQueueOutput {
	return o
}

// The access key of the SQS queue.
func (o MnqSqsQueueOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSqsQueue) pulumi.StringOutput { return v.AccessKey }).(pulumi.StringOutput)
}

// Specifies whether to enable content-based deduplication. Defaults to `false`.
func (o MnqSqsQueueOutput) ContentBasedDeduplication() pulumi.BoolOutput {
	return o.ApplyT(func(v *MnqSqsQueue) pulumi.BoolOutput { return v.ContentBasedDeduplication }).(pulumi.BoolOutput)
}

// The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `http://sqs-sns.mnq.{region}.scw.cloud`.
func (o MnqSqsQueueOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MnqSqsQueue) pulumi.StringPtrOutput { return v.Endpoint }).(pulumi.StringPtrOutput)
}

// Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
func (o MnqSqsQueueOutput) FifoQueue() pulumi.BoolOutput {
	return o.ApplyT(func(v *MnqSqsQueue) pulumi.BoolOutput { return v.FifoQueue }).(pulumi.BoolOutput)
}

// The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
func (o MnqSqsQueueOutput) MessageMaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MnqSqsQueue) pulumi.IntPtrOutput { return v.MessageMaxAge }).(pulumi.IntPtrOutput)
}

// The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
func (o MnqSqsQueueOutput) MessageMaxSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MnqSqsQueue) pulumi.IntPtrOutput { return v.MessageMaxSize }).(pulumi.IntPtrOutput)
}

// The unique name of the sqs queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
func (o MnqSqsQueueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSqsQueue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
func (o MnqSqsQueueOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSqsQueue) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the sqs is enabled for.
func (o MnqSqsQueueOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSqsQueue) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
func (o MnqSqsQueueOutput) ReceiveWaitTimeSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MnqSqsQueue) pulumi.IntPtrOutput { return v.ReceiveWaitTimeSeconds }).(pulumi.IntPtrOutput)
}

// `region`). The region in which sqs is enabled.
func (o MnqSqsQueueOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSqsQueue) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The secret key of the SQS queue.
func (o MnqSqsQueueOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSqsQueue) pulumi.StringOutput { return v.SecretKey }).(pulumi.StringOutput)
}

// The URL of the queue.
func (o MnqSqsQueueOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSqsQueue) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
func (o MnqSqsQueueOutput) VisibilityTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MnqSqsQueue) pulumi.IntPtrOutput { return v.VisibilityTimeoutSeconds }).(pulumi.IntPtrOutput)
}

type MnqSqsQueueArrayOutput struct{ *pulumi.OutputState }

func (MnqSqsQueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqSqsQueue)(nil)).Elem()
}

func (o MnqSqsQueueArrayOutput) ToMnqSqsQueueArrayOutput() MnqSqsQueueArrayOutput {
	return o
}

func (o MnqSqsQueueArrayOutput) ToMnqSqsQueueArrayOutputWithContext(ctx context.Context) MnqSqsQueueArrayOutput {
	return o
}

func (o MnqSqsQueueArrayOutput) Index(i pulumi.IntInput) MnqSqsQueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MnqSqsQueue {
		return vs[0].([]*MnqSqsQueue)[vs[1].(int)]
	}).(MnqSqsQueueOutput)
}

type MnqSqsQueueMapOutput struct{ *pulumi.OutputState }

func (MnqSqsQueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqSqsQueue)(nil)).Elem()
}

func (o MnqSqsQueueMapOutput) ToMnqSqsQueueMapOutput() MnqSqsQueueMapOutput {
	return o
}

func (o MnqSqsQueueMapOutput) ToMnqSqsQueueMapOutputWithContext(ctx context.Context) MnqSqsQueueMapOutput {
	return o
}

func (o MnqSqsQueueMapOutput) MapIndex(k pulumi.StringInput) MnqSqsQueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MnqSqsQueue {
		return vs[0].(map[string]*MnqSqsQueue)[vs[1].(string)]
	}).(MnqSqsQueueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MnqSqsQueueInput)(nil)).Elem(), &MnqSqsQueue{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqSqsQueueArrayInput)(nil)).Elem(), MnqSqsQueueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqSqsQueueMapInput)(nil)).Elem(), MnqSqsQueueMap{})
	pulumi.RegisterOutputType(MnqSqsQueueOutput{})
	pulumi.RegisterOutputType(MnqSqsQueueArrayOutput{})
	pulumi.RegisterOutputType(MnqSqsQueueMapOutput{})
}
