---
title: dnssec_dnssec_use_nsec3
page_id: schema-dnssec-dnssec-use-nsec3-3fffe8d7
path: schemas
description: |-
    If true, enables the use of NSEC3 together with DNSSEC on the zone.
    Combined with setting dnssec_presigned to true, this enables the use of
    NSEC3 records when transferring in from an external provider.
    If dnssec_presigned is instead set to false (default), NSEC3 records will be
    generated and signed at request time.

    See [DNSSEC with NSEC3](https://developers.cloudflare.com/dns/dnssec/enable-nsec3/) for details.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dnssec_dnssec_use_nsec3

If true, enables the use of NSEC3 together with DNSSEC on the zone.
Combined with setting dnssec_presigned to true, this enables the use of
NSEC3 records when transferring in from an external provider.
If dnssec_presigned is instead set to false (default), NSEC3 records will be
generated and signed at request time.

See [DNSSEC with NSEC3](https://developers.cloudflare.com/dns/dnssec/enable-nsec3/) for details.

```yaml
{"description": "If true, enables the use of NSEC3 together with DNSSEC on the zone.\nCombined with setting dnssec_presigned to true, this enables the use of\nNSEC3 records when transferring in from an external provider.\nIf dnssec_presigned is instead set to false (default), NSEC3 records will be\ngenerated and signed at request time.\n\nSee [DNSSEC with NSEC3](https://developers.cloudflare.com/dns/dnssec/enable-nsec3/) for details.", "type": "boolean", "example": false, "x-auditable": true}
```
