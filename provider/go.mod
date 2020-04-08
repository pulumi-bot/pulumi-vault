module github.com/pulumi/pulumi-vault/provider

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.4.1
	github.com/pulumi/pulumi-terraform-bridge v1.8.4
	github.com/pulumi/pulumi/sdk v1.14.1-0.20200402002223-a0f615ad0938
	github.com/terraform-providers/terraform-provider-vault v1.9.1-0.20200403230059-d29ffdba020f
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

replace github.com/pulumi/pulumi-terraform-bridge => ../../pulumi-terraform-bridge
