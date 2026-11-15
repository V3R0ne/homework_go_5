package interfaces

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	ID   int
	Name string
}

// TODO: String должен вернуть пользователя в формате "user#<ID>:<Name>".
func (u User) String() string {
	return fmt.Sprintf("user#%d:%s", u.ID, u.Name)
}

// TODO: FormatOne должен вызвать String у переданного значения. Если значение nil, верни пустую строку.
func FormatOne(value fmt.Stringer) string {
	if value == nil {
		return ""
	}
	return value.String()
}

// TODO: FormatMany должен вернуть строки для всех элементов в исходном порядке. nil элементы пропускаются.
func FormatMany(values []fmt.Stringer) []string {
	answer := make([]string, 0, len(values))
	for i := range values {
		if values[i] != nil {
			answer = append(answer, FormatOne(values[i]))
		}
	}
	return answer
}

type EmailNotifier struct {
	Address string
}

// TODO: Notify должен вернуть "email:<Address>:<message>".
func (n EmailNotifier) Notify(message string) string {
	return fmt.Sprintf("email:%s:%s", n.Address, message)
}

type PushNotifier struct {
	DeviceID string
}

// TODO: Notify должен вернуть "push:<DeviceID>:<message>".
func (n PushNotifier) Notify(message string) string {
	return fmt.Sprintf("push:%s:%s", n.DeviceID, message)
}

type Notifier interface {
	Notify(message string) string
}

// TODO: SendNotification должен отправить одно сообщение через Notifier. Если Notifier nil, верни пустую строку.
func SendNotification(n Notifier, message string) string {
	if n == nil {
		return ""
	}
	return n.Notify(message)
}

// TODO: SendBatch должен отправить все сообщения и вернуть результаты в том же порядке. nil Notifier даёт пустой результат.
func SendBatch(n Notifier, messages []string) []string {
	answer := make([]string, 0, len(messages))
	if n == nil {
		return []string{}
	}
	for i := range messages {
		answer = append(answer, n.Notify(messages[i]))
	}
	return answer
}

type StaticLoader struct {
	Value string
	Err   error
}

// TODO: Load должен вернуть поля Value и Err без дополнительной обработки.
func (l StaticLoader) Load() (string, error) {
	return l.Value, l.Err
}

type Loader interface {
	Load() (string, error)
}

// TODO: LoadUpper должен загрузить строку, привести её к верхнему регистру и вернуть ошибку без потери, если загрузка не удалась.
func LoadUpper(l Loader) (string, error) {
	if l == nil {
		return "", errors.New("nil loader")
	}
	result, err := l.Load()
	if err != nil {
		return "", err
	}
	return strings.ToUpper(result), nil
}

// TODO: Describe должен вернуть описание динамического типа: nil, string:<value>, user:<String>, notifier:<Notify("ping")>, loader:<loaded>, error:<text> или unknown:<T>.
func Describe(value any) string {
	switch v := value.(type) {
	case nil:
		return "nil"
	case string:
		return fmt.Sprintf("string:%s", v)
	case User:
		return fmt.Sprintf("user:%s", v.String())
	case Notifier:
		return fmt.Sprintf("notifier:%s", v.Notify("ping"))
	case Loader:
		result, err := v.Load()
		if err != nil {
			return fmt.Sprintf("loader-error:%s", err.Error())
		}
		return fmt.Sprintf("loader:%s", result)
	case error:
		return fmt.Sprintf("error:%s", v.Error())
	default:
		return fmt.Sprintf("unknown:%T", v)

	}
}

// TODO: OnlyErrors должен вернуть тексты только ненулевых ошибок в исходном порядке.
func OnlyErrors(values []error) []string {
	answer := make([]string, 0, len(values))
	for i := range values {
		if values[i] != nil {
			answer = append(answer, values[i].Error())
		}
	}
	return answer
}

// TODO: JoinStringers должен отформатировать ненулевые Stringer и объединить их через sep.
func JoinStringers(values []fmt.Stringer, sep string) string {
	strs := make([]string, 0, len(values))
	for i := range values {
		if values[i] != nil {
			strs = append(strs, values[i].String())
		}
	}
	return strings.Join(strs, sep)
}

// TODO: FirstNonEmptyStringer должен вернуть первую непустую строку среди Stringer. nil и пустые строки пропускаются.
func FirstNonEmptyStringer(values []fmt.Stringer) string {
	for i := range values {
		if values[i] != nil {
			if values[i].String() != "" {
				return values[i].String()
			}
		}
	}
	return ""
}

// TODO: CountStringers должен посчитать, сколько элементов реально реализуют fmt.Stringer.
func CountStringers(values []any) int {
	count := 0
	for i := range values {
		if _, ok := values[i].(fmt.Stringer); ok {
			count++
		}
	}
	return count
}

// TODO: NormalizeAndFormat должен обрезать пробелы вокруг String() и схлопнуть пустой результат в "empty".
func NormalizeAndFormat(value fmt.Stringer) string {
	if value == nil {
		return "empty"
	}
	result := strings.TrimSpace(value.String())
	if result == "" {
		return "empty"
	}
	return result
}

func Example() string {
	var out strings.Builder
	fmt.Fprintf(&out, "user: %s\n", FormatOne(User{ID: 7, Name: "Maria"}))
	fmt.Fprintf(&out, "email: %s\n", SendNotification(EmailNotifier{Address: "team@example.com"}, "hello"))
	fmt.Fprintf(&out, "push: %s\n", SendNotification(PushNotifier{DeviceID: "ios-1"}, "hello"))
	fmt.Fprintf(&out, "describe: %s", Describe(User{ID: 8, Name: "Alex"}))
	return out.String()
}
