# Copyright 2023 Google LLC. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
from connector import channel
from google3.cloud.graphite.mmv2.services.google.container_aws import node_pool_pb2
from google3.cloud.graphite.mmv2.services.google.container_aws import node_pool_pb2_grpc

from typing import List


class NodePool(object):
    def __init__(
        self,
        name: str = None,
        version: str = None,
        config: dict = None,
        autoscaling: dict = None,
        subnet_id: str = None,
        state: str = None,
        uid: str = None,
        reconciling: bool = None,
        create_time: str = None,
        update_time: str = None,
        etag: str = None,
        annotations: dict = None,
        max_pods_constraint: dict = None,
        project: str = None,
        location: str = None,
        cluster: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.version = version
        self.config = config
        self.autoscaling = autoscaling
        self.subnet_id = subnet_id
        self.annotations = annotations
        self.max_pods_constraint = max_pods_constraint
        self.project = project
        self.location = location
        self.cluster = cluster
        self.service_account_file = service_account_file

    def apply(self):
        stub = node_pool_pb2_grpc.ContainerawsNodePoolServiceStub(channel.Channel())
        request = node_pool_pb2.ApplyContainerawsNodePoolRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if NodePoolConfig.to_proto(self.config):
            request.resource.config.CopyFrom(NodePoolConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if NodePoolAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                NodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if Primitive.to_proto(self.subnet_id):
            request.resource.subnet_id = Primitive.to_proto(self.subnet_id)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            request.resource.max_pods_constraint.CopyFrom(
                NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            request.resource.ClearField("max_pods_constraint")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.cluster):
            request.resource.cluster = Primitive.to_proto(self.cluster)

        request.service_account_file = self.service_account_file

        response = stub.ApplyContainerawsNodePool(request)
        self.name = Primitive.from_proto(response.name)
        self.version = Primitive.from_proto(response.version)
        self.config = NodePoolConfig.from_proto(response.config)
        self.autoscaling = NodePoolAutoscaling.from_proto(response.autoscaling)
        self.subnet_id = Primitive.from_proto(response.subnet_id)
        self.state = NodePoolStateEnum.from_proto(response.state)
        self.uid = Primitive.from_proto(response.uid)
        self.reconciling = Primitive.from_proto(response.reconciling)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.etag = Primitive.from_proto(response.etag)
        self.annotations = Primitive.from_proto(response.annotations)
        self.max_pods_constraint = NodePoolMaxPodsConstraint.from_proto(
            response.max_pods_constraint
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.cluster = Primitive.from_proto(response.cluster)

    def delete(self):
        stub = node_pool_pb2_grpc.ContainerawsNodePoolServiceStub(channel.Channel())
        request = node_pool_pb2.DeleteContainerawsNodePoolRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if NodePoolConfig.to_proto(self.config):
            request.resource.config.CopyFrom(NodePoolConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if NodePoolAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                NodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if Primitive.to_proto(self.subnet_id):
            request.resource.subnet_id = Primitive.to_proto(self.subnet_id)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            request.resource.max_pods_constraint.CopyFrom(
                NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            request.resource.ClearField("max_pods_constraint")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.cluster):
            request.resource.cluster = Primitive.to_proto(self.cluster)

        response = stub.DeleteContainerawsNodePool(request)

    @classmethod
    def list(self, project, location, cluster, service_account_file=""):
        stub = node_pool_pb2_grpc.ContainerawsNodePoolServiceStub(channel.Channel())
        request = node_pool_pb2.ListContainerawsNodePoolRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.Cluster = cluster

        return stub.ListContainerawsNodePool(request).items

    def to_proto(self):
        resource = node_pool_pb2.ContainerawsNodePool()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.version):
            resource.version = Primitive.to_proto(self.version)
        if NodePoolConfig.to_proto(self.config):
            resource.config.CopyFrom(NodePoolConfig.to_proto(self.config))
        else:
            resource.ClearField("config")
        if NodePoolAutoscaling.to_proto(self.autoscaling):
            resource.autoscaling.CopyFrom(
                NodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            resource.ClearField("autoscaling")
        if Primitive.to_proto(self.subnet_id):
            resource.subnet_id = Primitive.to_proto(self.subnet_id)
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            resource.max_pods_constraint.CopyFrom(
                NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            resource.ClearField("max_pods_constraint")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.cluster):
            resource.cluster = Primitive.to_proto(self.cluster)
        return resource


class NodePoolConfig(object):
    def __init__(
        self,
        instance_type: str = None,
        root_volume: dict = None,
        taints: list = None,
        labels: dict = None,
        tags: dict = None,
        iam_instance_profile: str = None,
        config_encryption: dict = None,
        ssh_config: dict = None,
        security_group_ids: list = None,
        proxy_config: dict = None,
        autoscaling_metrics_collection: dict = None,
    ):
        self.instance_type = instance_type
        self.root_volume = root_volume
        self.taints = taints
        self.labels = labels
        self.tags = tags
        self.iam_instance_profile = iam_instance_profile
        self.config_encryption = config_encryption
        self.ssh_config = ssh_config
        self.security_group_ids = security_group_ids
        self.proxy_config = proxy_config
        self.autoscaling_metrics_collection = autoscaling_metrics_collection

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerawsNodePoolConfig()
        if Primitive.to_proto(resource.instance_type):
            res.instance_type = Primitive.to_proto(resource.instance_type)
        if NodePoolConfigRootVolume.to_proto(resource.root_volume):
            res.root_volume.CopyFrom(
                NodePoolConfigRootVolume.to_proto(resource.root_volume)
            )
        else:
            res.ClearField("root_volume")
        if NodePoolConfigTaintsArray.to_proto(resource.taints):
            res.taints.extend(NodePoolConfigTaintsArray.to_proto(resource.taints))
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        if Primitive.to_proto(resource.tags):
            res.tags = Primitive.to_proto(resource.tags)
        if Primitive.to_proto(resource.iam_instance_profile):
            res.iam_instance_profile = Primitive.to_proto(resource.iam_instance_profile)
        if NodePoolConfigConfigEncryption.to_proto(resource.config_encryption):
            res.config_encryption.CopyFrom(
                NodePoolConfigConfigEncryption.to_proto(resource.config_encryption)
            )
        else:
            res.ClearField("config_encryption")
        if NodePoolConfigSshConfig.to_proto(resource.ssh_config):
            res.ssh_config.CopyFrom(
                NodePoolConfigSshConfig.to_proto(resource.ssh_config)
            )
        else:
            res.ClearField("ssh_config")
        if Primitive.to_proto(resource.security_group_ids):
            res.security_group_ids.extend(
                Primitive.to_proto(resource.security_group_ids)
            )
        if NodePoolConfigProxyConfig.to_proto(resource.proxy_config):
            res.proxy_config.CopyFrom(
                NodePoolConfigProxyConfig.to_proto(resource.proxy_config)
            )
        else:
            res.ClearField("proxy_config")
        if NodePoolConfigAutoscalingMetricsCollection.to_proto(
            resource.autoscaling_metrics_collection
        ):
            res.autoscaling_metrics_collection.CopyFrom(
                NodePoolConfigAutoscalingMetricsCollection.to_proto(
                    resource.autoscaling_metrics_collection
                )
            )
        else:
            res.ClearField("autoscaling_metrics_collection")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConfig(
            instance_type=Primitive.from_proto(resource.instance_type),
            root_volume=NodePoolConfigRootVolume.from_proto(resource.root_volume),
            taints=NodePoolConfigTaintsArray.from_proto(resource.taints),
            labels=Primitive.from_proto(resource.labels),
            tags=Primitive.from_proto(resource.tags),
            iam_instance_profile=Primitive.from_proto(resource.iam_instance_profile),
            config_encryption=NodePoolConfigConfigEncryption.from_proto(
                resource.config_encryption
            ),
            ssh_config=NodePoolConfigSshConfig.from_proto(resource.ssh_config),
            security_group_ids=Primitive.from_proto(resource.security_group_ids),
            proxy_config=NodePoolConfigProxyConfig.from_proto(resource.proxy_config),
            autoscaling_metrics_collection=NodePoolConfigAutoscalingMetricsCollection.from_proto(
                resource.autoscaling_metrics_collection
            ),
        )


class NodePoolConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConfig.from_proto(i) for i in resources]


class NodePoolConfigRootVolume(object):
    def __init__(
        self,
        size_gib: int = None,
        volume_type: str = None,
        iops: int = None,
        kms_key_arn: str = None,
    ):
        self.size_gib = size_gib
        self.volume_type = volume_type
        self.iops = iops
        self.kms_key_arn = kms_key_arn

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerawsNodePoolConfigRootVolume()
        if Primitive.to_proto(resource.size_gib):
            res.size_gib = Primitive.to_proto(resource.size_gib)
        if NodePoolConfigRootVolumeVolumeTypeEnum.to_proto(resource.volume_type):
            res.volume_type = NodePoolConfigRootVolumeVolumeTypeEnum.to_proto(
                resource.volume_type
            )
        if Primitive.to_proto(resource.iops):
            res.iops = Primitive.to_proto(resource.iops)
        if Primitive.to_proto(resource.kms_key_arn):
            res.kms_key_arn = Primitive.to_proto(resource.kms_key_arn)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConfigRootVolume(
            size_gib=Primitive.from_proto(resource.size_gib),
            volume_type=NodePoolConfigRootVolumeVolumeTypeEnum.from_proto(
                resource.volume_type
            ),
            iops=Primitive.from_proto(resource.iops),
            kms_key_arn=Primitive.from_proto(resource.kms_key_arn),
        )


class NodePoolConfigRootVolumeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConfigRootVolume.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConfigRootVolume.from_proto(i) for i in resources]


class NodePoolConfigTaints(object):
    def __init__(self, key: str = None, value: str = None, effect: str = None):
        self.key = key
        self.value = value
        self.effect = effect

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerawsNodePoolConfigTaints()
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if NodePoolConfigTaintsEffectEnum.to_proto(resource.effect):
            res.effect = NodePoolConfigTaintsEffectEnum.to_proto(resource.effect)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConfigTaints(
            key=Primitive.from_proto(resource.key),
            value=Primitive.from_proto(resource.value),
            effect=NodePoolConfigTaintsEffectEnum.from_proto(resource.effect),
        )


class NodePoolConfigTaintsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConfigTaints.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConfigTaints.from_proto(i) for i in resources]


class NodePoolConfigConfigEncryption(object):
    def __init__(self, kms_key_arn: str = None):
        self.kms_key_arn = kms_key_arn

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerawsNodePoolConfigConfigEncryption()
        if Primitive.to_proto(resource.kms_key_arn):
            res.kms_key_arn = Primitive.to_proto(resource.kms_key_arn)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConfigConfigEncryption(
            kms_key_arn=Primitive.from_proto(resource.kms_key_arn),
        )


class NodePoolConfigConfigEncryptionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConfigConfigEncryption.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConfigConfigEncryption.from_proto(i) for i in resources]


