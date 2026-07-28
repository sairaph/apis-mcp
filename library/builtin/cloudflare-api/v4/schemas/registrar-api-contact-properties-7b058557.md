---
title: registrar-api_contact_properties
page_id: schema-registrar-api-contact-properties-7b058557
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api_contact_properties

```yaml
{"type": "object", "properties": {"address": {"$ref": "#/components/schemas/registrar-api_address"}, "address2": {"$ref": "#/components/schemas/registrar-api_address2"}, "city": {"$ref": "#/components/schemas/registrar-api_city"}, "country": {"$ref": "#/components/schemas/registrar-api_country"}, "email": {"$ref": "#/components/schemas/registrar-api_email"}, "fax": {"$ref": "#/components/schemas/registrar-api_fax"}, "first_name": {"$ref": "#/components/schemas/registrar-api_first_name"}, "id": {"$ref": "#/components/schemas/registrar-api_contact_identifier"}, "last_name": {"$ref": "#/components/schemas/registrar-api_last_name"}, "organization": {"$ref": "#/components/schemas/registrar-api_organization"}, "phone": {"$ref": "#/components/schemas/registrar-api_telephone"}, "state": {"$ref": "#/components/schemas/registrar-api_state"}, "zip": {"$ref": "#/components/schemas/registrar-api_zipcode"}}, "required": ["first_name", "last_name", "address", "city", "state", "zip", "country", "phone", "organization"]}
```
