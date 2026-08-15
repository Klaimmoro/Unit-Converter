package src

import (
	"fmt"
	"unit-converter/src/kernel"
)

type LenConverter struct {
	Value    float64
	UnitFrom string
	UnitTo   string
}

func (lc *LenConverter) ConvertToMeters(value float64, from string) (float64, error) {
	switch from {
	case "mm":
		return value / 1000.0, nil
	case "cm":
		return value / 100.0, nil
	case "m":
		return value, nil
	case "km":
		return value * 1000.0, nil
	case "in":
		return value * 0.0254, nil
	case "ft":
		return value * 0.3048, nil
	case "yd":
		return value * 0.9144, nil
	case "mi":
		return value * 1609.344, nil
	default:
		return 0.0, fmt.Errorf("Unknown length unit: %s\n", from)
	}
}

func (lc *LenConverter) MetersToUnit(value float64, to string) (float64, error) {
	switch to {
	case "mm":
		return value * 1000.0, nil
	case "cm":
		return value * 100.0, nil
	case "m":
		return value, nil
	case "km":
		return value / 1000.0, nil
	case "in":
		return value / 0.0254, nil
	case "ft":
		return value / 0.3048, nil
	case "yd":
		return value / 0.9144, nil
	case "mi":
		return value / 1609.344, nil
	default:
		return 0.0, fmt.Errorf("Unknown length unit: %s\n", to)
	}
}

type WeightConverter struct {
	Value    float64
	UnitFrom string
	UnitTo   string
}

func (wc *WeightConverter) ConvertToKg(value float64, from string) (float64, error) {
	switch from {
	case "kg":
		return value, nil
	case "g":
		return value / 1000.0, nil
	case "lb":
		return value * 0.45359237, nil
	case "oz":
		return value * 0.0283495231, nil
	case "t":
		return value * 1000.0, nil
	default:
		return 0.0, fmt.Errorf("Unknown weight unit: %s\n", from)
	}
}

func (wc *WeightConverter) KgToUnit(value float64, to string) (float64, error) {
	switch to {
	case "kg":
		return value, nil
	case "g":
		return value * 1000.0, nil
	case "lb":
		return value * 0.45359237, nil
	case "oz":
		return value * 0.0283495231, nil
	case "t":
		return value / 1000.0, nil
	default:
		return 0.0, fmt.Errorf("Unknown weight unit: %s\n", to)
	}
}

type TempConverter struct {
	Value    float64
	UnitFrom string
	UnitTo   string
}

func (tc *TempConverter) ConvertToCel(value float64, from string) (float64, error) {
	switch from {
	case "c":
		return value, nil
	case "f":
		return (value - 32) * 5.0 / 9.0, nil
	case "k":
		return value - 273.15, nil
	default:
		return 0.0, fmt.Errorf("Unknown temperature unit: %s\n", from)
	}
}

func (tc *TempConverter) CelToUnit(value float64, to string) (float64, error) {
	switch to {
	case "c":
		return value, nil
	case "f":
		return (value * 9.0 / 5.0) + 32.0, nil
	case "k":
		return value + 273.15, nil
	default:
		return 0.0, fmt.Errorf("Unknown temperature unit: %s\n", to)
	}
}

type UnitConverter struct {
	LenConverter    *LenConverter
	WeightConverter *WeightConverter
	TempConverter   *TempConverter
}

func (uc *UnitConverter) Convert(request_to_convert *kernel.ConvertRequest) (float64, error) {
	switch request_to_convert.Category {
	case "length":
		meters, err := uc.LenConverter.ConvertToMeters(request_to_convert.Value, request_to_convert.UnitFrom)
		if err != nil {
			return 0.0, err
		}
		res, err := uc.LenConverter.MetersToUnit(meters, request_to_convert.UnitTo)
		if err != nil {
			return 0.0, err
		}
		return res, nil
	case "weight":
		kg, err := uc.WeightConverter.ConvertToKg(request_to_convert.Value, request_to_convert.UnitFrom)
		if err != nil {
			return 0.0, err
		}
		res, err := uc.WeightConverter.KgToUnit(kg, request_to_convert.UnitTo)
		if err != nil {
			return 0.0, err
		}
		return res, nil
	case "temperature":
		kg, err := uc.TempConverter.ConvertToCel(request_to_convert.Value, request_to_convert.UnitFrom)
		if err != nil {
			return 0.0, err
		}
		res, err := uc.TempConverter.CelToUnit(kg, request_to_convert.UnitTo)
		if err != nil {
			return 0.0, err
		}
		return res, nil
	default:
		return 0.0, fmt.Errorf("Unknown category: %s\n", request_to_convert.Category)
	}
}
