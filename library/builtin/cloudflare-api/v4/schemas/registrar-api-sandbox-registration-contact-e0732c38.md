---
title: registrar-api-sandbox_registration_contact
page_id: schema-registrar-api-sandbox-registration-contact-e0732c38
path: schemas
description: |-
    Contact data for the domain registration. This information
    is submitted to the domain registry and, depending on extension and
    privacy settings, may appear in public WHOIS records.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api-sandbox_registration_contact

Contact data for the domain registration. This information
is submitted to the domain registry and, depending on extension and
privacy settings, may appear in public WHOIS records.

```yaml
{"description": "Contact data for the domain registration. This information\nis submitted to the domain registry and, depending on extension and\nprivacy settings, may appear in public WHOIS records.\n", "type": "object", "properties": {"email": {"description": "Email address for the registrant. Used for domain-related\ncommunications from the registry, including ownership verification\nand renewal notices.\n", "type": "string", "format": "email", "example": "ada@example.com"}, "fax": {"description": "Fax number in E.164 format (e.g., `+1.5555555555`). Optional.\nMost registrations do not require a fax number.\n", "type": "string", "example": "+1.5555555555"}, "phone": {"description": "Phone number in E.164 format: `+{country_code}.{number}` with no\nspaces or dashes. Examples: `+1.5555555555` (US), `+44.2071234567`\n(UK), `+81.312345678` (Japan).\n", "type": "string", "example": "+1.5555555555"}, "postal_info": {"$ref": "#/components/schemas/registrar-api-sandbox_registration_contact_postal_info"}}, "required": ["phone", "email", "postal_info"]}
```
