// [1,2,3,4,5,6,7], what add up to 9 -> 4,5
// input: arrays, up to value
// output: two indices

function addUpToNumber(arrays, upToValue) {
    hastable = {}
    for (let i = 0; i < arrays.length; i++) {
        currentValue = arrays[i]
        remainingRequired = upToValue - currentValue

        if (remainingRequired < 0) {
            throw Error('input invalid')
        }

        if (hastable[currentValue] != undefined) {
            return [i, hastable[currentValue]]
        } else {
            hastable[(remainingRequired)] = i
        }
    }
    reuturn -1
}

console.log(addUpToNumber([1,2,3,4,5,6], 9))