---
title: Plan account routing rule changes
page_id: operation-post-accounts-account-id-email-routing-rules-plan-230356b6
path: operations/email-routing-routing-rules
description: Computes the Email Routing rule changes that would be needed to reconcile a Wrangler-managed desired ruleset. This endpoint is read-only and does not create, update, or delete rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email/routing/rules/plan
operation_ids:
    - email-routing-routing-rules-plan-account-routing-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Plan account routing rule changes

`POST /accounts/{account_id}/email/routing/rules/plan`

Operation ID: `email-routing-routing-rules-plan-account-routing-rules`

Computes the Email Routing rule changes that would be needed to reconcile a Wrangler-managed desired ruleset. This endpoint is read-only and does not create, update, or delete rules.

## Definition

```yaml
{"operationId": "email-routing-routing-rules-plan-account-routing-rules", "summary": "Plan account routing rule changes", "description": "Computes the Email Routing rule changes that would be needed to reconcile a Wrangler-managed desired ruleset. This endpoint is read-only and does not create, update, or delete rules.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_account_rules_plan_request"}}}}, "responses": {"200": {"description": "Account routing rules plan response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_account_rules_plan_response_single"}}}}, "400": {"description": "Error 2062: Unknown Email Routing domain. A desired address or catch-all target does not resolve to an Email Routing-enabled domain in this account."}, "403": {"description": "Authentication error or missing account rule list permission."}, "422": {"description": "Invalid plan input, such as malformed matchers/actions, duplicate targets, or an owner_worker_tag that does not resolve to a known Worker script."}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing routing rules"], "x-api-token-group": null, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.rules", "x-fern-sdk-method-name": "plan", "x-forge-hidden": true}
```
