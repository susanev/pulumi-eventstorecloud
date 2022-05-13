// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventstorecloud

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new scheduled backup.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-eventstorecloud/sdk/go/eventstorecloud"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := eventstorecloud.NewScheduledBackup(ctx, "daily", &eventstorecloud.ScheduledBackupArgs{
// 			ProjectId:         pulumi.Any(eventstorecloud_project.Example.Id),
// 			Schedule:          pulumi.String("0 12 * * */1"),
// 			Description:       pulumi.String("Creates a backup once a day at 12:00"),
// 			SourceClusterId:   pulumi.Any(eventstorecloud_managed_cluster.Example.Id),
// 			BackupDescription: pulumi.String("{cluster} Daily Backup {datetime:RFC3339}"),
// 			MaxBackupCount:    pulumi.Int(3),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ```sh
//  $ pulumi import eventstorecloud:index/scheduledBackup:ScheduledBackup daily project_id:backup_id
// ```
type ScheduledBackup struct {
	pulumi.CustomResourceState

	// backup_description
	BackupDescription pulumi.StringOutput `pulumi:"backupDescription"`
	// Human readable description of the job
	Description pulumi.StringOutput `pulumi:"description"`
	// The maximum number of backups to keep for this job
	MaxBackupCount pulumi.IntOutput `pulumi:"maxBackupCount"`
	// ID of the project in which the backup exists
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Schedule for the backup, defined using restricted subset of cron
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// the ID of the cluster to back up
	SourceClusterId pulumi.StringOutput `pulumi:"sourceClusterId"`
}

// NewScheduledBackup registers a new resource with the given unique name, arguments, and options.
func NewScheduledBackup(ctx *pulumi.Context,
	name string, args *ScheduledBackupArgs, opts ...pulumi.ResourceOption) (*ScheduledBackup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupDescription == nil {
		return nil, errors.New("invalid value for required argument 'BackupDescription'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.MaxBackupCount == nil {
		return nil, errors.New("invalid value for required argument 'MaxBackupCount'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	if args.SourceClusterId == nil {
		return nil, errors.New("invalid value for required argument 'SourceClusterId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ScheduledBackup
	err := ctx.RegisterResource("eventstorecloud:index/scheduledBackup:ScheduledBackup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduledBackup gets an existing ScheduledBackup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduledBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledBackupState, opts ...pulumi.ResourceOption) (*ScheduledBackup, error) {
	var resource ScheduledBackup
	err := ctx.ReadResource("eventstorecloud:index/scheduledBackup:ScheduledBackup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduledBackup resources.
type scheduledBackupState struct {
	// backup_description
	BackupDescription *string `pulumi:"backupDescription"`
	// Human readable description of the job
	Description *string `pulumi:"description"`
	// The maximum number of backups to keep for this job
	MaxBackupCount *int `pulumi:"maxBackupCount"`
	// ID of the project in which the backup exists
	ProjectId *string `pulumi:"projectId"`
	// Schedule for the backup, defined using restricted subset of cron
	Schedule *string `pulumi:"schedule"`
	// the ID of the cluster to back up
	SourceClusterId *string `pulumi:"sourceClusterId"`
}

type ScheduledBackupState struct {
	// backup_description
	BackupDescription pulumi.StringPtrInput
	// Human readable description of the job
	Description pulumi.StringPtrInput
	// The maximum number of backups to keep for this job
	MaxBackupCount pulumi.IntPtrInput
	// ID of the project in which the backup exists
	ProjectId pulumi.StringPtrInput
	// Schedule for the backup, defined using restricted subset of cron
	Schedule pulumi.StringPtrInput
	// the ID of the cluster to back up
	SourceClusterId pulumi.StringPtrInput
}

func (ScheduledBackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledBackupState)(nil)).Elem()
}

type scheduledBackupArgs struct {
	// backup_description
	BackupDescription string `pulumi:"backupDescription"`
	// Human readable description of the job
	Description string `pulumi:"description"`
	// The maximum number of backups to keep for this job
	MaxBackupCount int `pulumi:"maxBackupCount"`
	// ID of the project in which the backup exists
	ProjectId string `pulumi:"projectId"`
	// Schedule for the backup, defined using restricted subset of cron
	Schedule string `pulumi:"schedule"`
	// the ID of the cluster to back up
	SourceClusterId string `pulumi:"sourceClusterId"`
}

// The set of arguments for constructing a ScheduledBackup resource.
type ScheduledBackupArgs struct {
	// backup_description
	BackupDescription pulumi.StringInput
	// Human readable description of the job
	Description pulumi.StringInput
	// The maximum number of backups to keep for this job
	MaxBackupCount pulumi.IntInput
	// ID of the project in which the backup exists
	ProjectId pulumi.StringInput
	// Schedule for the backup, defined using restricted subset of cron
	Schedule pulumi.StringInput
	// the ID of the cluster to back up
	SourceClusterId pulumi.StringInput
}

func (ScheduledBackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledBackupArgs)(nil)).Elem()
}

type ScheduledBackupInput interface {
	pulumi.Input

	ToScheduledBackupOutput() ScheduledBackupOutput
	ToScheduledBackupOutputWithContext(ctx context.Context) ScheduledBackupOutput
}

func (*ScheduledBackup) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledBackup)(nil)).Elem()
}

