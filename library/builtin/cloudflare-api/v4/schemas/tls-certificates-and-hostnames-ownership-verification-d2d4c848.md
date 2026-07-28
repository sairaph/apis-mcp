---
title: tls-certificates-and-hostnames_ownership_verification
page_id: schema-tls-certificates-and-hostnames-ownership-verification-d2d4c848
path: schemas
description: This is a record which can be placed to activate a hostname.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_ownership_verification

This is a record which can be placed to activate a hostname.

```yaml
{"description": "This is a record which can be placed to activate a hostname.", "type": "object", "oneOf": [{"properties": {"name": {"description": "DNS Name for record.", "type": "string", "example": "_cf-custom-hostname.app.example.com", "x-auditable": true}, "type": {"description": "DNS Record type.", "type": "string", "example": "txt", "enum": ["txt"], "x-auditable": true}, "value": {"description": "Content for the record.", "type": "string", "example": "5cc07c04-ea62-4a5a-95f0-419334a875a4"}}, "type": "object"}]}
```
