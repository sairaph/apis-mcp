---
title: email-security_UpdateUrlIgnorePattern
page_id: schema-email-security-updateurlignorepattern-93650683
path: schemas
description: Updates a URL rewrite ignore pattern; modifies only the provided fields.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_UpdateUrlIgnorePattern

Updates a URL rewrite ignore pattern; modifies only the provided fields.

```yaml
{"description": "Updates a URL rewrite ignore pattern; modifies only the provided fields.", "type": "object", "properties": {"comments": {"description": "Optional note describing the reason for the ignore pattern.", "type": "string", "example": "Trusted internal redirect service", "maxLength": 1024, "nullable": true, "x-auditable": true}, "pattern": {"description": "Regular expression identifying URLs to exempt from rewriting.", "type": "string", "example": "https://example\\.com/.*", "maxLength": 1024, "minLength": 1, "x-auditable": true}}, "additionalProperties": false, "minProperties": 1}
```
