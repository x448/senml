# Disco SenML is cisco/senml with bugfixes
[![](https://github.com/x448/senml/workflows/ci/badge.svg)](https://github.com/x448/senml/blob/master/.github/workflows/ci.yml)
[![cover ≥80.7%](https://github.com/x448/senml/workflows/cover%20%E2%89%A580.7%25/badge.svg)](https://github.com/x448/senml/blob/master/.github/workflows/cover.yml)

## SenML
[RFC 8428 Sensor Measurement Lists (SenML)](https://tools.ietf.org/html/rfc8428) defines a format for representing simple sensor measurements and device parameters in Sensor Measurement Lists.  

## Disco SenML
Disco SenML is a fork of [cisco/senml](https://github.com/cisco/senml) that resolves various issues.  Programs are 4 MB smaller, CBOR representation no longer violates RFC 8428, and all unit tests pass. 

All features are the same as cisco/senml except MessagePack support is removed.

Disco SenML was created on Nov 23, 2019 using cisco/senml (4d43ea8) dated Oct 10, 2019.

__Special thanks__:

* Cullen Jennings for his work on RFC 8428 (SenML) and cisco/senml. He did the heavy lifting in those, so changes made by this project are trivial by comparison.
* Faye Amacker for adding requested features to her [CBOR library](https://github.com/fxamacker/cbor) that made this easy.

## Release Notes
Disco SenML initial release on Nov 24, 2019.

This project fixes open issues in cisco/senml (d5a3c66, Dec 11, 2019):

* __cisco/senml #2 (2016)__ "CBOR does not encode or decode numeric field names". <-- RFC 8428 violation.
* __cisco/senml #18 (2017)__ "Base Value and Base Sum missing from the model.
* __cisco/senml #22 (2019)__ "CBOR support uses go-codec which adds bloat to the binary.
* __cisco/senml #25 (2019)__ "cisco/senml does not pass unit tests"

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

## Limitations
Known limitations:

* There are no unit tests for senmlCat or senmlServer. Using senmlCat or senmlServer is not recommended.
* Security Audit: I didn't conduct a security audit of cisco/senml or this project.  A security audit is recommended.
* Code Review and Refactoring:  I didn't perform any code review or refactoring beyond the minimum changes required to resolve cisco/senml issues 2, 18, 22 and 25.  Code review and refactoring is recommended.

## Requirements

Go 1.12+ is required.

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

# License
Copyright 2019-present Montgomery Edwards⁴⁴⁸ (github.com/x448)  
Copyright 2016-2019 Cullen Jennings

x448/senml is licensed under the BSD 2-Clause "Simplified" License.  See [LICENSE](LICENSE) for the full license text.
