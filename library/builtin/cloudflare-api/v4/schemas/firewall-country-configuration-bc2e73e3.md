---
title: firewall_country_configuration
page_id: schema-firewall-country-configuration-bc2e73e3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_country_configuration

```yaml
{"type": "object", "properties": {"target": {"description": "The configuration target. You must set the target to `country` when specifying a country code in the rule.", "type": "string", "example": "country", "enum": ["country"]}, "value": {"description": "The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).", "type": "string", "example": "US", "x-auditable": true}}, "title": "A country configuration."}
```
