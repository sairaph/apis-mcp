---
title: payment_pages_checkout_session_customer_details
page_id: schema-payment-pages-checkout-session-customer-details-ecc580fb
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_customer_details

```yaml
{"title": "PaymentPagesCheckoutSessionCustomerDetails", "type": "object", "properties": {"address": {"description": "The customer's address after a completed Checkout Session. Note: This property is populated only for sessions on or after March 30, 2022.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/address"}]}, "business_name": {"maxLength": 150, "type": "string", "description": "The customer's business name after a completed Checkout Session.", "nullable": true}, "email": {"maxLength": 5000, "type": "string", "description": "The email associated with the Customer, if one exists, on the Checkout Session after a completed Checkout Session or at time of session expiry.\nOtherwise, if the customer has consented to promotional content, this value is the most recent valid email provided by the customer on the Checkout form.", "nullable": true}, "individual_name": {"maxLength": 150, "type": "string", "description": "The customer's individual name after a completed Checkout Session.", "nullable": true}, "name": {"maxLength": 5000, "type": "string", "description": "The customer's name after a completed Checkout Session. Note: This property is populated only for sessions on or after March 30, 2022.", "nullable": true}, "phone": {"maxLength": 5000, "type": "string", "description": "The customer's phone number after a completed Checkout Session.", "nullable": true}, "tax_exempt": {"type": "string", "description": "The customer’s tax exempt status after a completed Checkout Session.", "nullable": true, "enum": ["exempt", "none", "reverse"]}, "tax_ids": {"type": "array", "description": "The customer’s tax IDs after a completed Checkout Session.", "nullable": true, "items": {"$ref": "#/components/schemas/payment_pages_checkout_session_tax_id"}}}, "description": "", "x-expandableFields": ["address", "tax_ids"]}
```
