---
title: Get daily report
page_id: operation-get-accounts-account-id-botnet-feed-asn-asn-id-day-report-8a44ed9a
path: operations/botnet-threat-feed
description: Gets all the data the botnet tracking database has for a given ASN registered to user account for given date. If no date is given, it will return results for the previous day.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/botnet_feed/asn/{asn_id}/day_report
operation_ids:
    - botnet-threat-feed-get-day-report
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get daily report

`GET /accounts/{account_id}/botnet_feed/asn/{asn_id}/day_report`

Operation ID: `botnet-threat-feed-get-day-report`

Gets all the data the botnet tracking database has for a given ASN registered to user account for given date. If no date is given, it will return results for the previous day.

## Definition

```yaml
{"operationId": "botnet-threat-feed-get-day-report", "summary": "Get daily report", "description": "Gets all the data the botnet tracking database has for a given ASN registered to user account for given date. If no date is given, it will return results for the previous day.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "asn_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dos_asn"}}, {"name": "date", "in": "query", "schema": {"$ref": "#/components/schemas/dos_timestamp"}}], "responses": {"200": {"description": "Get botnet feed report for day", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dos_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"cidr": {"type": "string", "example": "1.1.1.1/32"}, "date": {"type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z"}, "offense_count": {"type": "integer", "example": 1000}}}}, "type": "object"}]}}}}, "4XX": {"description": "Get botnet feed report for day response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Botnet Threat Feed"], "x-api-token-group": ["DDoS Botnet Feed Write", "DDoS Botnet Feed Read"]}
```
