// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "scaleway:index/accountSshKey:AccountSshKey":
		r = &AccountSshKey{}
	case "scaleway:index/appleSliconValleyServer:AppleSliconValleyServer":
		r = &AppleSliconValleyServer{}
	case "scaleway:index/baremetalServer:BaremetalServer":
		r = &BaremetalServer{}
	case "scaleway:index/container:Container":
		r = &Container{}
	case "scaleway:index/containerCron:ContainerCron":
		r = &ContainerCron{}
	case "scaleway:index/containerNamespace:ContainerNamespace":
		r = &ContainerNamespace{}
	case "scaleway:index/containerToken:ContainerToken":
		r = &ContainerToken{}
	case "scaleway:index/database:Database":
		r = &Database{}
	case "scaleway:index/databaseAcl:DatabaseAcl":
		r = &DatabaseAcl{}
	case "scaleway:index/databaseBackup:DatabaseBackup":
		r = &DatabaseBackup{}
	case "scaleway:index/databaseInstance:DatabaseInstance":
		r = &DatabaseInstance{}
	case "scaleway:index/databasePrivilege:DatabasePrivilege":
		r = &DatabasePrivilege{}
	case "scaleway:index/databaseReadReplica:DatabaseReadReplica":
		r = &DatabaseReadReplica{}
	case "scaleway:index/databaseUser:DatabaseUser":
		r = &DatabaseUser{}
	case "scaleway:index/domainRecord:DomainRecord":
		r = &DomainRecord{}
	case "scaleway:index/domainZone:DomainZone":
		r = &DomainZone{}
	case "scaleway:index/flexibleIp:FlexibleIp":
		r = &FlexibleIp{}
	case "scaleway:index/function:Function":
		r = &Function{}
	case "scaleway:index/functionCron:FunctionCron":
		r = &FunctionCron{}
	case "scaleway:index/functionNamespace:FunctionNamespace":
		r = &FunctionNamespace{}
	case "scaleway:index/instanceImage:InstanceImage":
		r = &InstanceImage{}
	case "scaleway:index/instanceIp:InstanceIp":
		r = &InstanceIp{}
	case "scaleway:index/instanceIpReverseDns:InstanceIpReverseDns":
		r = &InstanceIpReverseDns{}
	case "scaleway:index/instancePlacementGroup:InstancePlacementGroup":
		r = &InstancePlacementGroup{}
	case "scaleway:index/instancePrivateNic:InstancePrivateNic":
		r = &InstancePrivateNic{}
	case "scaleway:index/instanceSecurityGroup:InstanceSecurityGroup":
		r = &InstanceSecurityGroup{}
	case "scaleway:index/instanceSecurityGroupRules:InstanceSecurityGroupRules":
		r = &InstanceSecurityGroupRules{}
	case "scaleway:index/instanceServer:InstanceServer":
		r = &InstanceServer{}
	case "scaleway:index/instanceSnapshot:InstanceSnapshot":
		r = &InstanceSnapshot{}
	case "scaleway:index/instanceUserData:InstanceUserData":
		r = &InstanceUserData{}
	case "scaleway:index/instanceVolume:InstanceVolume":
		r = &InstanceVolume{}
	case "scaleway:index/iotDevice:IotDevice":
		r = &IotDevice{}
	case "scaleway:index/iotHub:IotHub":
		r = &IotHub{}
	case "scaleway:index/iotNetwork:IotNetwork":
		r = &IotNetwork{}
	case "scaleway:index/iotRoute:IotRoute":
		r = &IotRoute{}
	case "scaleway:index/kubernetesCluster:KubernetesCluster":
		r = &KubernetesCluster{}
	case "scaleway:index/kubernetesNodePool:KubernetesNodePool":
		r = &KubernetesNodePool{}
	case "scaleway:index/loadbalancer:Loadbalancer":
		r = &Loadbalancer{}
	case "scaleway:index/loadbalancerBackend:LoadbalancerBackend":
		r = &LoadbalancerBackend{}
	case "scaleway:index/loadbalancerCertificate:LoadbalancerCertificate":
		r = &LoadbalancerCertificate{}
	case "scaleway:index/loadbalancerFrontend:LoadbalancerFrontend":
		r = &LoadbalancerFrontend{}
	case "scaleway:index/loadbalancerIp:LoadbalancerIp":
		r = &LoadbalancerIp{}
	case "scaleway:index/loadbalancerRoute:LoadbalancerRoute":
		r = &LoadbalancerRoute{}
	case "scaleway:index/objectBucket:ObjectBucket":
		r = &ObjectBucket{}
	case "scaleway:index/objectBucketAcl:ObjectBucketAcl":
		r = &ObjectBucketAcl{}
	case "scaleway:index/objectBucketPolicy:ObjectBucketPolicy":
		r = &ObjectBucketPolicy{}
	case "scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration":
		r = &ObjectBucketWebsiteConfiguration{}
	case "scaleway:index/objectItem:ObjectItem":
		r = &ObjectItem{}
	case "scaleway:index/redisCluster:RedisCluster":
		r = &RedisCluster{}
	case "scaleway:index/registryNamespace:RegistryNamespace":
		r = &RegistryNamespace{}
	case "scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork":
		r = &VpcGatewayNetwork{}
	case "scaleway:index/vpcPrivateNetwork:VpcPrivateNetwork":
		r = &VpcPrivateNetwork{}
	case "scaleway:index/vpcPublicGateway:VpcPublicGateway":
		r = &VpcPublicGateway{}
	case "scaleway:index/vpcPublicGatewayDhcp:VpcPublicGatewayDhcp":
		r = &VpcPublicGatewayDhcp{}
	case "scaleway:index/vpcPublicGatewayDhcpReservation:VpcPublicGatewayDhcpReservation":
		r = &VpcPublicGatewayDhcpReservation{}
	case "scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp":
		r = &VpcPublicGatewayIp{}
	case "scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule":
		r = &VpcPublicGatewayPatRule{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:scaleway" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/accountSshKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/appleSliconValleyServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/baremetalServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/container",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/containerCron",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/containerNamespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/containerToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/database",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/databaseAcl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/databaseBackup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/databaseInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/databasePrivilege",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/databaseReadReplica",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/databaseUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/domainRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/domainZone",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/flexibleIp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/function",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/functionCron",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/functionNamespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceImage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceIp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceIpReverseDns",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instancePlacementGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instancePrivateNic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceSecurityGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceSecurityGroupRules",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceSnapshot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceUserData",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceVolume",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/iotDevice",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/iotHub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/iotNetwork",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/iotRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/kubernetesCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/kubernetesNodePool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/loadbalancer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/loadbalancerBackend",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/loadbalancerCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/loadbalancerFrontend",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/loadbalancerIp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/loadbalancerRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/objectBucket",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/objectBucketAcl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/objectBucketPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/objectBucketWebsiteConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/objectItem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/redisCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/registryNamespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpcGatewayNetwork",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpcPrivateNetwork",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpcPublicGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpcPublicGatewayDhcp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpcPublicGatewayDhcpReservation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpcPublicGatewayIp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpcPublicGatewayPatRule",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"scaleway",
		&pkg{version},
	)
}
