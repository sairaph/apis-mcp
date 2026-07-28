---
title: rulesets_UrlNormalization
page_id: schema-rulesets-urlnormalization-03386b21
path: schemas
description: A URL Normalization object.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_UrlNormalization

A URL Normalization object.

```yaml
{"description": "A URL Normalization object.", "type": "object", "properties": {"scope": {"description": "The scope of the URL normalization.", "type": "string", "example": "incoming", "enum": ["incoming", "both", "none"], "title": "Scope"}, "type": {"description": "The type of URL normalization performed by Cloudflare.", "type": "string", "example": "cloudflare", "enum": ["cloudflare", "rfc3986"], "title": "Type"}}, "required": ["type", "scope"], "title": "URL Normalization"}
```
