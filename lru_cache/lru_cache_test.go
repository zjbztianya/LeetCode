package lru_cache

import "testing"

func TestLRUCache(t *testing.T) {
	lru:=Constructor(2)
	lru.Put(1,1)
	lru.Put(2,2)
	if got:=lru.Get(1); got!=1{
		t.Errorf("Get(%d)=%d want:%d",1,got,1)
	}
	lru.Put(3,3)
	if got:=lru.Get(2); got!=-1{
		t.Errorf("Get(%d)=%d want:%d",2,got,-1)
	}
	lru.Put(4,4)
	if got:=lru.Get(1); got!=-1{
		t.Errorf("Get(%d)=%d want:%d",1,got,-1)
	}
	if got:=lru.Get(3); got!=3{
		t.Errorf("Get(%d)=%d want:%d",3,got,3)
	}
	if got:=lru.Get(4); got!=4{
		t.Errorf("Get(%d)=%d want:%d",4,got,4)
	}
}