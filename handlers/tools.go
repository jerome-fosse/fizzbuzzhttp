package handlers

import (
	"errors"

	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
)

func parseParamToInt(url *url.URL, name string, defval int) int {
	val, err := parseParam(url, name)
	if err != nil {
		logrus.Debugf("Query parameter %s not found. Value set to default %d", name, defval)
		return defval
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		logrus.Debugf("Query parameter %s not a number. Value set to default %d", name, defval)
		return defval
	}

	return i
}

func parseParamToString(url *url.URL, name string, defval string) string {
	val, err := parseParam(url, name)
	if err != nil {
		logrus.Debugf("Query parameter %s not found. Value set to default %s", name, defval)
		return defval
	}
	return val
}

func parseParam(url *url.URL, name string) (string, error) {
	params, ok := url.Query()[name]
	if !ok || len(params) < 1 {
		return "", errors.New("not found")
	}

	return params[0], nil
}
