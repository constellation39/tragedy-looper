package validate

import (
	"fmt"
)

// ValidationResult 表示单个验证结果
type ValidationResult struct {
	Valid bool
	Error error
}

// Validate 验证一组验证结果，返回第一个错误
func Validate(results ...ValidationResult) error {
	for _, result := range results {
		if !result.Valid {
			return result.Error
		}
	}
	return nil
}

// StringRequired 验证字符串是否非空
func StringRequired(name, value string) ValidationResult {
	if value == "" {
		return ValidationResult{
			Valid: false,
			Error: fmt.Errorf("%s不能为空", name),
		}
	}
	return ValidationResult{Valid: true}
}

// StringMinLength 验证字符串最小长度
func StringMinLength(name, value string, minLen int) ValidationResult {
	if len(value) < minLen {
		return ValidationResult{
			Valid: false,
			Error: fmt.Errorf("%s长度必须至少为%d个字符", name, minLen),
		}
	}
	return ValidationResult{Valid: true}
}

// InRange 验证整数是否在指定范围内
func InRange(name string, value, min, max int) ValidationResult {
	if value < min || value > max {
		return ValidationResult{
			Valid: false,
			Error: fmt.Errorf("%s必须在%d到%d之间", name, min, max),
		}
	}
	return ValidationResult{Valid: true}
}
