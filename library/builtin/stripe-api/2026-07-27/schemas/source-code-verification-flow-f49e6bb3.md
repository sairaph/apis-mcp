---
title: source_code_verification_flow
page_id: schema-source-code-verification-flow-f49e6bb3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# source_code_verification_flow

```yaml
{"title": "SourceCodeVerificationFlow", "required": ["attempts_remaining", "status"], "type": "object", "properties": {"attempts_remaining": {"type": "integer", "description": "The number of attempts remaining to authenticate the source object with a verification code."}, "status": {"maxLength": 5000, "type": "string", "description": "The status of the code verification, either `pending` (awaiting verification, `attempts_remaining` should be greater than 0), `succeeded` (successful verification) or `failed` (failed verification, cannot be verified anymore as `attempts_remaining` should be 0)."}}, "description": "", "x-expandableFields": []}
```
