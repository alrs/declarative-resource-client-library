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
// gen_go_data -package beta -var YAML_monitored_project blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/beta/monitored_project.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/beta/monitored_project.yaml
var YAML_monitored_project = []byte("info:\n  title: Monitoring/MonitoredProject\n  description: Monitored Project allows you to set a project as monitored by a _metrics\n    scope_, which is a term for a project used to group the metrics of multiple projects,\n    potentially across multiple organizations.  This enables you to view these groups\n    in the Monitoring page of the cloud console.\n  x-dcl-struct-name: MonitoredProject\n  x-dcl-has-iam: false\n  x-dcl-ref:\n    text: REST API\n    url: https://cloud.google.com/monitoring/api/ref_v3/rest/v1/locations.global.metricsScopes\n  x-dcl-guides:\n  - text: Understanding metrics scopes\n    url: https://cloud.google.com/monitoring/settings#concept-scope\n  - text: API notes\n    url: https://cloud.google.com/monitoring/settings/manage-api\npaths:\n  get:\n    description: The function used to get information about a MonitoredProject\n    parameters:\n    - name: monitoredProject\n      required: true\n      description: A full instance of a MonitoredProject\n  apply:\n    description: The function used to apply information about a MonitoredProject\n    parameters:\n    - name: monitoredProject\n      required: true\n      description: A full instance of a MonitoredProject\n  delete:\n    description: The function used to delete a MonitoredProject\n    parameters:\n    - name: monitoredProject\n      required: true\n      description: A full instance of a MonitoredProject\n  deleteAll:\n    description: The function used to delete all MonitoredProject\n    parameters:\n    - name: metricsScope\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many MonitoredProject\n    parameters:\n    - name: metricsScope\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    MonitoredProject:\n      title: MonitoredProject\n      x-dcl-id: locations/global/metricsScopes/{{metrics_scope}}/projects/{{name}}\n      x-dcl-locations:\n      - global\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - metricsScope\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time when this `MonitoredProject` was created.\n          x-kubernetes-immutable: true\n        metricsScope:\n          type: string\n          x-dcl-go-name: MetricsScope\n          description: 'Required. The resource name of the existing Metrics Scope\n            that will monitor this project. Example: locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}'\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Immutable. The resource name of the `MonitoredProject`. On\n            input, the resource name includes the scoping project ID and monitored\n            project ID. On output, it contains the equivalent project numbers. Example:\n            `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`'\n          x-kubernetes-immutable: true\n")

// 3225 bytes
// MD5: a2568c724b961c0d484ec49d4439ae8b
