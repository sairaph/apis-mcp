---
title: email-security_PhishGuardReport
page_id: schema-email-security-phishguardreport-18826079
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_PhishGuardReport

```yaml
{"type": "object", "properties": {"content": {"type": "string"}, "created_at": {"type": "string", "format": "date-time", "nullable": true, "readOnly": true}, "disposition": {"$ref": "#/components/schemas/email-security_DispositionLabel"}, "fields": {"type": "object", "properties": {"from": {"type": "string", "nullable": true}, "occurred_at": {"type": "string", "format": "date-time", "readOnly": true}, "postfix_id": {"type": "string", "nullable": true}, "to": {"type": "array", "items": {"type": "string"}}, "ts": {"description": "Deprecated, use `occurred_at` instead.", "type": "string", "format": "date-time", "deprecated": true, "readOnly": true, "x-stainless-deprecation-message": "Use `occurred_at` instead."}}, "required": ["to"]}, "id": {"type": "integer"}, "priority": {"type": "string"}, "tags": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_PhishGuardReportTag"}, "nullable": true}, "title": {"type": "string"}, "ts": {"description": "Deprecated, use `created_at` instead.", "type": "string", "format": "date-time", "deprecated": true, "readOnly": true, "x-stainless-deprecation-message": "Use `created_at` instead."}, "updated_at": {"type": "string", "format": "date-time", "nullable": true, "readOnly": true}}, "required": ["id", "title", "content", "fields", "priority", "disposition"]}
```
