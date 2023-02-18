// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface BaremetalServerIp {
    /**
     * The address of the IP.
     */
    address?: pulumi.Input<string>;
    /**
     * The id of the private network to attach.
     */
    id?: pulumi.Input<string>;
    /**
     * The reverse of the IP.
     */
    reverse?: pulumi.Input<string>;
    version?: pulumi.Input<string>;
}

export interface BaremetalServerOption {
    /**
     * The auto expiration date for compatible options
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The id of the private network to attach.
     */
    id: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    name?: pulumi.Input<string>;
}

export interface BaremetalServerPrivateNetwork {
    /**
     * The date and time of the creation of the private network.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The id of the private network to attach.
     */
    id: pulumi.Input<string>;
    /**
     * The private network status.
     */
    status?: pulumi.Input<string>;
    /**
     * The date and time of the last update of the private network.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The VLAN ID associated to the private network.
     */
    vlan?: pulumi.Input<number>;
}

export interface DatabaseAclAclRule {
    /**
     * A text describing this rule. Default description: `IP allowed`
     */
    description?: pulumi.Input<string>;
    /**
     * The ip range to whitelist in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)
     */
    ip: pulumi.Input<string>;
}

export interface DatabaseInstanceLoadBalancer {
    /**
     * The ID of the endpoint of the private network.
     */
    endpointId?: pulumi.Input<string>;
    /**
     * Name of the endpoint.
     */
    hostname?: pulumi.Input<string>;
    /**
     * IP of the endpoint.
     */
    ip?: pulumi.Input<string>;
    /**
     * The name of the Database Instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Port of the endpoint.
     */
    port?: pulumi.Input<number>;
}

export interface DatabaseInstancePrivateNetwork {
    /**
     * The ID of the endpoint of the private network.
     */
    endpointId?: pulumi.Input<string>;
    /**
     * Name of the endpoint.
     */
    hostname?: pulumi.Input<string>;
    /**
     * IP of the endpoint.
     */
    ip?: pulumi.Input<string>;
    ipNet: pulumi.Input<string>;
    /**
     * The name of the Database Instance.
     */
    name?: pulumi.Input<string>;
    pnId: pulumi.Input<string>;
    /**
     * Port of the endpoint.
     */
    port?: pulumi.Input<number>;
    zone?: pulumi.Input<string>;
}

export interface DatabaseInstanceReadReplica {
    /**
     * IP of the endpoint.
     */
    ip?: pulumi.Input<string>;
    /**
     * The name of the Database Instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Port of the endpoint.
     */
    port?: pulumi.Input<number>;
}

export interface DatabaseReadReplicaDirectAccess {
    /**
     * The ID of the endpoint of the read replica.
     */
    endpointId?: pulumi.Input<string>;
    /**
     * Hostname of the endpoint. Only one of ip and hostname may be set.
     */
    hostname?: pulumi.Input<string>;
    /**
     * IPv4 address of the endpoint (IP address). Only one of ip and hostname may be set.
     */
    ip?: pulumi.Input<string>;
    /**
     * Name of the endpoint.
     */
    name?: pulumi.Input<string>;
    /**
     * TCP port of the endpoint.
     */
    port?: pulumi.Input<number>;
}

export interface DatabaseReadReplicaPrivateNetwork {
    /**
     * The ID of the endpoint of the read replica.
     */
    endpointId?: pulumi.Input<string>;
    /**
     * Hostname of the endpoint. Only one of ip and hostname may be set.
     */
    hostname?: pulumi.Input<string>;
    /**
     * IPv4 address of the endpoint (IP address). Only one of ip and hostname may be set.
     */
    ip?: pulumi.Input<string>;
    /**
     * Name of the endpoint.
     */
    name?: pulumi.Input<string>;
    /**
     * TCP port of the endpoint.
     */
    port?: pulumi.Input<number>;
    /**
     * UUID of the private network to be connected to the read replica.
     */
    privateNetworkId: pulumi.Input<string>;
    /**
     * Endpoint IPv4 address with a CIDR notation. Check documentation about IP and subnet limitations. (IP network).
     */
    serviceIp: pulumi.Input<string>;
    zone?: pulumi.Input<string>;
}

