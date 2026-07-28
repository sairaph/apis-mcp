---
title: spectrum-config_origin_dns
page_id: schema-spectrum-config-origin-dns-65157400
path: schemas
description: The name and type of DNS record for the Spectrum application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# spectrum-config_origin_dns

The name and type of DNS record for the Spectrum application.

```yaml
{"description": "The name and type of DNS record for the Spectrum application.", "type": "object", "properties": {"name": {"$ref": "#/components/schemas/spectrum-config_origin_dns_name"}, "ttl": {"$ref": "#/components/schemas/spectrum-config_dns_ttl"}, "type": {"$ref": "#/components/schemas/spectrum-config_origin_dns_type"}}}
```
