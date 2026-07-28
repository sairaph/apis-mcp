---
title: ChatDebugOptions
page_id: schema-chatdebugoptions-5f173082
path: schemas
description: Debug options for inspecting request transformations (streaming only)
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatDebugOptions

Debug options for inspecting request transformations (streaming only)

```yaml
{"description": "Debug options for inspecting request transformations (streaming only)", "example": {"echo_upstream_body": true}, "properties": {"echo_upstream_body": {"description": "If true, includes the transformed upstream request body in a debug chunk at the start of the stream. Only works with streaming mode.", "example": true, "type": "boolean"}}, "type": "object"}
```
