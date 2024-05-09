// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ram

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	ram_sdkv1 "github.com/aws/aws-sdk-go/service/ram"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceResourceShare,
			TypeName: "aws_ram_resource_share",
			Name:     "Resource Shared",
			Tags:     &types.ServicePackageResourceTags{},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourcePrincipalAssociation,
			TypeName: "aws_ram_principal_association",
			Name:     "Principal Association",
		},
		{
			Factory:  resourceResourceAssociation,
			TypeName: "aws_ram_resource_association",
			Name:     "Resource Association",
		},
		{
			Factory:  resourceResourceShare,
			TypeName: "aws_ram_resource_share",
			Name:     "Resource Share",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceResourceShareAccepter,
			TypeName: "aws_ram_resource_share_accepter",
			Name:     "Resource Share Accepter",
		},
		{
			Factory:  resourceSharingWithOrganization,
			TypeName: "aws_ram_sharing_with_organization",
			Name:     "Sharing With Organization",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.RAM
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*ram_sdkv1.RAM, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return ram_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config[names.AttrEndpoint].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
