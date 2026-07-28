---
title: load-balancing_load-balancer_components-schemas-usage_response
page_id: schema-load-balancing-load-balancer-components-schemas-usage-response-5c993c48
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_load-balancer_components-schemas-usage_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/load-balancing_api-response-single"}, {"properties": {"result": {"description": "Load balancer usage counts for an account.", "type": "object", "properties": {"load_balancers": {"description": "The number of configured load balancers.", "type": "integer", "format": "int64", "example": 5}, "max_origins_per_pool": {"description": "The largest number of origins in any single pool.", "type": "integer", "format": "int64", "example": 5}, "minimum_monitor_interval": {"description": "The smallest monitor interval (in seconds) configured across all monitors.", "type": "integer", "format": "int64", "example": 60}, "monitors": {"description": "The number of configured monitors.", "type": "integer", "format": "int64", "example": 3}, "origins": {"description": "The number of configured origins across all pools.", "type": "integer", "format": "int64", "example": 12}, "pools": {"description": "The number of configured pools.", "type": "integer", "format": "int64", "example": 4}}, "required": ["load_balancers", "monitors", "origins", "pools", "max_origins_per_pool", "minimum_monitor_interval"]}}, "type": "object"}]}
```
