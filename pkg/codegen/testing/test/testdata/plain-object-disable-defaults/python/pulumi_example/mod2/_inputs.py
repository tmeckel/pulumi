# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import mod1 as _mod1

__all__ = [
    'TypArgs',
]

@pulumi.input_type
class TypArgs:
    def __init__(__self__, *,
                 mod1: pulumi.Input[Optional['_mod1.TypArgs']] = None,
                 val: pulumi.Input[Optional[str]] = None):
        """
        A test for namespaces (mod 2)
        """
        if mod1 is not None:
            pulumi.set(__self__, "mod1", mod1)
        if val is None:
            val = 'mod2'
        if val is not None:
            pulumi.set(__self__, "val", val)

    @property
    @pulumi.getter
    def mod1(self) -> pulumi.Input[Optional['_mod1.TypArgs']]:
        return pulumi.get(self, "mod1")

    @mod1.setter
    def mod1(self, value: pulumi.Input[Optional['_mod1.TypArgs']]):
        pulumi.set(self, "mod1", value)

    @property
    @pulumi.getter
    def val(self) -> pulumi.Input[Optional[str]]:
        return pulumi.get(self, "val")

    @val.setter
    def val(self, value: pulumi.Input[Optional[str]]):
        pulumi.set(self, "val", value)


