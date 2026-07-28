---
title: Table
page_id: operation-get-accounts-account-id-dns-firewall-dns-firewall-id-dns-analytics-repor-73b10f35
path: operations/dns-firewall-analytics
description: |-
    Retrieves a list of summarised aggregate metrics over a given time period.

    See [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/) for detailed information about the available query parameters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dns_firewall/{dns_firewall_id}/dns_analytics/report
operation_ids:
    - dns-firewall-analytics-table
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Table

`GET /accounts/{account_id}/dns_firewall/{dns_firewall_id}/dns_analytics/report`

Operation ID: `dns-firewall-analytics-table`

Retrieves a list of summarised aggregate metrics over a given time period.

See [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/) for detailed information about the available query parameters.

## Definition

```yaml
{"operationId": "dns-firewall-analytics-table", "summary": "Table", "description": "Retrieves a list of summarised aggregate metrics over a given time period.\n\nSee [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/) for detailed information about the available query parameters.", "parameters": [{"name": "dns_firewall_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-analytics_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-analytics_identifier"}}, {"name": "metrics", "in": "query", "schema": {"$ref": "#/components/schemas/dns-analytics_metrics"}}, {"name": "dimensions", "in": "query", "schema": {"$ref": "#/components/schemas/dns-analytics_dimensions"}}, {"name": "since", "in": "query", "schema": {"$ref": "#/components/schemas/dns-analytics_since"}}, {"name": "until", "in": "query", "schema": {"$ref": "#/components/schemas/dns-analytics_until"}}, {"name": "limit", "in": "query", "schema": {"$ref": "#/components/schemas/dns-analytics_limit"}}, {"name": "sort", "in": "query", "schema": {"$ref": "#/components/schemas/dns-analytics_sort"}}, {"name": "filters", "in": "query", "schema": {"$ref": "#/components/schemas/dns-analytics_filters"}}], "responses": {"200": {"description": "Table response", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-analytics_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dns-analytics_report"}}, "type": "object"}]}}}}, "4XX": {"description": "Table response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-analytics_api-response-common-failure"}, {"properties": {"result": {"$ref": "#/components/schemas/dns-analytics_report"}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Firewall Analytics"], "x-api-token-group": ["DNS Firewall Write", "DNS Firewall Read"]}
```
