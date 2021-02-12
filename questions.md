## Почему на обычном мьютексе вылетает а на таком нет и чем они отличаются
var mu sync.RWMutex
mu.RLock()
mu.RUnlock()