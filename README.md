# Rana Ora package test
CTRL+C signal causes program to crash on Linux and OS X

Step to reproduce the issue:

1)build the test program (https://github.com/jusongchen/rana_ora_test)


2) run the programm , pass in the connection string as the first parameter. e.g:

	rana_ora_test scott/tiger@//hostname/oracle_sid

3) Enter CTRL+C after the prompt


4) CTRL+C signal won't cause any issue before a DB operation. However, after a DB operation, it will crash the programm with message

"llegal instruction: 4" -- on Mac OS X

or "Trace/breakpoint trap" --on Ubuntu Linux.


