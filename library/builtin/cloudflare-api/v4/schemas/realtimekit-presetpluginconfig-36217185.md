---
title: realtimekit_PresetPluginConfig
page_id: schema-realtimekit-presetpluginconfig-36217185
path: schemas
description: Plugin configuration keyed by plugin UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_PresetPluginConfig

Plugin configuration keyed by plugin UUID.

```yaml
{"description": "Plugin configuration keyed by plugin UUID.", "type": "object", "additionalProperties": {"additionalProperties": true, "properties": {"access_control": {"type": "string", "enum": ["FULL_ACCESS", "VIEW_ONLY"]}, "handles_view_only": {"type": "boolean"}}, "type": "object"}}
```
