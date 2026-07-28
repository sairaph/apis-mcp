---
title: secondary-dns_peer
page_id: schema-secondary-dns-peer-1266fefd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# secondary-dns_peer

```yaml
{"type": "object", "properties": {"id": {"$ref": "#/components/schemas/secondary-dns_identifier-3"}, "ip": {"$ref": "#/components/schemas/secondary-dns_ip"}, "ixfr_enable": {"$ref": "#/components/schemas/secondary-dns_ixfr_enable"}, "name": {"$ref": "#/components/schemas/secondary-dns_name-3"}, "port": {"$ref": "#/components/schemas/secondary-dns_port"}, "tsig_id": {"$ref": "#/components/schemas/secondary-dns_tsig_id"}}, "required": ["id", "name"]}
```
