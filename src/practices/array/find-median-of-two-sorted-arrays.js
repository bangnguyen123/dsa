// We have two arrays with the same size, both are sorted
// Find the median of the array obtained after merging the about 2 arrays

function getMax(arr, partition) {
    if (partition == 0) {
        return Number.NEGATIVE_INFINITY
    } else {
        return arr[partition - 1]
    }
}

function getMin(arr, partition) {
    if (partition == arr.length) {
        return Number.POSITIVE_INFINITY
    } else {
        return arr[partition]
    }
}

function findMedianTwoSortedArrays(arr1, arr2) {
    if (arr1.length > arr2.length) {
        let temp = arr1
        arr1 = arr2
        arr2 = temp
    }

    let low = 0
    let high = arr1.length
    let combinedLength = arr1.length + arr2.length

    while (low <= high) {
        let partition1 = Math.floor((low + high) / 2)
        let partition2 = Math.floor((combinedLength )/2) - partition1

        let leftElement1 = getMax(arr1, partition1)
        let rightElement1 = getMin(arr1, partition1)
        let leftElement2 = getMax(arr2, partition2)
        let rightElement2 = getMin(arr2, partition2)

        if (leftElement1 <= rightElement2 <= rightElement1) {
            if (combinedLength % 2 == 0) {
                return (Math.max(leftElement1, leftElement2) + Math.min(rightElement1, rightElement2)) / 2.0
            }
            
            return Math.max(leftElement1, leftElement2)
        }
        if (leftElement1 > rightElement2) {
            high = partition1 - 1
        } else {
            low = partition1 + 1
        }
    }
    return -1
}

x = findMedianTwoSortedArrays([1], [4, 5])
console.log(x)