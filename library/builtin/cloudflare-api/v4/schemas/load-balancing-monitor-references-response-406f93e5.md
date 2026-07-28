---
title: load-balancing_monitor-references-response
page_id: schema-load-balancing-monitor-references-response-406f93e5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_monitor-references-response

```yaml
{"allOf": [{"$ref": "#/components/schemas/load-balancing_api-response-common"}, {"properties": {"result": {"description": "List of resources that reference a given monitor.", "type": "array", "items": {"properties": {"reference_type": {"type": "string", "enum": ["*", "referral", "referrer"], "x-auditable": true}, "resource_id": {"type": "string", "x-auditable": true}, "resource_name": {"type": "string", "x-auditable": true}, "resource_type": {"type": "string", "x-auditable": true}}, "type": "object"}, "example": [{"reference_type": "referrer", "resource_id": "17b5962d775c646f3f9725cbc7a53df4", "resource_name": "primary-dc-1", "resource_type": "pool"}]}}, "type": "object"}]}
```
