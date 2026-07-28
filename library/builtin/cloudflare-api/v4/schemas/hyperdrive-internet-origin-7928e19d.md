---
title: hyperdrive_internet-origin
page_id: schema-hyperdrive-internet-origin-7928e19d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# hyperdrive_internet-origin

```yaml
{"type": "object", "properties": {"host": {"description": "Defines the host (hostname or IP) of your origin database.", "type": "string", "example": "database.example.com", "x-auditable": true}, "port": {"description": "Defines the port of your origin database. Defaults to 5432 for PostgreSQL or 3306 for MySQL if not specified.", "type": "integer", "example": 5432, "x-auditable": true}}, "required": ["host", "port"], "title": "Public Database"}
```
