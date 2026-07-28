---
title: Update indicator feed data
page_id: operation-put-accounts-account-id-intel-indicator-feeds-feed-id-snapshot-e155d555
path: operations/custom-indicator-feeds
description: |-
    Revises the raw data entries in a custom threat indicator feed.

    Accepts both plain and gzipped STIX2/CRDF bodies. Gzip is
    detected by RFC 1952 magic bytes (`0x1f 0x8b`) and/or a `.gz`
    filename suffix (case-insensitive) — either signal alone is
    sufficient to trigger the gzip path; if the body is not valid
    gzip, the upload fails fast. Customers are encouraged to gzip
    larger uploads — the api-gateway 500 MB body cap applies to
    the on-the-wire (compressed) size, so gzip lets a single
    upload carry several GiB of decompressed STIX.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/intel/indicator-feeds/{feed_id}/snapshot
operation_ids:
    - custom-indicator-feeds-update-indicator-feed-data
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update indicator feed data

`PUT /accounts/{account_id}/intel/indicator-feeds/{feed_id}/snapshot`

Operation ID: `custom-indicator-feeds-update-indicator-feed-data`

Revises the raw data entries in a custom threat indicator feed.

Accepts both plain and gzipped STIX2/CRDF bodies. Gzip is
detected by RFC 1952 magic bytes (`0x1f 0x8b`) and/or a `.gz`
filename suffix (case-insensitive) — either signal alone is
sufficient to trigger the gzip path; if the body is not valid
gzip, the upload fails fast. Customers are encouraged to gzip
larger uploads — the api-gateway 500 MB body cap applies to
the on-the-wire (compressed) size, so gzip lets a single
upload carry several GiB of decompressed STIX.

## Definition

```yaml
{"operationId": "custom-indicator-feeds-update-indicator-feed-data", "summary": "Update indicator feed data", "description": "Revises the raw data entries in a custom threat indicator feed.\n\nAccepts both plain and gzipped STIX2/CRDF bodies. Gzip is\ndetected by RFC 1952 magic bytes (`0x1f 0x8b`) and/or a `.gz`\nfilename suffix (case-insensitive) — either signal alone is\nsufficient to trigger the gzip path; if the body is not valid\ngzip, the upload fails fast. Customers are encouraged to gzip\nlarger uploads — the api-gateway 500 MB body cap applies to\nthe on-the-wire (compressed) size, so gzip lets a single\nupload carry several GiB of decompressed STIX.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-indicator-feeds_identifier"}}, {"name": "feed_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-indicator-feeds_feed_id"}}], "requestBody": {"required": true, "content": {"multipart/form-data": {"schema": {"type": "object", "properties": {"source": {"description": "The file to upload. Either a plain STIX2/CRDF body\nor a gzipped one (recognised by `0x1f 0x8b` magic\nbytes or a `.gz` filename suffix).\n", "type": "string", "example": "@/Users/me/test.stix2.gz"}}}}}}, "responses": {"200": {"description": "Get indicator feed metadata", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-indicator-feeds_update_feed_response"}}}}, "413": {"description": "Decompressed upload exceeds the maximum allowed size. The\nserver caps the decompressed body at a configured limit\n(default 6 GiB) as a gzip-bomb defence; this response is\nreturned before any feed state changes. Resubmit a smaller\n(or non-pathologically-compressed) body.\n", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-indicator-feeds_api_response_common_failure"}]}}}}, "4XX": {"description": "Get indicator feeds response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-indicator-feeds_api_response_common_failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom Indicator Feeds"]}
```
