---
title: vuln_scanner_target-type
page_id: schema-vuln-scanner-target-type-ac6b750b
path: schemas
description: |-
    Identifies the Cloudflare asset to scan. Uses a `type` discriminator.
    Currently the service supports only `zone` targets.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_target-type

Identifies the Cloudflare asset to scan. Uses a `type` discriminator.
Currently the service supports only `zone` targets.

```yaml
{"description": "Identifies the Cloudflare asset to scan. Uses a `type` discriminator.\nCurrently the service supports only `zone` targets.\n", "discriminator": {"mapping": {"zone": "#/components/schemas/vuln_scanner_zone-target"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/vuln_scanner_zone-target"}], "x-auditable": true}
```
