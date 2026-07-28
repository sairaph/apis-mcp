---
title: pay-per-crawl_DaricConfig
page_id: schema-pay-per-crawl-daricconfig-a4243629
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# pay-per-crawl_DaricConfig

```yaml
{"type": "object", "properties": {"bot_overrides": {"type": "object", "additionalProperties": {"$ref": "#/components/schemas/pay-per-crawl_BotAccessMode"}}, "enabled": {"type": "boolean"}, "price_usd_microcents": {"description": "Price in microcents 1 USD = 100,000,000 microcents. Must be 0 or a multiple of 100,000 $0.001. Range: $0.001–$9,999.999.", "type": "integer", "maximum": 999999900000, "minimum": 0}}}
```
