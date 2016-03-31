package counters

// 私有类型
type alertCounter int

// 提供获取私有类型的方法
func New(value int) alertCounter {
    return alertCounter(value)
}
