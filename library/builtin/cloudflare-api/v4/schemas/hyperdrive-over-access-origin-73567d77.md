---
title: hyperdrive_over-access-origin
page_id: schema-hyperdrive-over-access-origin-73567d77
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# hyperdrive_over-access-origin

```yaml
{"type": "object", "properties": {"access_client_id": {"description": "Defines the Client ID of the Access token to use when connecting to the origin database.", "type": "string", "example": "0123456789abcdef0123456789abcdef.access", "x-auditable": true}, "access_client_secret": {"description": "Defines the Client Secret of the Access Token to use when connecting to the origin database. The API never returns this write-only value.", "type": "string", "example": "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef", "writeOnly": true, "x-sensitive": true}, "host": {"description": "Defines the host (hostname or IP) of your origin database.", "type": "string", "example": "database.example.com", "x-auditable": true}}, "required": ["host", "access_client_id", "access_client_secret"], "title": "Access-protected Database behind Cloudflare Tunnel"}
```
