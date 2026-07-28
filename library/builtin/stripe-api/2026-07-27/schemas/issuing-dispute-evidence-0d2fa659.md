---
title: issuing_dispute_evidence
page_id: schema-issuing-dispute-evidence-0d2fa659
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_dispute_evidence

```yaml
{"title": "IssuingDisputeEvidence", "required": ["reason"], "type": "object", "properties": {"canceled": {"$ref": "#/components/schemas/issuing_dispute_canceled_evidence"}, "duplicate": {"$ref": "#/components/schemas/issuing_dispute_duplicate_evidence"}, "fraudulent": {"$ref": "#/components/schemas/issuing_dispute_fraudulent_evidence"}, "merchandise_not_as_described": {"$ref": "#/components/schemas/issuing_dispute_merchandise_not_as_described_evidence"}, "no_valid_authorization": {"$ref": "#/components/schemas/issuing_dispute_no_valid_authorization_evidence"}, "not_received": {"$ref": "#/components/schemas/issuing_dispute_not_received_evidence"}, "other": {"$ref": "#/components/schemas/issuing_dispute_other_evidence"}, "reason": {"type": "string", "description": "The reason for filing the dispute. Its value will match the field containing the evidence.", "enum": ["canceled", "duplicate", "fraudulent", "merchandise_not_as_described", "no_valid_authorization", "not_received", "other", "service_not_as_described"], "x-stripeBypassValidation": true}, "service_not_as_described": {"$ref": "#/components/schemas/issuing_dispute_service_not_as_described_evidence"}}, "description": "", "x-expandableFields": ["canceled", "duplicate", "fraudulent", "merchandise_not_as_described", "no_valid_authorization", "not_received", "other", "service_not_as_described"]}
```
