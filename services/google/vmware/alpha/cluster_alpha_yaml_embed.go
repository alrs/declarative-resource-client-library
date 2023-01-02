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
// gen_go_data -package alpha -var YAML_cluster blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vmware/alpha/cluster.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vmware/alpha/cluster.yaml
var YAML_cluster = []byte("info:\n  title: Vmware/Cluster\n  description: The Vmware Cluster resource\n  x-dcl-struct-name: Cluster\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Cluster\n    parameters:\n    - name: cluster\n      required: true\n      description: A full instance of a Cluster\n  apply:\n    description: The function used to apply information about a Cluster\n    parameters:\n    - name: cluster\n      required: true\n      description: A full instance of a Cluster\n  delete:\n    description: The function used to delete a Cluster\n    parameters:\n    - name: cluster\n      required: true\n      description: A full instance of a Cluster\n  deleteAll:\n    description: The function used to delete all Cluster\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: privateCloud\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Cluster\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: privateCloud\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Cluster:\n      title: Cluster\n      x-dcl-id: projects/{{project}}/locations/{{location}}/privateClouds/{{private_cloud}}/clusters/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 9600\n      x-dcl-delete-timeout: 7200\n      type: object\n      required:\n      - name\n      - nodeTypeConfigs\n      - project\n      - location\n      - privateCloud\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Creation time of this resource.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        management:\n          type: boolean\n          x-dcl-go-name: Management\n          readOnly: true\n          description: Output only. True if the cluster is a management cluster; false\n            otherwise. There can only be one management cluster in a private cloud\n            and it has to be the first one.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Output only. The resource name of this cluster. Resource names\n            are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.\n            For example: `projects/my-project/locations/us-west1-a/privateClouds/my-cloud/clusters/my-cluster`'\n          x-kubernetes-immutable: true\n        nodeTypeConfigs:\n          type: object\n          additionalProperties:\n            type: object\n            x-dcl-go-type: ClusterNodeTypeConfigs\n            required:\n            - nodeCount\n            properties:\n              customCoreCount:\n                type: integer\n                format: int64\n                x-dcl-go-name: CustomCoreCount\n                description: Optional. Customized number of cores available to each\n                  node of the type. This number must always be one of `nodeType.availableCustomCoreCounts`.\n                  If zero is provided max value from `nodeType.availableCustomCoreCounts`\n                  will be used.\n              nodeCount:\n                type: integer\n                format: int64\n                x-dcl-go-name: NodeCount\n                description: Required. The number of nodes of this type in the cluster\n          x-dcl-go-name: NodeTypeConfigs\n          description: Required. The map of cluster node types in this cluster, where\n            the key is canonical identifier of the node type (corresponds to the `NodeType`).\n        privateCloud:\n          type: string\n          x-dcl-go-name: PrivateCloud\n          description: The private_cloud for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Vmwareengine/PrivateCloud\n            field: name\n            parent: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: ClusterStateEnum\n          readOnly: true\n          description: 'Output only. State of the resource. Possible values: STATE_UNSPECIFIED,\n            ACTIVE, CREATING, UPDATING, DELETING, REPAIRING'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - ACTIVE\n          - CREATING\n          - UPDATING\n          - DELETING\n          - REPAIRING\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. System-generated unique identifier for the resource.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Last update time of this resource.\n          x-kubernetes-immutable: true\n")

// 5655 bytes
// MD5: 249e1762421dc366398a979ddc05103e
