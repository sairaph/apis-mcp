---
title: email-security_TrustedDomain
page_id: schema-email-security-trusteddomain-150b122f
path: schemas
description: A trusted email domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_TrustedDomain

A trusted email domain.

```yaml
{"description": "A trusted email domain.", "type": "object", "properties": {"comments": {"type": "string", "example": "Trusted partner domain", "maxLength": 1024, "nullable": true}, "created_at": {"allOf": [{"$ref": "#/components/schemas/email-security_timestamp"}], "readOnly": true}, "id": {"allOf": [{"$ref": "#/components/schemas/email-security_TrustedDomainId"}], "readOnly": true}, "is_recent": {"description": "Select to prevent recently registered domains from triggering a Suspicious or Malicious disposition.", "type": "boolean", "example": true, "x-auditable": true}, "is_regex": {"type": "boolean", "example": false, "x-auditable": true}, "is_similarity": {"description": "Select for partner or other approved domains that have similar spelling to your connected domains. Prevents listed domains from triggering a Spoof disposition.", "type": "boolean", "example": false, "x-auditable": true}, "last_modified": {"description": "Deprecated, use `modified_at` instead. End of life: November 1, 2026.", "allOf": [{"$ref": "#/components/schemas/email-security_timestamp"}], "deprecated": true, "readOnly": true, "x-stainless-deprecation-message": "Use `modified_at` instead."}, "modified_at": {"allOf": [{"$ref": "#/components/schemas/email-security_timestamp"}], "readOnly": true}, "pattern": {"type": "string", "example": "example.com", "maxLength": 1024, "minLength": 1, "x-auditable": true}}}
```
