---
title: issuing_personalization_design_rejection_reasons
page_id: schema-issuing-personalization-design-rejection-reasons-19f6849d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_personalization_design_rejection_reasons

```yaml
{"title": "IssuingPersonalizationDesignRejectionReasons", "type": "object", "properties": {"card_logo": {"type": "array", "description": "The reason(s) the card logo was rejected.", "nullable": true, "items": {"type": "string", "enum": ["geographic_location", "inappropriate", "network_name", "non_binary_image", "non_fiat_currency", "other", "other_entity", "promotional_material"]}}, "carrier_text": {"type": "array", "description": "The reason(s) the carrier text was rejected.", "nullable": true, "items": {"type": "string", "enum": ["geographic_location", "inappropriate", "network_name", "non_fiat_currency", "other", "other_entity", "promotional_material"]}}}, "description": "", "x-expandableFields": []}
```
