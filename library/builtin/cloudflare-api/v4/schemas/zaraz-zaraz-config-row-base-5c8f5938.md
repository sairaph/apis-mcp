---
title: zaraz_zaraz-config-row-base
page_id: schema-zaraz-zaraz-config-row-base-5c8f5938
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_zaraz-config-row-base

```yaml
{"type": "object", "properties": {"createdAt": {"description": "Date and time the configuration was created.", "type": "string", "format": "date-time", "x-auditable": true}, "id": {"description": "ID of the configuration.", "type": "integer", "x-auditable": true}, "updatedAt": {"description": "Date and time the configuration was last updated.", "type": "string", "format": "date-time", "x-auditable": true}, "userId": {"description": "Alpha-numeric ID of the account user who published the configuration.", "type": "string", "x-auditable": true}}, "required": ["id", "createdAt", "updatedAt", "userId"]}
```
