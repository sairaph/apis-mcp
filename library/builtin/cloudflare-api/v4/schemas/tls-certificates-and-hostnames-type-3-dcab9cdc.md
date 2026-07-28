---
title: tls-certificates-and-hostnames_type-3
page_id: schema-tls-certificates-and-hostnames-type-3-dcab9cdc
path: schemas
description: The type of the certificate, indicating how it was created and who manages it.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_type-3

The type of the certificate, indicating how it was created and who manages it.

```yaml
{"description": "The type of the certificate, indicating how it was created and who manages it.", "type": "string", "example": "custom", "enum": ["custom", "gateway_managed", "access_managed"], "readOnly": true, "x-auditable": true}
```
