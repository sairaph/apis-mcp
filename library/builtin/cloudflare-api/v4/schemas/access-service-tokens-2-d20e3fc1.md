---
title: access_service-tokens-2
page_id: schema-access-service-tokens-2-d20e3fc1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_service-tokens-2

```yaml
{"type": "object", "properties": {"client_id": {"$ref": "#/components/schemas/access_client_id"}, "created_at": {"$ref": "#/components/schemas/access_timestamp"}, "duration": {"$ref": "#/components/schemas/access_duration-2"}, "expires_at": {"$ref": "#/components/schemas/access_timestamp"}, "id": {"allOf": [{"description": "The ID of the service token."}, {"$ref": "#/components/schemas/access_uuid"}]}, "last_seen_at": {"$ref": "#/components/schemas/access_timestamp"}, "name": {"$ref": "#/components/schemas/access_name-17"}, "updated_at": {"$ref": "#/components/schemas/access_timestamp"}}}
```
