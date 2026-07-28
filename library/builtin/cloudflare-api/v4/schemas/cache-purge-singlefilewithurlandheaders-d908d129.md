---
title: cache-purge_SingleFileWithUrlAndHeaders
page_id: schema-cache-purge-singlefilewithurlandheaders-d908d129
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-purge_SingleFileWithUrlAndHeaders

```yaml
{"type": "object", "properties": {"files": {"description": "For more information on purging files with URL and headers, please refer to [purge by single-file documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).", "type": "array", "items": {"properties": {"headers": {"type": "object", "example": {"CF-Device-Type": "desktop", "CF-IPCountry": "US"}, "additionalProperties": {"type": "string", "x-auditable": true}}, "url": {"type": "string", "example": "http://www.example.com/cat_picture.jpg", "x-auditable": true}}, "type": "object"}, "example": [{"headers": {"Accept-Language": "zh-CN", "CF-Device-Type": "desktop", "CF-IPCountry": "US"}, "url": "http://www.example.com/cat_picture.jpg"}, {"headers": {"Accept-Language": "en-US", "CF-Device-Type": "mobile", "CF-IPCountry": "EU"}, "url": "http://www.example.com/dog_picture.jpg"}]}}, "title": "Purge files with URL and headers"}
```
