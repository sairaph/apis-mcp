---
title: issuing_dispute_no_valid_authorization_evidence
page_id: schema-issuing-dispute-no-valid-authorization-evidence-3d6ad671
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_dispute_no_valid_authorization_evidence

```yaml
{"title": "IssuingDisputeNoValidAuthorizationEvidence", "type": "object", "properties": {"additional_documentation": {"description": "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}, "explanation": {"maxLength": 5000, "type": "string", "description": "Explanation of why the cardholder is disputing this transaction.", "nullable": true}}, "description": "", "x-expandableFields": ["additional_documentation"]}
```
