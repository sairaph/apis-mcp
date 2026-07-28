---
title: mandate_payment_method_details
page_id: schema-mandate-payment-method-details-979e4af4
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# mandate_payment_method_details

```yaml
{"title": "mandate_payment_method_details", "required": ["type"], "type": "object", "properties": {"acss_debit": {"$ref": "#/components/schemas/mandate_acss_debit"}, "amazon_pay": {"$ref": "#/components/schemas/mandate_amazon_pay"}, "au_becs_debit": {"$ref": "#/components/schemas/mandate_au_becs_debit"}, "bacs_debit": {"$ref": "#/components/schemas/mandate_bacs_debit"}, "card": {"$ref": "#/components/schemas/card_mandate_payment_method_details"}, "cashapp": {"$ref": "#/components/schemas/mandate_cashapp"}, "kakao_pay": {"$ref": "#/components/schemas/mandate_kakao_pay"}, "klarna": {"$ref": "#/components/schemas/mandate_klarna"}, "kr_card": {"$ref": "#/components/schemas/mandate_kr_card"}, "link": {"$ref": "#/components/schemas/mandate_link"}, "naver_pay": {"$ref": "#/components/schemas/mandate_naver_pay"}, "nz_bank_account": {"$ref": "#/components/schemas/mandate_nz_bank_account"}, "paypal": {"$ref": "#/components/schemas/mandate_paypal"}, "payto": {"$ref": "#/components/schemas/mandate_payto"}, "pix": {"$ref": "#/components/schemas/mandate_pix"}, "revolut_pay": {"$ref": "#/components/schemas/mandate_revolut_pay"}, "sepa_debit": {"$ref": "#/components/schemas/mandate_sepa_debit"}, "twint": {"$ref": "#/components/schemas/mandate_twint"}, "type": {"maxLength": 5000, "type": "string", "description": "This mandate corresponds with a specific payment method type. The `payment_method_details` includes an additional hash with the same name and contains mandate information that's specific to that payment method."}, "upi": {"$ref": "#/components/schemas/mandate_upi"}, "us_bank_account": {"$ref": "#/components/schemas/mandate_us_bank_account"}}, "description": "", "x-expandableFields": ["acss_debit", "amazon_pay", "au_becs_debit", "bacs_debit", "card", "cashapp", "kakao_pay", "klarna", "kr_card", "link", "naver_pay", "nz_bank_account", "paypal", "payto", "pix", "revolut_pay", "sepa_debit", "twint", "upi", "us_bank_account"]}
```
