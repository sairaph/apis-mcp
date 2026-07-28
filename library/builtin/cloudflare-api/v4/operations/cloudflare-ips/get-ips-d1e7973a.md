---
title: Cloudflare/JD Cloud IP Details
page_id: operation-get-ips-37ee8c85
path: operations/cloudflare-ips
description: Get IPs used on the Cloudflare/JD Cloud network, see https://www.cloudflare.com/ips for Cloudflare IPs or https://developers.cloudflare.com/china-network/reference/infrastructure/ for JD Cloud IPs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /ips
operation_ids:
    - cloudflare-ips-cloudflare-ip-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Cloudflare/JD Cloud IP Details

`GET /ips`

Operation ID: `cloudflare-ips-cloudflare-ip-details`

Get IPs used on the Cloudflare/JD Cloud network, see https://www.cloudflare.com/ips for Cloudflare IPs or https://developers.cloudflare.com/china-network/reference/infrastructure/ for JD Cloud IPs.

## Definition

```yaml
{"operationId": "cloudflare-ips-cloudflare-ip-details", "summary": "Cloudflare/JD Cloud IP Details", "description": "Get IPs used on the Cloudflare/JD Cloud network, see https://www.cloudflare.com/ips for Cloudflare IPs or https://developers.cloudflare.com/china-network/reference/infrastructure/ for JD Cloud IPs.", "parameters": [{"name": "networks", "in": "query", "description": "Specified as `jdcloud` to list IPs used by JD Cloud data centers.", "schema": {"type": "string"}}], "responses": {"200": {"description": "Cloudflare IP Details response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/public-ip_api-response-single"}, {"properties": {"result": {"oneOf": [{"$ref": "#/components/schemas/public-ip_ips"}, {"$ref": "#/components/schemas/public-ip_ips_jdcloud"}]}}}]}}}}, "4XX": {"description": "Cloudflare IP Details response failure", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/public-ip_api-response-single"}, {"properties": {"result": {"oneOf": [{"$ref": "#/components/schemas/public-ip_ips"}, {"$ref": "#/components/schemas/public-ip_ips_jdcloud"}]}}}]}, {"$ref": "#/components/schemas/public-ip_api-response-common-failure"}]}}}}}, "security": [{}], "tags": ["Cloudflare IPs"]}
```
