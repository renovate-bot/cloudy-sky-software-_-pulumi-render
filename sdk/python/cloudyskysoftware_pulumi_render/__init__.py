# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import cloudyskysoftware_pulumi_render.config as __config
    config = __config
    import cloudyskysoftware_pulumi_render.services as __services
    services = __services
else:
    config = _utilities.lazy_import('cloudyskysoftware_pulumi_render.config')
    services = _utilities.lazy_import('cloudyskysoftware_pulumi_render.services')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "render",
  "mod": "services",
  "fqn": "cloudyskysoftware_pulumi_render.services",
  "classes": {
   "render:services:CustomDomain": "CustomDomain",
   "render:services:Deploy": "Deploy",
   "render:services:Scale": "Scale",
   "render:services:StaticSite": "StaticSite",
   "render:services:Suspend": "Suspend",
   "render:services:WebService": "WebService"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "render",
  "token": "pulumi:providers:render",
  "fqn": "cloudyskysoftware_pulumi_render",
  "class": "Provider"
 }
]
"""
)
