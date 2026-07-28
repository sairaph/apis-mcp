---
title: payment_method_ideal
page_id: schema-payment-method-ideal-db50c871
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_ideal

```yaml
{"title": "payment_method_ideal", "type": "object", "properties": {"bank": {"type": "string", "description": "The customer's bank, if provided. Can be one of `abn_amro`, `adyen`, `asn_bank`, `bunq`, `buut`, `finom`, `handelsbanken`, `ing`, `knab`, `mollie`, `moneyou`, `n26`, `nn`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, `van_lanschot`, or `yoursafe`.", "nullable": true, "enum": ["abn_amro", "adyen", "asn_bank", "bunq", "buut", "finom", "handelsbanken", "ing", "knab", "mollie", "moneyou", "n26", "nn", "rabobank", "regiobank", "revolut", "sns_bank", "triodos_bank", "van_lanschot", "yoursafe"]}, "bic": {"type": "string", "description": "The Bank Identifier Code of the customer's bank, if the bank was provided.", "nullable": true, "enum": ["ABNANL2A", "ADYBNL2A", "ASNBNL21", "BITSNL2A", "BUNQNL2A", "BUUTNL2A", "FNOMNL22", "FVLBNL22", "HANDNL2A", "INGBNL2A", "KNABNL2H", "MLLENL2A", "MOYONL21", "NNBANL2G", "NTSBDEB1", "RABONL2U", "RBRBNL21", "REVOIE23", "REVOLT21", "SNSBNL2A", "TRIONL2U"]}}, "description": "", "x-expandableFields": []}
```
