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

// Creates and manages Scaleway object storage objects.
// For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).
//
// ## Import
//
// Objects can be imported using the `{region}/{bucketName}/{objectKey}` identifier, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/objectItem:ObjectItem some_object fr-par/some-bucket/some-file
//
// ```
type ObjectItem struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The content of the file to upload. Only one of `file`, `content` or `contentBase64` can be defined.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// The base64-encoded content of the file to upload. Only one of `file`, `content` or `contentBase64` can be defined.
	ContentBase64 pulumi.StringPtrOutput `pulumi:"contentBase64"`
	// The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `contentBase64` can be defined.
	File pulumi.StringPtrOutput `pulumi:"file"`
	// Hash of the file, used to trigger upload on file change
	Hash pulumi.StringPtrOutput `pulumi:"hash"`
	// The path of the object.
	Key pulumi.StringOutput `pulumi:"key"`
	// Map of metadata used for the object, keys must be lowercase
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Scaleway region this bucket resides in.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
	StorageClass pulumi.StringPtrOutput `pulumi:"storageClass"`
	// Map of tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Visibility of the object, `public-read` or `private`
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewObjectItem registers a new resource with the given unique name, arguments, and options.
func NewObjectItem(ctx *pulumi.Context,
	name string, args *ObjectItemArgs, opts ...pulumi.ResourceOption) (*ObjectItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ObjectItem
	err := ctx.RegisterResource("scaleway:index/objectItem:ObjectItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectItem gets an existing ObjectItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectItemState, opts ...pulumi.ResourceOption) (*ObjectItem, error) {
	var resource ObjectItem
	err := ctx.ReadResource("scaleway:index/objectItem:ObjectItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectItem resources.
type objectItemState struct {
	// The name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// The content of the file to upload. Only one of `file`, `content` or `contentBase64` can be defined.
	Content *string `pulumi:"content"`
	// The base64-encoded content of the file to upload. Only one of `file`, `content` or `contentBase64` can be defined.
	ContentBase64 *string `pulumi:"contentBase64"`
	// The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `contentBase64` can be defined.
	File *string `pulumi:"file"`
	// Hash of the file, used to trigger upload on file change
	Hash *string `pulumi:"hash"`
	// The path of the object.
	Key *string `pulumi:"key"`
	// Map of metadata used for the object, keys must be lowercase
	Metadata map[string]string `pulumi:"metadata"`
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The Scaleway region this bucket resides in.
	Region *string `pulumi:"region"`
	// Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
	StorageClass *string `pulumi:"storageClass"`
	// Map of tags
	Tags map[string]string `pulumi:"tags"`
	// Visibility of the object, `public-read` or `private`
	Visibility *string `pulumi:"visibility"`
}

type ObjectItemState struct {
	// The name of the bucket.
	Bucket pulumi.StringPtrInput
	// The content of the file to upload. Only one of `file`, `content` or `contentBase64` can be defined.
	Content pulumi.StringPtrInput
	// The base64-encoded content of the file to upload. Only one of `file`, `content` or `contentBase64` can be defined.
	ContentBase64 pulumi.StringPtrInput
	// The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `contentBase64` can be defined.
	File pulumi.StringPtrInput
	// Hash of the file, used to trigger upload on file change
	Hash pulumi.StringPtrInput
	// The path of the object.
	Key pulumi.StringPtrInput
	// Map of metadata used for the object, keys must be lowercase
	Metadata pulumi.StringMapInput
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId pulumi.StringPtrInput
	// The Scaleway region this bucket resides in.
	Region pulumi.StringPtrInput
	// Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
	StorageClass pulumi.StringPtrInput
	// Map of tags
	Tags pulumi.StringMapInput
	// Visibility of the object, `public-read` or `private`
	Visibility pulumi.StringPtrInput
}

func (ObjectItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectItemState)(nil)).Elem()
}

type objectItemArgs struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// The content of the file to upload. Only one of `file`, `content` or `contentBase64` can be defined.
	Content *string `pulumi:"content"`
	// The base64-encoded content of the file to upload. Only one of `file`, `content` or `contentBase64` can be defined.
	ContentBase64 *string `pulumi:"contentBase64"`
	// The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `contentBase64` can be defined.
	File *string `pulumi:"file"`
	// Hash of the file, used to trigger upload on file change
	Hash *string `pulumi:"hash"`
	// The path of the object.
	Key string `pulumi:"key"`
	// Map of metadata used for the object, keys must be lowercase
	Metadata map[string]string `pulumi:"metadata"`
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The Scaleway region this bucket resides in.
	Region *string `pulumi:"region"`
	// Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
	StorageClass *string `pulumi:"storageClass"`
	// Map of tags
	Tags map[string]string `pulumi:"tags"`
	// Visibility of the object, `public-read` or `private`
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a ObjectItem resource.
type ObjectItemArgs struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// The content of the file to upload. Only one of `file`, `content` or `contentBase64` can be defined.
	Content pulumi.StringPtrInput
	// The base64-encoded content of the file to upload. Only one of `file`, `content` or `contentBase64` can be defined.
	ContentBase64 pulumi.StringPtrInput
	// The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `contentBase64` can be defined.
	File pulumi.StringPtrInput
	// Hash of the file, used to trigger upload on file change
	Hash pulumi.StringPtrInput
	// The path of the object.
	Key pulumi.StringInput
	// Map of metadata used for the object, keys must be lowercase
	Metadata pulumi.StringMapInput
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId pulumi.StringPtrInput
	// The Scaleway region this bucket resides in.
	Region pulumi.StringPtrInput
	// Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
	StorageClass pulumi.StringPtrInput
	// Map of tags
	Tags pulumi.StringMapInput
	// Visibility of the object, `public-read` or `private`
	Visibility pulumi.StringPtrInput
}

