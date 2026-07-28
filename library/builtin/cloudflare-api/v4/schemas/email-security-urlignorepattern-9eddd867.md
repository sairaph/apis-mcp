---
title: email-security_UrlIgnorePattern
page_id: schema-email-security-urlignorepattern-9eddd867
path: schemas
description: A URL ignore pattern that exempts matching URLs from Email Security's URL rewriting.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_UrlIgnorePattern

A URL ignore pattern that exempts matching URLs from Email Security's URL rewriting.

```yaml
{"description": "A URL ignore pattern that exempts matching URLs from Email Security's URL rewriting.", "type": "object", "properties": {"comments": {"description": "Optional note describing the reason for the ignore pattern.", "type": "string", "example": "Trusted internal redirect service", "maxLength": 1024, "nullable": true, "x-auditable": true}, "created_at": {"type": "string", "format": "date-time", "readOnly": true, "x-auditable": true}, "id": {"allOf": [{"$ref": "#/components/schemas/email-security_UrlIgnorePatternId"}], "readOnly": true}, "last_modified": {"description": "Deprecated, use `modified_at` instead. End of life: November 1, 2026.", "type": "string", "format": "date-time", "deprecated": true, "readOnly": true, "x-auditable": true, "x-stainless-deprecation-message": "Use `modified_at` instead."}, "modified_at": {"type": "string", "format": "date-time", "readOnly": true, "x-auditable": true}, "pattern": {"description": "Regular expression identifying URLs to exempt from rewriting.", "type": "string", "example": "https://example\\.com/.*", "maxLength": 1024, "minLength": 1, "x-auditable": true}}, "required": ["id", "created_at", "pattern"]}
```
