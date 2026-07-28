---
title: person_additional_tos_acceptance
page_id: schema-person-additional-tos-acceptance-971bbb70
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# person_additional_tos_acceptance

```yaml
{"title": "PersonAdditionalTOSAcceptance", "type": "object", "properties": {"date": {"type": "integer", "description": "The Unix timestamp marking when the legal guardian accepted the service agreement.", "format": "unix-time", "nullable": true}, "ip": {"maxLength": 5000, "type": "string", "description": "The IP address from which the legal guardian accepted the service agreement.", "nullable": true}, "user_agent": {"maxLength": 5000, "type": "string", "description": "The user agent of the browser from which the legal guardian accepted the service agreement.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
