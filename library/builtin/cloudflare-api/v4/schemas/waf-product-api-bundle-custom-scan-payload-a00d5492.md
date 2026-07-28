---
title: waf-product-api-bundle_custom-scan-payload
page_id: schema-waf-product-api-bundle-custom-scan-payload-a00d5492
path: schemas
description: Defines the ruleset expression to use in matching content objects.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-product-api-bundle_custom-scan-payload

Defines the ruleset expression to use in matching content objects.

```yaml
{"description": "Defines the ruleset expression to use in matching content objects.", "type": "string", "example": "lookup_json_string(http.request.body.raw, \"file\")", "x-auditable": true}
```
