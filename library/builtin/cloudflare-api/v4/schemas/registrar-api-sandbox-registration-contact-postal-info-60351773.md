---
title: registrar-api-sandbox_registration_contact_postal_info
page_id: schema-registrar-api-sandbox-registration-contact-postal-info-60351773
path: schemas
description: |-
    Postal/mailing information for the contact. The `name` field is the
    complete contact name in one string. Some registries require a complete
    personal name, including a family or last name where applicable, but this
    API does not accept separate first-name and last-name fields for
    registration contacts.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api-sandbox_registration_contact_postal_info

Postal/mailing information for the contact. The `name` field is the
complete contact name in one string. Some registries require a complete
personal name, including a family or last name where applicable, but this
API does not accept separate first-name and last-name fields for
registration contacts.

```yaml
{"description": "Postal/mailing information for the contact. The `name` field is the\ncomplete contact name in one string. Some registries require a complete\npersonal name, including a family or last name where applicable, but this\nAPI does not accept separate first-name and last-name fields for\nregistration contacts.\n", "type": "object", "properties": {"address": {"$ref": "#/components/schemas/registrar-api-sandbox_registration_contact_address"}, "name": {"description": "Full legal name of the contact, including all required name components\nfor an individual or authorized representative. Some registries require\na complete personal name that includes a family or last name where\napplicable. Provide the complete name in this single field, for example\n`Ada Lovelace`; do not send separate first-name or last-name fields.\n", "type": "string", "example": "Ada Lovelace"}, "organization": {"description": "Organization or company name. Optional for individual registrants.", "type": "string", "example": "Example Inc"}}, "required": ["name", "address"]}
```
