---
title: firewall_products
page_id: schema-firewall-products-f12c846e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_products

```yaml
{"type": "array", "items": {"description": "A list of products to bypass for a request when using the `bypass` action.", "enum": ["zoneLockdown", "uaBlock", "bic", "hot", "securityLevel", "rateLimit", "waf"], "example": "waf", "type": "string"}}
```
