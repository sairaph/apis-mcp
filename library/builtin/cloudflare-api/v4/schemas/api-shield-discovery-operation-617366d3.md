---
title: api-shield_discovery_operation
page_id: schema-api-shield-discovery-operation-617366d3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_discovery_operation

```yaml
{"allOf": [{"properties": {"features": {"$ref": "#/components/schemas/api-shield_traffic_stats"}, "id": {"$ref": "#/components/schemas/api-shield_uuid-2"}, "last_updated": {"$ref": "#/components/schemas/api-shield_timestamp-2"}, "origin": {"description": "API discovery engine(s) that discovered this operation", "type": "array", "items": {"$ref": "#/components/schemas/api-shield_api_discovery_origin"}}, "state": {"$ref": "#/components/schemas/api-shield_api_discovery_state"}}, "required": ["id", "last_updated", "state", "origin"], "type": "object"}, {"$ref": "#/components/schemas/api-shield_basic_operation"}]}
```
