---
title: registrar-api-sandbox_registration_contact_address
page_id: schema-registrar-api-sandbox-registration-contact-address-1c65e102
path: schemas
description: Physical mailing address for the registrant contact.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api-sandbox_registration_contact_address

Physical mailing address for the registrant contact.

```yaml
{"description": "Physical mailing address for the registrant contact.", "type": "object", "properties": {"city": {"description": "City or locality name.", "type": "string", "example": "Austin"}, "country_code": {"description": "Two-letter country code per ISO 3166-1 alpha-2 (e.g., `US`, `GB`, `CA`, `DE`).", "type": "string", "example": "US"}, "postal_code": {"description": "Postal or ZIP code.", "type": "string", "example": "78701"}, "state": {"description": "State, province, or region. Use the standard abbreviation where applicable (e.g., `TX` for Texas, `ON` for Ontario).", "type": "string", "example": "TX"}, "street": {"description": "Street address including building/suite number.", "type": "string", "example": "123 Main St"}}, "required": ["street", "city", "state", "postal_code", "country_code"]}
```
