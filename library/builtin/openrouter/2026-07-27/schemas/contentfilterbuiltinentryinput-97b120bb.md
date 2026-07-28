---
title: ContentFilterBuiltinEntryInput
page_id: schema-contentfilterbuiltinentryinput-97b120bb
path: schemas
description: A builtin content filter entry for create/update requests. Labels are system-assigned and cannot be set by the caller.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ContentFilterBuiltinEntryInput

A builtin content filter entry for create/update requests. Labels are system-assigned and cannot be set by the caller.

```yaml
{"description": "A builtin content filter entry for create/update requests. Labels are system-assigned and cannot be set by the caller.", "example": {"action": "redact", "slug": "email"}, "properties": {"action": {"$ref": "#/components/schemas/ContentFilterBuiltinAction"}, "label": {"deprecated": true, "description": "Deprecated: labels are system-assigned and cannot be set by the caller. Accepted for backward compatibility but silently ignored.", "maxLength": 100, "type": "string"}, "scan_scope": {"$ref": "#/components/schemas/PromptInjectionScanScope"}, "slug": {"$ref": "#/components/schemas/ContentFilterBuiltinSlug"}}, "required": ["slug", "action"], "type": "object"}
```
