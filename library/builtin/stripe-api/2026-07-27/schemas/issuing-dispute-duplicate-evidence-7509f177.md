---
title: issuing_dispute_duplicate_evidence
page_id: schema-issuing-dispute-duplicate-evidence-7509f177
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_dispute_duplicate_evidence

```yaml
{"title": "IssuingDisputeDuplicateEvidence", "type": "object", "properties": {"additional_documentation": {"description": "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}, "card_statement": {"description": "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Copy of the card statement showing that the product had already been paid for.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}, "cash_receipt": {"description": "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Copy of the receipt showing that the product had been paid for in cash.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}, "check_image": {"description": "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Image of the front and back of the check that was used to pay for the product.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}, "explanation": {"maxLength": 5000, "type": "string", "description": "Explanation of why the cardholder is disputing this transaction.", "nullable": true}, "original_transaction": {"maxLength": 5000, "type": "string", "description": "Transaction (e.g., ipi_...) that the disputed transaction is a duplicate of. Of the two or more transactions that are copies of each other, this is original undisputed one.", "nullable": true}}, "description": "", "x-expandableFields": ["additional_documentation", "card_statement", "cash_receipt", "check_image"]}
```
