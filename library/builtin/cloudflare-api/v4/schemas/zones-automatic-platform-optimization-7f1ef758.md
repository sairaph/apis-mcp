---
title: zones_automatic_platform_optimization
page_id: schema-zones-automatic-platform-optimization-7f1ef758
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_automatic_platform_optimization

```yaml
{"type": "object", "properties": {"cache_by_device_type": {"description": "Indicates whether or not [cache by device type](https://developers.cloudflare.com/automatic-platform-optimization/reference/cache-device-type/) is enabled.", "type": "boolean", "example": false}, "cf": {"description": "Indicates whether or not Cloudflare proxy is enabled.", "type": "boolean", "example": true, "default": false}, "enabled": {"description": "Indicates whether or not Automatic Platform Optimization is enabled.", "type": "boolean", "example": true, "default": false}, "hostnames": {"description": "An array of hostnames where Automatic Platform Optimization for WordPress is activated.", "type": "array", "items": {"format": "hostname", "type": "string"}, "example": ["www.example.com", "example.com", "shop.example.com"]}, "wordpress": {"description": "Indicates whether or not site is powered by WordPress.", "type": "boolean", "example": true, "default": false}, "wp_plugin": {"description": "Indicates whether or not [Cloudflare for WordPress plugin](https://wordpress.org/plugins/cloudflare/) is installed.", "type": "boolean", "example": true, "default": false}}, "required": ["enabled", "cf", "wordpress", "wp_plugin", "hostnames", "cache_by_device_type"]}
```
