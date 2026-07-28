---
title: issuing_dispute_canceled_evidence
page_id: schema-issuing-dispute-canceled-evidence-11c28fb6
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_dispute_canceled_evidence

```yaml
{"title": "IssuingDisputeCanceledEvidence", "type": "object", "properties": {"additional_documentation": {"description": "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}, "canceled_at": {"type": "integer", "description": "Date when order was canceled.", "format": "unix-time", "nullable": true}, "cancellation_policy_provided": {"type": "boolean", "description": "Whether the cardholder was provided with a cancellation policy.", "nullable": true}, "cancellation_reason": {"maxLength": 5000, "type": "string", "description": "Reason for canceling the order.", "nullable": true}, "expected_at": {"type": "integer", "description": "Date when the cardholder expected to receive the product.", "format": "unix-time", "nullable": true}, "explanation": {"maxLength": 5000, "type": "string", "description": "Explanation of why the cardholder is disputing this transaction.", "nullable": true}, "product_description": {"maxLength": 5000, "type": "string", "description": "Description of the merchandise or service that was purchased.", "nullable": true}, "product_type": {"type": "string", "description": "Whether the product was a merchandise or service.", "nullable": true, "enum": ["merchandise", "service"]}, "return_status": {"type": "string", "description": "Result of cardholder's attempt to return the product.", "nullable": true, "enum": ["merchant_rejected", "successful"]}, "returned_at": {"type": "integer", "description": "Date when the product was returned or attempted to be returned.", "format": "unix-time", "nullable": true}}, "description": "", "x-expandableFields": ["additional_documentation"]}
```