func (i *ScheduledBackup) ToScheduledBackupOutput() ScheduledBackupOutput {
	return i.ToScheduledBackupOutputWithContext(context.Background())
}

func (i *ScheduledBackup) ToScheduledBackupOutputWithContext(ctx context.Context) ScheduledBackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledBackupOutput)
}

// ScheduledBackupArrayInput is an input type that accepts ScheduledBackupArray and ScheduledBackupArrayOutput values.
// You can construct a concrete instance of `ScheduledBackupArrayInput` via:
//
//          ScheduledBackupArray{ ScheduledBackupArgs{...} }
type ScheduledBackupArrayInput interface {
	pulumi.Input

	ToScheduledBackupArrayOutput() ScheduledBackupArrayOutput
	ToScheduledBackupArrayOutputWithContext(context.Context) ScheduledBackupArrayOutput
}

type ScheduledBackupArray []ScheduledBackupInput

func (ScheduledBackupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScheduledBackup)(nil)).Elem()
}

func (i ScheduledBackupArray) ToScheduledBackupArrayOutput() ScheduledBackupArrayOutput {
	return i.ToScheduledBackupArrayOutputWithContext(context.Background())
}

func (i ScheduledBackupArray) ToScheduledBackupArrayOutputWithContext(ctx context.Context) ScheduledBackupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledBackupArrayOutput)
}

// ScheduledBackupMapInput is an input type that accepts ScheduledBackupMap and ScheduledBackupMapOutput values.
// You can construct a concrete instance of `ScheduledBackupMapInput` via:
//
//          ScheduledBackupMap{ "key": ScheduledBackupArgs{...} }
type ScheduledBackupMapInput interface {
	pulumi.Input

	ToScheduledBackupMapOutput() ScheduledBackupMapOutput
	ToScheduledBackupMapOutputWithContext(context.Context) ScheduledBackupMapOutput
}

type ScheduledBackupMap map[string]ScheduledBackupInput

func (ScheduledBackupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScheduledBackup)(nil)).Elem()
}

func (i ScheduledBackupMap) ToScheduledBackupMapOutput() ScheduledBackupMapOutput {
	return i.ToScheduledBackupMapOutputWithContext(context.Background())
}

func (i ScheduledBackupMap) ToScheduledBackupMapOutputWithContext(ctx context.Context) ScheduledBackupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledBackupMapOutput)
}

type ScheduledBackupOutput struct{ *pulumi.OutputState }

func (ScheduledBackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledBackup)(nil)).Elem()
}

func (o ScheduledBackupOutput) ToScheduledBackupOutput() ScheduledBackupOutput {
	return o
}

func (o ScheduledBackupOutput) ToScheduledBackupOutputWithContext(ctx context.Context) ScheduledBackupOutput {
	return o
}

type ScheduledBackupArrayOutput struct{ *pulumi.OutputState }

func (ScheduledBackupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScheduledBackup)(nil)).Elem()
}

func (o ScheduledBackupArrayOutput) ToScheduledBackupArrayOutput() ScheduledBackupArrayOutput {
	return o
}

func (o ScheduledBackupArrayOutput) ToScheduledBackupArrayOutputWithContext(ctx context.Context) ScheduledBackupArrayOutput {
	return o
}

func (o ScheduledBackupArrayOutput) Index(i pulumi.IntInput) ScheduledBackupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScheduledBackup {
		return vs[0].([]*ScheduledBackup)[vs[1].(int)]
	}).(ScheduledBackupOutput)
}

type ScheduledBackupMapOutput struct{ *pulumi.OutputState }

func (ScheduledBackupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScheduledBackup)(nil)).Elem()
}

func (o ScheduledBackupMapOutput) ToScheduledBackupMapOutput() ScheduledBackupMapOutput {
	return o
}

func (o ScheduledBackupMapOutput) ToScheduledBackupMapOutputWithContext(ctx context.Context) ScheduledBackupMapOutput {
	return o
}

func (o ScheduledBackupMapOutput) MapIndex(k pulumi.StringInput) ScheduledBackupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScheduledBackup {
		return vs[0].(map[string]*ScheduledBackup)[vs[1].(string)]
	}).(ScheduledBackupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledBackupInput)(nil)).Elem(), &ScheduledBackup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledBackupArrayInput)(nil)).Elem(), ScheduledBackupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledBackupMapInput)(nil)).Elem(), ScheduledBackupMap{})
	pulumi.RegisterOutputType(ScheduledBackupOutput{})
	pulumi.RegisterOutputType(ScheduledBackupArrayOutput{})
	pulumi.RegisterOutputType(ScheduledBackupMapOutput{})
}
