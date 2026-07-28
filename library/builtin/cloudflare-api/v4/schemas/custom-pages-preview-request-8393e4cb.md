---
title: custom-pages_preview_request
page_id: schema-custom-pages-preview-request-8393e4cb
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# custom-pages_preview_request

```yaml
{"type": "object", "properties": {"act": {"description": "The preview action type. Required for request parsing but not used in token generation. Typically set to \"preview\".", "type": "string", "example": "preview"}, "target": {"description": "The target custom page type to preview (e.g. \"block:waf\"). Encoded as the \"endpoint\" claim in the resulting JWT.", "type": "string", "example": "block:waf"}, "url": {"description": "The URL of the custom page content to preview. Encoded as the \"zone\" claim in the resulting JWT.", "type": "string", "format": "uri", "example": "https://example.com/error.html"}}, "required": ["act", "target", "url"]}
```
