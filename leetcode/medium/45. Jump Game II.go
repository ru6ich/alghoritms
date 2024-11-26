package medium

// Function to calculate the minimum number of jumps [2,3,1,1,4]
func jump(nums []int) int {
    n := len(nums)
    if n == 1 {
        return 0 // Если массив из одного элемента, прыжки не нужны
    }
    
    jumps := 0       // Количество прыжков
    currentEnd := 0  // Текущий конец текущего прыжка
    farthest := 0    // Дальше всего, куда можно прыгнуть

    for i := 0; i < n-1; i++ {
        farthest = max(farthest, i+nums[i]) // Обновляем дальность прыжка
        
        // Если мы достигаем конца текущего прыжка
        if i == currentEnd {
            jumps++                     // Увеличиваем счетчик прыжков
            currentEnd = farthest       // Обновляем текущий конец
        }
 
        // Если мы достигли последнего элемента, можно выйти
        if currentEnd >= n-1 {
            break
        }
    }

    return jumps
}

// Utility function to get the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}