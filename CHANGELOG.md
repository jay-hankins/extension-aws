# Changelog

## v2.2.19

- Update dependencies

## v2.2.18

- Update dependencies
- fixed enrichment for jvm instances

## v2.2.17

- ignore erroneous roleArns during initialization

## v2.2.16

- use discovery_kit_sdk for discoveries
- add `aws.zone.id` to ec2- and rds-instances
- added `aws.zone.id` to all enrichment rules

## v2.2.15

- Added `pprof` endpoints for debugging purposes
- Update dependencies
- Enrichment for kubernetes-statefulsets, -daemonsets, -nodes and -pods

## v2.2.14

- Possibility to exclude attributes from discovery

## v2.2.13

- Fix: `slice index out of range` error in discovery

## v2.2.12

- Fix: `DiscoveryDisabled*`-Properties are working again

## v2.2.11

- Make Discovery Intervals configurable
- Keep a copy of current targets in discoveries and call aws apis not in the context of the agent request.
- Allow parallel API calls using a configurable amount of worker threads via `STEADYBIT_EXTENSION_WORKER_THREADS`

## v2.2.10

- Add enrichment rules for kubernetes deployments
- Make targets to recieve EC2 data configurable via `STEADYBIT_EXTENSION_ENRICH_EC2_DATA_FOR_TARGET_TYPES`

## v2.2.9

- user service icons for actions

## v2.2.8

- collect EC2 instance state in discovery

## v2.2.7

- update dependencies

## v2.2.5

- migration to new unified steadybit actionIds and targetTypes
- added hint to aws account lookup of the agent in case of an error

## v2.2.4

- update dependencies

## v2.2.3

- replace RDS instance downtime attack with RDS instance stop attack

## v2.2.2

 - add RDS instance downtime attack
 - add RDS cluster failover attack
 - add RDS cluster discovery

## v2.2.1

- Add linux package build

## v2.2.0

- Prefix label attributes at ec2 target discovery with `aws-ec2.`.

## v2.1.1

- Support read only file system
- Update dependencies

## v2.1.0

- Support Readiness & Liveness probes (requires helm chart version >= 2.0.0)
- Refactored to use `action_kit_sdk` and thus use the extended rollback safety while having connection issues
- Added Lambda discovery & actions (requires new permissions)

## v2.0.0

- Renamed attack `ec2-instance.state` to `com.github.steadybit.extension_aws.ec2_instance.state`
- Added EC2-Instance discovery
- Added Zone-Discovery and Availability Zone Blackhole attack
- Added AWS FIS-Experiment discovery and AWS FIS-Experiment action

## v1.8.0

- Print build information on extension startup.

## v1.7.1

 - Add missing `kind` field to both actions.

## v1.7.0

 - Support creation of a TLS server through the environment variables `STEADYBIT_EXTENSION_TLS_SERVER_CERT` and `STEADYBIT_EXTENSION_TLS_SERVER_KEY`. Both environment variables must refer to files containing the certificate and key in PEM format.
 - Support mutual TLS through the environment variable `STEADYBIT_EXTENSION_TLS_CLIENT_CAS`. The environment must refer to a comma-separated list of files containing allowed clients' CA certificates in PEM format.

## v1.6.0

 - Support for AWS role assumption. This permits one extension instance from gathering data from multiple AWS accounts. To configure this, you must set the `STEADYBIT_EXTENSION_ASSUME_ROLES` environment variable to a comma-separated list of role ARNs. Example: `STEADYBIT_EXTENSION_ASSUME_ROLES='arn:aws:iam::1111111111:role/steadybit-extension-aws,arn:aws:iam::22222222:role/steadybit-extension-aws'`.

## v1.5.0

 - Support for the `STEADYBIT_LOG_FORMAT` env variable. When set to `json`, extensions will log JSON lines to stderr.

## v1.4.0

 - Restrict discovery execution to AWS agents to avoid common issues.
 - The log level can now be configured through the `STEADYBIT_LOG_LEVEL` environment variable.

## v1.3.0

 - Expose AWS RDS instance status in target table

## v1.2.0

 - Report AWS RDS instance status

## v1.1.0

 - EC2 instance state attacks, i.e., EC2 instance stop, reboot, hibernate and terminate.

## v1.0.0

 - Initial release
