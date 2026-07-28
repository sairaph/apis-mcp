---
title: cache-rules_variants
page_id: schema-cache-rules-variants-3c3f0cda
path: schemas
description: 'Variant support enables caching variants of images with certain file extensions in addition to the original. This only applies when the origin server sends the ''Vary: Accept'' response header. If the origin server sends ''Vary: Accept'' but does not serve the variant requested, the response will not be cached. This will be indicated with BYPASS cache status in the response headers.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_variants

Variant support enables caching variants of images with certain file extensions in addition to the original. This only applies when the origin server sends the 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but does not serve the variant requested, the response will not be cached. This will be indicated with BYPASS cache status in the response headers.

```yaml
{"description": "Variant support enables caching variants of images with certain file extensions in addition to the original. This only applies when the origin server sends the 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but does not serve the variant requested, the response will not be cached. This will be indicated with BYPASS cache status in the response headers.", "type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "type": "string", "example": "variants", "enum": ["variants"], "x-auditable": true}}, "type": "object"}], "title": "Variants Caching"}
```
