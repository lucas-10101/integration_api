package services

import (
	"net/url"
	"strconv"
	"time"
)

type ParamName string

const (
	PARAM_ASSET_ID   ParamName = "asset_id"
	PARAM_START_DATE ParamName = "start_date"
	PARAM_END_DATE   ParamName = "end_date"
	PARAM_PAGE       ParamName = "page"
	PARAM_PAGE_SIZE  ParamName = "page_size"
)

var (
	DefaultPageSize  = 100
	DefaultPage      = 1
	DefaultStartDate = time.Now()
	DefaultEndDate   = time.Now().AddDate(0, -6, 0)
	DefaultAssetID   = -1
	DefaultDeviceID  = -1
	DefaultLanguage  = LANG_EN
)

type UrlQueryParser struct {
	Values url.Values
}

// *************************
// *************************
// *************************
// Specific getters
// *************************
// *************************
// *************************

func (p *UrlQueryParser) GetPage() int {
	if value, ok := p.Values[string(PARAM_PAGE)]; ok && len(value) > 0 {
		if page, err := strconv.Atoi(value[0]); err == nil && page > 0 {
			return page
		}
	}
	return DefaultPage
}

func (p *UrlQueryParser) GetPageSize() int {
	if value, ok := p.Values[string(PARAM_PAGE_SIZE)]; ok && len(value) > 0 {
		if pageSize, err := strconv.Atoi(value[0]); err == nil && pageSize > 1 {
			return pageSize
		}
	}
	return DefaultPageSize
}

func (p *UrlQueryParser) GetStartDate() time.Time {
	if value, ok := p.Values[string(PARAM_START_DATE)]; ok && len(value) > 0 {
		if startDate, err := time.Parse(time.RFC3339, value[0]); err == nil {
			return startDate
		}
	}
	return DefaultStartDate
}

func (p *UrlQueryParser) GetEndDate() time.Time {
	if value, ok := p.Values[string(PARAM_END_DATE)]; ok && len(value) > 0 {
		if endDate, err := time.Parse(time.RFC3339, value[0]); err == nil {
			return endDate
		}
	}
	return DefaultEndDate
}

func (p *UrlQueryParser) GetDateRange() (time.Time, time.Time) {
	startDate := p.GetStartDate()
	endDate := p.GetEndDate()

	if endDate.Before(startDate) {
		return startDate, endDate
	}
	return DefaultStartDate, DefaultEndDate
}

// *************************
// *************************
// *************************
// Generic getters
// *************************
// *************************
// *************************

func (p *UrlQueryParser) GetString(param ParamName) string {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		return value[0]
	}
	return ""
}

func (p *UrlQueryParser) GetBool(param ParamName) bool {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		if boolValue, err := strconv.ParseBool(value[0]); err == nil {
			return boolValue
		}
	}
	return false
}

func (p *UrlQueryParser) GetFloat64(param ParamName) float64 {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		if floatValue, err := strconv.ParseFloat(value[0], 64); err == nil {
			return floatValue
		}
	}
	return 0.0
}

func (p *UrlQueryParser) GetFloat32(param ParamName) float32 {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		if floatValue, err := strconv.ParseFloat(value[0], 32); err == nil {
			return float32(floatValue)
		}
	}
	return 0.0
}

func (p *UrlQueryParser) GetUint64(param ParamName) uint64 {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		if uint64Value, err := strconv.ParseUint(value[0], 10, 64); err == nil {
			return uint64Value
		}
	}
	return 0
}

func (p *UrlQueryParser) GetUint32(param ParamName) uint32 {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		if uint32Value, err := strconv.ParseUint(value[0], 10, 32); err == nil {
			return uint32(uint32Value)
		}
	}
	return 0
}

func (p *UrlQueryParser) GetUint16(param ParamName) uint16 {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		if uint16Value, err := strconv.ParseUint(value[0], 10, 16); err == nil {
			return uint16(uint16Value)
		}
	}
	return 0
}

func (p *UrlQueryParser) GetUint8(param ParamName) uint8 {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		if uint8Value, err := strconv.ParseUint(value[0], 10, 8); err == nil {
			return uint8(uint8Value)
		}
	}
	return 0
}

func (p *UrlQueryParser) GetUint(param ParamName) uint {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		if uintValue, err := strconv.ParseUint(value[0], 10, 32); err == nil {
			return uint(uintValue)
		}
	}
	return 0
}

func (p *UrlQueryParser) GetInt64(param ParamName) int64 {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		if int64Value, err := strconv.ParseInt(value[0], 10, 64); err == nil {
			return int64Value
		}
	}
	return 0
}

func (p *UrlQueryParser) GetInt32(param ParamName) int32 {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		if int32Value, err := strconv.ParseInt(value[0], 10, 32); err == nil {
			return int32(int32Value)
		}
	}
	return 0
}

func (p *UrlQueryParser) GetInt16(param ParamName) int16 {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		if int16Value, err := strconv.ParseInt(value[0], 10, 16); err == nil {
			return int16(int16Value)
		}
	}
	return 0
}

func (p *UrlQueryParser) GetInt8(param ParamName) int8 {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		if int8Value, err := strconv.ParseInt(value[0], 10, 8); err == nil {
			return int8(int8Value)
		}
	}
	return 0
}

func (p *UrlQueryParser) GetInt(param ParamName) int {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		if intValue, err := strconv.Atoi(value[0]); err == nil {
			return intValue
		}
	}
	return 0
}

func (p *UrlQueryParser) GetIntList(param ParamName) []int {

	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		intList := make([]int, 0, len(value))
		for _, v := range value {
			if intValue, err := strconv.Atoi(v); err == nil {
				intList = append(intList, intValue)
			}
		}
		return intList
	}
	return []int{}
}

func (p *UrlQueryParser) GetUintList(param ParamName) []uint {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		uintList := make([]uint, 0, len(value))
		for _, v := range value {
			if uintValue, err := strconv.ParseUint(v, 10, 32); err == nil {
				uintList = append(uintList, uint(uintValue))
			}
		}
		return uintList
	}
	return []uint{}
}

func (p *UrlQueryParser) GetStringList(param ParamName) []string {
	if value, ok := p.Values[string(param)]; ok && len(value) > 0 {
		return value
	}
	return []string{}
}
