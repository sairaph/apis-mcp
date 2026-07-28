---
title: ContentFilterBuiltinEntry
page_id: schema-contentfilterbuiltinentry-68459346
path: schemas
description: A builtin content filter entry. Builtin filters include PII detectors and the regex-based prompt injection detector.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ContentFilterBuiltinEntry

A builtin content filter entry. Builtin filters include PII detectors and the regex-based prompt injection detector.

```yaml
{"description": "A builtin content filter entry. Builtin filters include PII detectors and the regex-based prompt injection detector.", "example": {"action": "redact", "label": "[EMAIL]", "slug": "email"}, "properties": {"action": {"$ref": "#/components/schemas/ContentFilterBuiltinAction"}, "label": {"description": "Read-only, system-assigned redaction placeholder derived from the slug (e.g. \"[EMAIL]\", \"[PHONE]\"). Not settable by the caller.", "example": "[EMAIL]", "maxLength": 100, "type": "string"}, "scan_scope": {"$ref": "#/components/schemas/PromptInjectionScanScope"}, "slug": {"$ref": "#/components/schemas/ContentFilterBuiltinSlug"}}, "required": ["slug", "action"], "type": "object"}
```
