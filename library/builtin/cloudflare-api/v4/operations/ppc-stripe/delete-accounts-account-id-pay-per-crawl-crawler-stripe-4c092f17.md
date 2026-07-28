---
title: Deletes the stripe config for a crawler
page_id: operation-delete-accounts-account-id-pay-per-crawl-crawler-stripe-6985b2a7
path: operations/ppc-stripe
description: Deletes the stripe config for a crawler.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/pay-per-crawl/crawler/stripe
operation_ids:
    - pay-per-crawl.crawlerDeleteStripeConfig
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Deletes the stripe config for a crawler

`DELETE /accounts/{account_id}/pay-per-crawl/crawler/stripe`

Operation ID: `pay-per-crawl.crawlerDeleteStripeConfig`

Deletes the stripe config for a crawler.

## Definition

```yaml
{"operationId": "pay-per-crawl.crawlerDeleteStripeConfig", "summary": "Deletes the stripe config for a crawler", "description": "Deletes the stripe config for a crawler.", "parameters": [{"name": "account_id", "in": "path", "description": "account id", "required": true, "schema": {"type": "string"}, "x-auditable": true}], "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pay-per-crawl_apiNoResultResponse"}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pay-per-crawl_apiErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["ppc_stripe"], "x-api-token-group": ["Account Settings Write"]}
```
