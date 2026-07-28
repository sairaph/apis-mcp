---
title: abuse-reports_RegistrarWhoIsFields
page_id: schema-abuse-reports-registrarwhoisfields-e0fdc58c
path: schemas
description: RDP-mandated fields for registrar WHOIS data disclosure requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_RegistrarWhoIsFields

RDP-mandated fields for registrar WHOIS data disclosure requests.

```yaml
{"description": "RDP-mandated fields for registrar WHOIS data disclosure requests.", "type": "object", "properties": {"reg_who_authorization_statement": {"description": "Optional authorization statement or power of attorney per RDP 10.2.1.3.", "type": "string", "maxLength": 5000}, "reg_who_good_faith_affirmation": {"description": "Affirmation that the request is made in good faith per RDP 10.2.4. Must be true.", "type": "boolean"}, "reg_who_lawful_processing_agreement": {"description": "Agreement to process data lawfully per RDP 10.2.5. Must be true.", "type": "boolean"}, "reg_who_legal_basis": {"description": "Legal rights and rationale for the request per RDP 10.2.3. Required for all WHOIS requests.", "type": "string", "maxLength": 5000}, "reg_who_request_type": {"description": "The type of WHOIS data request per RDP procedure.", "type": "string", "enum": ["disclosure", "invalid_whois"]}, "reg_who_requested_data_elements": {"description": "The specific WHOIS data elements being requested per RDP 10.2.2. Required for all WHOIS requests.", "type": "array", "items": {"enum": ["registrant_name", "registrant_organization", "registrant_email", "registrant_phone", "registrant_address", "registrant_address_country", "registrant_address_postal_code", "admin_name", "admin_organization", "admin_email", "admin_phone", "admin_address", "tech_name", "tech_organization", "tech_email", "tech_phone", "tech_address"], "type": "string"}, "maxItems": 17, "minItems": 1}, "reg_who_requestor_type": {"description": "The nature of the requestor per RDP 10.2.1.2.", "type": "string", "enum": ["government", "corporation", "individual"]}}, "required": ["reg_who_request_type", "reg_who_requested_data_elements", "reg_who_legal_basis", "reg_who_good_faith_affirmation", "reg_who_lawful_processing_agreement"]}
```
