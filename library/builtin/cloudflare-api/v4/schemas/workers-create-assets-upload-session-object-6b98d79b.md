---
title: workers_create-assets-upload-session-object
page_id: schema-workers-create-assets-upload-session-object-6b98d79b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_create-assets-upload-session-object

```yaml
{"type": "object", "properties": {"manifest": {"description": "A manifest ([path]: {hash, size}) map of files to upload. As an example, `/blog/hello-world.html` would be a valid path key.", "type": "object", "additionalProperties": {"$ref": "#/components/schemas/workers_manifest-value"}}}, "required": ["manifest"]}
```
