---
title: load-balancing_preview_result
page_id: schema-load-balancing-preview-result-3dea3c55
path: schemas
description: Resulting health data from a preview operation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_preview_result

Resulting health data from a preview operation.

```yaml
{"description": "Resulting health data from a preview operation.", "type": "object", "example": {"abwlnp5jbqn45ecgxd03erbgtxtqai0d": {"healthy": true, "origins": [{"originone.example.com.": {"failure_reason": "No failures", "healthy": true, "response_code": 200, "rtt": "66ms"}}]}}, "additionalProperties": {"properties": {"healthy": {"type": "boolean", "x-auditable": true}, "origins": {"type": "array", "items": {"additionalProperties": {"$ref": "#/components/schemas/load-balancing_origin_health_data"}, "maxProperties": 1, "minProperties": 1, "type": "object"}}}, "type": "object"}}
```
