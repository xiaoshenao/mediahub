package redis

import "strings"

const ServicePrefix = "shorturl_"

func GetKey(key string, parts ...string) string {
	key = ServicePrefix + key
	if len(parts) == 0 {
		return key
	}
	key += "_" + strings.Join(parts, "_")
	return key
}
