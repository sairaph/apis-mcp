---
title: Get ASN Subnets
page_id: operation-get-accounts-account-id-intel-asn-asn-subnets-7f8b9b13
path: operations/asn-intelligence
description: Get ASN Subnets.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/asn/{asn}/subnets
operation_ids:
    - asn-intelligence-get-asn-subnets
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get ASN Subnets

`GET /accounts/{account_id}/intel/asn/{asn}/subnets`

Operation ID: `asn-intelligence-get-asn-subnets`

Get ASN Subnets.

## Definition

```yaml
{"operationId": "asn-intelligence-get-asn-subnets", "summary": "Get ASN Subnets", "description": "Get ASN Subnets.", "parameters": [{"name": "asn", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/intel_asn"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/intel_identifier"}}], "responses": {"200": {"description": "Get ASN Subnets response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"asn": {"$ref": "#/components/schemas/intel_asn"}, "count": {"$ref": "#/components/schemas/intel_count"}, "ip_count_total": {"type": "integer"}, "page": {"$ref": "#/components/schemas/intel_page"}, "per_page": {"$ref": "#/components/schemas/intel_per_page"}, "subnets": {"type": "array", "items": {"type": "string"}, "example": ["192.0.2.0/24", "2001:DB8::/32"]}}}}}}, "4XX": {"description": "Get ASN Subnets response failure.", "content": {"application/json": {"schema": {"allOf": [{"properties": {"asn": {"$ref": "#/components/schemas/intel_asn"}, "count": {"$ref": "#/components/schemas/intel_count"}, "ip_count_total": {"type": "integer"}, "page": {"$ref": "#/components/schemas/intel_page"}, "per_page": {"$ref": "#/components/schemas/intel_per_page"}, "subnets": {"type": "array", "items": {"type": "string"}, "example": ["192.0.2.0/24", "2001:DB8::/32"]}}, "type": "object"}, {"$ref": "#/components/schemas/intel_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["ASN Intelligence"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
