---
title: mcn_create_catalog_sync_request
page_id: schema-mcn-create-catalog-sync-request-6ecb6ab5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_create_catalog_sync_request

```yaml
{"type": "object", "properties": {"description": {"type": "string"}, "destination_type": {"$ref": "#/components/schemas/mcn_catalog_sync_destination_type"}, "name": {"type": "string"}, "policy": {"type": "string"}, "update_mode": {"$ref": "#/components/schemas/mcn_catalog_sync_update_mode"}}, "required": ["name", "update_mode", "destination_type"]}
```
