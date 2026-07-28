---
title: cloudflare-pipelines_PipelineEdge
page_id: schema-cloudflare-pipelines-pipelineedge-fca042a0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudflare-pipelines_PipelineEdge

```yaml
{"type": "object", "properties": {"dest_id": {"type": "integer", "format": "int32", "minimum": 0}, "edge_type": {"type": "string"}, "key_type": {"type": "string"}, "src_id": {"type": "integer", "format": "int32", "minimum": 0}, "value_type": {"type": "string"}}, "required": ["src_id", "dest_id", "key_type", "value_type", "edge_type"]}
```
