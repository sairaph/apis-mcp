---
title: Download command output file
page_id: operation-get-accounts-account-id-dex-commands-command-id-downloads-filename-3a5b721d
path: operations/dex-remote-commands
description: Downloads artifacts for an executed command. Bulk downloads are not supported
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/commands/{command_id}/downloads/{filename}
operation_ids:
    - get-commands-command-id-downloads-filename
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Download command output file

`GET /accounts/{account_id}/dex/commands/{command_id}/downloads/{filename}`

Operation ID: `get-commands-command-id-downloads-filename`

Downloads artifacts for an executed command. Bulk downloads are not supported

## Definition

```yaml
{"operationId": "get-commands-command-id-downloads-filename", "summary": "Download command output file", "description": "Downloads artifacts for an executed command. Bulk downloads are not supported", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "command_id", "in": "path", "description": "Unique identifier for command.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_command_id"}}, {"name": "filename", "in": "path", "description": "The name of the file to be downloaded, including the `.zip` extension.", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Get command artifacts response.", "headers": {"Accept-Ranges": {"description": "Indicates that the file supports byte-range requests.", "schema": {"type": "string", "example": "bytes"}}, "Content-Disposition": {"description": "Indicates that the file should be treated as an attachment for downloading", "schema": {"type": "string", "example": "attachment; filename*=example.zip"}}, "Content-Encoding": {"description": "Specifies the encoding of the file content, if any.", "schema": {"type": "string", "example": "gzip"}}, "Content-Length": {"description": "The size of the file in bytes.", "schema": {"type": "integer"}}, "Content-Type": {"description": "Specifies the media type of the file.", "schema": {"type": "string", "example": "application/zip"}}, "ETag": {"description": "The entity tag of the file for cache validation.", "schema": {"type": "string", "example": "34f9b5e9c8a6"}}, "Last-Modified": {"description": "The last modification date of the file.", "schema": {"type": "string", "format": "date-time", "example": "Wed, 21 Oct 2023 07:28:00 GMT"}}}, "content": {"application/zip": {"schema": {"type": "string", "format": "binary"}}}}, "4XX": {"description": "Get downloaded commands failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Remote Commands"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.commands.results", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
