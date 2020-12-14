package main

import "testing"

//测试store
func TestStore(t *testing.T) {
	monster := &Monster{
		Name:  "红孩子",
		Age:   10,
		Skill: "吐火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() 错误了，希望是%v，结果为%v", true, res)
	}
	t.Logf("monster.Store() success")
}

func TestRestore(t *testing.T) {
	var monster Monster
	res := monster.Restore()
	if !res {
		t.Fatalf("monster.Store() 错误了，希望是%v，结果为%v", true, res)
	}
	if monster.Name != "红孩子" {
		t.Fatalf("monster.Store() 错误了，希望是%v，结果为%v", "红孩子",
			monster.Name)
	}
	t.Logf("monster.Restore() success")
}
