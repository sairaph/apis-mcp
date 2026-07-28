---
title: access_country_rule
page_id: schema-access-country-rule-3a066431
path: schemas
description: Matches a specific country
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_country_rule

Matches a specific country

```yaml
{"description": "Matches a specific country", "type": "object", "properties": {"geo": {"type": "object", "properties": {"country_code": {"description": "The country code that should be matched.", "type": "string", "example": "US"}}, "required": ["country_code"]}}, "required": ["geo"], "title": "Country"}
```
