package mongobson

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	mongobson "go.mongodb.org/mongo-driver/bson"
)

type MongoBsonSerializer struct{}

func (m MongoBsonSerializer) TimePrecision() time.Duration {
	return time.Millisecond
}

func (m MongoBsonSerializer) Marshal(o interface{}) ([]byte, error) {
	return mongobson.Marshal(o)
}

func (m MongoBsonSerializer) Unmarshal(d []byte, o interface{}) error {
	return mongobson.Unmarshal(d, o)
}

func NewMongoBSONSerializer() goserbench.Serializer {
	return MongoBsonSerializer{}
}
