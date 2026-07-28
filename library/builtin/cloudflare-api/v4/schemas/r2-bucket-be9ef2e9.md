---
title: r2_bucket
page_id: schema-r2-bucket-be9ef2e9
path: schemas
description: A single R2 bucket.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_bucket

A single R2 bucket.

```yaml
{"description": "A single R2 bucket.", "type": "object", "properties": {"creation_date": {"description": "Creation timestamp.", "type": "string"}, "jurisdiction": {"$ref": "#/components/schemas/r2_jurisdiction"}, "location": {"$ref": "#/components/schemas/r2_bucket_location"}, "name": {"$ref": "#/components/schemas/r2_bucket_name"}, "storage_class": {"$ref": "#/components/schemas/r2_storage_class"}}}
```
