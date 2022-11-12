#!/usr/bin/python
#package main

from bandal import Lexer, Parser, IL

code = """
fn main(a int, b int) (int, int)
	c int = 1
	d int = 2
	print(c)
	print(d)
	return a, b
end

main(1, 2)
"""

l = Lexer().lex(code)
# print('\n'.join(map(str, l)))

p = Parser().parse(l)
# print(p)

i = IL().inter(p)
# print(i)