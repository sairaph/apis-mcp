---
title: PromptInjectionScanScope
page_id: schema-promptinjectionscanscope-5445596d
path: schemas
description: Which message roles to scan for prompt injection. Only applies to the regex-prompt-injection builtin. Defaults to all_messages.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# PromptInjectionScanScope

Which message roles to scan for prompt injection. Only applies to the regex-prompt-injection builtin. Defaults to all_messages.

```yaml
{"description": "Which message roles to scan for prompt injection. Only applies to the regex-prompt-injection builtin. Defaults to all_messages.", "enum": ["user_only", "all_messages"], "example": "user_only", "type": "string", "x-speakeasy-unknown-values": "allow"}
```
