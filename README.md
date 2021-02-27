# goa-snmp

Goa micro service for snmp
[![CircleCI](https://circleci.com/gh/rshade/goasnmp.svg?style=svg)](https://circleci.com/gh/rshade/goasnmp)

## Design

This Service will take an ip address or hostname, and snmp public/private, walk the tree and save
the mib/oid information for http retrieval.

`http://localhost/goa-snmp/device/{mame/ip}`
`http://localhost/goa-snmp/device/{name}/mibs`