---
title: issuing_network_token_network_data
page_id: schema-issuing-network-token-network-data-6b7655f7
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_network_token_network_data

```yaml
{"title": "IssuingNetworkTokenNetworkData", "required": ["type"], "type": "object", "properties": {"device": {"$ref": "#/components/schemas/issuing_network_token_device"}, "mastercard": {"$ref": "#/components/schemas/issuing_network_token_mastercard"}, "type": {"type": "string", "description": "The network that the token is associated with. An additional hash is included with a name matching this value, containing tokenization data specific to the card network.", "enum": ["mastercard", "visa"]}, "visa": {"$ref": "#/components/schemas/issuing_network_token_visa"}, "wallet_provider": {"$ref": "#/components/schemas/issuing_network_token_wallet_provider"}}, "description": "", "x-expandableFields": ["device", "mastercard", "visa", "wallet_provider"]}
```