class NodePoolConfigSshConfig(object):
    def __init__(self, ec2_key_pair: str = None):
        self.ec2_key_pair = ec2_key_pair

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerawsNodePoolConfigSshConfig()
        if Primitive.to_proto(resource.ec2_key_pair):
            res.ec2_key_pair = Primitive.to_proto(resource.ec2_key_pair)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConfigSshConfig(
            ec2_key_pair=Primitive.from_proto(resource.ec2_key_pair),
        )


class NodePoolConfigSshConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConfigSshConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConfigSshConfig.from_proto(i) for i in resources]


class NodePoolConfigProxyConfig(object):
    def __init__(self, secret_arn: str = None, secret_version: str = None):
        self.secret_arn = secret_arn
        self.secret_version = secret_version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerawsNodePoolConfigProxyConfig()
        if Primitive.to_proto(resource.secret_arn):
            res.secret_arn = Primitive.to_proto(resource.secret_arn)
        if Primitive.to_proto(resource.secret_version):
            res.secret_version = Primitive.to_proto(resource.secret_version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConfigProxyConfig(
            secret_arn=Primitive.from_proto(resource.secret_arn),
            secret_version=Primitive.from_proto(resource.secret_version),
        )


class NodePoolConfigProxyConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConfigProxyConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConfigProxyConfig.from_proto(i) for i in resources]


