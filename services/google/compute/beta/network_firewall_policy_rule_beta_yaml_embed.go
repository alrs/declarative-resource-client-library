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
// gen_go_data -package beta -var YAML_network_firewall_policy_rule blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/network_firewall_policy_rule.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/network_firewall_policy_rule.yaml
var YAML_network_firewall_policy_rule = []byte("info:\n  title: Compute/NetworkFirewallPolicyRule\n  description: The Compute NetworkFirewallPolicyRule resource\n  x-dcl-struct-name: NetworkFirewallPolicyRule\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a NetworkFirewallPolicyRule\n    parameters:\n    - name: networkFirewallPolicyRule\n      required: true\n      description: A full instance of a NetworkFirewallPolicyRule\n  apply:\n    description: The function used to apply information about a NetworkFirewallPolicyRule\n    parameters:\n    - name: networkFirewallPolicyRule\n      required: true\n      description: A full instance of a NetworkFirewallPolicyRule\n  delete:\n    description: The function used to delete a NetworkFirewallPolicyRule\n    parameters:\n    - name: networkFirewallPolicyRule\n      required: true\n      description: A full instance of a NetworkFirewallPolicyRule\n  deleteAll:\n    description: The function used to delete all NetworkFirewallPolicyRule\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: firewallPolicy\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many NetworkFirewallPolicyRule\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: firewallPolicy\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    NetworkFirewallPolicyRule:\n      title: NetworkFirewallPolicyRule\n      x-dcl-id: projects/{{project}}/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}\n      x-dcl-locations:\n      - global\n      - region\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - priority\n      - match\n      - action\n      - direction\n      - firewallPolicy\n      - project\n      properties:\n        action:\n          type: string\n          x-dcl-go-name: Action\n          description: The Action to perform when the client connection triggers the\n            rule. Can currently be either \"allow\" or \"deny()\" where valid values for\n            status are 403, 404, and 502.\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description for this resource.\n        direction:\n          type: string\n          x-dcl-go-name: Direction\n          x-dcl-go-type: NetworkFirewallPolicyRuleDirectionEnum\n          description: 'The direction in which this rule applies. Possible values:\n            INGRESS, EGRESS'\n          enum:\n          - INGRESS\n          - EGRESS\n        disabled:\n          type: boolean\n          x-dcl-go-name: Disabled\n          description: Denotes whether the firewall policy rule is disabled. When\n            set to true, the firewall policy rule is not enforced and traffic behaves\n            as if it did not exist. If this is unspecified, the firewall policy rule\n            will be enabled.\n        enableLogging:\n          type: boolean\n          x-dcl-go-name: EnableLogging\n          description: 'Denotes whether to enable logging for a particular rule. If\n            logging is enabled, logs will be exported to the configured export destination\n            in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you\n            cannot enable logging on \"goto_next\" rules.'\n        firewallPolicy:\n          type: string\n          x-dcl-go-name: FirewallPolicy\n          description: The firewall policy of the resource.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Compute/NetworkFirewallPolicy\n            field: name\n            parent: true\n        kind:\n          type: string\n          x-dcl-go-name: Kind\n          readOnly: true\n          description: Type of the resource. Always `compute#firewallPolicyRule` for\n            firewall policy rules\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location of this resource.\n          x-kubernetes-immutable: true\n        match:\n          type: object\n          x-dcl-go-name: Match\n          x-dcl-go-type: NetworkFirewallPolicyRuleMatch\n          description: A match condition that incoming traffic is evaluated against.\n            If it evaluates to true, the corresponding 'action' is enforced.\n          required:\n          - layer4Configs\n          properties:\n            destFqdns:\n              type: array\n              x-dcl-go-name: DestFqdns\n              description: Domain names that will be used to match against the resolved\n                domain name of destination of traffic. Can only be specified if DIRECTION\n                is egress.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            destIPRanges:\n              type: array\n              x-dcl-go-name: DestIPRanges\n              description: CIDR IP address range. Maximum number of destination CIDR\n                IP ranges allowed is 5000.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            destRegionCodes:\n              type: array\n              x-dcl-go-name: DestRegionCodes\n              description: The Unicode country codes whose IP addresses will be used\n                to match against the source of traffic. Can only be specified if DIRECTION\n                is egress.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            destThreatIntelligences:\n              type: array\n              x-dcl-go-name: DestThreatIntelligences\n              description: Name of the Google Cloud Threat Intelligence list.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            layer4Configs:\n              type: array\n              x-dcl-go-name: Layer4Configs\n              description: Pairs of IP protocols and ports that the rule should match.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: NetworkFirewallPolicyRuleMatchLayer4Configs\n                required:\n                - ipProtocol\n                properties:\n                  ipProtocol:\n                    type: string\n                    x-dcl-go-name: IPProtocol\n                    description: The IP protocol to which this rule applies. The protocol\n                      type is required when creating a firewall rule. This value can\n                      either be one of the following well known protocol strings (`tcp`,\n                      `udp`, `icmp`, `esp`, `ah`, `ipip`, `sctp`), or the IP protocol\n                      number.\n                  ports:\n                    type: array\n                    x-dcl-go-name: Ports\n                    description: 'An optional list of ports to which this rule applies.\n                      This field is only applicable for UDP or TCP protocol. Each\n                      entry must be either an integer or a range. If not specified,\n                      this rule applies to connections through any port. Example inputs\n                      include: ``.'\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: string\n                      x-dcl-go-type: string\n            srcFqdns:\n              type: array\n              x-dcl-go-name: SrcFqdns\n              description: Domain names that will be used to match against the resolved\n                domain name of source of traffic. Can only be specified if DIRECTION\n                is ingress.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            srcIPRanges:\n              type: array\n              x-dcl-go-name: SrcIPRanges\n              description: CIDR IP address range. Maximum number of source CIDR IP\n                ranges allowed is 5000.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            srcRegionCodes:\n              type: array\n              x-dcl-go-name: SrcRegionCodes\n              description: The Unicode country codes whose IP addresses will be used\n                to match against the source of traffic. Can only be specified if DIRECTION\n                is ingress.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            srcSecureTags:\n              type: array\n              x-dcl-go-name: SrcSecureTags\n              description: List of secure tag values, which should be matched at the\n                source of the traffic. For INGRESS rule, if all the <code>srcSecureTag</code>\n                are INEFFECTIVE, and there is no <code>srcIpRange</code>, this rule\n                will be ignored. Maximum number of source tag values allowed is 256.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: NetworkFirewallPolicyRuleMatchSrcSecureTags\n                required:\n                - name\n                properties:\n                  name:\n                    type: string\n                    x-dcl-go-name: Name\n                    description: Name of the secure tag, created with TagManager's\n                      TagValue API. @pattern tagValues/[0-9]+\n                    x-dcl-references:\n                    - resource: Cloudresourcemanager/TagValue\n                      field: namespacedName\n                  state:\n                    type: string\n                    x-dcl-go-name: State\n                    x-dcl-go-type: NetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum\n                    readOnly: true\n                    description: '[Output Only] State of the secure tag, either `EFFECTIVE`\n                      or `INEFFECTIVE`. A secure tag is `INEFFECTIVE` when it is deleted\n                      or its network is deleted.'\n                    enum:\n                    - EFFECTIVE\n                    - INEFFECTIVE\n            srcThreatIntelligences:\n              type: array\n              x-dcl-go-name: SrcThreatIntelligences\n              description: Name of the Google Cloud Threat Intelligence list.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n        priority:\n          type: integer\n          format: int64\n          x-dcl-go-name: Priority\n          description: An integer indicating the priority of a rule in the list. The\n            priority must be a positive value between 0 and 2147483647. Rules are\n            evaluated from highest to lowest priority where 0 is the highest priority\n            and 2147483647 is the lowest prority.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        ruleName:\n          type: string\n          x-dcl-go-name: RuleName\n          description: An optional name for the rule. This field is not a unique identifier\n            and can be updated.\n        ruleTupleCount:\n          type: integer\n          format: int64\n          x-dcl-go-name: RuleTupleCount\n          readOnly: true\n          description: Calculation of the complexity of a single firewall policy rule.\n          x-kubernetes-immutable: true\n        targetSecureTags:\n          type: array\n          x-dcl-go-name: TargetSecureTags\n          description: A list of secure tags that controls which instances the firewall\n            rule applies to. If <code>targetSecureTag</code> are specified, then the\n            firewall rule applies only to instances in the VPC network that have one\n            of those EFFECTIVE secure tags, if all the target_secure_tag are in INEFFECTIVE\n            state, then this rule will be ignored. <code>targetSecureTag</code> may\n            not be set at the same time as <code>targetServiceAccounts</code>. If\n            neither <code>targetServiceAccounts</code> nor <code>targetSecureTag</code>\n            are specified, the firewall rule applies to all instances on the specified\n            network. Maximum number of target label tags allowed is 256.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: NetworkFirewallPolicyRuleTargetSecureTags\n            required:\n            - name\n            properties:\n              name:\n                type: string\n                x-dcl-go-name: Name\n                description: Name of the secure tag, created with TagManager's TagValue\n                  API. @pattern tagValues/[0-9]+\n                x-dcl-references:\n                - resource: Cloudresourcemanager/TagValue\n                  field: namespacedName\n              state:\n                type: string\n                x-dcl-go-name: State\n                x-dcl-go-type: NetworkFirewallPolicyRuleTargetSecureTagsStateEnum\n                readOnly: true\n                description: '[Output Only] State of the secure tag, either `EFFECTIVE`\n                  or `INEFFECTIVE`. A secure tag is `INEFFECTIVE` when it is deleted\n                  or its network is deleted.'\n                enum:\n                - EFFECTIVE\n                - INEFFECTIVE\n        targetServiceAccounts:\n          type: array\n          x-dcl-go-name: TargetServiceAccounts\n          description: A list of service accounts indicating the sets of instances\n            that are applied with this rule.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Iam/ServiceAccount\n              field: name\n")

// 14926 bytes
// MD5: 0bb3c3391f1c252a5044d3395b1ec451
