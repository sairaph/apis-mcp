---
title: hyperdrive_hyperdrive-database
page_id: schema-hyperdrive-hyperdrive-database-b0a84ea5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# hyperdrive_hyperdrive-database

```yaml
{"type": "object", "properties": {"database": {"description": "Set the name of your origin database.", "type": "string", "example": "postgres", "x-auditable": true}, "password": {"description": "Set the password needed to access your origin database. The API never returns this write-only value.", "type": "string", "writeOnly": true, "x-sensitive": true}, "scheme": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-scheme"}, "user": {"description": "Set the user of your origin database.", "type": "string", "example": "postgres", "x-auditable": true}}, "title": "Connection Options"}
```
