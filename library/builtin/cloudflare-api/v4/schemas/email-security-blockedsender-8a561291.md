---
title: email-security_BlockedSender
page_id: schema-email-security-blockedsender-8a561291
path: schemas
description: A blocked sender pattern.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_BlockedSender

A blocked sender pattern.

```yaml
{"description": "A blocked sender pattern.", "type": "object", "properties": {"comments": {"type": "string", "example": "Block sender with email test@example.com", "maxLength": 1024, "nullable": true}, "created_at": {"allOf": [{"$ref": "#/components/schemas/email-security_timestamp"}], "readOnly": true}, "id": {"allOf": [{"$ref": "#/components/schemas/email-security_BlockedSenderId"}], "readOnly": true}, "is_regex": {"type": "boolean", "example": false, "x-auditable": true}, "last_modified": {"description": "Deprecated, use `modified_at` instead. End of life: November 1, 2026.", "allOf": [{"$ref": "#/components/schemas/email-security_timestamp"}], "deprecated": true, "readOnly": true, "x-stainless-deprecation-message": "Use `modified_at` instead."}, "modified_at": {"allOf": [{"$ref": "#/components/schemas/email-security_timestamp"}], "readOnly": true}, "pattern": {"description": "The pattern value to match. The format depends on `pattern_type`: a valid email address for EMAIL (e.g. `user@example.com`), a valid domain name for DOMAIN (e.g. `example.com`), or a plain IPv4 address or IPv4 CIDR block for IP (e.g. `1.2.3.4` or `1.2.3.0/24`); the API accepts only globally reachable IP addresses and rejects private, loopback, link-local, and unspecified addresses.", "type": "string", "example": "test@example.com", "maxLength": 1024, "minLength": 1, "x-auditable": true}, "pattern_type": {"$ref": "#/components/schemas/email-security_PatternType"}}}
```
