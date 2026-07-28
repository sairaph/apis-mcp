---
title: Create DNS Firewall Cluster
page_id: operation-post-accounts-account-id-dns-firewall-f49722d8
path: operations/dns-firewall
description: Create a DNS Firewall cluster
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dns_firewall
operation_ids:
    - dns-firewall-create-dns-firewall-cluster
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create DNS Firewall Cluster

`POST /accounts/{account_id}/dns_firewall`

Operation ID: `dns-firewall-create-dns-firewall-cluster`

Create a DNS Firewall cluster

## Definition

```yaml
{"operationId": "dns-firewall-create-dns-firewall-cluster", "summary": "Create DNS Firewall Cluster", "description": "Create a DNS Firewall cluster", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-firewall_dns-firewall-cluster-post"}}}}, "responses": {"200": {"description": "Create DNS Firewall Cluster response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-firewall_dns_firewall_single_response"}}}}, "4XX": {"description": "Create DNS Firewall Cluster response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-firewall_dns_firewall_single_response"}, {"$ref": "#/components/schemas/dns-firewall_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Firewall"], "x-api-token-group": ["DNS Firewall Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns-firewall", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
