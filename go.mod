module github.com/pulumi/pulumi-tencentcloud

go 1.12

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/hashicorp/terraform-plugin-sdk v1.14.0
	github.com/pulumi/pulumi v1.8.0
	github.com/pulumi/pulumi-terraform-bridge v1.5.2
	github.com/stretchr/testify v1.5.1
	github.com/tencentcloudstack/terraform-provider-tencentcloud v1.56.8
	github.com/terraform-providers/terraform-provider-tencentcloud v1.38.3 // indirect
)
