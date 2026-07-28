---
title: dns-custom-nameservers_zone_metadata
page_id: schema-dns-custom-nameservers-zone-metadata-46c6b0a2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-custom-nameservers_zone_metadata

```yaml
{"type": "object", "properties": {"enabled": {"description": "Whether zone uses account-level custom nameservers.", "type": "boolean", "example": true, "x-auditable": true}, "ns_set": {"description": "The number of the name server set to assign to the zone.", "type": "number", "example": 1, "default": 1, "maximum": 5, "minimum": 1, "x-auditable": true}}}
```
