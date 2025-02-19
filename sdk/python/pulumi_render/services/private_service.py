# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['PrivateServiceArgs', 'PrivateService']

@pulumi.input_type
class PrivateServiceArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 owner_id: pulumi.Input[str],
                 repo: pulumi.Input[str],
                 auto_deploy: Optional[pulumi.Input['ServiceAutoDeploy']] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 env_vars: Optional[pulumi.Input[Sequence[pulumi.Input['EnvVarKeyValueArgs']]]] = None,
                 notify_on_fail: Optional[pulumi.Input['ServiceNotifyOnFail']] = None,
                 secret_files: Optional[pulumi.Input[Sequence[pulumi.Input['SecretFileArgs']]]] = None,
                 service_details: Optional[pulumi.Input['PrivateServiceDetailsArgs']] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 suspended: Optional[pulumi.Input['ServiceSuspended']] = None,
                 suspenders: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PrivateService resource.
        :param pulumi.Input[str] owner_id: The id of the owner (user/team).
        :param pulumi.Input[str] repo: Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
        :param pulumi.Input['ServiceAutoDeploy'] auto_deploy: Whether to auto deploy the service or not upon git push.
        :param pulumi.Input[str] branch: If left empty, this will fall back to the default branch of the repository.
        :param pulumi.Input['ServiceNotifyOnFail'] notify_on_fail: The notification setting for this service upon deployment failure.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "owner_id", owner_id)
        pulumi.set(__self__, "repo", repo)
        if auto_deploy is None:
            auto_deploy = 'no'
        if auto_deploy is not None:
            pulumi.set(__self__, "auto_deploy", auto_deploy)
        if branch is not None:
            pulumi.set(__self__, "branch", branch)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if env_vars is not None:
            pulumi.set(__self__, "env_vars", env_vars)
        if notify_on_fail is not None:
            pulumi.set(__self__, "notify_on_fail", notify_on_fail)
        if secret_files is not None:
            pulumi.set(__self__, "secret_files", secret_files)
        if service_details is not None:
            pulumi.set(__self__, "service_details", service_details)
        if slug is not None:
            pulumi.set(__self__, "slug", slug)
        if suspended is not None:
            pulumi.set(__self__, "suspended", suspended)
        if suspenders is not None:
            pulumi.set(__self__, "suspenders", suspenders)
        if type is None:
            type = 'private_service'
        if type is not None:
            pulumi.set(__self__, "type", type)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Input[str]:
        """
        The id of the owner (user/team).
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter
    def repo(self) -> pulumi.Input[str]:
        """
        Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
        """
        return pulumi.get(self, "repo")

    @repo.setter
    def repo(self, value: pulumi.Input[str]):
        pulumi.set(self, "repo", value)

    @property
    @pulumi.getter(name="autoDeploy")
    def auto_deploy(self) -> Optional[pulumi.Input['ServiceAutoDeploy']]:
        """
        Whether to auto deploy the service or not upon git push.
        """
        return pulumi.get(self, "auto_deploy")

    @auto_deploy.setter
    def auto_deploy(self, value: Optional[pulumi.Input['ServiceAutoDeploy']]):
        pulumi.set(self, "auto_deploy", value)

    @property
    @pulumi.getter
    def branch(self) -> Optional[pulumi.Input[str]]:
        """
        If left empty, this will fall back to the default branch of the repository.
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="envVars")
    def env_vars(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EnvVarKeyValueArgs']]]]:
        return pulumi.get(self, "env_vars")

    @env_vars.setter
    def env_vars(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EnvVarKeyValueArgs']]]]):
        pulumi.set(self, "env_vars", value)

    @property
    @pulumi.getter(name="notifyOnFail")
    def notify_on_fail(self) -> Optional[pulumi.Input['ServiceNotifyOnFail']]:
        """
        The notification setting for this service upon deployment failure.
        """
        return pulumi.get(self, "notify_on_fail")

    @notify_on_fail.setter
    def notify_on_fail(self, value: Optional[pulumi.Input['ServiceNotifyOnFail']]):
        pulumi.set(self, "notify_on_fail", value)

    @property
    @pulumi.getter(name="secretFiles")
    def secret_files(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecretFileArgs']]]]:
        return pulumi.get(self, "secret_files")

    @secret_files.setter
    def secret_files(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecretFileArgs']]]]):
        pulumi.set(self, "secret_files", value)

    @property
    @pulumi.getter(name="serviceDetails")
    def service_details(self) -> Optional[pulumi.Input['PrivateServiceDetailsArgs']]:
        return pulumi.get(self, "service_details")

    @service_details.setter
    def service_details(self, value: Optional[pulumi.Input['PrivateServiceDetailsArgs']]):
        pulumi.set(self, "service_details", value)

    @property
    @pulumi.getter
    def slug(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "slug")

    @slug.setter
    def slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slug", value)

    @property
    @pulumi.getter
    def suspended(self) -> Optional[pulumi.Input['ServiceSuspended']]:
        return pulumi.get(self, "suspended")

    @suspended.setter
    def suspended(self, value: Optional[pulumi.Input['ServiceSuspended']]):
        pulumi.set(self, "suspended", value)

    @property
    @pulumi.getter
    def suspenders(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "suspenders")

    @suspenders.setter
    def suspenders(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "suspenders", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class PrivateService(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_deploy: Optional[pulumi.Input['ServiceAutoDeploy']] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 env_vars: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnvVarKeyValueArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_on_fail: Optional[pulumi.Input['ServiceNotifyOnFail']] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 repo: Optional[pulumi.Input[str]] = None,
                 secret_files: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretFileArgs']]]]] = None,
                 service_details: Optional[pulumi.Input[pulumi.InputType['PrivateServiceDetailsArgs']]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 suspended: Optional[pulumi.Input['ServiceSuspended']] = None,
                 suspenders: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A private service

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['ServiceAutoDeploy'] auto_deploy: Whether to auto deploy the service or not upon git push.
        :param pulumi.Input[str] branch: If left empty, this will fall back to the default branch of the repository.
        :param pulumi.Input['ServiceNotifyOnFail'] notify_on_fail: The notification setting for this service upon deployment failure.
        :param pulumi.Input[str] owner_id: The id of the owner (user/team).
        :param pulumi.Input[str] repo: Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A private service

        :param str resource_name: The name of the resource.
        :param PrivateServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_deploy: Optional[pulumi.Input['ServiceAutoDeploy']] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 env_vars: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnvVarKeyValueArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_on_fail: Optional[pulumi.Input['ServiceNotifyOnFail']] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 repo: Optional[pulumi.Input[str]] = None,
                 secret_files: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretFileArgs']]]]] = None,
                 service_details: Optional[pulumi.Input[pulumi.InputType['PrivateServiceDetailsArgs']]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 suspended: Optional[pulumi.Input['ServiceSuspended']] = None,
                 suspenders: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateServiceArgs.__new__(PrivateServiceArgs)

            if auto_deploy is None:
                auto_deploy = 'no'
            __props__.__dict__["auto_deploy"] = auto_deploy
            __props__.__dict__["branch"] = branch
            __props__.__dict__["created_at"] = created_at
            __props__.__dict__["env_vars"] = env_vars
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["notify_on_fail"] = notify_on_fail
            if owner_id is None and not opts.urn:
                raise TypeError("Missing required property 'owner_id'")
            __props__.__dict__["owner_id"] = owner_id
            if repo is None and not opts.urn:
                raise TypeError("Missing required property 'repo'")
            __props__.__dict__["repo"] = repo
            __props__.__dict__["secret_files"] = secret_files
            __props__.__dict__["service_details"] = service_details
            __props__.__dict__["slug"] = slug
            __props__.__dict__["suspended"] = suspended
            __props__.__dict__["suspenders"] = suspenders
            if type is None:
                type = 'private_service'
            __props__.__dict__["type"] = type
            __props__.__dict__["updated_at"] = updated_at
        super(PrivateService, __self__).__init__(
            'render:services:PrivateService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PrivateService':
        """
        Get an existing PrivateService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PrivateServiceArgs.__new__(PrivateServiceArgs)

        __props__.__dict__["auto_deploy"] = None
        __props__.__dict__["branch"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["env_vars"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["notify_on_fail"] = None
        __props__.__dict__["owner_id"] = None
        __props__.__dict__["repo"] = None
        __props__.__dict__["secret_files"] = None
        __props__.__dict__["service_details"] = None
        __props__.__dict__["slug"] = None
        __props__.__dict__["suspended"] = None
        __props__.__dict__["suspenders"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["updated_at"] = None
        return PrivateService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoDeploy")
    def auto_deploy(self) -> pulumi.Output[Optional['ServiceAutoDeploy']]:
        """
        Whether to auto deploy the service or not upon git push.
        """
        return pulumi.get(self, "auto_deploy")

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Output[Optional[str]]:
        """
        If left empty, this will fall back to the default branch of the repository.
        """
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="envVars")
    def env_vars(self) -> pulumi.Output[Optional[Sequence['outputs.EnvVarKeyValue']]]:
        return pulumi.get(self, "env_vars")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notifyOnFail")
    def notify_on_fail(self) -> pulumi.Output[Optional['ServiceNotifyOnFail']]:
        """
        The notification setting for this service upon deployment failure.
        """
        return pulumi.get(self, "notify_on_fail")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[Optional[str]]:
        """
        The id of the owner (user/team).
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def repo(self) -> pulumi.Output[Optional[str]]:
        """
        Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
        """
        return pulumi.get(self, "repo")

    @property
    @pulumi.getter(name="secretFiles")
    def secret_files(self) -> pulumi.Output[Optional[Sequence['outputs.SecretFile']]]:
        return pulumi.get(self, "secret_files")

    @property
    @pulumi.getter(name="serviceDetails")
    def service_details(self) -> pulumi.Output[Optional['outputs.PrivateServiceDetails']]:
        return pulumi.get(self, "service_details")

    @property
    @pulumi.getter
    def slug(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter
    def suspended(self) -> pulumi.Output[Optional['ServiceSuspended']]:
        return pulumi.get(self, "suspended")

    @property
    @pulumi.getter
    def suspenders(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "suspenders")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "updated_at")

