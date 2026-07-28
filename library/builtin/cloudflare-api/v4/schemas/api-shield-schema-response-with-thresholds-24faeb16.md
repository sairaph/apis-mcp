---
title: api-shield_schema-response-with-thresholds
page_id: schema-api-shield-schema-response-with-thresholds-24faeb16
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_schema-response-with-thresholds

```yaml
{"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"schemas": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_openapi-with-thresholds"}}, "timestamp": {"type": "string", "x-auditable": true}}}}, "required": ["result"], "type": "object"}]}
```
