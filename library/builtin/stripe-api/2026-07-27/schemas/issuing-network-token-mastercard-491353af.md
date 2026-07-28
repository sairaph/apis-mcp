---
title: issuing_network_token_mastercard
page_id: schema-issuing-network-token-mastercard-491353af
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_network_token_mastercard

```yaml
{"title": "IssuingNetworkTokenMastercard", "required": ["token_reference_id", "token_requestor_id"], "type": "object", "properties": {"card_reference_id": {"maxLength": 5000, "type": "string", "description": "A unique reference ID from MasterCard to represent the card account number."}, "token_reference_id": {"maxLength": 5000, "type": "string", "description": "The network-unique identifier for the token."}, "token_requestor_id": {"maxLength": 5000, "type": "string", "description": "The ID of the entity requesting tokenization, specific to MasterCard."}, "token_requestor_name": {"maxLength": 5000, "type": "string", "description": "The name of the entity requesting tokenization, if known. This is directly provided from MasterCard."}}, "description": "", "x-expandableFields": []}
```
