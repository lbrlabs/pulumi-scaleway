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

// Creates and manages Scaleway Messaging and Queuing queues.
//
// For more information about MNQ, see [the documentation](https://www.scaleway.com/en/developers/api/messaging-and-queuing/).
//
// > NOTE: This resource refers to the old version of the MNQ API. You should use new resources dedicated to your protocol. SQS.
//
// ## Examples
//
// ### NATS
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
//			mainMnqNamespace, err := scaleway.NewMnqNamespace(ctx, "mainMnqNamespace", &scaleway.MnqNamespaceArgs{
//				Protocol: pulumi.String("nats"),
//			})
//			if err != nil {
//				return err
//			}
//			mainMnqCredential, err := scaleway.NewMnqCredential(ctx, "mainMnqCredential", &scaleway.MnqCredentialArgs{
//				NamespaceId: mainMnqNamespace.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewMnqQueue(ctx, "myQueue", &scaleway.MnqQueueArgs{
//				NamespaceId: mainMnqNamespace.ID(),
//				Nats: &scaleway.MnqQueueNatsArgs{
//					Credentials: mainMnqCredential.NatsCredentials.ApplyT(func(natsCredentials scaleway.MnqCredentialNatsCredentials) (*string, error) {
//						return &natsCredentials.Content, nil
//					}).(pulumi.StringPtrOutput),
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
// <!--End PulumiCodeChooser -->
//
// ### SQS
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
//			mainMnqNamespace, err := scaleway.NewMnqNamespace(ctx, "mainMnqNamespace", &scaleway.MnqNamespaceArgs{
//				Protocol: pulumi.String("sqs_sns"),
//			})
//			if err != nil {
//				return err
//			}
//			mainMnqCredential, err := scaleway.NewMnqCredential(ctx, "mainMnqCredential", &scaleway.MnqCredentialArgs{
//				NamespaceId: mainMnqNamespace.ID(),
//				SqsSnsCredentials: &scaleway.MnqCredentialSqsSnsCredentialsArgs{
//					Permissions: &scaleway.MnqCredentialSqsSnsCredentialsPermissionsArgs{
//						CanPublish: pulumi.Bool(true),
//						CanReceive: pulumi.Bool(true),
//						CanManage:  pulumi.Bool(true),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewMnqQueue(ctx, "myQueue", &scaleway.MnqQueueArgs{
//				NamespaceId: mainMnqNamespace.ID(),
//				Sqs: &scaleway.MnqQueueSqsArgs{
//					AccessKey: mainMnqCredential.SqsSnsCredentials.ApplyT(func(sqsSnsCredentials scaleway.MnqCredentialSqsSnsCredentials) (*string, error) {
//						return &sqsSnsCredentials.AccessKey, nil
//					}).(pulumi.StringPtrOutput),
//					SecretKey: mainMnqCredential.SqsSnsCredentials.ApplyT(func(sqsSnsCredentials scaleway.MnqCredentialSqsSnsCredentials) (*string, error) {
//						return &sqsSnsCredentials.SecretKey, nil
//					}).(pulumi.StringPtrOutput),
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
// <!--End PulumiCodeChooser -->
//
// ### Argument Reference
//
// The following arguments are supported:
//
// * `namespaceId` - (Required) The ID of the Namespace associated to.
//
// * `name` - (Optional) The name of the queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
//
// * `namePrefix` - (Optional) Creates a unique name beginning with the specified prefix. Conflicts with `name`.
//
// * `messageMaxAge` - (Optional) The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
//
// * `messageMaxSize` - (Optional) The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
//
// * `sqs` - (Optional) The SQS attributes of the queue. Conflicts with `nats`.
//   - `endpoint` - (Optional) The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `http://sqs-sns.mnq.{region}.scw.cloud`.
//   - `accessKey` - (Required) The access key of the SQS queue.
//   - `secretKey` - (Required) The secret key of the SQS queue.
//   - `fifoQueue` - (Optional) Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
//   - `contentBasedDeduplication` - (Optional) Specifies whether to enable content-based deduplication. Defaults to `false`.
//   - `receiveWaitTimeSeconds` - (Optional) The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
//   - `visibilityTimeoutSeconds` - (Optional) The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
//   - For more information about the SQS limitations, see [the documentation](https://www.scaleway.com/en/developers/api/messaging-and-queuing/#technical-limitations).
//
// * `nats` - (Optional) The NATS attributes of the queue. Conflicts with `sqs`.
//   - `endpoint` - (Optional) The endpoint of the NATS queue. Can contain a {region} placeholder. Defaults to `nats://nats.mnq.{region}.scw.cloud:4222`.
//   - `credentials` - (Required) Line jump separated key and seed.
//   - `retentionPolicy` - (Optional) The retention policy of the queue. See https://docs.nats.io/nats-concepts/jetstream/streams#retentionpolicy for more information. Defaults to `workqueue`.
//
// ### Attribute Reference
//
// In addition to all arguments above, the following attributes are exported:
//
//   - `sqs` - The SQS attributes of the queue.
//     ~ `url` - The URL of the queue.
//
// ### Import
//
// Queues can be imported using the `{region}/{namespace-id}/{queue-name}` format:
type MnqQueue struct {
	pulumi.CustomResourceState

	// The number of seconds the queue retains a message.
	MessageMaxAge pulumi.IntPtrOutput `pulumi:"messageMaxAge"`
	// The maximum size of a message. Should be in bytes.
	MessageMaxSize pulumi.IntPtrOutput `pulumi:"messageMaxSize"`
	// The name of the queue. Conflicts with name_prefix.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// The ID of the Namespace associated to
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	// The NATS attributes of the queue
	Nats MnqQueueNatsPtrOutput `pulumi:"nats"`
	// The SQS attributes of the queue
	Sqs MnqQueueSqsPtrOutput `pulumi:"sqs"`
}

