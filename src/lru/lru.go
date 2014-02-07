package main 

import (
	"fmt"
	"errors"
	"bufio"
	"strconv"
	"strings"
	"os"
	"sort"
)

type dllist struct{
	head, tail *listItem
	size int
}

func (this *dllist) AddFirst(li *listItem){
	this.size ++
	if this.head == nil{
		this.head = li
		this.tail = li
		
		li.prev = nil
		li.next = nil
		return
	}

	this.head.prev = li
	li.next = this.head
	li.prev = nil
	this.head = li
}

func (this *dllist) remove(li *listItem){
	if li == nil{
		return
	}
	
	if li.prev == nil{
		this.head = li.next
		if this.head != nil{
			this.head.prev = nil
		}
	} 
	
	if li.next == nil{
		this.tail = li.prev
		if this.tail != nil{
			this.tail.next = nil
		}
	} 
	
	if li.prev != nil && li.next != nil{
		li.prev.next = li.next
		li.next.prev = li.prev
		
		li.prev = nil
		li.next = nil
	}
	
	this.size--
}


func (this *dllist) removeLast() *listItem{
	if this.tail == nil{
		return nil
	}
	
	this.size--
	
	item := this.tail 
	prev := item.prev
	if prev == nil{
		this.head = nil
		this.tail = nil
		return item
	}
	
	prev.next = nil
	this.tail = prev

	item.prev = nil
	return item
}

type listItem struct{
	k, v string
	prev *listItem
	next *listItem
}

func testDl(){
	list := &dllist{}
	a := &listItem{k:"a", v:"1"}
	b := &listItem{k:"b", v:"2"}
	c := &listItem{k:"c", v:"3"}
	
	list.AddFirst(a)
	list.AddFirst(b)
	list.AddFirst(c)
	
	dumpList := func(){
		fmt.Println("size", list.size)
		for i := list.head; i != nil; i = i.next{
			fmt.Println("--", i)
		}
		fmt.Println()
	}
	
	dumpList()
	list.remove(a)
	dumpList()
	list.remove(c)
	dumpList()
	list.remove(b)
	dumpList()
	
	fmt.Println("Second batch of testing")
	list.AddFirst(a)
	list.AddFirst(b)
	list.AddFirst(c)
	dumpList()
	list.remove(b)
	dumpList()
}

type cache struct{
	list dllist
	lru map[string]*listItem
	size int
}

func (this *cache) bound(size int){
	for this.list.size > size{
		item := this.list.removeLast()
		if item != nil{
			delete(this.lru, item.k)
		}
	}
	this.size = size
}

func (this *cache) get(key string) (string, error){
	if item, ok := this.lru[key]; ok{
		this.list.remove(item)
		this.list.AddFirst(item)
		
		return item.v, nil
	}
	
	return "", errors.New("Not found")
}

func (this *cache) peek(key string) (string, error){
	if item, ok := this.lru[key]; ok{
		return item.v, nil
	}
	
	return "", errors.New("Not found")
}

func (this *cache) put(key, value string){
	if item, ok := this.lru[key]; ok{
		item.v = value
		this.list.remove(item)
		this.list.AddFirst(item)
		this.lru[key] = item
	} else{
		if this.list.size >= this.size{
			item = this.list.removeLast()
			if item != nil{
				delete(this.lru, item.k)
			}
		}
		item = &listItem{k: key, v:value}
		this.list.AddFirst(item)
		this.lru[key] = item
	}
}

func (this *cache) dump(){
	if this.list.size == 0{
		return
	}
	
	keys := make([]string, this.list.size)
	idx := 0
	for i:= this.list.head; i != nil; i = i.next{
		keys[idx] = i.k
		idx++
	}
	sort.Strings(keys)
	
	for _, k := range keys{
		fmt.Println(fmt.Sprintf("%s %s", k, this.lru[k].v))
	}
}

func main() {
	var num int
	fmt.Scanf("%d", &num)
	var command string
	lrucache := &cache{list: dllist{}, lru: make(map[string]*listItem)}
	inr := bufio.NewReader(os.Stdin)
	for i := 0; i < num; i++{
		bytes,_,_ := inr.ReadLine()
		command = string(bytes)
		tokens := strings.Split(command, " ")
		switch tokens[0]{
		case "BOUND":
			size,_ := strconv.Atoi(tokens[1])
			lrucache.bound(size)
		case "SET":
			lrucache.put(tokens[1], tokens[2])
		case "GET":
			v, err := lrucache.get(tokens[1])
			if err == nil{
				fmt.Println(v)
			} else{
				fmt.Println("NULL")
			}
		case "PEEK":
			v, err := lrucache.peek(tokens[1])
			if err == nil{
				fmt.Println(v)
			} else{
				fmt.Println("NULL")
			}
		case "DUMP":
			lrucache.dump()
		}
	}
	
}