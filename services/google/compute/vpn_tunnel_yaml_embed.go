// Copyright 2022 Google LLC. All Rights Reserved.
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
// gen_go_data -package compute -var YAML_vpn_tunnel blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/vpn_tunnel.yaml

package compute

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/vpn_tunnel.yaml
var YAML_vpn_tunnel = []byte("info:\n  title: Compute/VpnTunnel\n  description: The Compute VpnTunnel resource\n  x-dcl-struct-name: VpnTunnel\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a VpnTunnel\n    parameters:\n    - name: VpnTunnel\n      required: true\n      description: A full instance of a VpnTunnel\n  apply:\n    description: The function used to apply information about a VpnTunnel\n    parameters:\n    - name: VpnTunnel\n      required: true\n      description: A full instance of a VpnTunnel\n  delete:\n    description: The function used to delete a VpnTunnel\n    parameters:\n    - name: VpnTunnel\n      required: true\n      description: A full instance of a VpnTunnel\n  deleteAll:\n    description: The function used to delete all VpnTunnel\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many VpnTunnel\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    VpnTunnel:\n      title: VpnTunnel\n      x-dcl-id: projects/{{project}}/regions/{{location}}/vpnTunnels/{{name}}\n      x-dcl-uses-state-hint: true\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - sharedSecret\n      - project\n      properties:\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description of this resource. Provide this property\n            when you create the resource.\n          x-kubernetes-immutable: true\n        detailedStatus:\n          type: string\n          x-dcl-go-name: DetailedStatus\n          readOnly: true\n          description: Detailed status message for the VPN tunnel.\n          x-kubernetes-immutable: true\n        id:\n          type: integer\n          format: int64\n          x-dcl-go-name: Id\n          readOnly: true\n          description: The unique identifier for the resource. This identifier is\n            defined by the server.\n          x-kubernetes-immutable: true\n        ikeVersion:\n          type: integer\n          format: int64\n          x-dcl-go-name: IkeVersion\n          description: IKE protocol version to use when establishing the VPN tunnel\n            with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default\n            version is 2.\n          x-kubernetes-immutable: true\n          default: 2\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Labels for this VPN Tunnel\n          x-kubernetes-immutable: true\n        localTrafficSelector:\n          type: array\n          x-dcl-go-name: LocalTrafficSelector\n          description: 'Local traffic selector to use when establishing the VPN tunnel\n            with the peer VPN gateway. The value should be a CIDR formatted string,\n            for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is\n            supported.'\n          x-kubernetes-immutable: true\n          x-dcl-server-default: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: set\n          items:\n            type: string\n            x-dcl-go-type: string\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: Name of the region where the VPN tunnel resides.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Name of the resource. Provided by the client when the resource\n            is created. The name must be 1-63 characters long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt)\n            Specifically, the name must be 1-63 characters long and match the regular\n            expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character\n            must be a lowercase letter, and all following characters must be a dash,\n            lowercase letter, or digit, except the last character, which cannot be\n            a dash.\n          x-kubernetes-immutable: true\n        peerExternalGateway:\n          type: string\n          x-dcl-go-name: PeerExternalGateway\n          description: URL of the peer side external VPN gateway to which this VPN\n            tunnel is connected. Provided by the client when the VPN tunnel is created.\n            This field is exclusive with the field peerGcpGateway.\n          x-kubernetes-immutable: true\n        peerExternalGatewayInterface:\n          type: integer\n          format: int64\n          x-dcl-go-name: PeerExternalGatewayInterface\n          description: The interface ID of the external VPN gateway to which this\n            VPN tunnel is connected. Provided by the client when the VPN tunnel is\n            created.\n          x-kubernetes-immutable: true\n        peerGcpGateway:\n          type: string\n          x-dcl-go-name: PeerGcpGateway\n          description: URL of the peer side HA GCP VPN gateway to which this VPN tunnel\n            is connected. Provided by the client when the VPN tunnel is created. This\n            field can be used when creating highly available VPN from VPC network\n            to VPC network, the field is exclusive with the field peerExternalGateway.\n            If provided, the VPN tunnel will automatically use the same vpnGatewayInterface\n            ID in the peer GCP VPN gateway.\n          x-kubernetes-immutable: true\n        peerIP:\n          type: string\n          x-dcl-go-name: PeerIP\n          description: IP address of the peer VPN gateway. Only IPv4 is supported.\n          x-kubernetes-immutable: true\n          x-dcl-server-default: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        remoteTrafficSelector:\n          type: array\n          x-dcl-go-name: RemoteTrafficSelector\n          description: 'Remote traffic selectors to use when establishing the VPN\n            tunnel with the peer VPN gateway. The value should be a CIDR formatted\n            string, for example: 192.168.0.0/16. The ranges should be disjoint. Only\n            IPv4 is supported.'\n          x-kubernetes-immutable: true\n          x-dcl-server-default: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: set\n          items:\n            type: string\n            x-dcl-go-type: string\n        router:\n          type: string\n          x-dcl-go-name: Router\n          description: URL of the router resource to be used for dynamic routing.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Compute/Router\n            field: selfLink\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: Server-defined URL for the resource.\n          x-kubernetes-immutable: true\n        sharedSecret:\n          type: string\n          x-dcl-go-name: SharedSecret\n          description: Shared secret used to set the secure session between the Cloud\n            VPN gateway and the peer VPN gateway.\n          x-kubernetes-immutable: true\n          x-dcl-sensitive: true\n          x-dcl-mutable-unreadable: true\n        sharedSecretHash:\n          type: string\n          x-dcl-go-name: SharedSecretHash\n          readOnly: true\n          description: Hash of the shared secret.\n          x-kubernetes-immutable: true\n        status:\n          type: string\n          x-dcl-go-name: Status\n          x-dcl-go-type: VpnTunnelStatusEnum\n          readOnly: true\n          description: 'The status of the VPN tunnel, which can be one of the following:  *\n            PROVISIONING: Resource is being allocated for the VPN tunnel.  * WAITING_FOR_FULL_CONFIG:\n            Waiting to receive all VPN-related configs from   the user. Network, TargetVpnGateway,\n            VpnTunnel, ForwardingRule, and Route   resources are needed to setup the\n            VPN tunnel.  * FIRST_HANDSHAKE: Successful first handshake with the peer\n            VPN.  * ESTABLISHED: Secure session is successfully established with the\n            peer VPN.  * NETWORK_ERROR: Deprecated, replaced by NO_INCOMING_PACKETS  *\n            AUTHORIZATION_ERROR: Auth error (for example, bad shared secret).  * NEGOTIATION_FAILURE:\n            Handshake failed.  * DEPROVISIONING: Resources are being deallocated for\n            the VPN tunnel.  * FAILED: Tunnel creation has failed and the tunnel is\n            not ready to be used.  * NO_INCOMING_PACKETS: No incoming packets from\n            peer.  * REJECTED: Tunnel configuration was rejected, can be result of\n            being blocklisted.  * ALLOCATING_RESOURCES: Cloud VPN is in the process\n            of allocating all required resources.  * STOPPED: Tunnel is stopped due\n            to its Forwarding Rules being deleted for Classic VPN tunnels or the project\n            is in frozen state.  * PEER_IDENTITY_MISMATCH: Peer identity does not\n            match peer IP, probably behind NAT.  * TS_NARROWING_NOT_ALLOWED: Traffic\n            selector narrowing not allowed for an HA-VPN tunnel.'\n          x-kubernetes-immutable: true\n          enum:\n          - PROVISIONING\n          - WAITING_FOR_FULL_CONFIG\n          - FIRST_HANDSHAKE\n          - ESTABLISHED\n          - NO_INCOMING_PACKETS\n          - AUTHORIZATION_ERROR\n          - NEGOTIATION_FAILURE\n          - DEPROVISIONING\n          - FAILED\n          - REJECTED\n          - ALLOCATING_RESOURCES\n          - STOPPED\n          - PEER_IDENTITY_MISMATCH\n          - TS_NARROWING_NOT_ALLOWED\n        targetVpnGateway:\n          type: string\n          x-dcl-go-name: TargetVpnGateway\n          description: URL of the Target VPN gateway with which this VPN tunnel is\n            associated. Provided by the client when the VPN tunnel is created.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Compute/TargetVpnGateway\n            field: selfLink\n        vpnGateway:\n          type: string\n          x-dcl-go-name: VpnGateway\n          description: URL of the VPN gateway with which this VPN tunnel is associated.\n            Provided by the client when the VPN tunnel is created. This must be used\n            (instead of target_vpn_gateway) if a High Availability VPN gateway resource\n            is created.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Compute/VpnGateway\n            field: selfLink\n        vpnGatewayInterface:\n          type: integer\n          format: int64\n          x-dcl-go-name: VpnGatewayInterface\n          description: The interface ID of the VPN gateway with which this VPN tunnel\n            is associated.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n")

// 11275 bytes
// MD5: 38a121c484b2c2983ca2a74df3e5b088
