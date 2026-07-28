---
title: cc_ListContainerInstances
page_id: schema-cc-listcontainerinstances-c287a6a3
path: schemas
description: Response body for listing container instances.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ListContainerInstances

Response body for listing container instances.

```yaml
{"description": "Response body for listing container instances.", "type": "object", "properties": {"instances": {"type": "array", "items": {"$ref": "#/components/schemas/cc_ContainerInstance"}}}, "required": ["instances"]}
```
