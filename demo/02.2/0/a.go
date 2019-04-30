package a
 
type RedisStorage struct {
}

func (s *RedisStorage) Set(key, val string) error {
	//save to redis
	return nil
}

func (s *RedisStorage) Get(key string) (val string,  err error) {
	//get from redis
	return "", nil
}

type A struct {
	 storage *RedisStorage
}

func NewA(stor * RedisStorage) *A {
	return &A{storage: stor}
}

func (a *A) Save(key, val string)  error {
	return  a.storage.Set(key, val)
}

func (a *A) Get(key string) (string, error) {
	return  a.storage.Get(key)
}
