<h1 align="center">DecBrute</h1>

!It is just a proof of concept, do not use for malicious purposes! 

DecBrute is a C2 (Command and Control Server) which provide target to suffer a decentralized brute force attacks.

Computers acting as part of a botnet are victims of an XSS attack, this JS code injected on any vulnerable web app that makes the browser to request the username and password from DecBrute server, and then test those credentials on the against.

- Bypass Ratelimit
- Bypass IP block

![decbrute](https://user-images.githubusercontent.com/62824857/137057973-be675382-1567-4a28-a2d1-41835b7fc5a8.png)

```
$ go run main.go -h
 
COMMAND  DESCRIPTION      DEFAULT CONF  REQUIRED
-------  -----------      ------------  --------
-h       Help menu.       false         NO
-H       Local Host.      0.0.0.0       NO
-P       Local Port.      4444          NO
-u       User.            admin         YES
-l       Wordlist.        NO            YES
-e       Enable service.  false         YES
-v       Verbose Mode.    false         NO
exit status 1
```