class NodePoolConfigAutoscalingMetricsCollection(object):
    def __init__(self, granularity: str = None, metrics: list = None):
        self.granularity = granularity
        self.metrics = metrics

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerawsNodePoolConfigAutoscalingMetricsCollection()
        if Primitive.to_proto(resource.granularity):
            res.granularity = Primitive.to_proto(resource.granularity)
        if Primitive.to_proto(resource.metrics):
            res.metrics.extend(Primitive.to_proto(resource.metrics))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConfigAutoscalingMetricsCollection(
            granularity=Primitive.from_proto(resource.granularity),
            metrics=Primitive.from_proto(resource.metrics),
        )


class NodePoolConfigAutoscalingMetricsCollectionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            NodePoolConfigAutoscalingMetricsCollection.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            NodePoolConfigAutoscalingMetricsCollection.from_proto(i) for i in resources
        ]


class NodePoolAutoscaling(object):
    def __init__(self, min_node_count: int = None, max_node_count: int = None):
        self.min_node_count = min_node_count
        self.max_node_count = max_node_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerawsNodePoolAutoscaling()
        if Primitive.to_proto(resource.min_node_count):
            res.min_node_count = Primitive.to_proto(resource.min_node_count)
        if Primitive.to_proto(resource.max_node_count):
            res.max_node_count = Primitive.to_proto(resource.max_node_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolAutoscaling(
            min_node_count=Primitive.from_proto(resource.min_node_count),
            max_node_count=Primitive.from_proto(resource.max_node_count),
        )


class NodePoolAutoscalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolAutoscaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolAutoscaling.from_proto(i) for i in resources]


