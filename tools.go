package main

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

func parseParamToInt(url *url.URL, name string) (int, error) {
	s, err := parseParam(url, name)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("%s is not a numeric value", s)
	}

	return i, nil
}

func parseParam(url *url.URL, name string) (string, error) {
	params, ok := url.Query()[name]
	if !ok || len(params) < 1 {
		return "", errors.New("not found")
	}

	return params[0], nil
}
