---
title: cloudflare-pipelines_PipelineGraph
page_id: schema-cloudflare-pipelines-pipelinegraph-f2fc6f03
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudflare-pipelines_PipelineGraph

```yaml
{"type": "object", "properties": {"edges": {"type": "array", "items": {"$ref": "#/components/schemas/cloudflare-pipelines_PipelineEdge"}}, "nodes": {"type": "array", "items": {"$ref": "#/components/schemas/cloudflare-pipelines_PipelineNode"}}}, "required": ["nodes", "edges"]}
```
