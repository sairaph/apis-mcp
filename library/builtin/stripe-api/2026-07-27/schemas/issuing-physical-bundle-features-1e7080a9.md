---
title: issuing_physical_bundle_features
page_id: schema-issuing-physical-bundle-features-1e7080a9
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_physical_bundle_features

```yaml
{"title": "IssuingPhysicalBundleFeatures", "required": ["card_logo", "carrier_text", "second_line"], "type": "object", "properties": {"card_logo": {"type": "string", "description": "The policy for how to use card logo images in a card design with this physical bundle.", "enum": ["optional", "required", "unsupported"]}, "carrier_text": {"type": "string", "description": "The policy for how to use carrier letter text in a card design with this physical bundle.", "enum": ["optional", "required", "unsupported"]}, "second_line": {"type": "string", "description": "The policy for how to use a second line on a card with this physical bundle.", "enum": ["optional", "required", "unsupported"]}}, "description": "", "x-expandableFields": []}
```
