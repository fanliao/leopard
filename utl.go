package orm

import (
    "reflect"
)

//// BKDR Hash Function
//unsigned int BKDRHash(char *str)
//{
//    unsigned int seed = 131; // 31 131 1313 13131 131313 etc..
//    unsigned int hash = 0;

//    while (*str)
//    {
//        hash = hash * seed + (*str++);
//    }

//    return (hash & 0x7FFFFFFF);
//}

func BKDRHash(s string) uint {
	var seed uint32 = 131
	var hash uint32 = 0

	cs := []int(s)
	for _, c = range cs {
		hash = hash*seed + c
	}

	return (hash & 0x7FFFFFFF)
}

//如果对象是一个pointer,则返回被指向对象的类型，如果是slice，返回slice元素的类型，否则出错
func GetStructType(obj interface{}) reflect.Type, error {
    return nil, nil
}

