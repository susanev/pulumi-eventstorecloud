module github.com/EventStore/pulumi-eventstorecloud/provider

go 1.16

replace github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e

require (
	github.com/EventStore/terraform-provider-eventstorecloud v1.5.5
	github.com/hashicorp/hcl/v2 v2.3.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.1 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.2.1
	github.com/pulumi/pulumi/sdk/v3 v3.4.0
	github.com/terraform-providers/terraform-provider-archive v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20201203163018-be400aefbc4c // indirect
	gopkg.in/src-d/go-git.v4 v4.13.1 // indirect
)
