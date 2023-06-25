let arr = [3, 2, 5, 6, 7, 2, 1, 9, 10, 5]

function mergeSort(arr) {
  if (arr.length < 2) return arr

  const firstHalf = arr.slice(0, Math.floor(arr.length / 2))
  const secondHalf = arr.slice(Math.floor(arr.length / 2))

  const sortedFirstHalf = mergeSort(firstHalf)
  const sortedSecondHalf = mergeSort(secondHalf)

  arr = []
  let i = 0
  let j = 0
  while (sortedFirstHalf.length > i || sortedSecondHalf.length > j) {
    if (
      j === sortedSecondHalf.length ||
      (i < sortedFirstHalf.length && sortedFirstHalf[i] <= sortedSecondHalf[j])
    ) {
      arr.push(sortedFirstHalf[i])
      i++
    } else {
      arr.push(sortedSecondHalf[j])
      j++
    }
  }

  return arr
}

console.log(mergeSort(arr))