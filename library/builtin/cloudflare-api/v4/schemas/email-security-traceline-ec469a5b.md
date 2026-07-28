---
title: email-security_TraceLine
page_id: schema-email-security-traceline-ec469a5b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_TraceLine

```yaml
{"type": "object", "properties": {"lineno": {"description": "Line number in the trace log.", "type": "integer"}, "logged_at": {"type": "string", "format": "date-time", "nullable": true, "readOnly": true}, "message": {"type": "string"}, "ts": {"description": "Deprecated, use `logged_at` instead. End of life: November 1, 2026.", "type": "string", "deprecated": true, "readOnly": true, "x-stainless-deprecation-message": "Use `logged_at` instead."}}}
```
