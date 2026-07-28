---
title: issuing_cardholder_verification
page_id: schema-issuing-cardholder-verification-46502968
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_cardholder_verification

```yaml
{"title": "IssuingCardholderVerification", "type": "object", "properties": {"document": {"description": "An identifying document, either a passport or local ID card.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_cardholder_id_document"}]}}, "description": "", "x-expandableFields": ["document"]}
```
