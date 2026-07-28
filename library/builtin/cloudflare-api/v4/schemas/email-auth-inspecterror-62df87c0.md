---
title: email-auth_InspectError
page_id: schema-email-auth-inspecterror-62df87c0
path: schemas
description: An error encountered during SPF inspection
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-auth_InspectError

An error encountered during SPF inspection

```yaml
{"description": "An error encountered during SPF inspection", "type": "object", "properties": {"code": {"description": "Error code. Known values:\n- `lookup_failed` — DNS TXT lookup failed\n- `spf_not_found` — no SPF record found\n- `invalid_spf` — record does not start with `v=spf1`\n- `invalid_domain` — PSL validation failed\n- `loop_detected` — include/redirect cycle detected\n- `invalid_mechanism` — unrecognised or malformed mechanism\n- `resource_limit_exceeded` — internal resource protection limits exceeded (recursion depth or query budget)\n- `max_lookups` — RFC 7208 10-lookup limit exceeded\n", "type": "string", "example": "max_lookups"}, "details": {"description": "Additional error-specific details (optional).\n- For `invalid_domain` errors: the invalid domain string\n- For `invalid_mechanism` errors: the invalid mechanism text (e.g., \"invalidmech123\")\n- For `loop_detected` errors: the domain that caused the loop\n- For other error types: not present\n", "type": "string", "example": "invalid"}, "domain": {"description": "Domain where the error occurred", "type": "string", "example": "example.com"}, "message": {"description": "Human-readable error message", "type": "string", "example": "RFC 7208 10-lookup limit exceeded"}}, "required": ["code", "message", "domain"]}
```
