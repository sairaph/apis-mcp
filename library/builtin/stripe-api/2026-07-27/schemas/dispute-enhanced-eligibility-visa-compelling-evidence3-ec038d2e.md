---
title: dispute_enhanced_eligibility_visa_compelling_evidence3
page_id: schema-dispute-enhanced-eligibility-visa-compelling-evidence3-ec038d2e
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# dispute_enhanced_eligibility_visa_compelling_evidence3

```yaml
{"title": "DisputeEnhancedEligibilityVisaCompellingEvidence3", "required": ["required_actions", "status"], "type": "object", "properties": {"required_actions": {"type": "array", "description": "List of actions required to qualify dispute for Visa Compelling Evidence 3.0 evidence submission.", "items": {"type": "string", "enum": ["missing_customer_identifiers", "missing_disputed_transaction_description", "missing_merchandise_or_services", "missing_prior_undisputed_transaction_description", "missing_prior_undisputed_transactions"]}}, "status": {"type": "string", "description": "Visa Compelling Evidence 3.0 eligibility status.", "enum": ["not_qualified", "qualified", "requires_action"]}}, "description": "", "x-expandableFields": []}
```
