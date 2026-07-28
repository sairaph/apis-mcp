---
title: custom-indicator-feeds_last_upload_summary
page_id: schema-custom-indicator-feeds-last-upload-summary-c11311d6
path: schemas
description: |-
    Summary of indicator counts from the last successful upload to this
    feed. Populated by the custom-threat-feeds loader at the end of each
    successful load. Absent (omitted) when no upload has completed
    successfully or the upload errored before the summary write.
    Surfaces silent-failure paths so operators can see when their
    indicators were dropped (popularity allowlist, expired valid_until,
    etc.) without reading loader logs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# custom-indicator-feeds_last_upload_summary

Summary of indicator counts from the last successful upload to this
feed. Populated by the custom-threat-feeds loader at the end of each
successful load. Absent (omitted) when no upload has completed
successfully or the upload errored before the summary write.
Surfaces silent-failure paths so operators can see when their
indicators were dropped (popularity allowlist, expired valid_until,
etc.) without reading loader logs.

```yaml
{"description": "Summary of indicator counts from the last successful upload to this\nfeed. Populated by the custom-threat-feeds loader at the end of each\nsuccessful load. Absent (omitted) when no upload has completed\nsuccessfully or the upload errored before the summary write.\nSurfaces silent-failure paths so operators can see when their\nindicators were dropped (popularity allowlist, expired valid_until,\netc.) without reading loader logs.\n", "type": "object", "properties": {"persisted": {"description": "Net delta applied to feed indicators by this upload. Snapshot\nuploads emit both *_added and *_removed; delta-add emits only\n*_added; delta-remove emits only *_removed.\n", "type": "object", "properties": {"domains_added": {"type": "integer", "x-auditable": true}, "domains_removed": {"type": "integer", "x-auditable": true}, "ips_added": {"type": "integer", "x-auditable": true}, "ips_removed": {"type": "integer", "x-auditable": true}, "urls_added": {"type": "integer", "x-auditable": true}, "urls_removed": {"type": "integer", "x-auditable": true}}}, "skipped": {"description": "Counts of indicators that were uploaded but did not reach\nQuickSilver, broken down by reason.\n", "type": "object", "properties": {"allowlisted_domains": {"description": "Domains filtered by the global popularity allowlist at QS\nprovisioning time. Popular domains (bing.com, naver.com,\netc.) are protected from custom-threat-feed enforcement.\n", "type": "integer", "x-auditable": true}, "expired_indicators": {"description": "Indicators in the upload whose valid_until is already in\nthe past. These are not added to QS; the expiration cron\nhandles cleanup.\n", "type": "integer", "x-auditable": true}, "invalid_indicators": {"description": "Reserved for future use. Currently always 0 — the unifier\naborts the entire upload on a single bad indicator.\n", "type": "integer", "x-auditable": true}}}, "uploaded": {"description": "Indicator counts from the unified file the loader received", "type": "object", "properties": {"domains": {"description": "Number of domain indicators in the upload", "type": "integer", "x-auditable": true}, "ips": {"description": "Number of IP indicators in the upload", "type": "integer", "x-auditable": true}, "urls": {"description": "Number of URL indicators in the upload", "type": "integer", "x-auditable": true}}}}, "x-auditable": true}
```
