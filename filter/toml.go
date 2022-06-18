package filter

import (
	"io"

	"github.com/BurntSushi/toml"
)

func Load(source io.Reader) (Filter, error) {
	dec := toml.NewDecoder(source)
	var filter Filter
	_, err := dec.Decode(&filter)
	return filter, err
}

func (filter Filter) Save(target io.Writer) error {
	enc := toml.NewEncoder(target)
	return enc.Encode(filter)
}
