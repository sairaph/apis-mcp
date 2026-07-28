---
title: List DNS Firewall Clusters
page_id: operation-get-accounts-account-id-dns-firewall-a97b5d02
path: operations/dns-firewall
description: List DNS Firewall clusters for an account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dns_firewall
operation_ids:
    - dns-firewall-list-dns-firewall-clusters
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List DNS Firewall Clusters

`GET /accounts/{account_id}/dns_firewall`

Operation ID: `dns-firewall-list-dns-firewall-clusters`

List DNS Firewall clusters for an account

## Definition

```yaml
{"operationId": "dns-firewall-list-dns-firewall-clusters", "summary": "List DNS Firewall Clusters", "description": "List DNS Firewall clusters for an account", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-firewall_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of clusters per page", "type": "number", "default": 20, "maximum": 100, "minimum": 1}}], "responses": {"200": {"description": "List DNS Firewall Clusters response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-firewall_dns_firewall_response_collection"}}}}, "4XX": {"description": "List DNS Firewall Clusters response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-firewall_dns_firewall_response_collection"}, {"$ref": "#/components/schemas/dns-firewall_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Firewall"], "x-api-token-group": ["DNS Firewall Write", "DNS Firewall Read"], "x-cfPermissionsRequired": {"enum": ["#dns_records:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns-firewall", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