class NodePoolMaxPodsConstraint(object):
    def __init__(self, max_pods_per_node: int = None):
        self.max_pods_per_node = max_pods_per_node

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerawsNodePoolMaxPodsConstraint()
        if Primitive.to_proto(resource.max_pods_per_node):
            res.max_pods_per_node = Primitive.to_proto(resource.max_pods_per_node)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolMaxPodsConstraint(
            max_pods_per_node=Primitive.from_proto(resource.max_pods_per_node),
        )


class NodePoolMaxPodsConstraintArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolMaxPodsConstraint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolMaxPodsConstraint.from_proto(i) for i in resources]


class NodePoolConfigRootVolumeVolumeTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return node_pool_pb2.ContainerawsNodePoolConfigRootVolumeVolumeTypeEnum.Value(
            "ContainerawsNodePoolConfigRootVolumeVolumeTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return node_pool_pb2.ContainerawsNodePoolConfigRootVolumeVolumeTypeEnum.Name(
            resource
        )[len("ContainerawsNodePoolConfigRootVolumeVolumeTypeEnum") :]


class NodePoolConfigTaintsEffectEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return node_pool_pb2.ContainerawsNodePoolConfigTaintsEffectEnum.Value(
            "ContainerawsNodePoolConfigTaintsEffectEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return node_pool_pb2.ContainerawsNodePoolConfigTaintsEffectEnum.Name(resource)[
            len("ContainerawsNodePoolConfigTaintsEffectEnum") :
        ]


class NodePoolStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return node_pool_pb2.ContainerawsNodePoolStateEnum.Value(
            "ContainerawsNodePoolStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return node_pool_pb2.ContainerawsNodePoolStateEnum.Name(resource)[
            len("ContainerawsNodePoolStateEnum") :
        ]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
