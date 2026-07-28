---
title: legal_entity_representative_declaration
page_id: schema-legal-entity-representative-declaration-03edbda5
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# legal_entity_representative_declaration

```yaml
{"title": "LegalEntityRepresentativeDeclaration", "type": "object", "properties": {"date": {"type": "integer", "description": "The Unix timestamp marking when the representative declaration attestation was made.", "format": "unix-time", "nullable": true}, "ip": {"maxLength": 5000, "type": "string", "description": "The IP address from which the representative declaration attestation was made.", "nullable": true}, "user_agent": {"maxLength": 5000, "type": "string", "description": "The user-agent string from the browser where the representative declaration attestation was made.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
