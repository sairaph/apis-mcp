---
title: card_issuing_account_terms_of_service
page_id: schema-card-issuing-account-terms-of-service-22e0437b
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# card_issuing_account_terms_of_service

```yaml
{"title": "CardIssuingAccountTermsOfService", "type": "object", "properties": {"date": {"type": "integer", "description": "The Unix timestamp marking when the account representative accepted the service agreement.", "nullable": true}, "ip": {"maxLength": 5000, "type": "string", "description": "The IP address from which the account representative accepted the service agreement.", "nullable": true}, "user_agent": {"maxLength": 5000, "type": "string", "description": "The user agent of the browser from which the account representative accepted the service agreement."}}, "description": "", "x-expandableFields": []}
```
