---
title: cache-rules_variants_value
page_id: schema-cache-rules-variants-value-1ee8bc53
path: schemas
description: Value of the zone setting.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_variants_value

Value of the zone setting.

```yaml
{"description": "Value of the zone setting.", "type": "object", "properties": {"avif": {"description": "List of strings with the MIME types of all the variants that should be served for avif.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["image/webp", "image/jpeg"], "uniqueItems": true}, "bmp": {"description": "List of strings with the MIME types of all the variants that should be served for bmp.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["image/webp", "image/jpeg"], "uniqueItems": true}, "gif": {"description": "List of strings with the MIME types of all the variants that should be served for gif.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["image/webp", "image/jpeg"], "uniqueItems": true}, "jp2": {"description": "List of strings with the MIME types of all the variants that should be served for jp2.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["image/webp", "image/avif"], "uniqueItems": true}, "jpeg": {"description": "List of strings with the MIME types of all the variants that should be served for jpeg.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["image/webp", "image/avif"], "uniqueItems": true}, "jpg": {"description": "List of strings with the MIME types of all the variants that should be served for jpg.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["image/webp", "image/avif"], "uniqueItems": true}, "jpg2": {"description": "List of strings with the MIME types of all the variants that should be served for jpg2.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["image/webp", "image/avif"], "uniqueItems": true}, "png": {"description": "List of strings with the MIME types of all the variants that should be served for png.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["image/webp", "image/avif"], "uniqueItems": true}, "tif": {"description": "List of strings with the MIME types of all the variants that should be served for tif.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["image/webp", "image/avif"], "uniqueItems": true}, "tiff": {"description": "List of strings with the MIME types of all the variants that should be served for tiff.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["image/webp", "image/avif"], "uniqueItems": true}, "webp": {"description": "List of strings with the MIME types of all the variants that should be served for webp.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["image/jpeg", "image/avif"], "uniqueItems": true}}}
```
