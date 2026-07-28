---
title: api-shield_selector
page_id: schema-api-shield-selector-8244fc1b
path: schemas
description: |-
    Select operations covered by this rule.

    For details on selectors, see the [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_selector

Select operations covered by this rule.

For details on selectors, see the [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).

```yaml
{"description": "Select operations covered by this rule.\n\nFor details on selectors, see the [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).\n", "type": "object", "properties": {"exclude": {"description": "Ignore operations that were otherwise included by `include`.", "type": "array", "items": {"$ref": "#/components/schemas/api-shield_selector-exclude"}, "nullable": true}, "include": {"description": "Select all matching operations.", "type": "array", "items": {"$ref": "#/components/schemas/api-shield_selector-include"}, "nullable": true}}}
```
