---
title: issuing_network_token_address
page_id: schema-issuing-network-token-address-00febea0
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_network_token_address

```yaml
{"title": "IssuingNetworkTokenAddress", "required": ["line1", "postal_code"], "type": "object", "properties": {"line1": {"maxLength": 5000, "type": "string", "description": "The street address of the cardholder tokenizing the card."}, "postal_code": {"maxLength": 5000, "type": "string", "description": "The postal code of the cardholder tokenizing the card."}}, "description": "", "x-expandableFields": []}
```
