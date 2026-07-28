---
title: tls-certificates-and-hostnames_certificate_status
page_id: schema-tls-certificates-and-hostnames-certificate-status-c1eb33b2
path: schemas
description: Current status of certificate.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_certificate_status

Current status of certificate.

```yaml
{"description": "Current status of certificate.", "type": "string", "example": "active", "enum": ["initializing", "authorizing", "active", "expired", "issuing", "timing_out", "pending_deployment"], "x-auditable": true}
```
