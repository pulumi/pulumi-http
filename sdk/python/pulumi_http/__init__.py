# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .get_http import *
from .provider import *
from ._inputs import *
from . import outputs
_utilities.register(
    resource_modules="""
[]
""",
    resource_packages="""
[
 {
  "pkg": "http",
  "token": "pulumi:providers:http",
  "fqn": "pulumi_http",
  "class": "Provider"
 }
]
"""
)
