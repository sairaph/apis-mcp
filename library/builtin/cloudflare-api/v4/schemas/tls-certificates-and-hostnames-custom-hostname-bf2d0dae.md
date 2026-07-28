---
title: tls-certificates-and-hostnames_custom-hostname
page_id: schema-tls-certificates-and-hostnames-custom-hostname-bf2d0dae
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_custom-hostname

```yaml
{"type": "object", "properties": {"hostname": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname"}, "id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}, "ssl": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_ssl"}}, "allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_customhostname"}], "required": ["id", "hostname"]}
```
