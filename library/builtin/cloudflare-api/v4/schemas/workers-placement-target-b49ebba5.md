---
title: workers_placement_target
page_id: schema-workers-placement-target-b49ebba5
path: schemas
description: A target to run your Worker near.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_placement_target

A target to run your Worker near.

```yaml
{"description": "A target to run your Worker near.", "type": "object", "oneOf": [{"additionalProperties": false, "properties": {"region": {"description": "Cloud region in format 'provider:region'.", "type": "string", "example": "aws:us-east-1"}}, "required": ["region"], "type": "object"}, {"additionalProperties": false, "properties": {"hostname": {"description": "HTTP hostname for targeted placement.", "type": "string", "example": "api.example.com"}}, "required": ["hostname"], "type": "object"}, {"additionalProperties": false, "properties": {"host": {"description": "TCP host:port for targeted placement.", "type": "string", "example": "db.example.com:5432"}}, "required": ["host"], "type": "object"}]}
```
