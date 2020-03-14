# Go Router

This module provides some common abstractions accross select router operating systems. 

## Functionality

The primary focus of the module is to abstract the non-volitle RAM (NVRAM) APIs and related settings, specifically as they pertain to headless reading and setting of networking settings. The same NVRAM APIs enable a range of dynamic runtime adjustments to the routers.

 Most "personal" router OSes run as a read-only image, so settings typically need to be persisted via NVRAM.

## Supported OSes

Support for the following is planned:

* OpenWRT
* DD-WRT
* RavPower OS
