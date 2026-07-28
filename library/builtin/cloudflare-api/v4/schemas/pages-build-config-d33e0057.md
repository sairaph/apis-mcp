---
title: pages_build_config
page_id: schema-pages-build-config-d33e0057
path: schemas
description: Configs for the project build process.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# pages_build_config

Configs for the project build process.

```yaml
{"description": "Configs for the project build process.", "type": "object", "properties": {"build_caching": {"description": "Enable build caching for the project.", "type": "boolean", "example": true, "nullable": true, "x-auditable": true}, "build_command": {"description": "Command used to build project.", "type": "string", "example": "npm run build", "nullable": true, "x-auditable": true}, "destination_dir": {"description": "Assets output directory of the build.", "type": "string", "example": "build", "nullable": true, "x-auditable": true}, "root_dir": {"description": "Directory to run the command.", "type": "string", "example": "/", "nullable": true, "x-auditable": true}, "web_analytics_tag": {"description": "The classifying tag for analytics.", "type": "string", "example": "cee1c73f6e4743d0b5e6bb1a0bcaabcc", "nullable": true, "x-auditable": true}, "web_analytics_token": {"description": "The auth token for analytics.", "type": "string", "example": "021e1057c18547eca7b79f2516f06o7x", "nullable": true, "x-sensitive": true}}, "required": ["web_analytics_tag", "web_analytics_token"]}
```
