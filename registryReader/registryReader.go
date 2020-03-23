package registryReader

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/windows/registry"
)

func GetStringFromLocalMachine(path string, key string) string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	s, _, err := k.GetStringValue(key)
	if err != nil {
		log.Fatal(err)
	}
	return s
}
