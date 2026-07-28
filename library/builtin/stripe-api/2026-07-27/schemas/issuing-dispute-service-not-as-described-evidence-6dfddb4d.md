---
title: issuing_dispute_service_not_as_described_evidence
page_id: schema-issuing-dispute-service-not-as-described-evidence-6dfddb4d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_dispute_service_not_as_described_evidence

```yaml
{"title": "IssuingDisputeServiceNotAsDescribedEvidence", "type": "object", "properties": {"additional_documentation": {"description": "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}, "canceled_at": {"type": "integer", "description": "Date when order was canceled.", "format": "unix-time", "nullable": true}, "cancellation_reason": {"maxLength": 5000, "type": "string", "description": "Reason for canceling the order.", "nullable": true}, "explanation": {"maxLength": 5000, "type": "string", "description": "Explanation of why the cardholder is disputing this transaction.", "nullable": true}, "received_at": {"type": "integer", "description": "Date when the product was received.", "format": "unix-time", "nullable": true}}, "description": "", "x-expandableFields": ["additional_documentation"]}
```
