---
title: BashServerToolConfig
page_id: schema-bashservertoolconfig-50ed9f87
path: schemas
description: Configuration for the openrouter:bash server tool
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BashServerToolConfig

Configuration for the openrouter:bash server tool

```yaml
{"description": "Configuration for the openrouter:bash server tool", "example": {"environment": {"type": "container_auto"}}, "properties": {"engine": {"$ref": "#/components/schemas/BashServerToolEngine"}, "environment": {"$ref": "#/components/schemas/BashServerToolEnvironment"}, "sleep_after_seconds": {"$ref": "#/components/schemas/SandboxSleepAfterSeconds"}}, "type": "object"}
```
