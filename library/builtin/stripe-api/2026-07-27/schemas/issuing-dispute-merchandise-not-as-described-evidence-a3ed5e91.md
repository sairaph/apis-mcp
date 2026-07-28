---
title: issuing_dispute_merchandise_not_as_described_evidence
page_id: schema-issuing-dispute-merchandise-not-as-described-evidence-a3ed5e91
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_dispute_merchandise_not_as_described_evidence

```yaml
{"title": "IssuingDisputeMerchandiseNotAsDescribedEvidence", "type": "object", "properties": {"additional_documentation": {"description": "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}, "explanation": {"maxLength": 5000, "type": "string", "description": "Explanation of why the cardholder is disputing this transaction.", "nullable": true}, "received_at": {"type": "integer", "description": "Date when the product was received.", "format": "unix-time", "nullable": true}, "return_description": {"maxLength": 5000, "type": "string", "description": "Description of the cardholder's attempt to return the product.", "nullable": true}, "return_status": {"type": "string", "description": "Result of cardholder's attempt to return the product.", "nullable": true, "enum": ["merchant_rejected", "successful"]}, "returned_at": {"type": "integer", "description": "Date when the product was returned or attempted to be returned.", "format": "unix-time", "nullable": true}}, "description": "", "x-expandableFields": ["additional_documentation"]}
```
