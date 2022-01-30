package syncutil

import "sync"

type Counter struct {
	Name string
	// エクスポートされないフィールドがある場合は空行を入れることが多い
	// ミューテックスを利用する際は対象となるフィールドらの先頭で定義することが多い
	m sync.RWMutex
	count int
}

func (c *Counter) Increment() int {
	c.m.Lock()
	defer c.m.Unlock()
	c.count++
	return c.count
}

func (c *Counter) View() int{
	c.m.RLock()
	defer c.m.RUnlock()
	return c.Counter
}

c := &syncutil.Counter {
	Name: "Access,"
}

fmt.Println(c.Increment())
fmt.Println(c.View())