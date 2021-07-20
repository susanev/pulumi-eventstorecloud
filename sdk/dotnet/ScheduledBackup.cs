// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Eventstorecloud
{
    /// <summary>
    /// Creates a new scheduled backup.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Eventstorecloud = Pulumi.Eventstorecloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var daily = new Eventstorecloud.ScheduledBackup("daily", new Eventstorecloud.ScheduledBackupArgs
    ///         {
    ///             ProjectId = eventstorecloud_project.Example.Id,
    ///             Schedule = "0 12 * * */1",
    ///             Description = "Creates a backup once a day at 12:00",
    ///             SourceClusterId = eventstorecloud_managed_cluster.Example.Id,
    ///             BackupDescription = "{cluster} Daily Backup {datetime:RFC3339}",
    ///             MaxBackupCount = 3,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [EventstorecloudResourceType("eventstorecloud:index/scheduledBackup:ScheduledBackup")]
    public partial class ScheduledBackup : Pulumi.CustomResource
    {
        /// <summary>
        /// backup_description
        /// </summary>
        [Output("backupDescription")]
        public Output<string> BackupDescription { get; private set; } = null!;

        /// <summary>
        /// Human readable description of the job
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The maximum number of backups to keep for this job
        /// </summary>
        [Output("maxBackupCount")]
        public Output<int> MaxBackupCount { get; private set; } = null!;

        /// <summary>
        /// ID of the project in which the backup exists
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Schedule for the backup, defined using restricted subset of cron
        /// </summary>
        [Output("schedule")]
        public Output<string> Schedule { get; private set; } = null!;

        /// <summary>
        /// the ID of the cluster to back up
        /// </summary>
        [Output("sourceClusterId")]
        public Output<string> SourceClusterId { get; private set; } = null!;


        /// <summary>
        /// Create a ScheduledBackup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScheduledBackup(string name, ScheduledBackupArgs args, CustomResourceOptions? options = null)
            : base("eventstorecloud:index/scheduledBackup:ScheduledBackup", name, args ?? new ScheduledBackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScheduledBackup(string name, Input<string> id, ScheduledBackupState? state = null, CustomResourceOptions? options = null)
            : base("eventstorecloud:index/scheduledBackup:ScheduledBackup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ScheduledBackup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScheduledBackup Get(string name, Input<string> id, ScheduledBackupState? state = null, CustomResourceOptions? options = null)
        {
            return new ScheduledBackup(name, id, state, options);
        }
    }

    public sealed class ScheduledBackupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// backup_description
        /// </summary>
        [Input("backupDescription", required: true)]
        public Input<string> BackupDescription { get; set; } = null!;

        /// <summary>
        /// Human readable description of the job
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The maximum number of backups to keep for this job
        /// </summary>
        [Input("maxBackupCount", required: true)]
        public Input<int> MaxBackupCount { get; set; } = null!;

        /// <summary>
        /// ID of the project in which the backup exists
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Schedule for the backup, defined using restricted subset of cron
        /// </summary>
        [Input("schedule", required: true)]
        public Input<string> Schedule { get; set; } = null!;

        /// <summary>
        /// the ID of the cluster to back up
        /// </summary>
        [Input("sourceClusterId", required: true)]
        public Input<string> SourceClusterId { get; set; } = null!;

        public ScheduledBackupArgs()
        {
        }
    }

    public sealed class ScheduledBackupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// backup_description
        /// </summary>
        [Input("backupDescription")]
        public Input<string>? BackupDescription { get; set; }

        /// <summary>
        /// Human readable description of the job
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The maximum number of backups to keep for this job
        /// </summary>
        [Input("maxBackupCount")]
        public Input<int>? MaxBackupCount { get; set; }

        /// <summary>
        /// ID of the project in which the backup exists
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Schedule for the backup, defined using restricted subset of cron
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// the ID of the cluster to back up
        /// </summary>
        [Input("sourceClusterId")]
        public Input<string>? SourceClusterId { get; set; }

        public ScheduledBackupState()
        {
        }
    }
}
