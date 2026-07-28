---
title: cloudflare-pipelines_PipelineNode
page_id: schema-cloudflare-pipelines-pipelinenode-de2c52db
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudflare-pipelines_PipelineNode

```yaml
{"type": "object", "properties": {"description": {"type": "string"}, "node_id": {"type": "integer", "format": "int32", "minimum": 0}, "operator": {"type": "string"}, "parallelism": {"type": "integer", "format": "int32", "minimum": 0}}, "required": ["node_id", "operator", "description", "parallelism"]}
```
