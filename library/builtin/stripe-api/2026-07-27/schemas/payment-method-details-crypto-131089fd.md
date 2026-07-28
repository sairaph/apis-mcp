---
title: payment_method_details_crypto
page_id: schema-payment-method-details-crypto-131089fd
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_crypto

```yaml
{"title": "payment_method_details_crypto", "type": "object", "properties": {"buyer_address": {"maxLength": 5000, "type": "string", "description": "The wallet address of the customer."}, "network": {"type": "string", "description": "The blockchain network that the transaction was sent on.", "enum": ["base", "ethereum", "polygon", "solana", "sui", "tempo"], "x-stripeBypassValidation": true}, "token_currency": {"type": "string", "description": "The token currency that the transaction was sent with.", "enum": ["phantom_cash", "usdc", "usdg", "usdp", "usdsui", "usdt"], "x-stripeBypassValidation": true}, "transaction_hash": {"maxLength": 5000, "type": "string", "description": "The blockchain transaction hash of the crypto payment."}}, "description": "", "x-expandableFields": []}
```
