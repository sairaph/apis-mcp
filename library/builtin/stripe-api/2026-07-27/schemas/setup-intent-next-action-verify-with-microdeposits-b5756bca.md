---
title: setup_intent_next_action_verify_with_microdeposits
page_id: schema-setup-intent-next-action-verify-with-microdeposits-b5756bca
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# setup_intent_next_action_verify_with_microdeposits

```yaml
{"title": "SetupIntentNextActionVerifyWithMicrodeposits", "required": ["arrival_date", "hosted_verification_url"], "type": "object", "properties": {"arrival_date": {"type": "integer", "description": "The timestamp when the microdeposits are expected to land.", "format": "unix-time"}, "hosted_verification_url": {"maxLength": 5000, "type": "string", "description": "The URL for the hosted verification page, which allows customers to verify their bank account."}, "microdeposit_type": {"type": "string", "description": "The type of the microdeposit sent to the customer. Used to distinguish between different verification methods.", "nullable": true, "enum": ["amounts", "descriptor_code"]}}, "description": "", "x-expandableFields": []}
```
