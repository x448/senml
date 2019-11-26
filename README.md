# Disco SenML is cisco/senml with bugfixes

## SenML
[RFC 8428 Sensor Measurement Lists (SenML)](https://tools.ietf.org/html/rfc8428) defines a format for representing simple sensor measurements and device parameters in Sensor Measurement Lists.  

## Disco SenML
Disco SenML is a fork of [cisco/senml](https://github.com/cisco/senml) that resolves various issues.  Programs are 4 MB smaller, CBOR representation no longer violates RFC 8428, and all unit tests pass. 

All features are the same as cisco/senml except MessagePack support is removed.

Disco SenML was created on Nov 23, 2019 using cisco/senml (4d43ea8) dated Oct 10, 2019.

__Special thanks to Cullen Jennings__ at Cisco Systems for his work on RFC 8428 and creating cisco/senml. He did the heavy lifting in those, so changes to his project were trivial by comparison. __And also to Faye Amacker__ for adding requested features to her [CBOR library](https://github.com/fxamacker/cbor) that made the changes easy.

## Release Notes
Disco SenML initial release on Nov 24, 2019.

This project fixes open issues in cisco/senml (4d43ea8, Oct 10, 2019):
* __cisco/senml #2 (2016)__ "CBOR does not encode or decode numeric field names". <-- RFC 8428 violation.
* __cisco/senml #18 (2017)__ "Base Value and Base Sum missing from the model.
* __cisco/senml #22 (2019)__ "CBOR support uses go-codec which adds bloat to the binary.

There are no changes to core cisco/senml features except removal of MessagePack.

MessagePack was removed because it:
* increased bloat (cisco/senml #22) and attack surface.
* isn't mentioned in SenML RFC 8428.
* prevented having a [CBOR library](https://github.com/fxamacker/cbor) (fxamacker/cbor) as the only external dependency.

Changes to cisco/senml (4d43ea8, Oct 10, 2019):
* Compiled programs are each 4 MB smaller (senmlCat and senmlServer).
* CBOR representation uses integers for labels, so it no longer violates SenML RFC 8428.
* Missing Base Value and Base Sum are added to the model.
* Fixed bad test data in unit tests and added new CBOR test using example from
  SenML RFC 8428 so all unit tests pass.
* Removed MessagePack feature for reasons cited in README.md.
* Replaced ugorji/go with fxamacker/cbor.
* Use Go modules and have at least one tagged release.
* Require Go 1.12
* Added name to LICENSE.
* Updated README.md with new name "Disco SenML"

## Limitations and Requirements
Known limitations:
* __Go 1.12__: Go 1.12 or newer is required.
* __Security Audit__: I didn't conduct a security audit of cisco/senml or this project.  I just tackled the most obvious low-hanging fruit while helping a CBOR library find non-hobby projects for program size comparisons.
* __Code Review and Refactoring__:  Only minimal changes were made to resolve cisco/senml issues 2, 18 and 22.  At a glance, code review and refactoring is highly recommended.

Possible limitations (not sure):
* __CBOR Tags (major type 6)__:  this might be a non-issue because I didn't see any used in cisco/senml.  If CBOR tags are present in SenML, they'll be ignored until this project upgrades fxamacker/cbor to v1.4 or v1.5 (when released).

# senmlCat
Tool to convert SenML between formats and act as gateway server to other services

# usage

## convert JSON SenML to XML 
senmlCat -json -i data.json > data.xml

## convert JSON SenML to CBOR
senmlCat.go -ijson -cbor data.json > data.cbor 

## convert to Excel spreadsheet CSV file
senmlCat -expand -ijsons -csv -print data.json > foo.csv

Note that this moves times to excel times that are days since 1900

## listen for posts of SenML in JSON and send to influxdb

This listens on port 880 then writes to an influx instance at localhost where to
the database called "junk"

The -expand is needed to expand base values into each line of the Line Protocol

senmlCat -ijsons -http 8880 -expand -linp -print -post http://localhost:8086/write?db=junk
