---
title: api-shield_schema_response_discovery
page_id: schema-api-shield-schema-response-discovery-566888fb
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_schema_response_discovery

```yaml
{"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"schemas": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_openapi"}}, "timestamp": {"$ref": "#/components/schemas/api-shield_timestamp-2"}}, "required": ["timestamp", "schemas"]}}, "required": ["result"], "type": "object"}]}
```
