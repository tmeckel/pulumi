# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'FooArgs',
]

@pulumi.input_type
class FooArgs:
    def __init__(__self__, *,
                 a: bool,
                 c: int,
                 e: str,
                 b: Optional[bool] = None,
                 d: Optional[int] = None,
                 f: Optional[str] = None):
        FooArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            a=a,
            c=c,
            e=e,
            b=b,
            d=d,
            f=f,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             b: Optional[bool] = None,
             d: Optional[int] = None,
             f: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'a' in kwargs:
            a = kwargs['a']
        if 'a' not in locals():
            raise TypeError("Missing required property 'a'")
        if 'c' in kwargs:
            c = kwargs['c']
        if 'c' not in locals():
            raise TypeError("Missing required property 'c'")
        if 'e' in kwargs:
            e = kwargs['e']
        if 'e' not in locals():
            raise TypeError("Missing required property 'e'")

        _setter("a", a)
        _setter("c", c)
        _setter("e", e)
        if b is not None:
            _setter("b", b)
        if d is not None:
            _setter("d", d)
        if f is not None:
            _setter("f", f)

    @property
    @pulumi.getter
    def a(self) -> bool:
        return pulumi.get(self, "a")

    @a.setter
    def a(self, value: bool):
        pulumi.set(self, "a", value)

    @property
    @pulumi.getter
    def c(self) -> int:
        return pulumi.get(self, "c")

    @c.setter
    def c(self, value: int):
        pulumi.set(self, "c", value)

    @property
    @pulumi.getter
    def e(self) -> str:
        return pulumi.get(self, "e")

    @e.setter
    def e(self, value: str):
        pulumi.set(self, "e", value)

    @property
    @pulumi.getter
    def b(self) -> Optional[bool]:
        return pulumi.get(self, "b")

    @b.setter
    def b(self, value: Optional[bool]):
        pulumi.set(self, "b", value)

    @property
    @pulumi.getter
    def d(self) -> Optional[int]:
        return pulumi.get(self, "d")

    @d.setter
    def d(self, value: Optional[int]):
        pulumi.set(self, "d", value)

    @property
    @pulumi.getter
    def f(self) -> Optional[str]:
        return pulumi.get(self, "f")

    @f.setter
    def f(self, value: Optional[str]):
        pulumi.set(self, "f", value)