func (ObjectItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectItemArgs)(nil)).Elem()
}

type ObjectItemInput interface {
	pulumi.Input

	ToObjectItemOutput() ObjectItemOutput
	ToObjectItemOutputWithContext(ctx context.Context) ObjectItemOutput
}

func (*ObjectItem) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectItem)(nil)).Elem()
}

func (i *ObjectItem) ToObjectItemOutput() ObjectItemOutput {
	return i.ToObjectItemOutputWithContext(context.Background())
}

func (i *ObjectItem) ToObjectItemOutputWithContext(ctx context.Context) ObjectItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectItemOutput)
}

// ObjectItemArrayInput is an input type that accepts ObjectItemArray and ObjectItemArrayOutput values.
// You can construct a concrete instance of `ObjectItemArrayInput` via:
//
//	ObjectItemArray{ ObjectItemArgs{...} }
type ObjectItemArrayInput interface {
	pulumi.Input

	ToObjectItemArrayOutput() ObjectItemArrayOutput
	ToObjectItemArrayOutputWithContext(context.Context) ObjectItemArrayOutput
}

type ObjectItemArray []ObjectItemInput

func (ObjectItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectItem)(nil)).Elem()
}

func (i ObjectItemArray) ToObjectItemArrayOutput() ObjectItemArrayOutput {
	return i.ToObjectItemArrayOutputWithContext(context.Background())
}

func (i ObjectItemArray) ToObjectItemArrayOutputWithContext(ctx context.Context) ObjectItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectItemArrayOutput)
}

// ObjectItemMapInput is an input type that accepts ObjectItemMap and ObjectItemMapOutput values.
// You can construct a concrete instance of `ObjectItemMapInput` via:
//
//	ObjectItemMap{ "key": ObjectItemArgs{...} }
type ObjectItemMapInput interface {
	pulumi.Input

	ToObjectItemMapOutput() ObjectItemMapOutput
	ToObjectItemMapOutputWithContext(context.Context) ObjectItemMapOutput
}

type ObjectItemMap map[string]ObjectItemInput

func (ObjectItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectItem)(nil)).Elem()
}

