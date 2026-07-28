---
title: r2_r2_object
page_id: schema-r2-r2-object-fe9055a2
path: schemas
description: Metadata for an R2 object.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_r2_object

Metadata for an R2 object.

```yaml
{"description": "Metadata for an R2 object.", "type": "object", "properties": {"custom_metadata": {"description": "Custom metadata key-value pairs associated with the object.", "type": "object", "example": {}, "additionalProperties": {"type": "string"}}, "etag": {"description": "The entity tag for the object. In JSON list/get responses this is the raw\nhex digest (without surrounding quotes). The HTTP `ETag` response header on\nGet Object follows RFC 7232 and IS wrapped in surrounding double-quotes.\n", "type": "string", "example": "d41d8cd98f00b204e9800998ecf8427e"}, "http_metadata": {"$ref": "#/components/schemas/r2_r2_object_http_metadata"}, "key": {"description": "The object key (name).", "type": "string", "example": "path/to/my-object.txt"}, "last_modified": {"description": "The date and time the object was last modified.", "type": "string", "format": "date-time", "example": "2024-01-15T10:30:00Z"}, "size": {"description": "The size of the object in bytes.", "type": "integer", "example": 1048576}, "ssec": {"description": "Whether the object is encrypted with a customer-supplied encryption key.", "type": "boolean", "example": false}, "storage_class": {"$ref": "#/components/schemas/r2_storage_class"}}}
```
