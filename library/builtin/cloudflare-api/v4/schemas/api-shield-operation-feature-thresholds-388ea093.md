---
title: api-shield_operation_feature_thresholds
page_id: schema-api-shield-operation-feature-thresholds-388ea093
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_operation_feature_thresholds

```yaml
{"type": "object", "properties": {"thresholds": {"type": "object", "properties": {"auth_id_tokens": {"$ref": "#/components/schemas/api-shield_auth_id_tokens"}, "data_points": {"$ref": "#/components/schemas/api-shield_data_points"}, "last_updated": {"$ref": "#/components/schemas/api-shield_timestamp"}, "p50": {"$ref": "#/components/schemas/api-shield_p50"}, "p90": {"$ref": "#/components/schemas/api-shield_p90"}, "p99": {"$ref": "#/components/schemas/api-shield_p99"}, "period_seconds": {"$ref": "#/components/schemas/api-shield_period_seconds"}, "requests": {"$ref": "#/components/schemas/api-shield_requests"}, "suggested_threshold": {"$ref": "#/components/schemas/api-shield_suggested_threshold"}}}}, "readOnly": true, "required": ["period_seconds", "suggested_threshold", "p50", "p90", "p99", "requests", "auth_id_tokens", "data_points", "last_updated"]}
```
