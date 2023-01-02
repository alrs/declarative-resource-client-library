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
// gen_go_data -package alpha -var YAML_fleet blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/gkehub/alpha/fleet.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/gkehub/alpha/fleet.yaml
var YAML_fleet = []byte("info:\n  title: GkeHub/Fleet\n  description: The GkeHub Fleet resource\n  x-dcl-struct-name: Fleet\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Fleet\n    parameters:\n    - name: fleet\n      required: true\n      description: A full instance of a Fleet\n  apply:\n    description: The function used to apply information about a Fleet\n    parameters:\n    - name: fleet\n      required: true\n      description: A full instance of a Fleet\n  delete:\n    description: The function used to delete a Fleet\n    parameters:\n    - name: fleet\n      required: true\n      description: A full instance of a Fleet\ncomponents:\n  schemas:\n    Fleet:\n      title: Fleet\n      x-dcl-id: projects/{{project}}/locations/{{location}}/fleets/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. When the Fleet was created.\n          x-kubernetes-immutable: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: 'Optional. A user-assigned display name of the Fleet. When\n            present, it must be between 4 to 30 characters. Allowed characters are:\n            lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote,\n            space, and exclamation point. Example: `Production Fleet`'\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        managedNamespaces:\n          type: boolean\n          x-dcl-go-name: ManagedNamespaces\n          description: Optional. If true, namespaces must be explicitly declared in\n            a `FleetNamespace` object in order to use Fleet Features.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. The full, unique resource name of this fleet in\n            the format of `projects/{project}/locations/{location}/fleets/{fleet}`.\n            Each GCP project can have at most one fleet resource, named \"default\".\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. Google-generated UUID for this resource. This\n            is unique across all Fleet resources. If a Fleet resource is deleted and\n            another resource with the same name is created, it gets a different uid.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. When the Fleet was last updated.\n          x-kubernetes-immutable: true\n")

// 3419 bytes
// MD5: bb86838055269d0297fbf1cfd9776083
