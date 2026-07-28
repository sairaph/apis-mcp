---
title: rulesets_SkipProducts
page_id: schema-rulesets-skipproducts-5635c217
path: schemas
description: A list of legacy security products to skip the execution of.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SkipProducts

A list of legacy security products to skip the execution of.

```yaml
{"description": "A list of legacy security products to skip the execution of.", "type": "array", "items": {"description": "The name of a legacy security product to skip the execution of.", "enum": ["bic", "hot", "rateLimit", "securityLevel", "uaBlock", "waf", "zoneLockdown"], "example": "bic", "title": "Product", "type": "string"}, "minItems": 1, "title": "Products", "uniqueItems": true}
```
