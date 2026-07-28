---
title: cache-purge_SingleFile
page_id: schema-cache-purge-singlefile-8bb8cd68
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-purge_SingleFile

```yaml
{"type": "object", "properties": {"files": {"description": "For more information on purging files, please refer to [purge by single-file documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["http://www.example.com/css/styles.css", "http://www.example.com/js/index.js"]}}, "title": "Purge files"}
```
