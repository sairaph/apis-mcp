---
title: issuing_network_token_visa
page_id: schema-issuing-network-token-visa-610dc76b
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_network_token_visa

```yaml
{"title": "IssuingNetworkTokenVisa", "required": ["token_reference_id", "token_requestor_id"], "type": "object", "properties": {"card_reference_id": {"maxLength": 5000, "type": "string", "description": "A unique reference ID from Visa to represent the card account number.", "nullable": true}, "token_reference_id": {"maxLength": 5000, "type": "string", "description": "The network-unique identifier for the token."}, "token_requestor_id": {"maxLength": 5000, "type": "string", "description": "The ID of the entity requesting tokenization, specific to Visa."}, "token_risk_score": {"maxLength": 5000, "type": "string", "description": "Degree of risk associated with the token between `01` and `99`, with higher number indicating higher risk. A `00` value indicates the token was not scored by Visa."}}, "description": "", "x-expandableFields": []}
```
