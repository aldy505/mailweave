# Testdata

This directory contains test data for the mailweave project.
The data is actually an actual DMARC and TLS-RPT report from a real mail server.
It is redacted in such way to protect the actual reports.

The domain of the report shall be `example.com`.
The successful IP address for DMARC and TLS-RPT sending shall be within the `192.0.2.0/24` network. Any other IP address shall be used as a failure.