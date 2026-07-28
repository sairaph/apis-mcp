---
title: Gets the stripe config for a crawler
page_id: operation-get-accounts-account-id-pay-per-crawl-crawler-stripe-f6628039
path: operations/ppc-stripe
description: Gets the stripe config for a crawler.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pay-per-crawl/crawler/stripe
operation_ids:
    - pay-per-crawl.crawlerGetStripeConfig
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Gets the stripe config for a crawler

`GET /accounts/{account_id}/pay-per-crawl/crawler/stripe`

Operation ID: `pay-per-crawl.crawlerGetStripeConfig`

Gets the stripe config for a crawler.

## Definition

```yaml
{"operationId": "pay-per-crawl.crawlerGetStripeConfig", "summary": "Gets the stripe config for a crawler", "description": "Gets the stripe config for a crawler.", "parameters": [{"name": "account_id", "in": "path", "description": "account id", "required": true, "schema": {"type": "string"}, "x-auditable": true}], "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pay-per-crawl_getStripeConfigResponse"}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pay-per-crawl_apiErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["ppc_stripe"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read", "DNS View Write", "DNS View Read", "SCIM Provisioning", "Load Balancers Account Write", "Load Balancers Account Read", "Zero Trust: PII Read", "DDoS Botnet Feed Write", "DDoS Botnet Feed Read", "Workers R2 Storage Write", "Workers R2 Storage Read", "DDoS Protection Write", "DDoS Protection Read", "Workers Tail Read", "Workers KV Storage Write", "Workers KV Storage Read", "Workers Scripts Write", "Workers Scripts Read", "Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read", "Account Firewall Access Rules Write", "Account Firewall Access Rules Read", "DNS Firewall Write", "DNS Firewall Read", "Billing Write", "Billing Read", "Account Settings Write", "Account Settings Read"]}
```
