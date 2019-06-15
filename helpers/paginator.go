package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
)

const defaultItemsPerPage = 15

type Paginator struct {
	start        int
	end          int
	itemsPerPage int
	data         []byte
}

type PaginatorIface interface {
	Paginate(page string) ([]map[string]interface{}, error)
}

func (p *Paginator) Paginate(page string) ([]map[string]interface{}, error) {
	if page == "" {
		page = "1"
	}

	pg, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}

	p.start = (pg - 1) * p.itemsPerPage
	p.end = p.start + p.itemsPerPage

	var decoded interface{}
	if err = json.NewDecoder(bytes.NewReader(p.data)).Decode(&decoded); err != nil {
		return nil, err
	}

	linput := ifaceLen(decoded)
	if linput == 0 {
		return nil, errors.New("data is not an array")
	}

	if p.start > linput {
		return make([]map[string]interface{}, 0), nil
	} else {
		if p.end > linput {
			p.end = linput
		}
	}

	s := reflect.ValueOf(decoded)
	var output []map[string]interface{}
	if p.end < p.itemsPerPage {
		output = make([]map[string]interface{}, p.end)
	} else {
		output = make([]map[string]interface{}, p.itemsPerPage)
	}

	j := 0
	for i := p.start; i < p.end; i++ {
		if v, ok := s.Index(i).Interface().(map[string]interface{}); !ok {
			return nil, errors.New("not interface")
		} else {
			b, _ := json.Marshal(v)
			_ = json.Unmarshal(b, &output[j])
			j++
		}
	}

	return output, nil
}

func isarraypointers (d interface{}) bool {
	t := reflect.TypeOf(d).Kind()
	return t == reflect.Slice
}

func length (d interface{}) int {
	p := reflect.ValueOf(d).Len()
	return p
}

func ifaceLen (d interface{}) int {
	if isarraypointers(d) {
		return length(d)
	}
	return 0
}

func NewPaginator(data []byte, itemsPerPage ...int) PaginatorIface {
	cfg := &Paginator{
		data: data,
	}

	if itemsPerPage == nil {
		cfg.itemsPerPage = defaultItemsPerPage
	}

	return cfg
}
