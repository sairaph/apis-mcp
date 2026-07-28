---
title: account_tos_acceptance
page_id: schema-account-tos-acceptance-d02e9d43
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_tos_acceptance

```yaml
{"title": "AccountTOSAcceptance", "type": "object", "properties": {"date": {"type": "integer", "description": "The Unix timestamp marking when the account representative accepted their service agreement", "format": "unix-time", "nullable": true}, "ip": {"maxLength": 5000, "type": "string", "description": "The IP address from which the account representative accepted their service agreement", "nullable": true}, "service_agreement": {"maxLength": 5000, "type": "string", "description": "The user's service agreement type"}, "user_agent": {"maxLength": 5000, "type": "string", "description": "The user agent of the browser from which the account representative accepted their service agreement", "nullable": true}}, "description": "", "x-expandableFields": []}
```
