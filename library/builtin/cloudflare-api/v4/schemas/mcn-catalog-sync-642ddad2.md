---
title: mcn_catalog_sync
page_id: schema-mcn-catalog-sync-642ddad2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_catalog_sync

```yaml
{"type": "object", "properties": {"description": {"type": "string"}, "destination_id": {"$ref": "#/components/schemas/mcn_catalog_sync_destination_id"}, "destination_type": {"$ref": "#/components/schemas/mcn_catalog_sync_destination_type"}, "errors": {"type": "object", "additionalProperties": {"$ref": "#/components/schemas/mcn_error"}}, "id": {"$ref": "#/components/schemas/mcn_catalog_sync_id"}, "includes_discoveries_until": {"type": "string"}, "last_attempted_update_at": {"type": "string"}, "last_successful_update_at": {"type": "string"}, "last_user_update_at": {"type": "string"}, "name": {"type": "string"}, "policy": {"type": "string"}, "update_mode": {"$ref": "#/components/schemas/mcn_catalog_sync_update_mode"}}, "required": ["id", "name", "description", "policy", "destination_type", "destination_id", "destination_url", "update_mode", "last_user_update_at"]}
```
