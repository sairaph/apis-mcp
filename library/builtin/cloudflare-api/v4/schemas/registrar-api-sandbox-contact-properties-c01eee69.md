---
title: registrar-api-sandbox_contact_properties
page_id: schema-registrar-api-sandbox-contact-properties-c01eee69
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api-sandbox_contact_properties

```yaml
{"type": "object", "properties": {"address": {"$ref": "#/components/schemas/registrar-api-sandbox_address"}, "address2": {"$ref": "#/components/schemas/registrar-api-sandbox_address2"}, "city": {"$ref": "#/components/schemas/registrar-api-sandbox_city"}, "country": {"$ref": "#/components/schemas/registrar-api-sandbox_country"}, "email": {"$ref": "#/components/schemas/registrar-api-sandbox_email"}, "fax": {"$ref": "#/components/schemas/registrar-api-sandbox_fax"}, "first_name": {"$ref": "#/components/schemas/registrar-api-sandbox_first_name"}, "id": {"$ref": "#/components/schemas/registrar-api-sandbox_contact_identifier"}, "last_name": {"$ref": "#/components/schemas/registrar-api-sandbox_last_name"}, "organization": {"$ref": "#/components/schemas/registrar-api-sandbox_organization"}, "phone": {"$ref": "#/components/schemas/registrar-api-sandbox_telephone"}, "state": {"$ref": "#/components/schemas/registrar-api-sandbox_state"}, "zip": {"$ref": "#/components/schemas/registrar-api-sandbox_zipcode"}}, "required": ["first_name", "last_name", "address", "city", "state", "zip", "country", "phone", "organization"]}
```
