---
title: dispute_evidence_details
page_id: schema-dispute-evidence-details-47eee214
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# dispute_evidence_details

```yaml
{"title": "DisputeEvidenceDetails", "required": ["enhanced_eligibility", "has_evidence", "past_due", "submission_count"], "type": "object", "properties": {"due_by": {"type": "integer", "description": "Date by which evidence must be submitted in order to successfully challenge dispute. Will be 0 if the customer's bank or credit card company doesn't allow a response for this particular dispute.", "format": "unix-time", "nullable": true}, "enhanced_eligibility": {"$ref": "#/components/schemas/dispute_enhanced_eligibility"}, "has_evidence": {"type": "boolean", "description": "Whether evidence has been staged for this dispute."}, "past_due": {"type": "boolean", "description": "Whether the last evidence submission was submitted past the due date. Defaults to `false` if no evidence submissions have occurred. If `true`, then delivery of the latest evidence is *not* guaranteed."}, "submission_count": {"type": "integer", "description": "The number of times evidence has been submitted. Typically, you may only submit evidence once."}}, "description": "", "x-expandableFields": ["enhanced_eligibility"]}
```
