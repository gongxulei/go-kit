/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/5
 * +----------------------------------------------------------------------
 * |Time: 3:46 下午
 * +----------------------------------------------------------------------
 */

package main

import (
	"sync"
)

// M
type M struct {
	Map map[string]string
}

// Set ...
func (m *M) Set(key, value string) {
	m.Map[key] = value
}

// Get ...
func (m *M) Get(key string) string {
	return m.Map[key]
}

// M1 加锁
type M1 struct {
	Map  map[string]string
	lock sync.RWMutex // 加锁
}

// Set ...
func (m *M1) Set(key, value string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.Map[key] = value
}

// Get ...
func (m *M1) Get(key string) string {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.Map[key]
}

// M2 线程安全map
type M2 struct {
	// Map    map[string]string
	Map sync.Map
}

// Set ...
func (m *M2) Set(key, value string) {
	m.Map.Store(key, value)
}

// Get ...
func (m *M2) Get(key string) string {
	inter, _ := m.Map.Load(key)
	str, _ := inter.(string)
	return str
}

// M3 加锁
type M3 struct {
	Map  map[string]string
	lock sync.Mutex // 加锁
}

// Set ...
func (m *M3) Set(key, value string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.Map[key] = value
}

// Get ...
func (m *M3) Get(key string) string {
	return m.Map[key]
}
