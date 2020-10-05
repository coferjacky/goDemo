package main

import "../packageDemo"

/*
编译失败，主要因为结构体的成员没有开头项小写导致没有导出
var _=backpack.T{a:1,b:2}
var _=backpack.T{1,2}
*/

var _ = packageDemo.T{Ai: 1, Bi: 2}
var _ = packageDemo.T{1, 2}
