---
title: cloudflare-pipelines_workers_pipelines_binding_source
page_id: schema-cloudflare-pipelines-workers-pipelines-binding-source-394f6104
path: schemas
description: '[DEPRECATED] Worker binding source configuration. Use the new streams API instead.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudflare-pipelines_workers_pipelines_binding_source

[DEPRECATED] Worker binding source configuration. Use the new streams API instead.

```yaml
{"description": "[DEPRECATED] Worker binding source configuration. Use the new streams API instead.", "type": "object", "properties": {"format": {"description": "Specifies the format of source data.", "type": "string", "enum": ["json"]}, "type": {"type": "string"}}, "deprecated": true, "required": ["type", "format"]}
```
