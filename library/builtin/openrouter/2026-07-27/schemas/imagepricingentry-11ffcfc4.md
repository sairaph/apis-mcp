---
title: ImagePricingEntry
page_id: schema-imagepricingentry-11ffcfc4
path: schemas
description: One billable pricing line for an image provider.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImagePricingEntry

One billable pricing line for an image provider.

```yaml
{"description": "One billable pricing line for an image provider.", "example": {"billable": "output_image", "cost_usd": 0.05, "unit": "image"}, "properties": {"billable": {"enum": ["output_image", "input_image", "input_font", "input_reference", "input_text"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "cost_usd": {"format": "double", "type": "number"}, "unit": {"enum": ["image", "megapixel", "token"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "variant": {"type": "string"}}, "required": ["billable", "unit", "cost_usd"], "type": "object"}
```
