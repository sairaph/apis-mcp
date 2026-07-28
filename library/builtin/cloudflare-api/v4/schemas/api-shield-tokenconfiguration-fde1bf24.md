---
title: api-shield_TokenConfiguration
page_id: schema-api-shield-tokenconfiguration-fde1bf24
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_TokenConfiguration

```yaml
{"type": "object", "properties": {"created_at": {"$ref": "#/components/schemas/api-shield_timestamp-2"}, "credentials": {"$ref": "#/components/schemas/api-shield_credentials"}, "description": {"$ref": "#/components/schemas/api-shield_description"}, "id": {"$ref": "#/components/schemas/api-shield_uuid-2"}, "last_updated": {"$ref": "#/components/schemas/api-shield_timestamp-2"}, "title": {"$ref": "#/components/schemas/api-shield_title"}, "token_sources": {"$ref": "#/components/schemas/api-shield_token_sources"}, "token_type": {"$ref": "#/components/schemas/api-shield_token_type"}}, "required": ["id", "title", "description", "token_sources", "token_type", "credentials", "created_at", "last_updated"]}
```
