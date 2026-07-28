---
title: observatory_labeled_region
page_id: schema-observatory-labeled-region-a509f100
path: schemas
description: A test region with a label.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# observatory_labeled_region

A test region with a label.

```yaml
{"description": "A test region with a label.", "type": "object", "properties": {"label": {"type": "string", "example": "Iowa, USA", "x-auditable": true}, "value": {"$ref": "#/components/schemas/observatory_region"}}}
```
