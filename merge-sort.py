def mergeSort(arr):
    if len(arr) < 2:
        return arr
    
    mid = len(arr) // 2
    firstHalf = arr[:mid]
    secondHalf = arr[mid:]

    sortedFirstHalf = mergeSort(firstHalf)
    sortedSecondHalf = mergeSort(secondHalf)

    sortedArr = []

    i = 0
    j = 0
    while i < len(sortedFirstHalf) and j < len(sortedSecondHalf):
        if sortedFirstHalf[i] <= sortedSecondHalf[j]:
            sortedArr.append(sortedFirstHalf[i])
            i += 1
        else:
            sortedArr.append(sortedSecondHalf[j])
            j += 1

    sortedArr.extend(sortedFirstHalf[i:])
    sortedArr.extend(sortedSecondHalf[j:])
    
    return sortedArr