package utils

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
	"time"
)

type DataItem struct {
	Time int64
	Key  string
	Data interface{}
}

func (data *DataItem) GetValue() interface{} {
	return data.Data
}

func CreateDataItem(key string, data interface{}) *DataItem {
	item := new(DataItem)
	item.Time = time.Now().Unix()
	item.Key = key
	item.Data = data
	return item
}

func CreateDataItemAutoKey(data interface{}) *DataItem {
	item := new(DataItem)
	item.Time = time.Now().Unix()
	item.Key = fmt.Sprintf("%d", time.Now().UnixNano())
	item.Data = data
	return item
}

var tFunc func(val interface{}) bool

type LRUQueue struct {
	list    *list.List
	datamap map[string]*list.Element
	Mutex   sync.RWMutex
	max     int
}

func NewQueue(maxlength int) *LRUQueue {
	que := new(LRUQueue)
	que.list = list.New()
	que.datamap = make(map[string]*list.Element)
	if maxlength <= 0 {
		maxlength = 1
	}
	que.max = maxlength
	return que
}

func (que *LRUQueue) AddElement(item *DataItem) error {
	if item == nil {
		return errors.New("data nil")
	}
	que.Mutex.Lock()
	defer que.Mutex.Unlock()
	if ele, ok := que.datamap[item.Key]; ok {
		que.list.Remove(ele)
		delete(que.datamap, item.Key)
		if que.list.Len() < que.max {
			newEle := que.list.PushBack(item)
			que.datamap[item.Key] = newEle
			return nil
		}
		return errors.New("add fail")
	} else {
		if que.list.Len() == que.max {
			oldValue := que.list.Remove(que.list.Front())
			if oldValue != nil {
				oldItem := oldValue.(*DataItem)
				delete(que.datamap, oldItem.Key)
			}
		}
		if que.list.Len() < que.max {
			newEle := que.list.PushBack(item)
			que.datamap[item.Key] = newEle
			return nil
		}
		return errors.New("add fail")
	}
}

func (que *LRUQueue) AddElementFast(item *DataItem) error {
	if item == nil {
		return errors.New("data nil")
	}
	if ele, ok := que.datamap[item.Key]; ok {
		que.list.Remove(ele)
		delete(que.datamap, item.Key)
		if que.list.Len() < que.max {
			newEle := que.list.PushBack(item)
			que.datamap[item.Key] = newEle
			return nil
		}
		return errors.New("add fail")
	} else {
		if que.list.Len() == que.max {
			oldValue := que.list.Remove(que.list.Front())
			if oldValue != nil {
				oldItem := oldValue.(*DataItem)
				delete(que.datamap, oldItem.Key)
			}
		}
		if que.list.Len() < que.max {
			newEle := que.list.PushBack(item)
			que.datamap[item.Key] = newEle
			return nil
		}
		return errors.New("add fail")
	}
}

func (que *LRUQueue) GetElement(key string) *DataItem {
	que.Mutex.RLock()
	defer que.Mutex.RUnlock()
	if ele, ok := que.datamap[key]; ok {
		return ele.Value.(*DataItem)
	} else {
		return nil
	}
}

func (que *LRUQueue) GetElementsList() (listeles []*DataItem) {
	que.Mutex.RLock()
	defer que.Mutex.RUnlock()
	for _, ele := range que.datamap {
		listeles = append(listeles, ele.Value.(*DataItem))
	}
	return listeles
}

func (que *LRUQueue) GetElementFast(key string) *DataItem {
	if ele, ok := que.datamap[key]; ok {
		return ele.Value.(*DataItem)
	} else {
		return nil
	}
}

func (que *LRUQueue) GetElementsListFast() (listeles []*DataItem) {
	for e := que.list.Front(); e != nil; e = e.Next() {
		listeles = append(listeles, e.Value.(*DataItem))
	}
	return listeles
}

func (que *LRUQueue) Size() int {
	que.Mutex.RLock()
	defer que.Mutex.RUnlock()
	return que.list.Len()
}

func (que *LRUQueue) SizeFast() int {
	return que.list.Len()
}

func (que *LRUQueue) DelElement(key string) error {
	if key == "" {
		return errors.New("delete error")
	}
	que.Mutex.Lock()
	defer que.Mutex.Unlock()
	if ele, ok := que.datamap[key]; ok {
		item := que.list.Remove(ele)
		if item != nil {
			delete(que.datamap, item.(*DataItem).Key)
			return nil
		} else {
			return errors.New("delete list error")
		}
	} else {
		return errors.New("no such key data")
	}
}

func (que *LRUQueue) UpdateElement(item *DataItem) error {
	if item == nil {
		return errors.New("data is error")
	} else {
		que.Mutex.Lock()
		defer que.Mutex.Unlock()
		if ele, ok := que.datamap[item.Key]; ok {
			newItem := que.list.Remove(ele)
			if newItem != nil {
				delete(que.datamap, item.Key)
				if que.list.Len() < que.max {
					newEle := que.list.PushBack(item)
					que.datamap[item.Key] = newEle
					return nil
				} else {
					return errors.New("list length greater than max")
				}
			} else {
				return errors.New("save list error")
			}
		}
		return errors.New("no find data")
	}
}
