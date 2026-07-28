---
title: dispute_visa_compelling_evidence3_prior_undisputed_transaction
page_id: schema-dispute-visa-compelling-evidence3-prior-undisputed-transaction-c1ac6e2b
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# dispute_visa_compelling_evidence3_prior_undisputed_transaction

```yaml
{"title": "DisputeVisaCompellingEvidence3PriorUndisputedTransaction", "required": ["charge"], "type": "object", "properties": {"charge": {"maxLength": 5000, "type": "string", "description": "Stripe charge ID for the Visa Compelling Evidence 3.0 eligible prior charge."}, "customer_account_id": {"maxLength": 5000, "type": "string", "description": "User Account ID used to log into business platform. Must be recognizable by the user.", "nullable": true}, "customer_device_fingerprint": {"maxLength": 5000, "type": "string", "description": "Unique identifier of the cardholder’s device derived from a combination of at least two hardware and software attributes. Must be at least 20 characters.", "nullable": true}, "customer_device_id": {"maxLength": 5000, "type": "string", "description": "Unique identifier of the cardholder’s device such as a device serial number (e.g., International Mobile Equipment Identity [IMEI]). Must be at least 15 characters.", "nullable": true}, "customer_email_address": {"maxLength": 5000, "type": "string", "description": "The email address of the customer.", "nullable": true}, "customer_purchase_ip": {"maxLength": 5000, "type": "string", "description": "The IP address that the customer used when making the purchase.", "nullable": true}, "product_description": {"maxLength": 150000, "type": "string", "description": "A description of the product or service that was sold.", "nullable": true}, "shipping_address": {"description": "The address to which a physical product was shipped. All fields are required for Visa Compelling Evidence 3.0 evidence submission.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/dispute_transaction_shipping_address"}]}}, "description": "", "x-expandableFields": ["shipping_address"]}
```
