---
title: email-security_CreateUrlIgnorePattern
page_id: schema-email-security-createurlignorepattern-6bb6cb75
path: schemas
description: Creates a URL ignore pattern that exempts matching URLs from rewriting.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_CreateUrlIgnorePattern

Creates a URL ignore pattern that exempts matching URLs from rewriting.

```yaml
{"description": "Creates a URL ignore pattern that exempts matching URLs from rewriting.", "type": "object", "properties": {"comments": {"description": "Optional note describing the reason for the ignore pattern.", "type": "string", "example": "Trusted internal redirect service", "maxLength": 1024, "nullable": true, "x-auditable": true}, "pattern": {"description": "Regular expression identifying URLs to exempt from rewriting.", "type": "string", "example": "https://example\\.com/.*", "maxLength": 1024, "minLength": 1, "x-auditable": true}}, "additionalProperties": false, "required": ["pattern"]}
```
