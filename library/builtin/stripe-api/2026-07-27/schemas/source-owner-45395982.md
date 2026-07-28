---
title: source_owner
page_id: schema-source-owner-45395982
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# source_owner

```yaml
{"title": "SourceOwner", "type": "object", "properties": {"address": {"description": "Owner's address.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/address"}]}, "email": {"maxLength": 5000, "type": "string", "description": "Owner's email address.", "nullable": true}, "name": {"maxLength": 5000, "type": "string", "description": "Owner's full name.", "nullable": true}, "phone": {"maxLength": 5000, "type": "string", "description": "Owner's phone number (including extension).", "nullable": true}, "verified_address": {"description": "Verified owner's address. Verified values are verified or provided by the payment method directly (and if supported) at the time of authorization or settlement. They cannot be set or mutated.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/address"}]}, "verified_email": {"maxLength": 5000, "type": "string", "description": "Verified owner's email address. Verified values are verified or provided by the payment method directly (and if supported) at the time of authorization or settlement. They cannot be set or mutated.", "nullable": true}, "verified_name": {"maxLength": 5000, "type": "string", "description": "Verified owner's full name. Verified values are verified or provided by the payment method directly (and if supported) at the time of authorization or settlement. They cannot be set or mutated.", "nullable": true}, "verified_phone": {"maxLength": 5000, "type": "string", "description": "Verified owner's phone number (including extension). Verified values are verified or provided by the payment method directly (and if supported) at the time of authorization or settlement. They cannot be set or mutated.", "nullable": true}}, "description": "", "x-expandableFields": ["address", "verified_address"]}
```
