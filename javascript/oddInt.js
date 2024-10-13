function findOddInt(arr) {
    return arr.reduce((acc, num) => acc ^ num, 0);
}

module.exports = findOddInt;

if (require.main === module) {
    const arr = [1, 2, 1, 3, 2];
    console.log("Odd integer:", findOddInt(arr)); // Output: 3
}
