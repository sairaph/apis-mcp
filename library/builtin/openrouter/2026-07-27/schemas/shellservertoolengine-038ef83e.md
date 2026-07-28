---
title: ShellServerToolEngine
page_id: schema-shellservertoolengine-038ef83e
path: schemas
description: Which shell engine to use. "openrouter" runs commands server-side in the OpenRouter sandbox. "auto" (default) keeps the provider's native hosted shell when available (OpenAI); on other providers the call is routed to the OpenRouter sandbox.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ShellServerToolEngine

Which shell engine to use. "openrouter" runs commands server-side in the OpenRouter sandbox. "auto" (default) keeps the provider's native hosted shell when available (OpenAI); on other providers the call is routed to the OpenRouter sandbox.

```yaml
{"description": "Which shell engine to use. \"openrouter\" runs commands server-side in the OpenRouter sandbox. \"auto\" (default) keeps the provider's native hosted shell when available (OpenAI); on other providers the call is routed to the OpenRouter sandbox.", "enum": ["auto", "openrouter"], "example": "openrouter", "type": "string", "x-speakeasy-unknown-values": "allow"}
```