// NewMnqQueue registers a new resource with the given unique name, arguments, and options.
func NewMnqQueue(ctx *pulumi.Context,
	name string, args *MnqQueueArgs, opts ...pulumi.ResourceOption) (*MnqQueue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MnqQueue
	err := ctx.RegisterResource("scaleway:index/mnqQueue:MnqQueue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMnqQueue gets an existing MnqQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMnqQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MnqQueueState, opts ...pulumi.ResourceOption) (*MnqQueue, error) {
	var resource MnqQueue
	err := ctx.ReadResource("scaleway:index/mnqQueue:MnqQueue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MnqQueue resources.
type mnqQueueState struct {
	// The number of seconds the queue retains a message.
	MessageMaxAge *int `pulumi:"messageMaxAge"`
	// The maximum size of a message. Should be in bytes.
	MessageMaxSize *int `pulumi:"messageMaxSize"`
	// The name of the queue. Conflicts with name_prefix.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	NamePrefix *string `pulumi:"namePrefix"`
	// The ID of the Namespace associated to
	NamespaceId *string `pulumi:"namespaceId"`
	// The NATS attributes of the queue
	Nats *MnqQueueNats `pulumi:"nats"`
	// The SQS attributes of the queue
	Sqs *MnqQueueSqs `pulumi:"sqs"`
}

type MnqQueueState struct {
	// The number of seconds the queue retains a message.
	MessageMaxAge pulumi.IntPtrInput
	// The maximum size of a message. Should be in bytes.
	MessageMaxSize pulumi.IntPtrInput
	// The name of the queue. Conflicts with name_prefix.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	NamePrefix pulumi.StringPtrInput
	// The ID of the Namespace associated to
	NamespaceId pulumi.StringPtrInput
	// The NATS attributes of the queue
	Nats MnqQueueNatsPtrInput
	// The SQS attributes of the queue
	Sqs MnqQueueSqsPtrInput
}

func (MnqQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqQueueState)(nil)).Elem()
}

type mnqQueueArgs struct {
	// The number of seconds the queue retains a message.
	MessageMaxAge *int `pulumi:"messageMaxAge"`
	// The maximum size of a message. Should be in bytes.
	MessageMaxSize *int `pulumi:"messageMaxSize"`
	// The name of the queue. Conflicts with name_prefix.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	NamePrefix *string `pulumi:"namePrefix"`
	// The ID of the Namespace associated to
	NamespaceId string `pulumi:"namespaceId"`
	// The NATS attributes of the queue
	Nats *MnqQueueNats `pulumi:"nats"`
	// The SQS attributes of the queue
	Sqs *MnqQueueSqs `pulumi:"sqs"`
}

// The set of arguments for constructing a MnqQueue resource.
type MnqQueueArgs struct {
	// The number of seconds the queue retains a message.
	MessageMaxAge pulumi.IntPtrInput
	// The maximum size of a message. Should be in bytes.
	MessageMaxSize pulumi.IntPtrInput
	// The name of the queue. Conflicts with name_prefix.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	NamePrefix pulumi.StringPtrInput
	// The ID of the Namespace associated to
	NamespaceId pulumi.StringInput
	// The NATS attributes of the queue
	Nats MnqQueueNatsPtrInput
	// The SQS attributes of the queue
	Sqs MnqQueueSqsPtrInput
}

func (MnqQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqQueueArgs)(nil)).Elem()
}

type MnqQueueInput interface {
	pulumi.Input

	ToMnqQueueOutput() MnqQueueOutput
	ToMnqQueueOutputWithContext(ctx context.Context) MnqQueueOutput
}

func (*MnqQueue) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqQueue)(nil)).Elem()
}

func (i *MnqQueue) ToMnqQueueOutput() MnqQueueOutput {
	return i.ToMnqQueueOutputWithContext(context.Background())
}

func (i *MnqQueue) ToMnqQueueOutputWithContext(ctx context.Context) MnqQueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqQueueOutput)
}

