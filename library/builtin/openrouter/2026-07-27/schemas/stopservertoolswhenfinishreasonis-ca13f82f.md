---
title: StopServerToolsWhenFinishReasonIs
page_id: schema-stopservertoolswhenfinishreasonis-ca13f82f
path: schemas
description: Stop when the upstream model emits this finish reason (e.g. `length`).
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# StopServerToolsWhenFinishReasonIs

Stop when the upstream model emits this finish reason (e.g. `length`).

```yaml
{"description": "Stop when the upstream model emits this finish reason (e.g. `length`).", "example": {"reason": "length", "type": "finish_reason_is"}, "properties": {"reason": {"minLength": 1, "type": "string"}, "type": {"enum": ["finish_reason_is"], "type": "string"}}, "required": ["type", "reason"], "type": "object"}
```
