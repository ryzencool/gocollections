package chainmap

type CmInterface interface {
	Len()
}

// ChainMap is map tools
type ChainMap struct {
	Maps map[interface{}]interface{}
}

func (c *ChainMap) Len() {

}

// NewChainMap go
func NewChainMap() *ChainMap {
	return &ChainMap{make(map[interface{}]interface{})}
}

// InitMaps go
func (c *ChainMap) InitMaps(maps ...interface{}) {
	// for _, v := range maps {
	// 	if reflect.Kind
	// }
}
