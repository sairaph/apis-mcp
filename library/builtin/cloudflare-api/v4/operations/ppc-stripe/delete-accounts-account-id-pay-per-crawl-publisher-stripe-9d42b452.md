---
title: Deletes the stripe config for a publisher
page_id: operation-delete-accounts-account-id-pay-per-crawl-publisher-stripe-75239270
path: operations/ppc-stripe
description: Deletes the stripe config for a publisher.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/pay-per-crawl/publisher/stripe
operation_ids:
    - pay-per-crawl.publisherDeleteStripeConfig
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Deletes the stripe config for a publisher

`DELETE /accounts/{account_id}/pay-per-crawl/publisher/stripe`

Operation ID: `pay-per-crawl.publisherDeleteStripeConfig`

Deletes the stripe config for a publisher.

## Definition

```yaml
{"operationId": "pay-per-crawl.publisherDeleteStripeConfig", "summary": "Deletes the stripe config for a publisher", "description": "Deletes the stripe config for a publisher.", "parameters": [{"name": "account_id", "in": "path", "description": "account id", "required": true, "schema": {"type": "string"}, "x-auditable": true}], "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pay-per-crawl_apiNoResultResponse"}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pay-per-crawl_apiErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["ppc_stripe"], "x-api-token-group": ["Account Settings Write"]}
```