export interface DomainRecordGeoIp {
    /**
     * The list of matches. *(Can be more than 1)*
     */
    matches: pulumi.Input<pulumi.Input<inputs.DomainRecordGeoIpMatch>[]>;
}

export interface DomainRecordGeoIpMatch {
    /**
     * List of continents (eg: `EU` for Europe, `NA` for North America, `AS` for Asia...). [List of all continents code](https://api.scaleway.com/domain-private/v2beta1/continents)
     */
    continents?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of countries (eg: `FR` for France, `US` for the United States, `GB` for Great Britain...). [List of all countries code](https://api.scaleway.com/domain-private/v2beta1/countries)
     */
    countries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The content of the record (an IPv4 for an `A`, a string for a `TXT`...).
     */
    data: pulumi.Input<string>;
}

export interface DomainRecordHttpService {
    /**
     * List of IPs to check
     */
    ips: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Text to search
     */
    mustContain: pulumi.Input<string>;
    /**
     * Strategy to return an IP from the IPs list. Can be `random` or `hashed`
     */
    strategy: pulumi.Input<string>;
    /**
     * URL to match the `mustContain` text to validate an IP
     */
    url: pulumi.Input<string>;
    /**
     * User-agent used when checking the URL
     */
    userAgent?: pulumi.Input<string>;
}

export interface DomainRecordView {
    /**
     * The content of the record (an IPv4 for an `A`, a string for a `TXT`...).
     */
    data: pulumi.Input<string>;
    /**
     * The subnet of the view
     */
    subnet: pulumi.Input<string>;
}

export interface DomainRecordWeighted {
    /**
     * The weighted IP
     */
    ip: pulumi.Input<string>;
    /**
     * The weight of the IP as an integer UInt32.
     */
    weight: pulumi.Input<number>;
}

export interface IamPolicyRule {
    /**
     * ID of organization scoped to the rule.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * Names of permission sets bound to the rule.
     */
    permissionSetNames: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of project IDs scoped to the rule.
     */
    projectIds?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface InstanceImageAdditionalVolume {
    /**
     * Date of the volume creation.
     */
    creationDate?: pulumi.Input<string>;
    /**
     * The export URI of the volume.
     */
    exportUri?: pulumi.Input<string>;
    /**
     * ID of the server containing the volume.
     */
    id?: pulumi.Input<string>;
    /**
     * Date of volume latest update.
     */
    modificationDate?: pulumi.Input<string>;
    /**
     * The name of the image. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization ID the volume is associated with.
     */
    organization?: pulumi.Input<string>;
    /**
     * ID of the project the volume is associated with
     */
    project?: pulumi.Input<string>;
    /**
     * Description of the server containing the volume (in case the image is a backup from a server).
     */
    server?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The size of the volume.
     */
    size?: pulumi.Input<number>;
    /**
     * State of the volume.
     */
    state?: pulumi.Input<string>;
    /**
     * A list of tags to apply to the image.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of volume, possible values are `lSsd` and `bSsd`.
     */
    volumeType?: pulumi.Input<string>;
    /**
     * The zone in which the image should be created.
     */
    zone?: pulumi.Input<string>;
}

export interface InstanceSecurityGroupInboundRule {
    /**
     * The action to take when rule match. Possible values are: `accept` or `drop`.
     */
    action: pulumi.Input<string>;
    /**
     * The ip this rule apply to. If no `ip` nor `ipRange` are specified, rule will apply to all ip. Only one of `ip` and `ipRange` should be specified.
     *
     * @deprecated Ip address is deprecated. Please use ip_range instead
     */
    ip?: pulumi.Input<string>;
    /**
     * The ip range (e.g `192.168.1.0/24`) this rule applies to. If no `ip` nor `ipRange` are specified, rule will apply to all ip. Only one of `ip` and `ipRange` should be specified.
     */
    ipRange?: pulumi.Input<string>;
    /**
     * The port this rule applies to. If no `port` nor `portRange` are specified, the rule will apply to all port. Only one of `port` and `portRange` should be specified.
     */
    port?: pulumi.Input<number>;
    portRange?: pulumi.Input<string>;
    /**
     * The protocol this rule apply to. Possible values are: `TCP`, `UDP`, `ICMP` or `ANY`.
     */
    protocol?: pulumi.Input<string>;
}

export interface InstanceSecurityGroupOutboundRule {
    /**
     * The action to take when rule match. Possible values are: `accept` or `drop`.
     */
    action: pulumi.Input<string>;
    /**
     * The ip this rule apply to. If no `ip` nor `ipRange` are specified, rule will apply to all ip. Only one of `ip` and `ipRange` should be specified.
     *
     * @deprecated Ip address is deprecated. Please use ip_range instead
     */
    ip?: pulumi.Input<string>;
    /**
     * The ip range (e.g `192.168.1.0/24`) this rule applies to. If no `ip` nor `ipRange` are specified, rule will apply to all ip. Only one of `ip` and `ipRange` should be specified.
     */
    ipRange?: pulumi.Input<string>;
    /**
     * The port this rule applies to. If no `port` nor `portRange` are specified, the rule will apply to all port. Only one of `port` and `portRange` should be specified.
     */
    port?: pulumi.Input<number>;
    portRange?: pulumi.Input<string>;
    /**
     * The protocol this rule apply to. Possible values are: `TCP`, `UDP`, `ICMP` or `ANY`.
     */
    protocol?: pulumi.Input<string>;
}

export interface InstanceSecurityGroupRulesInboundRule {
    /**
     * The action to take when rule match. Possible values are: `accept` or `drop`.
     */
    action: pulumi.Input<string>;
    /**
     * The ip this rule apply to. If no `ip` nor `ipRange` are specified, rule will apply to all ip. Only one of `ip` and `ipRange` should be specified.
     *
     * @deprecated Ip address is deprecated. Please use ip_range instead
     */
    ip?: pulumi.Input<string>;
    /**
     * The ip range (e.g `192.168.1.0/24`) this rule applies to. If no `ip` nor `ipRange` are specified, rule will apply to all ip. Only one of `ip` and `ipRange` should be specified.
     */
    ipRange?: pulumi.Input<string>;
    /**
     * The port this rule apply to. If no port is specified, rule will apply to all port.
     */
    port?: pulumi.Input<number>;
    portRange?: pulumi.Input<string>;
    /**
     * The protocol this rule apply to. Possible values are: `TCP`, `UDP`, `ICMP` or `ANY`.
     */
    protocol?: pulumi.Input<string>;
}

export interface InstanceSecurityGroupRulesOutboundRule {
    /**
     * The action to take when rule match. Possible values are: `accept` or `drop`.
     */
    action: pulumi.Input<string>;
    /**
     * The ip this rule apply to. If no `ip` nor `ipRange` are specified, rule will apply to all ip. Only one of `ip` and `ipRange` should be specified.
     *
     * @deprecated Ip address is deprecated. Please use ip_range instead
     */
    ip?: pulumi.Input<string>;
    /**
     * The ip range (e.g `192.168.1.0/24`) this rule applies to. If no `ip` nor `ipRange` are specified, rule will apply to all ip. Only one of `ip` and `ipRange` should be specified.
     */
    ipRange?: pulumi.Input<string>;
    /**
     * The port this rule apply to. If no port is specified, rule will apply to all port.
     */
    port?: pulumi.Input<number>;
    portRange?: pulumi.Input<string>;
    /**
     * The protocol this rule apply to. Possible values are: `TCP`, `UDP`, `ICMP` or `ANY`.
     */
    protocol?: pulumi.Input<string>;
}

export interface InstanceServerPrivateNetwork {
    macAddress?: pulumi.Input<string>;
    pnId: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the server should be created.
     */
    zone?: pulumi.Input<string>;
}

export interface InstanceServerRootVolume {
    boot?: pulumi.Input<boolean>;
    /**
     * Forces deletion of the root volume on instance termination.
     */
    deleteOnTermination?: pulumi.Input<boolean>;
    /**
     * The name of the server.
     */
    name?: pulumi.Input<string>;
    /**
     * Size of the root volume in gigabytes.
     * To find the right size use [this endpoint](https://api.scaleway.com/instance/v1/zones/fr-par-1/products/servers) and
     * check the `volumes_constraint.{min|max}_size` (in bytes) for your `commercialType`.
     * Updates to this field will recreate a new resource.
     */
    sizeInGb?: pulumi.Input<number>;
    /**
     * The volume ID of the root volume of the server, allows you to create server with an existing volume. If empty, will be computed to a created volume ID.
     */
    volumeId?: pulumi.Input<string>;
    /**
     * Volume type of root volume, can be `bSsd` or `lSsd`, default value depends on server type
     */
    volumeType?: pulumi.Input<string>;
}

export interface InstanceSnapshotImport {
    /**
     * Bucket name containing [qcow2](https://en.wikipedia.org/wiki/Qcow) to import
     */
    bucket: pulumi.Input<string>;
    /**
     * Key of the object to import
     */
    key: pulumi.Input<string>;
}

export interface IotDeviceCertificate {
    crt?: pulumi.Input<string>;
    /**
     * The private key of the device, in case it is generated by Scaleway.
     */
    key?: pulumi.Input<string>;
}

export interface IotDeviceMessageFilters {
    /**
     * Rules used to restrict topics the device can publish to.
     */
    publish?: pulumi.Input<inputs.IotDeviceMessageFiltersPublish>;
    /**
     * Rules used to restrict topics the device can subscribe to.
     */
    subscribe?: pulumi.Input<inputs.IotDeviceMessageFiltersSubscribe>;
}

export interface IotDeviceMessageFiltersPublish {
    /**
     * Same as publish rules.
     */
    policy?: pulumi.Input<string>;
    /**
     * Same as publish rules.
     */
    topics?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface IotDeviceMessageFiltersSubscribe {
    /**
     * Same as publish rules.
     */
    policy?: pulumi.Input<string>;
    /**
     * Same as publish rules.
     */
    topics?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface IotRouteDatabase {
    dbname: pulumi.Input<string>;
    host: pulumi.Input<string>;
    password: pulumi.Input<string>;
    port: pulumi.Input<number>;
    query: pulumi.Input<string>;
    username: pulumi.Input<string>;
}

export interface IotRouteRest {
    headers: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    uri: pulumi.Input<string>;
    verb: pulumi.Input<string>;
}

export interface IotRouteS3 {
    bucketName: pulumi.Input<string>;
    bucketRegion: pulumi.Input<string>;
    objectPrefix?: pulumi.Input<string>;
    strategy: pulumi.Input<string>;
}

export interface KubernetesClusterAutoUpgrade {
    /**
     * Set to `true` to enable Kubernetes patch version auto upgrades.
     * > **Important:** When enabling auto upgrades, the `version` field take a minor version like x.y (ie 1.18).
     */
    enable: pulumi.Input<boolean>;
    /**
     * The day of the auto upgrade maintenance window (`monday` to `sunday`, or `any`).
     */
    maintenanceWindowDay: pulumi.Input<string>;
    /**
     * The start hour (UTC) of the 2-hour auto upgrade maintenance window (0 to 23).
     */
    maintenanceWindowStartHour: pulumi.Input<number>;
}

export interface KubernetesClusterAutoscalerConfig {
    /**
     * Detect similar node groups and balance the number of nodes between them.
     */
    balanceSimilarNodeGroups?: pulumi.Input<boolean>;
    /**
     * Disables the scale down feature of the autoscaler.
     */
    disableScaleDown?: pulumi.Input<boolean>;
    /**
     * Type of resource estimator to be used in scale up.
     */
    estimator?: pulumi.Input<string>;
    /**
     * Type of node group expander to be used in scale up.
     */
    expander?: pulumi.Input<string>;
    /**
     * Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they don't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
     */
    expendablePodsPriorityCutoff?: pulumi.Input<number>;
    /**
     * Ignore DaemonSet pods when calculating resource utilization for scaling down.
     */
    ignoreDaemonsetsUtilization?: pulumi.Input<boolean>;
    /**
     * Maximum number of seconds the cluster autoscaler waits for pod termination when trying to scale down a node
     */
    maxGracefulTerminationSec?: pulumi.Input<number>;
    /**
     * How long after scale up that scale down evaluation resumes.
     */
    scaleDownDelayAfterAdd?: pulumi.Input<string>;
    /**
     * How long a node should be unneeded before it is eligible for scale down.
     */
    scaleDownUnneededTime?: pulumi.Input<string>;
    /**
     * Node utilization level, defined as sum of requested resources divided by capacity, below which a node can be considered for scale down
     */
    scaleDownUtilizationThreshold?: pulumi.Input<number>;
}

export interface KubernetesClusterKubeconfig {
    /**
     * The CA certificate of the Kubernetes API server.
     */
    clusterCaCertificate?: pulumi.Input<string>;
    /**
     * The raw kubeconfig file.
     */
    configFile?: pulumi.Input<string>;
    /**
     * The URL of the Kubernetes API server.
     */
    host?: pulumi.Input<string>;
    /**
     * The token to connect to the Kubernetes API server.
     */
    token?: pulumi.Input<string>;
}

export interface KubernetesClusterOpenIdConnectConfig {
    /**
     * A client id that all tokens must be issued for
     */
    clientId: pulumi.Input<string>;
    /**
     * JWT claim to use as the user's group
     */
    groupsClaims?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Prefix prepended to group claims
     */
    groupsPrefix?: pulumi.Input<string>;
    /**
     * URL of the provider which allows the API server to discover public signing keys
     */
    issuerUrl: pulumi.Input<string>;
    /**
     * Multiple key=value pairs that describes a required claim in the ID Token
     */
    requiredClaims?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * JWT claim to use as the user name
     */
    usernameClaim?: pulumi.Input<string>;
    /**
     * Prefix prepended to username
     */
    usernamePrefix?: pulumi.Input<string>;
}

export interface KubernetesNodePoolNode {
    /**
     * The name for the pool.
     * > **Important:** Updates to this field will recreate a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The public IPv4.
     */
    publicIp?: pulumi.Input<string>;
    /**
     * The public IPv6.
     */
    publicIpV6?: pulumi.Input<string>;
    /**
     * The status of the node.
     */
    status?: pulumi.Input<string>;
}

export interface KubernetesNodePoolUpgradePolicy {
    /**
     * The maximum number of nodes to be created during the upgrade
     */
    maxSurge?: pulumi.Input<number>;
    /**
     * The maximum number of nodes that can be not ready at the same time
     */
    maxUnavailable?: pulumi.Input<number>;
}

export interface LoadbalancerBackendHealthCheckHttp {
    /**
     * The expected HTTP status code.
     */
    code?: pulumi.Input<number>;
    /**
     * The HTTP method to use for HC requests.
     */
    method?: pulumi.Input<string>;
    /**
     * The HTTP endpoint URL to call for HC requests.
     */
    uri: pulumi.Input<string>;
}

export interface LoadbalancerBackendHealthCheckHttps {
    /**
     * The expected HTTP status code.
     */
    code?: pulumi.Input<number>;
    /**
     * The HTTP method to use for HC requests.
     */
    method?: pulumi.Input<string>;
    /**
     * The HTTP endpoint URL to call for HC requests.
     */
    uri: pulumi.Input<string>;
}

export interface LoadbalancerBackendHealthCheckTcp {
}

export interface LoadbalancerCertificateCustomCertificate {
    /**
     * Full PEM-formatted certificate chain.
     */
    certificateChain: pulumi.Input<string>;
}

export interface LoadbalancerCertificateLetsencrypt {
    /**
     * Main domain of the certificate. A new certificate will be created if this field is changed.
     */
    commonName: pulumi.Input<string>;
    /**
     * Array of alternative domain names.  A new certificate will be created if this field is changed.
     */
    subjectAlternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface LoadbalancerFrontendAcl {
    /**
     * Action to undertake when an ACL filter matches.
     */
    action: pulumi.Input<inputs.LoadbalancerFrontendAclAction>;
    /**
     * The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
     */
    match: pulumi.Input<inputs.LoadbalancerFrontendAclMatch>;
    /**
     * The ACL name. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
}

export interface LoadbalancerFrontendAclAction {
    /**
     * The action type. Possible values are: `allow` or `deny`.
     */
    type: pulumi.Input<string>;
}

export interface LoadbalancerFrontendAclMatch {
    /**
     * The HTTP filter to match. This filter is supported only if your backend protocol has an HTTP forward protocol.
     * It extracts the request's URL path, which starts at the first slash and ends before the question mark (without the host part).
     * Possible values are: `aclHttpFilterNone`, `pathBegin`, `pathEnd`, `httpHeaderMatch` or `regex`.
     */
    httpFilter?: pulumi.Input<string>;
    httpFilterOption?: pulumi.Input<string>;
    /**
     * A list of possible values to match for the given HTTP filter.
     * Keep in mind that in the case of `httpHeaderMatch` the HTTP header field name is case-insensitive.
     */
    httpFilterValues?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set to `true`, the condition will be of type "unless".
     */
    invert?: pulumi.Input<boolean>;
    /**
     * A list of IPs or CIDR v4/v6 addresses of the client of the session to match.
     */
    ipSubnets?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface LoadbalancerPrivateNetwork {
    /**
     * (Optional) Set to true if you want to let DHCP assign IP addresses. See below.
     */
    dhcpConfig?: pulumi.Input<boolean>;
    /**
     * (Required) The ID of the Private Network to associate.
     */
    privateNetworkId: pulumi.Input<string>;
    /**
     * (Optional) Define a local ip address of your choice for the load balancer instance. See below.
     */
    staticConfig?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the IP should be reserved.
     */
    zone?: pulumi.Input<string>;
}

export interface MnqCredentialNatsCredentials {
    /**
     * Raw content of the NATS credentials file.
     */
    content?: pulumi.Input<string>;
}

export interface MnqCredentialSqsSnsCredentials {
    /**
     * The ID of the key.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * List of permissions associated to this Credential. Only one of permissions may be set.
     */
    permissions?: pulumi.Input<inputs.MnqCredentialSqsSnsCredentialsPermissions>;
    /**
     * The Secret value of the key.
     */
    secretKey?: pulumi.Input<string>;
}

export interface MnqCredentialSqsSnsCredentialsPermissions {
    /**
     * . Defines if user can manage the associated resource(s).
     */
    canManage?: pulumi.Input<boolean>;
    /**
     * . Defines if user can publish messages to the service.
     */
    canPublish?: pulumi.Input<boolean>;
    /**
     * . Defines if user can receive messages from the service.
     */
    canReceive?: pulumi.Input<boolean>;
}

export interface ObjectBucketAclAccessControlPolicy {
    grants?: pulumi.Input<pulumi.Input<inputs.ObjectBucketAclAccessControlPolicyGrant>[]>;
    owner: pulumi.Input<inputs.ObjectBucketAclAccessControlPolicyOwner>;
}

export interface ObjectBucketAclAccessControlPolicyGrant {
    grantee?: pulumi.Input<inputs.ObjectBucketAclAccessControlPolicyGrantGrantee>;
    permission: pulumi.Input<string>;
}

export interface ObjectBucketAclAccessControlPolicyGrantGrantee {
    displayName?: pulumi.Input<string>;
    /**
     * The `region`,`bucket` and `acl` separated by (`/`).
     */
    id: pulumi.Input<string>;
    type: pulumi.Input<string>;
}

export interface ObjectBucketAclAccessControlPolicyOwner {
    displayName?: pulumi.Input<string>;
    /**
     * The `region`,`bucket` and `acl` separated by (`/`).
     */
    id: pulumi.Input<string>;
}

export interface ObjectBucketCorsRule {
    /**
     * Specifies which headers are allowed.
     */
    allowedHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies which methods are allowed. Can be `GET`, `PUT`, `POST`, `DELETE` or `HEAD`.
     */
    allowedMethods: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies which origins are allowed.
     */
    allowedOrigins: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies expose header in the response.
     */
    exposeHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies time in seconds that browser can cache the response for a preflight request.
     */
    maxAgeSeconds?: pulumi.Input<number>;
}

export interface ObjectBucketLifecycleRule {
    /**
     * Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
     */
    abortIncompleteMultipartUploadDays?: pulumi.Input<number>;
    /**
     * The element value can be either Enabled or Disabled. If a rule is disabled, Scaleway S3 doesn't perform any of the actions defined in the rule.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Specifies a period in the object's expire (documented below).
     */
    expiration?: pulumi.Input<inputs.ObjectBucketLifecycleRuleExpiration>;
    /**
     * Unique identifier for the rule. Must be less than or equal to 255 characters in length.
     */
    id?: pulumi.Input<string>;
    /**
     * Object key prefix identifying one or more objects to which the rule applies.
     */
    prefix?: pulumi.Input<string>;
    /**
     * Specifies object tags key and value.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies a period in the object's transitions (documented below).
     */
    transitions?: pulumi.Input<pulumi.Input<inputs.ObjectBucketLifecycleRuleTransition>[]>;
}

export interface ObjectBucketLifecycleRuleExpiration {
    /**
     * Specifies the number of days after object creation when the specific rule action takes effect.
     */
    days: pulumi.Input<number>;
}

export interface ObjectBucketLifecycleRuleTransition {
    /**
     * Specifies the number of days after object creation when the specific rule action takes effect.
     */
    days?: pulumi.Input<number>;
    /**
     * Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA`  to which you want the object to transition.
     */
    storageClass: pulumi.Input<string>;
}

export interface ObjectBucketLockConfigurationRule {
    defaultRetention: pulumi.Input<inputs.ObjectBucketLockConfigurationRuleDefaultRetention>;
}

export interface ObjectBucketLockConfigurationRuleDefaultRetention {
    days?: pulumi.Input<number>;
    mode: pulumi.Input<string>;
    years?: pulumi.Input<number>;
}

export interface ObjectBucketVersioning {
    /**
     * Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.
     */
    enabled?: pulumi.Input<boolean>;
}

export interface ObjectBucketWebsiteConfigurationErrorDocument {
    key: pulumi.Input<string>;
}

export interface ObjectBucketWebsiteConfigurationIndexDocument {
    suffix: pulumi.Input<string>;
}

export interface RedisClusterAcl {
    /**
     * A text describing this rule. Default description: `Allow IP`
     */
    description?: pulumi.Input<string>;
    /**
     * The UUID of the private network resource.
     */
    id?: pulumi.Input<string>;
    /**
     * The ip range to whitelist in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)
     */
    ip: pulumi.Input<string>;
}

export interface RedisClusterPrivateNetwork {
    endpointId?: pulumi.Input<string>;
    /**
     * The UUID of the private network resource.
     */
    id: pulumi.Input<string>;
    /**
     * Endpoint IPv4 addresses in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation). You must provide at least one IP per node.
     */
    serviceIps: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `zone`) The zone in which the Redis Cluster should be created.
     */
    zone?: pulumi.Input<string>;
}

export interface RedisClusterPublicNetwork {
    /**
     * The UUID of the private network resource.
     */
    id?: pulumi.Input<string>;
    ips?: pulumi.Input<pulumi.Input<string>[]>;
    port?: pulumi.Input<number>;
}
