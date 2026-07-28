---
title: r2_r2_put_object_result
page_id: schema-r2-r2-put-object-result-ae5724d9
path: schemas
description: Result of a successful object upload.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_r2_put_object_result

Result of a successful object upload.

```yaml
{"description": "Result of a successful object upload.", "type": "object", "properties": {"etag": {"description": "The entity tag for the uploaded object.", "type": "string", "example": "d41d8cd98f00b204e9800998ecf8427e"}, "key": {"description": "The key (name) of the uploaded object.", "type": "string", "example": "path/to/my-object.txt"}, "size": {"description": "The size of the uploaded object in bytes (as a string).", "type": "string", "example": "1048576"}, "storage_class": {"$ref": "#/components/schemas/r2_storage_class"}, "uploaded": {"description": "The date and time the object was uploaded.", "type": "string", "format": "date-time", "example": "2024-01-15T10:30:00Z"}, "version": {"description": "The version UUID of the uploaded object.", "type": "string", "example": "3fd5b4a8-1234-5678-abcd-ef0123456789"}}}
```
