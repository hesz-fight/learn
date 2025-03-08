package rotate

//从外到里 逐层旋转
func rotate(matrix [][]int) {
	// 计算层数
	layers := len(matrix) / 2
	for layer := 0; layer < layers; layer++ {
		rotateLayer(matrix, layer)
	}
}

func rotateLayer(matrix [][]int, layer int) {
	min, max := layer, len(matrix)-layer-1
	for i := 0; i < max-min; i++ {
		// 缓存左上角元素
		tmp := matrix[min][min+i]
		// 左下角元素赋值给左上角
		matrix[min][min+i] = matrix[max-i][min]
		// 右下角元素赋值给左下角
		matrix[max-i][min] = matrix[max][max-i]
		// 右上角元素赋值给右下角
		matrix[max][max-i] = matrix[min+i][max]
		// 左上角元素赋值给右上角
		matrix[min+i][max] = tmp
	}
}
