---
title: DNS Firewall Cluster Details
page_id: operation-get-accounts-account-id-dns-firewall-dns-firewall-id-b3536291
path: operations/dns-firewall
description: Show a single DNS Firewall cluster for an account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dns_firewall/{dns_firewall_id}
operation_ids:
    - dns-firewall-dns-firewall-cluster-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# DNS Firewall Cluster Details

`GET /accounts/{account_id}/dns_firewall/{dns_firewall_id}`

Operation ID: `dns-firewall-dns-firewall-cluster-details`

Show a single DNS Firewall cluster for an account

## Definition

```yaml
{"operationId": "dns-firewall-dns-firewall-cluster-details", "summary": "DNS Firewall Cluster Details", "description": "Show a single DNS Firewall cluster for an account", "parameters": [{"name": "dns_firewall_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-firewall_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-firewall_identifier"}}], "responses": {"200": {"description": "DNS Firewall Cluster Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-firewall_dns_firewall_single_response"}}}}, "4XX": {"description": "DNS Firewall Cluster Details response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-firewall_dns_firewall_single_response"}, {"$ref": "#/components/schemas/dns-firewall_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Firewall"], "x-api-token-group": ["DNS Firewall Write", "DNS Firewall Read"], "x-cfPermissionsRequired": {"enum": ["#dns_records:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns-firewall", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
