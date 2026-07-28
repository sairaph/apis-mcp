---
title: api-shield_operation_feature_confidence_intervals
page_id: schema-api-shield-operation-feature-confidence-intervals-ffab1dc7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_operation_feature_confidence_intervals

```yaml
{"type": "object", "properties": {"confidence_intervals": {"type": "object", "properties": {"last_updated": {"$ref": "#/components/schemas/api-shield_timestamp"}, "suggested_threshold": {"type": "object", "properties": {"confidence_intervals": {"type": "object", "properties": {"p90": {"$ref": "#/components/schemas/api-shield_confidence_intervals_bounds"}, "p95": {"$ref": "#/components/schemas/api-shield_confidence_intervals_bounds"}, "p99": {"$ref": "#/components/schemas/api-shield_confidence_intervals_bounds"}}}, "mean": {"description": "Suggested threshold.", "type": "number", "example": 25.5, "x-auditable": true}}}}}}, "example": {"confidence_intervals": {"last_updated": "2014-01-01T05:20:00.12345Z", "suggested_threshold": {"confidence_intervals": {"p90": {"lower": 23.1, "upper": 23.9}, "p95": {"lower": 22, "upper": 24.1}, "p99": {"lower": 20.2, "upper": 30}}, "mean": 23.5}}}, "readOnly": true}
```
