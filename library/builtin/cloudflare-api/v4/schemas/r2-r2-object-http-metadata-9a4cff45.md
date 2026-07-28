---
title: r2_r2_object_http_metadata
page_id: schema-r2-r2-object-http-metadata-9a4cff45
path: schemas
description: HTTP metadata associated with an R2 object.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_r2_object_http_metadata

HTTP metadata associated with an R2 object.

```yaml
{"description": "HTTP metadata associated with an R2 object.", "type": "object", "properties": {"cacheControl": {"description": "Specifies caching behavior for the object.", "type": "string", "example": "max-age=3600"}, "cacheExpiry": {"description": "The date and time at which the object's cache entry expires.", "type": "string", "format": "date-time", "example": "2024-12-31T23:59:59Z"}, "contentDisposition": {"description": "Specifies presentational information for the object.", "type": "string", "example": "attachment; filename=\"example.jpg\""}, "contentEncoding": {"description": "Specifies the content encoding applied to the object.", "type": "string", "example": "gzip"}, "contentLanguage": {"description": "The language of the object content.", "type": "string", "example": "en-US"}, "contentType": {"description": "The MIME type of the object.", "type": "string", "example": "image/jpeg"}}}
```
