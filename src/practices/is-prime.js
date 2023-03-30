// Testing whether a number is a prime number.
// primeNumber = 6 * x + 1

function isPrime(n) {
    if (n < 1) {
        return false
    }
    if (n == 3 || n == 2) {
        return true
    }

    for (let i = 1; i < Math.sqrt(n); i++) {
        if (n == 6 * i + 1) {
            return true
        }
    }
    return false
}

console.log(isPrime(29))