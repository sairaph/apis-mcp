---
title: secondary-dns_tsig
page_id: schema-secondary-dns-tsig-c84091e2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# secondary-dns_tsig

```yaml
{"type": "object", "properties": {"algo": {"$ref": "#/components/schemas/secondary-dns_algo"}, "id": {"$ref": "#/components/schemas/secondary-dns_identifier-2"}, "name": {"$ref": "#/components/schemas/secondary-dns_name-2"}, "secret": {"$ref": "#/components/schemas/secondary-dns_secret"}}, "required": ["id", "name", "secret", "algo"]}
```
