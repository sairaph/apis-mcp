---
title: ShellServerToolConfig
page_id: schema-shellservertoolconfig-4f261aad
path: schemas
description: Configuration for the openrouter:shell server tool
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ShellServerToolConfig

Configuration for the openrouter:shell server tool

```yaml
{"description": "Configuration for the openrouter:shell server tool", "example": {"engine": "openrouter", "environment": {"type": "container_auto"}}, "properties": {"engine": {"$ref": "#/components/schemas/ShellServerToolEngine"}, "environment": {"$ref": "#/components/schemas/ShellServerToolEnvironment"}, "sleep_after_seconds": {"$ref": "#/components/schemas/SandboxSleepAfterSeconds"}}, "type": "object"}
```
