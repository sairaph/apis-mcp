---
title: issuing_authorization_merchant_data
page_id: schema-issuing-authorization-merchant-data-8a34e56a
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_authorization_merchant_data

```yaml
{"title": "IssuingAuthorizationMerchantData", "required": ["category", "category_code", "network_id"], "type": "object", "properties": {"category": {"maxLength": 5000, "type": "string", "description": "A categorization of the seller's type of business. See our [merchant categories guide](https://docs.stripe.com/issuing/merchant-categories) for a list of possible values."}, "category_code": {"maxLength": 5000, "type": "string", "description": "The merchant category code for the seller’s business"}, "city": {"maxLength": 5000, "type": "string", "description": "City where the seller is located", "nullable": true}, "country": {"maxLength": 5000, "type": "string", "description": "Country where the seller is located", "nullable": true}, "name": {"maxLength": 5000, "type": "string", "description": "Name of the seller", "nullable": true}, "network_id": {"maxLength": 5000, "type": "string", "description": "Identifier assigned to the seller by the card network. Different card networks may assign different network_id fields to the same merchant."}, "postal_code": {"maxLength": 5000, "type": "string", "description": "Postal code where the seller is located", "nullable": true}, "state": {"maxLength": 5000, "type": "string", "description": "State where the seller is located", "nullable": true}, "tax_id": {"maxLength": 5000, "type": "string", "description": "The seller's tax identification number. Currently populated for French merchants only.", "nullable": true}, "terminal_id": {"maxLength": 5000, "type": "string", "description": "An ID assigned by the seller to the location of the sale.", "nullable": true}, "url": {"maxLength": 5000, "type": "string", "description": "URL provided by the merchant on a 3DS request", "nullable": true}}, "description": "", "x-expandableFields": []}
```
