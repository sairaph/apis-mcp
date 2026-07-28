---
title: dns-settings_internal_dns_base
page_id: schema-dns-settings-internal-dns-base-57b62778
path: schemas
description: Settings for this internal zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-settings_internal_dns_base

Settings for this internal zone.

```yaml
{"description": "Settings for this internal zone.", "type": "object", "properties": {"reference_zone_id": {"description": "The ID of the zone to fallback to.", "type": "string", "example": {"description": "Identifier.", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "type": "string", "x-auditable": true}}}}
```
