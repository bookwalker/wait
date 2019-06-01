# /wait

Simple utility to test the availability of TCP connections.
Intended for environments, where a certain startup order is important (eg. multi-container Docker applications).

## Usage

`Wait` tries to connect to a TCP connection within the given `timeout`.
The scripts exits with `0` on success, `1` otherwise.

    Usage of ./wait:   
      -address string     a hostname or ip-address including port. Eg. 127.0.0.1:80
      -timeout int        a timeout in seconds or 0 for no timeout (default 10)


**Examples**

    wait -address <hostname-or-ip>:<port> -timeout <timeout-in-seconds>
    wait -address nginx:443 -timeout 10
