---
title: TopProviderInfo
page_id: schema-topproviderinfo-421cd381
path: schemas
description: Information about the top provider for this model
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# TopProviderInfo

Information about the top provider for this model

```yaml
{"description": "Information about the top provider for this model", "example": {"context_length": 8192, "is_moderated": true, "max_completion_tokens": 4096}, "properties": {"context_length": {"description": "Context length from the top provider", "example": 8192, "type": ["integer", "null"], "nullable": true}, "is_moderated": {"description": "Whether the top provider moderates content", "example": true, "type": "boolean"}, "max_completion_tokens": {"description": "Maximum completion tokens from the top provider", "example": 4096, "type": ["integer", "null"], "nullable": true}}, "required": ["is_moderated"], "type": "object"}
```
