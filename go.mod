module github.com/pulumi/pulumi-vault

go 1.12

require (
	github.com/Azure/go-autorest/autorest/azure/auth v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.1.1
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.3.3
	github.com/pulumi/pulumi-terraform-bridge v1.2.1-0.20191030115615-68f8d85120cb
	github.com/terraform-providers/terraform-provider-vault v0.0.0-20191017211552-55806c9f74a4
	golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7 // indirect
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

replace github.com/pulumi/pulumi-terraform-bridge => ../pulumi-terraform-bridge
