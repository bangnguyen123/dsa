// Ugly number is the number that have factors are only prime numbers 
// such as [2,3,5]

const factors = [2, 3, 5]

function divideNumber(number, divider) {
    while (number % divider == 0) {
        number = number / divider
    }
    return number
}

function isUglyNumber(number) {
    for (let i of factors) {
        number = divideNumber(number, i)
    }
    return number == 1
}

console.log(isUglyNumber(10))