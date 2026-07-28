---
title: dispute_enhanced_evidence_visa_compelling_evidence3
page_id: schema-dispute-enhanced-evidence-visa-compelling-evidence3-b73132fb
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# dispute_enhanced_evidence_visa_compelling_evidence3

```yaml
{"title": "DisputeEnhancedEvidenceVisaCompellingEvidence3", "required": ["prior_undisputed_transactions"], "type": "object", "properties": {"disputed_transaction": {"description": "Disputed transaction details for Visa Compelling Evidence 3.0 evidence submission.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/dispute_visa_compelling_evidence3_disputed_transaction"}]}, "prior_undisputed_transactions": {"type": "array", "description": "List of exactly two prior undisputed transaction objects for Visa Compelling Evidence 3.0 evidence submission.", "items": {"$ref": "#/components/schemas/dispute_visa_compelling_evidence3_prior_undisputed_transaction"}}}, "description": "", "x-expandableFields": ["disputed_transaction", "prior_undisputed_transactions"]}
```
