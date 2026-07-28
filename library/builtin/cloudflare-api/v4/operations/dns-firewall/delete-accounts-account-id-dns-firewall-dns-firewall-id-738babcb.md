---
title: Delete DNS Firewall Cluster
page_id: operation-delete-accounts-account-id-dns-firewall-dns-firewall-id-6d06047a
path: operations/dns-firewall
description: Delete a DNS Firewall cluster
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dns_firewall/{dns_firewall_id}
operation_ids:
    - dns-firewall-delete-dns-firewall-cluster
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete DNS Firewall Cluster

`DELETE /accounts/{account_id}/dns_firewall/{dns_firewall_id}`

Operation ID: `dns-firewall-delete-dns-firewall-cluster`

Delete a DNS Firewall cluster

## Definition

```yaml
{"operationId": "dns-firewall-delete-dns-firewall-cluster", "summary": "Delete DNS Firewall Cluster", "description": "Delete a DNS Firewall cluster", "parameters": [{"name": "dns_firewall_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-firewall_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete DNS Firewall Cluster response", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-firewall_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/dns-firewall_identifier"}}}}, "type": "object"}]}}}}, "4XX": {"description": "Delete DNS Firewall Cluster response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-firewall_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/dns-firewall_identifier"}}}}, "type": "object"}, {"$ref": "#/components/schemas/dns-firewall_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Firewall"], "x-api-token-group": ["DNS Firewall Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns-firewall", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
