// Copyright 2023 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// GENERATED BY gen_go_data.go
// gen_go_data -package alpha -var YAML_private_cloud blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vmware/alpha/private_cloud.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vmware/alpha/private_cloud.yaml
var YAML_private_cloud = []byte("info:\n  title: Vmware/PrivateCloud\n  description: The Vmware PrivateCloud resource\n  x-dcl-struct-name: PrivateCloud\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a PrivateCloud\n    parameters:\n    - name: privateCloud\n      required: true\n      description: A full instance of a PrivateCloud\n  apply:\n    description: The function used to apply information about a PrivateCloud\n    parameters:\n    - name: privateCloud\n      required: true\n      description: A full instance of a PrivateCloud\n  delete:\n    description: The function used to delete a PrivateCloud\n    parameters:\n    - name: privateCloud\n      required: true\n      description: A full instance of a PrivateCloud\n  deleteAll:\n    description: The function used to delete all PrivateCloud\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many PrivateCloud\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    PrivateCloud:\n      title: PrivateCloud\n      x-dcl-id: projects/{{project}}/locations/{{location}}/privateClouds/{{name}}\n      x-dcl-uses-state-hint: true\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: true\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 9600\n      x-dcl-delete-timeout: 7200\n      type: object\n      required:\n      - name\n      - networkConfig\n      - managementCluster\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Creation time of this resource.\n          x-kubernetes-immutable: true\n        deleteTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: DeleteTime\n          readOnly: true\n          description: Output only. Time when the resource was scheduled for deletion.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: User-provided description for this private cloud.\n        expireTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: ExpireTime\n          readOnly: true\n          description: Output only. Time when the resource will be irreversibly deleted.\n          x-kubernetes-immutable: true\n        hcx:\n          type: object\n          x-dcl-go-name: Hcx\n          x-dcl-go-type: PrivateCloudHcx\n          readOnly: true\n          description: Output only. HCX appliance.\n          properties:\n            fqdn:\n              type: string\n              x-dcl-go-name: Fqdn\n              description: Fully qualified domain name of the appliance.\n            internalIP:\n              type: string\n              x-dcl-go-name: InternalIP\n              description: Internal IP address of the appliance.\n            state:\n              type: string\n              x-dcl-go-name: State\n              x-dcl-go-type: PrivateCloudHcxStateEnum\n              readOnly: true\n              description: 'Output only. The state of the appliance. Possible values:\n                STATE_UNSPECIFIED, ACTIVE, CREATING'\n              x-kubernetes-immutable: true\n              enum:\n              - STATE_UNSPECIFIED\n              - ACTIVE\n              - CREATING\n            version:\n              type: string\n              x-dcl-go-name: Version\n              description: Version of the appliance.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Locations/Location\n            field: name\n            parent: true\n        managementCluster:\n          type: object\n          x-dcl-go-name: ManagementCluster\n          x-dcl-go-type: PrivateCloudManagementCluster\n          description: 'Required. Input only. The management cluster for this private\n            cloud. This field is required during creation of the private cloud to\n            provide details for the default cluster. The following fields can''t be\n            changed after private cloud creation: `ManagementCluster.clusterId`, `ManagementCluster.nodeTypeId`.'\n          x-dcl-mutable-unreadable: true\n          required:\n          - clusterId\n          - nodeTypeConfigs\n          properties:\n            clusterId:\n              type: string\n              x-dcl-go-name: ClusterId\n              description: 'Required. The user-provided identifier of the new `Cluster`.\n                The identifier must meet the following requirements: * Only contains\n                1-63 alphanumeric characters and hyphens * Begins with an alphabetical\n                character * Ends with a non-hyphen character * Not formatted as a\n                UUID * Complies with [RFC 1034](https://datatracker.ietf.org/doc/html/rfc1034)\n                (section 3.5)'\n            nodeTypeConfigs:\n              type: object\n              additionalProperties:\n                type: object\n                x-dcl-go-type: PrivateCloudManagementClusterNodeTypeConfigs\n                required:\n                - nodeCount\n                properties:\n                  customCoreCount:\n                    type: integer\n                    format: int64\n                    x-dcl-go-name: CustomCoreCount\n                    description: Optional. Customized number of cores available to\n                      each node of the type. This number must always be one of `nodeType.availableCustomCoreCounts`.\n                      If zero is provided max value from `nodeType.availableCustomCoreCounts`\n                      will be used.\n                  nodeCount:\n                    type: integer\n                    format: int64\n                    x-dcl-go-name: NodeCount\n                    description: Required. The number of nodes of this type in the\n                      cluster\n              x-dcl-go-name: NodeTypeConfigs\n              description: Required. The map of cluster node types in this cluster,\n                where the key is canonical identifier of the node type (corresponds\n                to the `NodeType`).\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Output only. The resource name of this private cloud. Resource\n            names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.\n            For example: `projects/my-project/locations/us-west1-a/privateClouds/my-cloud`'\n          x-kubernetes-immutable: true\n        networkConfig:\n          type: object\n          x-dcl-go-name: NetworkConfig\n          x-dcl-go-type: PrivateCloudNetworkConfig\n          description: Required. Network configuration of the private cloud.\n          required:\n          - managementCidr\n          properties:\n            managementCidr:\n              type: string\n              x-dcl-go-name: ManagementCidr\n              description: Required. Management CIDR used by VMware management appliances.\n            managementIPAddressLayoutVersion:\n              type: integer\n              format: int64\n              x-dcl-go-name: ManagementIPAddressLayoutVersion\n              readOnly: true\n              description: 'Output only. The IP address layout version of the management\n                IP address range. Possible versions include: * `managementIpAddressLayoutVersion=1`:\n                Indicates the legacy IP address layout used by some existing private\n                clouds. This is no longer supported for new private clouds as it does\n                not support all features. * `managementIpAddressLayoutVersion=2`:\n                Indicates the latest IP address layout used by all newly created private\n                clouds. This version supports all current features.'\n            vmwareEngineNetwork:\n              type: string\n              x-dcl-go-name: VmwareEngineNetwork\n              description: 'Optional. The relative resource name of the VMware Engine\n                network attached to the private cloud. Specify the name in the following\n                form: `projects/{project}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}`\n                where `{project}` can either be a project number or a project ID.'\n              x-dcl-references:\n              - resource: Vmwareengine/VmwareEngineNetwork\n                field: name\n            vmwareEngineNetworkCanonical:\n              type: string\n              x-dcl-go-name: VmwareEngineNetworkCanonical\n              readOnly: true\n              description: 'Output only. The canonical name of the VMware Engine network\n                in the form: `projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}`'\n              x-dcl-references:\n              - resource: Vmwareengine/VmwareEngineNetwork\n                field: name\n        nsx:\n          type: object\n          x-dcl-go-name: Nsx\n          x-dcl-go-type: PrivateCloudNsx\n          readOnly: true\n          description: Output only. NSX appliance.\n          properties:\n            fqdn:\n              type: string\n              x-dcl-go-name: Fqdn\n              description: Fully qualified domain name of the appliance.\n            internalIP:\n              type: string\n              x-dcl-go-name: InternalIP\n              description: Internal IP address of the appliance.\n            state:\n              type: string\n              x-dcl-go-name: State\n              x-dcl-go-type: PrivateCloudNsxStateEnum\n              readOnly: true\n              description: 'Output only. The state of the appliance. Possible values:\n                STATE_UNSPECIFIED, ACTIVE, CREATING'\n              x-kubernetes-immutable: true\n              enum:\n              - STATE_UNSPECIFIED\n              - ACTIVE\n              - CREATING\n            version:\n              type: string\n              x-dcl-go-name: Version\n              description: Version of the appliance.\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: PrivateCloudStateEnum\n          readOnly: true\n          description: 'Output only. State of the resource. New values may be added\n            to this enum when appropriate. Possible values: STATE_UNSPECIFIED, ACTIVE,\n            CREATING, UPDATING, FAILED, DELETED, PURGING'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - ACTIVE\n          - CREATING\n          - UPDATING\n          - FAILED\n          - DELETED\n          - PURGING\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. System-generated unique identifier for the resource.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Last update time of this resource.\n          x-kubernetes-immutable: true\n        vcenter:\n          type: object\n          x-dcl-go-name: Vcenter\n          x-dcl-go-type: PrivateCloudVcenter\n          readOnly: true\n          description: Output only. Vcenter appliance.\n          properties:\n            fqdn:\n              type: string\n              x-dcl-go-name: Fqdn\n              description: Fully qualified domain name of the appliance.\n            internalIP:\n              type: string\n              x-dcl-go-name: InternalIP\n              description: Internal IP address of the appliance.\n            state:\n              type: string\n              x-dcl-go-name: State\n              x-dcl-go-type: PrivateCloudVcenterStateEnum\n              readOnly: true\n              description: 'Output only. The state of the appliance. Possible values:\n                STATE_UNSPECIFIED, ACTIVE, CREATING'\n              x-kubernetes-immutable: true\n              enum:\n              - STATE_UNSPECIFIED\n              - ACTIVE\n              - CREATING\n            version:\n              type: string\n              x-dcl-go-name: Version\n              description: Version of the appliance.\n")

// 12817 bytes
// MD5: 37420c7f1a4988fbedde879ce08d70b9