func (i ObjectItemMap) ToObjectItemMapOutput() ObjectItemMapOutput {
	return i.ToObjectItemMapOutputWithContext(context.Background())
}

func (i ObjectItemMap) ToObjectItemMapOutputWithContext(ctx context.Context) ObjectItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectItemMapOutput)
}

type ObjectItemOutput struct{ *pulumi.OutputState }

func (ObjectItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectItem)(nil)).Elem()
}

func (o ObjectItemOutput) ToObjectItemOutput() ObjectItemOutput {
	return o
}

func (o ObjectItemOutput) ToObjectItemOutputWithContext(ctx context.Context) ObjectItemOutput {
	return o
}

// The name of the bucket.
func (o ObjectItemOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectItem) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The content of the file to upload. Only one of `file`, `content` or `contentBase64` can be defined.
func (o ObjectItemOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectItem) pulumi.StringPtrOutput { return v.Content }).(pulumi.StringPtrOutput)
}

// The base64-encoded content of the file to upload. Only one of `file`, `content` or `contentBase64` can be defined.
func (o ObjectItemOutput) ContentBase64() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectItem) pulumi.StringPtrOutput { return v.ContentBase64 }).(pulumi.StringPtrOutput)
}

// The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `contentBase64` can be defined.
func (o ObjectItemOutput) File() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectItem) pulumi.StringPtrOutput { return v.File }).(pulumi.StringPtrOutput)
}

// Hash of the file, used to trigger upload on file change
func (o ObjectItemOutput) Hash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectItem) pulumi.StringPtrOutput { return v.Hash }).(pulumi.StringPtrOutput)
}

// The path of the object.
func (o ObjectItemOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectItem) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Map of metadata used for the object, keys must be lowercase
func (o ObjectItemOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ObjectItem) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// `projectId`) The ID of the project the bucket is associated with.
func (o ObjectItemOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectItem) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Scaleway region this bucket resides in.
func (o ObjectItemOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectItem) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
func (o ObjectItemOutput) StorageClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectItem) pulumi.StringPtrOutput { return v.StorageClass }).(pulumi.StringPtrOutput)
}

// Map of tags
func (o ObjectItemOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ObjectItem) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Visibility of the object, `public-read` or `private`
func (o ObjectItemOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectItem) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type ObjectItemArrayOutput struct{ *pulumi.OutputState }

func (ObjectItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectItem)(nil)).Elem()
}

func (o ObjectItemArrayOutput) ToObjectItemArrayOutput() ObjectItemArrayOutput {
	return o
}

func (o ObjectItemArrayOutput) ToObjectItemArrayOutputWithContext(ctx context.Context) ObjectItemArrayOutput {
	return o
}

func (o ObjectItemArrayOutput) Index(i pulumi.IntInput) ObjectItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObjectItem {
		return vs[0].([]*ObjectItem)[vs[1].(int)]
	}).(ObjectItemOutput)
}

type ObjectItemMapOutput struct{ *pulumi.OutputState }

func (ObjectItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectItem)(nil)).Elem()
}

func (o ObjectItemMapOutput) ToObjectItemMapOutput() ObjectItemMapOutput {
	return o
}

func (o ObjectItemMapOutput) ToObjectItemMapOutputWithContext(ctx context.Context) ObjectItemMapOutput {
	return o
}

func (o ObjectItemMapOutput) MapIndex(k pulumi.StringInput) ObjectItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObjectItem {
		return vs[0].(map[string]*ObjectItem)[vs[1].(string)]
	}).(ObjectItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectItemInput)(nil)).Elem(), &ObjectItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectItemArrayInput)(nil)).Elem(), ObjectItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectItemMapInput)(nil)).Elem(), ObjectItemMap{})
	pulumi.RegisterOutputType(ObjectItemOutput{})
	pulumi.RegisterOutputType(ObjectItemArrayOutput{})
	pulumi.RegisterOutputType(ObjectItemMapOutput{})
}
