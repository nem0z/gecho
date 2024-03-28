package payloads

import (
	"bytes"
	"encoding/binary"
)

type Version struct {
	version  uint32
	features uint32
	relay    bool
}

func NewVersion(version, features uint32, relay bool) *Version {
	return &Version{version: version, features: features, relay: relay}
}

func ParseVersion(data []byte) *Version {
	version := binary.BigEndian.Uint32(data[0:4])
	features := binary.BigEndian.Uint32(data[4:8])
	relay := data[9] == 1

	return NewVersion(version, features, relay)
}

func (v *Version) ToByte() ([]byte, error) {
	version := make([]byte, 4)
	features := make([]byte, 4)
	binary.BigEndian.PutUint32(version, uint32(v.version))
	binary.BigEndian.PutUint32(features, uint32(v.features))

	var relay byte
	if v.relay {
		relay = 1
	}

	return bytes.Join([][]byte{version, features, []byte{relay}}, []byte{}), nil
}
