package intersect

func IntersectSimple(nums1 []int, nums2 []int) []int {
	var res []int
	if len(nums1) != 0 && len(nums2) != 0 {
		for i := 0; i < len(nums1); i++ {
			var isFind bool
			for j := 0; j < len(nums2); j++ {
				if nums1[i] == nums2[j] {
					// если элементы найдены
					isFind = true
					res = append(res, nums1[i])
					// удаление найденного элемента из списков
					copy(nums2[j:], nums2[j+1:])
					nums2 = nums2[:len(nums2)-1]

					copy(nums1[i:], nums1[i+1:])
					nums1 = nums1[:len(nums1)-1]
					// выход из вложенного цикла
					break
				}
			}
			if isFind {
				// перезапуск поиска
				i = -1
			}
		}
	}
	return res
}
