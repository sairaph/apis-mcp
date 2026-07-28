---
title: Gets the can_be_enabled zone setting
page_id: operation-post-accounts-account-id-pay-per-crawl-zones-can-be-enabled-query-f4c1f601
path: operations/ppc-config
description: Provided a list of pay-per-crawl configured zones this method will return whether they can enable PPC or not.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/pay-per-crawl/zones_can_be_enabled/query
operation_ids:
    - pay-per-crawl.queryZonesCanBeEnabled
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Gets the can_be_enabled zone setting

`POST /accounts/{account_id}/pay-per-crawl/zones_can_be_enabled/query`

Operation ID: `pay-per-crawl.queryZonesCanBeEnabled`

Provided a list of pay-per-crawl configured zones this method will return whether they can enable PPC or not.

## Definition

```yaml
{"operationId": "pay-per-crawl.queryZonesCanBeEnabled", "summary": "Gets the can_be_enabled zone setting", "description": "Provided a list of pay-per-crawl configured zones this method will return whether they can enable PPC or not.", "parameters": [{"name": "account_id", "in": "path", "description": "account id", "required": true, "schema": {"type": "string"}, "x-auditable": true}], "requestBody": {"$ref": "#/components/requestBodies/pay-per-crawl_ZonesCanBeEnabledPayload"}, "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pay-per-crawl_queryZonesCanBeEnabledResponse"}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pay-per-crawl_apiErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["ppc_config"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read", "DNS View Write", "DNS View Read", "SCIM Provisioning", "Load Balancers Account Write", "Load Balancers Account Read", "Zero Trust: PII Read", "DDoS Botnet Feed Write", "DDoS Botnet Feed Read", "Workers R2 Storage Write", "Workers R2 Storage Read", "DDoS Protection Write", "DDoS Protection Read", "Workers Tail Read", "Workers KV Storage Write", "Workers KV Storage Read", "Workers Scripts Write", "Workers Scripts Read", "Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read", "Account Firewall Access Rules Write", "Account Firewall Access Rules Read", "DNS Firewall Write", "DNS Firewall Read", "Billing Write", "Billing Read", "Account Settings Write", "Account Settings Read"]}
```
