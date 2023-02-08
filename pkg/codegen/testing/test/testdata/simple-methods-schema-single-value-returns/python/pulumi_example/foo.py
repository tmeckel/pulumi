# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FooArgs', 'Foo']

@pulumi.input_type
class FooArgs:
    def __init__(__self__):
        """
        The set of arguments for constructing a Foo resource.
        """
        pass


class Foo(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 __props__=None):
        """
        Create a Foo resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FooArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Foo resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FooArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FooArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FooArgs.__new__(FooArgs)

        super(Foo, __self__).__init__(
            'example::Foo',
            resource_name,
            __props__,
            opts,
            remote=True)

    @pulumi.output_type
    class GetKubeconfigResult:
        def __init__(__self__, kubeconfig=None):
            if kubeconfig and not isinstance(kubeconfig, str):
                raise TypeError("Expected argument 'kubeconfig' to be a str")
            pulumi.set(__self__, "kubeconfig", kubeconfig)

        @property
        @pulumi.getter
        def kubeconfig(self) -> str:
            return pulumi.get(self, "kubeconfig")

    def get_kubeconfig(__self__, *,
                       profile_name: pulumi.Input[Optional[str]] = None,
                       role_arn: pulumi.Input[Optional[str]] = None) -> pulumi.Output['str']:
        __args__ = dict()
        __args__['__self__'] = __self__
        __args__['profileName'] = profile_name
        __args__['roleArn'] = role_arn
        __result__ = pulumi.runtime.call('example::Foo/getKubeconfig', __args__, res=__self__, typ=Foo.GetKubeconfigResult)
        return __result__.kubeconfig