// MnqQueueArrayInput is an input type that accepts MnqQueueArray and MnqQueueArrayOutput values.
// You can construct a concrete instance of `MnqQueueArrayInput` via:
//
//	MnqQueueArray{ MnqQueueArgs{...} }
type MnqQueueArrayInput interface {
	pulumi.Input

	ToMnqQueueArrayOutput() MnqQueueArrayOutput
	ToMnqQueueArrayOutputWithContext(context.Context) MnqQueueArrayOutput
}

type MnqQueueArray []MnqQueueInput

func (MnqQueueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqQueue)(nil)).Elem()
}

func (i MnqQueueArray) ToMnqQueueArrayOutput() MnqQueueArrayOutput {
	return i.ToMnqQueueArrayOutputWithContext(context.Background())
}

func (i MnqQueueArray) ToMnqQueueArrayOutputWithContext(ctx context.Context) MnqQueueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqQueueArrayOutput)
}

// MnqQueueMapInput is an input type that accepts MnqQueueMap and MnqQueueMapOutput values.
// You can construct a concrete instance of `MnqQueueMapInput` via:
//
//	MnqQueueMap{ "key": MnqQueueArgs{...} }
type MnqQueueMapInput interface {
	pulumi.Input

	ToMnqQueueMapOutput() MnqQueueMapOutput
	ToMnqQueueMapOutputWithContext(context.Context) MnqQueueMapOutput
}

type MnqQueueMap map[string]MnqQueueInput

func (MnqQueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqQueue)(nil)).Elem()
}

func (i MnqQueueMap) ToMnqQueueMapOutput() MnqQueueMapOutput {
	return i.ToMnqQueueMapOutputWithContext(context.Background())
}

func (i MnqQueueMap) ToMnqQueueMapOutputWithContext(ctx context.Context) MnqQueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqQueueMapOutput)
}

type MnqQueueOutput struct{ *pulumi.OutputState }

func (MnqQueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqQueue)(nil)).Elem()
}

func (o MnqQueueOutput) ToMnqQueueOutput() MnqQueueOutput {
	return o
}

func (o MnqQueueOutput) ToMnqQueueOutputWithContext(ctx context.Context) MnqQueueOutput {
	return o
}

// The number of seconds the queue retains a message.
func (o MnqQueueOutput) MessageMaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MnqQueue) pulumi.IntPtrOutput { return v.MessageMaxAge }).(pulumi.IntPtrOutput)
}

// The maximum size of a message. Should be in bytes.
func (o MnqQueueOutput) MessageMaxSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MnqQueue) pulumi.IntPtrOutput { return v.MessageMaxSize }).(pulumi.IntPtrOutput)
}

// The name of the queue. Conflicts with name_prefix.
func (o MnqQueueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqQueue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified prefix. Conflicts with name.
func (o MnqQueueOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqQueue) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// The ID of the Namespace associated to
func (o MnqQueueOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqQueue) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

// The NATS attributes of the queue
func (o MnqQueueOutput) Nats() MnqQueueNatsPtrOutput {
	return o.ApplyT(func(v *MnqQueue) MnqQueueNatsPtrOutput { return v.Nats }).(MnqQueueNatsPtrOutput)
}

// The SQS attributes of the queue
func (o MnqQueueOutput) Sqs() MnqQueueSqsPtrOutput {
	return o.ApplyT(func(v *MnqQueue) MnqQueueSqsPtrOutput { return v.Sqs }).(MnqQueueSqsPtrOutput)
}

type MnqQueueArrayOutput struct{ *pulumi.OutputState }

func (MnqQueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqQueue)(nil)).Elem()
}

func (o MnqQueueArrayOutput) ToMnqQueueArrayOutput() MnqQueueArrayOutput {
	return o
}

func (o MnqQueueArrayOutput) ToMnqQueueArrayOutputWithContext(ctx context.Context) MnqQueueArrayOutput {
	return o
}

func (o MnqQueueArrayOutput) Index(i pulumi.IntInput) MnqQueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MnqQueue {
		return vs[0].([]*MnqQueue)[vs[1].(int)]
	}).(MnqQueueOutput)
}

type MnqQueueMapOutput struct{ *pulumi.OutputState }

func (MnqQueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqQueue)(nil)).Elem()
}

func (o MnqQueueMapOutput) ToMnqQueueMapOutput() MnqQueueMapOutput {
	return o
}

func (o MnqQueueMapOutput) ToMnqQueueMapOutputWithContext(ctx context.Context) MnqQueueMapOutput {
	return o
}

func (o MnqQueueMapOutput) MapIndex(k pulumi.StringInput) MnqQueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MnqQueue {
		return vs[0].(map[string]*MnqQueue)[vs[1].(string)]
	}).(MnqQueueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MnqQueueInput)(nil)).Elem(), &MnqQueue{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqQueueArrayInput)(nil)).Elem(), MnqQueueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqQueueMapInput)(nil)).Elem(), MnqQueueMap{})
	pulumi.RegisterOutputType(MnqQueueOutput{})
	pulumi.RegisterOutputType(MnqQueueArrayOutput{})
	pulumi.RegisterOutputType(MnqQueueMapOutput{})
}
