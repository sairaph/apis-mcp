---
title: payments_primitives_payment_records_resource_address
page_id: schema-payments-primitives-payment-records-resource-address-50a89b4c
path: schemas
description: A representation of a physical address.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payments_primitives_payment_records_resource_address

A representation of a physical address.

```yaml
{"title": "PaymentsPrimitivesPaymentRecordsResourceAddress", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string", "description": "City, district, suburb, town, or village.", "nullable": true}, "country": {"maxLength": 5000, "type": "string", "description": "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).", "nullable": true}, "line1": {"maxLength": 5000, "type": "string", "description": "Address line 1, such as the street, PO Box, or company name.", "nullable": true}, "line2": {"maxLength": 5000, "type": "string", "description": "Address line 2, such as the apartment, suite, unit, or building.", "nullable": true}, "postal_code": {"maxLength": 5000, "type": "string", "description": "ZIP or postal code.", "nullable": true}, "state": {"maxLength": 5000, "type": "string", "description": "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).", "nullable": true}}, "description": "A representation of a physical address.", "x-expandableFields": []}
```
