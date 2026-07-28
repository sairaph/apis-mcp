---
title: issuing_cardholder_user_terms_acceptance
page_id: schema-issuing-cardholder-user-terms-acceptance-07c1db8d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_cardholder_user_terms_acceptance

```yaml
{"title": "IssuingCardholderUserTermsAcceptance", "type": "object", "properties": {"date": {"type": "integer", "description": "The Unix timestamp marking when the cardholder accepted the Authorized User Terms.", "format": "unix-time", "nullable": true}, "ip": {"maxLength": 5000, "type": "string", "description": "The IP address from which the cardholder accepted the Authorized User Terms.", "nullable": true}, "user_agent": {"maxLength": 5000, "type": "string", "description": "The user agent of the browser from which the cardholder accepted the Authorized User Terms.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
