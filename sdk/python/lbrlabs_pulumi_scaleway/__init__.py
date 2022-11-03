# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .account_ssh_key import *
from .apple_slicon_valley_server import *
from .baremetal_server import *
from .container import *
from .container_cron import *
from .container_namespace import *
from .container_token import *
from .database import *
from .database_acl import *
from .database_backup import *
from .database_instance import *
from .database_privilege import *
from .database_read_replica import *
from .database_user import *
from .domain_record import *
from .domain_zone import *
from .flexible_ip import *
from .function import *
from .function_cron import *
from .function_namespace import *
from .get_account_ssh_key import *
from .get_baremetal_offer import *
from .get_baremetal_os import *
from .get_baremetal_server import *
from .get_container import *
from .get_container_namespace import *
from .get_database import *
from .get_database_acl import *
from .get_database_backup import *
from .get_database_instance import *
from .get_database_privilege import *
from .get_domain_record import *
from .get_domain_zone import *
from .get_flexible_ip import *
from .get_function import *
from .get_function_namespace import *
from .get_instance_image import *
from .get_instance_ip import *
from .get_instance_security_group import *
from .get_instance_server import *
from .get_instance_servers import *
from .get_instance_volume import *
from .get_iot_device import *
from .get_iot_hub import *
from .get_kubernetes_cluster import *
from .get_kubernetes_node_pool import *
from .get_loadbalancer import *
from .get_loadbalancer_certificate import *
from .get_loadbalancer_ip import *
from .get_marketplace_image import *
from .get_object_bucket import *
from .get_redis_cluster import *
from .get_registry_image import *
from .get_registry_namespace import *
from .get_vpc_gateway_network import *
from .get_vpc_private_network import *
from .get_vpc_public_gateway import *
from .get_vpc_public_gateway_dhcp import *
from .get_vpc_public_gateway_dhcp_reservation import *
from .get_vpc_public_gateway_ip import *
from .get_vpc_public_pat_rule import *
from .instance_image import *
from .instance_ip import *
from .instance_ip_reverse_dns import *
from .instance_placement_group import *
from .instance_private_nic import *
from .instance_security_group import *
from .instance_security_group_rules import *
from .instance_server import *
from .instance_snapshot import *
from .instance_user_data import *
from .instance_volume import *
from .iot_device import *
from .iot_hub import *
from .iot_network import *
from .iot_route import *
from .kubernetes_cluster import *
from .kubernetes_node_pool import *
from .loadbalancer import *
from .loadbalancer_backend import *
from .loadbalancer_certificate import *
from .loadbalancer_frontend import *
from .loadbalancer_ip import *
from .loadbalancer_route import *
from .object_bucket import *
from .object_bucket_acl import *
from .object_bucket_policy import *
from .object_bucket_website_configuration import *
from .object_item import *
from .provider import *
from .redis_cluster import *
from .registry_namespace import *
from .vpc_gateway_network import *
from .vpc_private_network import *
from .vpc_public_gateway import *
from .vpc_public_gateway_dhcp import *
from .vpc_public_gateway_dhcp_reservation import *
from .vpc_public_gateway_ip import *
from .vpc_public_gateway_pat_rule import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import lbrlabs_pulumi_scaleway.config as __config
    config = __config
