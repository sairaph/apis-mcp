---
title: teams-devices_version_config
page_id: schema-teams-devices-version-config-01008a4f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_version_config

```yaml
{"type": "object", "properties": {"target_environment": {"description": "The target environment for the client version (e.g., windows, macos).", "type": "string", "example": "windows", "nullable": true, "x-auditable": true}, "version": {"description": "The specific client version to deploy.", "type": "string", "example": "2026.6.234.0", "x-auditable": true}}, "required": ["version", "target_environment"]}
```
