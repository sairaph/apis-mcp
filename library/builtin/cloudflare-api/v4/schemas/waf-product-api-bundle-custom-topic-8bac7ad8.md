---
title: waf-product-api-bundle_custom-topic
page_id: schema-waf-product-api-bundle-custom-topic-8bac7ad8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-product-api-bundle_custom-topic

```yaml
{"type": "object", "properties": {"label": {"description": "Unique label identifier. Must contain only lowercase letters (a–z), digits (0–9), and hyphens.", "type": "string", "example": "credit-cards", "maxLength": 20, "minLength": 2, "pattern": "^[a-z0-9-]+$"}, "topic": {"description": "Description of the topic category. Must contain only printable ASCII characters.", "type": "string", "example": "credit card numbers", "maxLength": 50, "minLength": 2}}, "required": ["label", "topic"]}
```