else:
    config = _utilities.lazy_import('lbrlabs_pulumi_scaleway.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "scaleway",
  "mod": "index/accountSshKey",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/accountSshKey:AccountSshKey": "AccountSshKey"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/appleSliconValleyServer",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/appleSliconValleyServer:AppleSliconValleyServer": "AppleSliconValleyServer"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/baremetalServer",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/baremetalServer:BaremetalServer": "BaremetalServer"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/container",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/container:Container": "Container"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/containerCron",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/containerCron:ContainerCron": "ContainerCron"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/containerNamespace",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/containerNamespace:ContainerNamespace": "ContainerNamespace"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/containerToken",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/containerToken:ContainerToken": "ContainerToken"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/database",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/database:Database": "Database"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/databaseAcl",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/databaseAcl:DatabaseAcl": "DatabaseAcl"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/databaseBackup",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/databaseBackup:DatabaseBackup": "DatabaseBackup"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/databaseInstance",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/databaseInstance:DatabaseInstance": "DatabaseInstance"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/databasePrivilege",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/databasePrivilege:DatabasePrivilege": "DatabasePrivilege"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/databaseReadReplica",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/databaseReadReplica:DatabaseReadReplica": "DatabaseReadReplica"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/databaseUser",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/databaseUser:DatabaseUser": "DatabaseUser"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/domainRecord",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/domainRecord:DomainRecord": "DomainRecord"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/domainZone",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/domainZone:DomainZone": "DomainZone"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/flexibleIp",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/flexibleIp:FlexibleIp": "FlexibleIp"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/function",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/function:Function": "Function"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/functionCron",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/functionCron:FunctionCron": "FunctionCron"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/functionNamespace",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/functionNamespace:FunctionNamespace": "FunctionNamespace"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceImage",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/instanceImage:InstanceImage": "InstanceImage"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceIp",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/instanceIp:InstanceIp": "InstanceIp"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceIpReverseDns",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/instanceIpReverseDns:InstanceIpReverseDns": "InstanceIpReverseDns"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instancePlacementGroup",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/instancePlacementGroup:InstancePlacementGroup": "InstancePlacementGroup"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instancePrivateNic",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/instancePrivateNic:InstancePrivateNic": "InstancePrivateNic"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceSecurityGroup",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/instanceSecurityGroup:InstanceSecurityGroup": "InstanceSecurityGroup"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceSecurityGroupRules",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/instanceSecurityGroupRules:InstanceSecurityGroupRules": "InstanceSecurityGroupRules"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceServer",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/instanceServer:InstanceServer": "InstanceServer"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceSnapshot",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/instanceSnapshot:InstanceSnapshot": "InstanceSnapshot"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceUserData",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/instanceUserData:InstanceUserData": "InstanceUserData"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceVolume",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/instanceVolume:InstanceVolume": "InstanceVolume"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iotDevice",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/iotDevice:IotDevice": "IotDevice"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iotHub",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/iotHub:IotHub": "IotHub"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iotNetwork",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/iotNetwork:IotNetwork": "IotNetwork"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iotRoute",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/iotRoute:IotRoute": "IotRoute"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/kubernetesCluster",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/kubernetesCluster:KubernetesCluster": "KubernetesCluster"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/kubernetesNodePool",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/kubernetesNodePool:KubernetesNodePool": "KubernetesNodePool"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/loadbalancer",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/loadbalancer:Loadbalancer": "Loadbalancer"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/loadbalancerBackend",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/loadbalancerBackend:LoadbalancerBackend": "LoadbalancerBackend"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/loadbalancerCertificate",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/loadbalancerCertificate:LoadbalancerCertificate": "LoadbalancerCertificate"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/loadbalancerFrontend",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/loadbalancerFrontend:LoadbalancerFrontend": "LoadbalancerFrontend"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/loadbalancerIp",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/loadbalancerIp:LoadbalancerIp": "LoadbalancerIp"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/loadbalancerRoute",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/loadbalancerRoute:LoadbalancerRoute": "LoadbalancerRoute"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/objectBucket",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/objectBucket:ObjectBucket": "ObjectBucket"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/objectBucketAcl",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/objectBucketAcl:ObjectBucketAcl": "ObjectBucketAcl"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/objectBucketPolicy",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/objectBucketPolicy:ObjectBucketPolicy": "ObjectBucketPolicy"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/objectBucketWebsiteConfiguration",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration": "ObjectBucketWebsiteConfiguration"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/objectItem",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/objectItem:ObjectItem": "ObjectItem"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/redisCluster",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/redisCluster:RedisCluster": "RedisCluster"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/registryNamespace",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/registryNamespace:RegistryNamespace": "RegistryNamespace"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpcGatewayNetwork",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork": "VpcGatewayNetwork"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpcPrivateNetwork",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/vpcPrivateNetwork:VpcPrivateNetwork": "VpcPrivateNetwork"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpcPublicGateway",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/vpcPublicGateway:VpcPublicGateway": "VpcPublicGateway"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpcPublicGatewayDhcp",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/vpcPublicGatewayDhcp:VpcPublicGatewayDhcp": "VpcPublicGatewayDhcp"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpcPublicGatewayDhcpReservation",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/vpcPublicGatewayDhcpReservation:VpcPublicGatewayDhcpReservation": "VpcPublicGatewayDhcpReservation"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpcPublicGatewayIp",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp": "VpcPublicGatewayIp"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpcPublicGatewayPatRule",
  "fqn": "lbrlabs_pulumi_scaleway",
  "classes": {
   "scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule": "VpcPublicGatewayPatRule"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "scaleway",
  "token": "pulumi:providers:scaleway",
  "fqn": "lbrlabs_pulumi_scaleway",
  "class": "Provider"
 }
]
"""
)
