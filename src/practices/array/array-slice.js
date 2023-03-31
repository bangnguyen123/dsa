// .slice() from scratch
// if only the begin index is found -> set endIndex is the last index of the array
// if both indice are found => return this 

function sliceArray(arrays, beginIndex, endIndex) {
    if (!beginIndex & !endIndex) {
        return arrays
    }

    if (!endIndex) {
        endIndex = arrays.length
    }

    newArr = []
    for (let i = beginIndex; i < endIndex; i++) {
        newArr.push(arrays[i])
    }
    return newArr
}

console.log(sliceArray([8,9,10,11,12,13,14], 1))