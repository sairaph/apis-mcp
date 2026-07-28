---
title: source_redirect_flow
page_id: schema-source-redirect-flow-ceecd775
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# source_redirect_flow

```yaml
{"title": "SourceRedirectFlow", "required": ["return_url", "status", "url"], "type": "object", "properties": {"failure_reason": {"maxLength": 5000, "type": "string", "description": "The failure reason for the redirect, either `user_abort` (the customer aborted or dropped out of the redirect flow), `declined` (the authentication failed or the transaction was declined), or `processing_error` (the redirect failed due to a technical error). Present only if the redirect status is `failed`.", "nullable": true}, "return_url": {"maxLength": 5000, "type": "string", "description": "The URL you provide to redirect the customer to after they authenticated their payment."}, "status": {"maxLength": 5000, "type": "string", "description": "The status of the redirect, either `pending` (ready to be used by your customer to authenticate the transaction), `succeeded` (successful authentication, cannot be reused) or `not_required` (redirect should not be used) or `failed` (failed authentication, cannot be reused)."}, "url": {"maxLength": 2048, "type": "string", "description": "The URL provided to you to redirect a customer to as part of a `redirect` authentication flow."}}, "description": "", "x-expandableFields": []}
```
