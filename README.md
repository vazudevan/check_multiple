`multi-check` aims to function as a cross platform tool that perform multiple network level connectivity checks similar to monitoring-plugins.

The idea behind this tool is to be able to check for loss of network or service dependencies from a perspective of local environment of an application/server.  This is in contrast to remote monitoring polls.  This would help avoid monitoring blind spots such as inter-server connections / dependencies.

`multi-check` differs from monitoring-plugins by these aspects.

* No thresholds for checks
* Requires a config file to define / declare checks
* Only 2 output states (OK and PROBLEM)
