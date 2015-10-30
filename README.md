# go-cmd-h
h is a unix line filter that makes large integers readable by inserting k,m,g,t,p... before groups of 3 digits.
Usage:
	lvd$ cmd_that_generates_text --with-large=numbers | h | less
E.g.
	echo 123456789 | h
Produces: 123m456k789
