---
title: workers_containers
page_id: schema-workers-containers-e79f088d
path: schemas
description: List of containers attached to a Worker. Containers can only be attached to Durable Object classes of this Worker script.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_containers

List of containers attached to a Worker. Containers can only be attached to Durable Object classes of this Worker script.

```yaml
{"description": "List of containers attached to a Worker. Containers can only be attached to Durable Object classes of this Worker script.", "type": "array", "items": {"$ref": "#/components/schemas/workers_container_item"}, "example": [{"class_name": "MyDurableObject"}], "x-stainless-collection-type": "set"}
```
