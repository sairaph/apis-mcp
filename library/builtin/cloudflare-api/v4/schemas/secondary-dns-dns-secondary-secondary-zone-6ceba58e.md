---
title: secondary-dns_dns-secondary-secondary-zone
page_id: schema-secondary-dns-dns-secondary-secondary-zone-6ceba58e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# secondary-dns_dns-secondary-secondary-zone

```yaml
{"type": "object", "properties": {"auto_refresh_seconds": {"$ref": "#/components/schemas/secondary-dns_auto_refresh_seconds"}, "id": {"$ref": "#/components/schemas/secondary-dns_identifier"}, "name": {"$ref": "#/components/schemas/secondary-dns_name"}, "peers": {"$ref": "#/components/schemas/secondary-dns_peers"}}, "required": ["id", "name", "peers", "auto_refresh_seconds"]}
```
