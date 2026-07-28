---
title: Update DNS Firewall Cluster
page_id: operation-patch-accounts-account-id-dns-firewall-dns-firewall-id-b4d72d66
path: operations/dns-firewall
description: Modify the configuration of a DNS Firewall cluster
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/dns_firewall/{dns_firewall_id}
operation_ids:
    - dns-firewall-update-dns-firewall-cluster
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update DNS Firewall Cluster

`PATCH /accounts/{account_id}/dns_firewall/{dns_firewall_id}`

Operation ID: `dns-firewall-update-dns-firewall-cluster`

Modify the configuration of a DNS Firewall cluster

## Definition

```yaml
{"operationId": "dns-firewall-update-dns-firewall-cluster", "summary": "Update DNS Firewall Cluster", "description": "Modify the configuration of a DNS Firewall cluster", "parameters": [{"name": "dns_firewall_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-firewall_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-firewall_dns-firewall-cluster-patch"}}}}, "responses": {"200": {"description": "Update DNS Firewall Cluster response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-firewall_dns_firewall_single_response"}}}}, "4XX": {"description": "Update DNS Firewall Cluster response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-firewall_dns_firewall_single_response"}, {"$ref": "#/components/schemas/dns-firewall_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Firewall"], "x-api-token-group": ["DNS Firewall Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns-firewall", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
