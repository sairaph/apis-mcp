---
title: dispute_enhanced_evidence_visa_compliance
page_id: schema-dispute-enhanced-evidence-visa-compliance-b5e58aad
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# dispute_enhanced_evidence_visa_compliance

```yaml
{"title": "DisputeEnhancedEvidenceVisaCompliance", "required": ["fee_acknowledged"], "type": "object", "properties": {"fee_acknowledged": {"type": "boolean", "description": "A field acknowledging the fee incurred when countering a Visa compliance dispute. If this field is set to true, evidence can be submitted for the compliance dispute. Stripe collects a 500 USD (or local equivalent) amount to cover the network costs associated with resolving compliance disputes. Stripe refunds the 500 USD network fee if you win the dispute."}}, "description": "", "x-expandableFields": []}
```
