---
title: tls-certificates-and-hostnames_verification_errors
page_id: schema-tls-certificates-and-hostnames-verification-errors-84fedf60
path: schemas
description: These are errors that were encountered while trying to activate a hostname.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_verification_errors

These are errors that were encountered while trying to activate a hostname.

```yaml
{"description": "These are errors that were encountered while trying to activate a hostname.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["None of the A or AAAA records are owned by this account and the pre-generated ownership verification token was not found."]}
```
