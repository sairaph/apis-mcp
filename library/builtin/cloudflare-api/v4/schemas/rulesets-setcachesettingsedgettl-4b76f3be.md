---
title: rulesets_SetCacheSettingsEdgeTTL
page_id: schema-rulesets-setcachesettingsedgettl-4b76f3be
path: schemas
description: How long the Cloudflare edge network should cache the response.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsEdgeTTL

How long the Cloudflare edge network should cache the response.

```yaml
{"description": "How long the Cloudflare edge network should cache the response.", "type": "object", "properties": {"default": {"description": "The edge TTL (in seconds) if you choose the \"override_origin\" mode.", "type": "integer", "example": 60, "minimum": 0, "title": "Default TTL"}, "mode": {"description": "The edge TTL mode.", "type": "string", "example": "override_origin", "enum": ["respect_origin", "bypass_by_default", "override_origin"], "title": "TTL Mode"}, "status_code_ttl": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsStatusCodeTTL"}}, "required": ["mode"], "title": "Edge TTL"}
```
