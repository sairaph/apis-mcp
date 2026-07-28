---
title: BashServerToolEngine
page_id: schema-bashservertoolengine-e5c36d59
path: schemas
description: Which bash engine to use. "openrouter" runs commands server-side in the OpenRouter sandbox. "auto" (default) and "native" use native passthrough, returning the tool call to your application to run client-side; OpenRouter does not execute the commands.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BashServerToolEngine

Which bash engine to use. "openrouter" runs commands server-side in the OpenRouter sandbox. "auto" (default) and "native" use native passthrough, returning the tool call to your application to run client-side; OpenRouter does not execute the commands.

```yaml
{"description": "Which bash engine to use. \"openrouter\" runs commands server-side in the OpenRouter sandbox. \"auto\" (default) and \"native\" use native passthrough, returning the tool call to your application to run client-side; OpenRouter does not execute the commands.", "enum": ["auto", "native", "openrouter"], "example": "auto", "type": "string", "x-speakeasy-unknown-values": "allow"}
```
