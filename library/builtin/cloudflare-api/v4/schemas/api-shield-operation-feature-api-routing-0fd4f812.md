---
title: api-shield_operation_feature_api_routing
page_id: schema-api-shield-operation-feature-api-routing-0fd4f812
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_operation_feature_api_routing

```yaml
{"type": "object", "properties": {"api_routing": {"description": "API Routing settings on endpoint.", "type": "object", "properties": {"last_updated": {"$ref": "#/components/schemas/api-shield_timestamp"}, "route": {"description": "Target route.", "type": "string", "example": "https://api.example.com/api/service", "x-auditable": true}}}}, "example": {"api_routing": {"last_updated": "2014-01-01T05:20:00.12345Z", "route": "https://api.example.com/api/service"}}, "readOnly": true}
```
