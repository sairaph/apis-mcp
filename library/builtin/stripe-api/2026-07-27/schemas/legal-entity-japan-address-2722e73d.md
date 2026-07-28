---
title: legal_entity_japan_address
page_id: schema-legal-entity-japan-address-2722e73d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# legal_entity_japan_address

```yaml
{"title": "LegalEntityJapanAddress", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string", "description": "City/Ward.", "nullable": true}, "country": {"maxLength": 5000, "type": "string", "description": "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).", "nullable": true}, "line1": {"maxLength": 5000, "type": "string", "description": "Block/Building number.", "nullable": true}, "line2": {"maxLength": 5000, "type": "string", "description": "Building details.", "nullable": true}, "postal_code": {"maxLength": 5000, "type": "string", "description": "ZIP or postal code.", "nullable": true}, "state": {"maxLength": 5000, "type": "string", "description": "Prefecture.", "nullable": true}, "town": {"maxLength": 5000, "type": "string", "description": "Town/cho-me.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
